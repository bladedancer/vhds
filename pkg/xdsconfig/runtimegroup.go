package xdsconfig

import (
	"fmt"

	"github.com/bladedancer/vhds/pkg/central"
	uuid "github.com/satori/go.uuid"

	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	route "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
)

// RuntimeGroup representation of a runtime group in envoy.
type RuntimeGroup struct {
	ID           string
	Name         string
	InstanceName string
	ListenerID   string // TODO 1:N RTG->Listener
	VirtualHosts []*route.VirtualHost
}

// MakeRuntimeGroup create a RuntimeGroup.
func MakeRuntimeGroup(centralListener *central.Listener) *RuntimeGroup {
	return &RuntimeGroup{
		ID:           centralListener.RuntimeGroupID,
		Name:         centralListener.Name,
		InstanceName: centralListener.InstanceName,
		ListenerID:   centralListener.ID,
		VirtualHosts: makeVirtualHosts(centralListener),
	}
}

func makeVirtualHosts(centralListener *central.Listener) []*route.VirtualHost {
	if !centralListener.Activated {
		return []*route.VirtualHost{}
	}

	domains := []string{}
	for _, env := range centralListener.VirtualHosts {
		domains = append(
			domains,
			fmt.Sprintf("%s-%s.%s", env, centralListener.InstanceName, config.Domain),
		)
	}

	vhost := &route.VirtualHost{
		Name:    centralListener.ID,
		Domains: domains,
		Routes:  []*route.Route{makeRoute(centralListener.ID)},
		RequestHeadersToAdd: []*core.HeaderValueOption{
			&core.HeaderValueOption{
				Header: &core.HeaderValue{
					Key:   "X-Axway-Tenant-Id",
					Value: centralListener.TenantID,
				},
			},
			&core.HeaderValueOption{
				Header: &core.HeaderValue{
					Key:   "X-Axway-Runtime-Group-Id",
					Value: centralListener.RuntimeGroupID,
				},
			},
			&core.HeaderValueOption{
				Header: &core.HeaderValue{
					Key:   "X-Axway-Instance-Id",
					Value: centralListener.InstanceName,
				},
			},
			&core.HeaderValueOption{
				Header: &core.HeaderValue{
					Key:   "X-Axway-Transaction-Id",
					Value: uuid.NewV4().String(), // TODO: Move to standard obs/tracing
				},
			},
			&core.HeaderValueOption{ // DELETE THIS AND UPDATE BROKER TO USE X-FORWARDED-FOR
				Header: &core.HeaderValue{
					Key:   "forwarded",
					Value: "kindaNotUsed",
				},
			},
		},
	}

	return []*route.VirtualHost{vhost}
}

// makeRoute everything is routed to API Broker in this scenario.
func makeRoute(id string) *route.Route {
	return &route.Route{
		Name: fmt.Sprintf("%s-broker", id),
		Match: &route.RouteMatch{
			PathSpecifier: &route.RouteMatch_Prefix{
				Prefix: "/",
			},
		},
		Action: &route.Route_Route{
			Route: &route.RouteAction{
				ClusterSpecifier: &route.RouteAction_Cluster{
					Cluster: "service_apibroker",
				},
			},
		},
	}
}
