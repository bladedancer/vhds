package xdsconfig

import (
	api "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	route "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	"github.com/envoyproxy/go-control-plane/pkg/cache"
)

// BackendShard is a configuration fragment for tenants on a single envoy node.
type BackendShard struct {
	Name          string
	RuntimeGroups []*RuntimeGroup
	Gateway       *Gateway
}

// MakeBackendShard create a backend shard.
func MakeBackendShard(name string) *BackendShard {
	return &BackendShard{
		Name:          name,
		RuntimeGroups: []*RuntimeGroup{},
		Gateway:       MakeBackendGateway(),
	}
}

// GetName Get the shard name
func (s *BackendShard) GetName() string {
	return s.Name
}

// GetXDS Convert the configuration to resources
func (s *BackendShard) GetXDS() *XDS {
	return &XDS{
		LDS: s.getListenerResources(),
		RDS: s.getRouteResources(),
	}
}

// getListenerResources Get the listener configuration data
func (s *BackendShard) getListenerResources() []cache.Resource {
	resources := []cache.Resource{s.Gateway.Listener}
	return resources
}

// getRouteResources Get the route configuration data
func (s *BackendShard) getRouteResources() []cache.Resource {
	var vhosts []*route.VirtualHost

	for _, rtg := range s.RuntimeGroups {
		vhosts = append(vhosts, rtg.VirtualHosts...)
	}

	// Create the Route Configuration for the shard
	config := &api.RouteConfiguration{
		Name:         "local_route",
		VirtualHosts: vhosts,
	}
	resources := []cache.Resource{config}
	return resources
}
