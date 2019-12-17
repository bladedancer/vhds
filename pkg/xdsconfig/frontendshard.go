package xdsconfig

import (
	"fmt"
	"time"

	api "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	endpoint "github.com/envoyproxy/go-control-plane/envoy/api/v2/endpoint"
	route "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	"github.com/envoyproxy/go-control-plane/pkg/cache"
	"github.com/golang/protobuf/ptypes"
)

// FrontendShard is a configuration fragment for frontend routing.
type FrontendShard struct {
	Name    string
	Gateway *Gateway
}

// MakeFrontendShard create a frontend shard.
func MakeFrontendShard(name string) *FrontendShard {
	return &FrontendShard{
		Name:    name,
		Gateway: MakeFrontendGateway(),
	}
}

// GetName Get the shard name
func (s *FrontendShard) GetName() string {
	return s.Name
}

// GetXDS Convert the configuration to resources
func (s *FrontendShard) GetXDS() *XDS {
	return &XDS{
		LDS: s.getListenerResources(),
		CDS: s.getClusterResources(),
		RDS: s.getRouteResources(),
	}
}

// getListenerResources Get the listener configuration data
func (s *FrontendShard) getListenerResources() []cache.Resource {
	resources := []cache.Resource{s.Gateway.Listener}
	return resources
}

// getClusterResources Get cluster configuration data
func (s *FrontendShard) getClusterResources() []cache.Resource {
	var resources []cache.Resource

	for i := 0; i < config.NumShards; i++ {
		config := makeRoutingCluster(fmt.Sprintf("back-%d", i)) // embedding "back" not ideal
		resource := []cache.Resource{config}
		resources = append(resources, resource...)
	}

	return resources
}

// getRouteResources Get the route configuration data
func (s *FrontendShard) getRouteResources() []cache.Resource {
	// At the moment this uses clusterheader to decide and has no
	// business logic. The listener should have a filter configured
	// to decide the correct shard and then route based on that.
	config := &api.RouteConfiguration{
		Name: "local_route",
		VirtualHosts: []*route.VirtualHost{
			&route.VirtualHost{
				Name:    "front",
				Domains: []string{"*"},
				Routes: []*route.Route{
					&route.Route{
						Name: "front",
						Match: &route.RouteMatch{
							PathSpecifier: &route.RouteMatch_Prefix{
								Prefix: "/",
							},
						},
						Action: &route.Route_Route{
							Route: &route.RouteAction{
								ClusterSpecifier: &route.RouteAction_ClusterHeader{
									ClusterHeader: "x-shard", // TODO
								},
							},
						},
					},
				},
			},
		},
	}
	resources := []cache.Resource{config}
	return resources
}

// Create a cluster for the proxy
func makeRoutingCluster(shardName string) *api.Cluster {
	//log.Infof("creating cluster for proxy: %s", proxy.Name)
	address := &core.Address{Address: &core.Address_SocketAddress{
		SocketAddress: &core.SocketAddress{
			Address: fmt.Sprintf("%s.back", shardName), // TODO shardName === pod name === clustername, not ideal
			PortSpecifier: &core.SocketAddress_PortValue{
				PortValue: 80, // TODO
			},
		},
	}}

	return &api.Cluster{
		Name:                 shardName,
		ConnectTimeout:       ptypes.DurationProto(5 * time.Second),
		ClusterDiscoveryType: &api.Cluster_Type{Type: api.Cluster_LOGICAL_DNS},
		DnsLookupFamily:      api.Cluster_V4_ONLY,
		RespectDnsTtl:        config.RespectDNSTTL,
		DnsRefreshRate:       ptypes.DurationProto(time.Duration(config.DNSRefreshRate) * time.Millisecond),
		LbPolicy:             api.Cluster_ROUND_ROBIN,
		LoadAssignment: &api.ClusterLoadAssignment{
			ClusterName: shardName,
			Endpoints: []*endpoint.LocalityLbEndpoints{
				&endpoint.LocalityLbEndpoints{
					LbEndpoints: []*endpoint.LbEndpoint{
						&endpoint.LbEndpoint{
							HostIdentifier: &endpoint.LbEndpoint_Endpoint{
								Endpoint: &endpoint.Endpoint{
									Address: address,
								},
							},
						},
					},
				},
			},
		},
		TlsContext:      nil,
		TransportSocket: nil,
	}
}
