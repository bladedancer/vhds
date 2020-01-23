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
	config := makeRoutingCluster(config.NumShards) // should be discovering backends
	resource := []cache.Resource{config}
	resources = append(resources, resource...)
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
								ClusterSpecifier: &route.RouteAction_Cluster{
									Cluster: "back",
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
func makeRoutingCluster(shards int) *api.Cluster {
	clusterName := "back"
	var endpoints []*endpoint.LocalityLbEndpoints

	for i := 0; i < shards; i++ {
		address := &core.Address{Address: &core.Address_SocketAddress{
			SocketAddress: &core.SocketAddress{
				Address: fmt.Sprintf("back-%d.back", i), // TODO discover endpoints
				PortSpecifier: &core.SocketAddress_PortValue{
					PortValue: 80, // TODO
				},
			},
		}}
		endpoint := &endpoint.LocalityLbEndpoints{
			LbEndpoints: []*endpoint.LbEndpoint{
				&endpoint.LbEndpoint{
					HostIdentifier: &endpoint.LbEndpoint_Endpoint{
						Endpoint: &endpoint.Endpoint{
							Address: address,
						},
					},
				},
			},
		}
		endpoints = append(endpoints, endpoint)
	}

	return &api.Cluster{
		Name:                 clusterName,
		ConnectTimeout:       ptypes.DurationProto(5 * time.Second),
		ClusterDiscoveryType: &api.Cluster_Type{Type: api.Cluster_STRICT_DNS},
		DnsLookupFamily:      api.Cluster_V4_ONLY,
		RespectDnsTtl:        config.RespectDNSTTL,
		DnsRefreshRate:       ptypes.DurationProto(time.Duration(config.DNSRefreshRate) * time.Millisecond),
		LbPolicy:             api.Cluster_MAGLEV,
		LoadAssignment: &api.ClusterLoadAssignment{
			ClusterName: clusterName,
			Endpoints:   endpoints,
		},
		TlsContext:      nil,
		TransportSocket: nil,
	}
}
