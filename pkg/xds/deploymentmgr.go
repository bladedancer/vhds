package xds

import (
	"fmt"

	"github.com/bladedancer/xdsing/pkg/central"
	"github.com/bladedancer/xdsing/pkg/xdsconfig"
)

// Deployment deployment of a listener on a shard
type Deployment struct {
	shardName       string
	listener        *central.Listener
	xdsRuntimeGroup *xdsconfig.RuntimeGroup
}

// DeploymentManager maps tenants to clusters
type DeploymentManager struct {
	byVhostInstance map[string]*Deployment
	nextShard       int
	shards          map[string]*xdsconfig.BackendShard
	OnChange        chan []*xdsconfig.BackendShard
}

// MakeDeploymentManager create a DeplyomentManager.
func MakeDeploymentManager() *DeploymentManager {
	return &DeploymentManager{
		byVhostInstance: make(map[string]*Deployment),
		nextShard:       0,
		shards:          make(map[string]*xdsconfig.BackendShard),
		OnChange:        make(chan []*xdsconfig.BackendShard),
	}
}

// GetShardName gets the shard that the tenant is deployed on
func (dm *DeploymentManager) GetShardName(vhostInstance string) string {
	dep := dm.byVhostInstance[vhostInstance]
	shard := ""
	if dep != nil {
		shard = dep.shardName
	}
	return shard
}

// GetShard return the named shard.
func (dm *DeploymentManager) GetShard(name string) *xdsconfig.BackendShard {
	return dm.shards[name]
}

// GetShards return the shards for the deployed tenants.
func (dm *DeploymentManager) GetShards() []*xdsconfig.BackendShard {
	backendShards := make([]*xdsconfig.BackendShard, 0, len(dm.shards))
	for _, value := range dm.shards {
		backendShards = append(backendShards, value)
	}
	return backendShards
}

// updateShard updates the backend configuration for the specified shard.
func (dm *DeploymentManager) updateShard(shardName string) {
	if _, exists := dm.shards[shardName]; !exists {
		dm.shards[shardName] = xdsconfig.MakeBackendShard(shardName)
	}

	runtimeGroups := []*xdsconfig.RuntimeGroup{}
	for id := range dm.byVhostInstance {
		if dm.byVhostInstance[id].shardName == shardName {
			runtimeGroups = append(runtimeGroups, dm.byVhostInstance[id].xdsRuntimeGroup)
		}
	}
	dm.shards[shardName].RuntimeGroups = runtimeGroups
}

// RemoveDeployment removes the deployment.
func (dm *DeploymentManager) RemoveDeployment(vhost string, instanceName string) {
	key := fmt.Sprintf("%s-%s", vhost, instanceName)
	if deployment, exists := dm.byVhostInstance[key]; exists {
		delete(dm.byVhostInstance, key)
		dm.updateShard(deployment.shardName)
		dm.OnChange <- ([]*xdsconfig.BackendShard{dm.shards[deployment.shardName]})
	}
}

// AddListeners adds the listeners to the deployment.
func (dm *DeploymentManager) AddListeners(listeners ...*central.Listener) {
	dirtyShards := make(map[string]bool)

	for _, listener := range listeners {
		var shardName string

		// Is this an update of an already deployed listener, if so don't move shards
		for _, vhost := range listener.VirtualHosts {
			key := fmt.Sprintf("%s-%s", vhost, listener.InstanceName)
			if curdep, exists := dm.byVhostInstance[key]; exists {
				shardName = curdep.shardName
				break
			}
		}
		// If not then assign to next available shard
		if shardName == "" {
			shardName = dm.getNextShard()
		}
		dirtyShards[shardName] = true

		deployment := &Deployment{
			shardName:       shardName,
			listener:        listener,
			xdsRuntimeGroup: xdsconfig.MakeRuntimeGroup(listener),
		}

		for _, vhost := range listener.VirtualHosts {
			key := fmt.Sprintf("%s-%s", vhost, listener.InstanceName)
			log.Infof("Adding deployment: %s=%+v", key, deployment)
			dm.byVhostInstance[key] = deployment
		}
	}

	// Update the dirty shards
	for shardName := range dirtyShards {
		dm.updateShard(shardName)
		dm.OnChange <- ([]*xdsconfig.BackendShard{dm.shards[shardName]})
	}
}

// getNextShard Get the name of the next shard to assign a tenant too.
func (dm *DeploymentManager) getNextShard() string {
	shardName := fmt.Sprintf("back-%d", dm.nextShard)
	dm.nextShard = (dm.nextShard + 1) % config.NumShards
	return shardName
}
