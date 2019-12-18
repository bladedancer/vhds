package xds

import (
	"fmt"
	"time"

	"github.com/bladedancer/vhds/pkg/xdsconfig"
	"github.com/envoyproxy/go-control-plane/pkg/cache"
)

var frontend *xdsconfig.FrontendShard
var backends []*xdsconfig.BackendShard

/*
  There's a number of stateful aspects here that will need to be fixed to support replicas of XDS.
  The version is generated at the time of the snapshot and not linked to it's content - so different
  replicas will have different versions and cause thrashing. Note even if this is resolved - need to
  ensure that when two replicas are running with different versions that it doesn't mess things up.
  Also the decision on where to deploy the listener is handled in memory - a different replica could
  in theory decide differently.
*/

func version() string {
	return fmt.Sprintf("%d", time.Now().UnixNano()) // good enough for now
}

func watch(snapshotCache cache.SnapshotCache) {
	frontend = xdsconfig.MakeFrontendShard("front")
	for i := 0; i < config.NumShards; i++ {
		backends = append(backends, xdsconfig.MakeBackendShard(fmt.Sprintf("back-%d", i)))
	}

	updateShard(snapshotCache, frontend)
	for _, shard := range backends {
		updateShard(snapshotCache, shard)
	}
}

// func watchCentral(snapshotCache cache.SnapshotCache) {
// 	// TODO: Readiness probably should be composite - both xds and sync
// 	listenerUpdateChan := central.Watch(config.ReadinessChan)

// 	// Frontend is static for now
// 	frontend = xdsconfig.MakeFrontendShard("front")
// 	updateShard(snapshotCache, frontend)

// 	// Tenants and shard contents are dynamic so listen for changes and update accordingly
// 	go func() {
// 		for shards := range deploymentManager.OnChange {
// 			for i := 0; i < len(shards); i++ {
// 				updateShard(snapshotCache, shards[i])
// 			}
// 		}
// 	}()

// 	// On a central update, update deployment
// 	go func() {
// 		for listeners := range listenerUpdateChan {
// 			deploymentManager.AddListeners(listeners...)
// 		}
// 	}()
// }

// updateShard Update the snapshot cache with the shard details.
func updateShard(snapshotCache cache.SnapshotCache, shard xdsconfig.Shard) error {
	xds := shard.GetXDS()
	log.Infof("Updating shard %s (%d:%d:%d)", shard.GetName(), len(xds.CDS), len(xds.RDS), len(xds.LDS))
	err := snapshotCache.SetSnapshot(shard.GetName(), cache.NewSnapshot(version(), nil, xds.CDS, xds.RDS, xds.LDS, nil))
	if err != nil {
		log.Error(err)
	}

	return err
}
