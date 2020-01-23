package vhds

import (
	"encoding/json"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/guiguan/caster"

	"github.com/bladedancer/vhds/pkg/central"
	"github.com/bladedancer/vhds/pkg/xdsconfig"
	api "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/golang/protobuf/ptypes"
	any "github.com/golang/protobuf/ptypes/any"
)

// CentralResource the details of the central resource.
type CentralResource struct {
	Name     string
	Version  string
	Resource *any.Any
}

// Subscription details
type Subscription struct {
	Name    string
	Version string
}

// AxwayVirtualHostDiscoveryServiceServer implements the Stream Access Logs endpoint.
type AxwayVirtualHostDiscoveryServiceServer struct {
	resources map[string]*CentralResource // Persist in memory ... what could go wrong
	caster    *caster.Caster
}

// NewAxwayVirtualHostDiscoveryServiceServer construct a new AxwayVirtualHostDiscoveryServiceServer.
func NewAxwayVirtualHostDiscoveryServiceServer() *AxwayVirtualHostDiscoveryServiceServer {
	var svc = &AxwayVirtualHostDiscoveryServiceServer{}
	svc.resources = make(map[string]*CentralResource)
	svc.caster = caster.New(nil)
	svc.watchCentral()
	return svc
}

// watchCentral for changes.
func (svc *AxwayVirtualHostDiscoveryServiceServer) watchCentral() {
	listenerUpdateChan := central.Watch(config.ReadinessChan)

	// Listen for changes in central and broadcast them to subscribed clients
	go func() {
		for listeners := range listenerUpdateChan {
			// Convert Updates to Resources
			var updated []string
			for _, listener := range listeners {
				rtg := xdsconfig.MakeRuntimeGroup(listener)
				for _, vhost := range rtg.VirtualHosts {
					vh, _ := ptypes.MarshalAny(vhost)
					for _, domain := range vhost.GetDomains() {
						resource := &CentralResource{
							Name:     fmt.Sprintf("local_route/%s", domain), // local_route is the RouteConfiguration name
							Version:  "todo",                                // TODO
							Resource: vh,
						}
						svc.resources[resource.Name] = resource
						log.Infof("Available route: %s", resource.Name)
						updated = append(updated, resource.Name)
					}
				}
			}
			// Emit Updates to subscribers
			svc.notifySubscribers(updated)
		}
	}()
}

func (svc *AxwayVirtualHostDiscoveryServiceServer) notifySubscribers(changed []string) {
	if len(changed) > 0 {
		svc.caster.Pub(changed)
	}
}

// DeltaVirtualHosts serves the virtual host delta.
func (svc *AxwayVirtualHostDiscoveryServiceServer) DeltaVirtualHosts(srv api.VirtualHostDiscoveryService_DeltaVirtualHostsServer) (err error) {
	reqCh := make(chan *api.DeltaDiscoveryRequest)
	reqStop := int32(0)

	// Just keep the list of subscriptions for the connection in the closure.
	subscriptions := make(map[string]*Subscription)

	// TODO: Refactor
	// List for incoming requests
	go func() {
		for {
			req, err := srv.Recv()
			if atomic.LoadInt32(&reqStop) != 0 {
				return
			}
			if err != nil {
				close(reqCh)
				return
			}
			reqCh <- req
		}
	}()

	// Subscribe for central resource updates
	updateCh, _ := svc.caster.Sub(srv.Context(), 1)

	// Process requests
	for {
		done := false
		select {
		case <-srv.Context().Done():
			done = true

		case msg := <-updateCh:
			names := msg.([]string)
			// TODO: Honor version
			var filteredNames []string
			for _, name := range names {
				if _, exists := subscriptions[name]; exists {
					filteredNames = append(filteredNames, name)
				}
			}

			err = svc.sendCentralResources(srv, filteredNames)
			if err != nil {
				// WHAT TO DO???
				log.Error(err.Error())
				break
			}

		case req := <-reqCh:
			reqStr, err := json.Marshal(req)
			log.Infof("Msg: %s", reqStr)
			if err != nil {
				log.Error(err.Error())
				break
			}

			// ACK/NACK?
			if req.GetErrorDetail() != nil {
				// NACK
				// We could lookup the nonce, figure out what failed, figure out what to do about it etc.
				// ... but probably will just resend.
				log.Error("NACK")
			} else {
				// ACK
				log.Info("ACK")

				var updates []string
				if len(req.GetInitialResourceVersions()) > 0 {
					// XDS could have restarted, the client is telling us the versions
					// it currently has.
					for name, version := range req.GetInitialResourceVersions() {
						log.Infof("Node: %+v is subscribing to: %s", req.GetNode(), name)
						subscriptions[name] = &Subscription{Name: name, Version: version}
						updates = append(updates, name)
					}
				}

				if len(req.GetResourceNamesSubscribe()) > 0 {
					// Client is subscribing
					// Could have rules on what you're allowed to subscribe too...
					for _, name := range req.GetResourceNamesSubscribe() {
						if _, subscribed := subscriptions[name]; !subscribed {
							log.Infof("Node: %+v is subscribing to: %s", req.GetNode(), name)
							subscriptions[name] = &Subscription{Name: name}
							updates = append(updates, name)
						}
					}
				}

				if len(req.GetResourceNamesUnsubscribe()) > 0 {
					// Client is unsubscribing
					for _, name := range req.GetResourceNamesUnsubscribe() {
						log.Infof("Node: %+v is unsubscribing from: %s", req.GetNode(), name)
						delete(subscriptions, name)
					}
				}

				err = svc.sendCentralResources(srv, updates)
				if err != nil {
					// WHAT TO DO???
					log.Error(err.Error())
					break
				}
			}
		}

		if done {
			break
		}
	}

	atomic.StoreInt32(&reqStop, 1)
	return err
}

// sendCentralResources deliver the resources to the listening clients if they're subscribed.
func (svc *AxwayVirtualHostDiscoveryServiceServer) sendCentralResources(srv api.VirtualHostDiscoveryService_DeltaVirtualHostsServer, names []string) (err error) {
	var updates []*api.Resource
	for _, name := range names {
		if resource, exists := svc.resources[name]; exists {
			updates = append(updates, &api.Resource{
				Name:     resource.Name,
				Version:  resource.Version,
				Aliases:  []string{},
				Resource: resource.Resource,
			})
		}
	}

	// Send them all
	if len(updates) > 0 {
		err = srv.Send(&api.DeltaDiscoveryResponse{
			TypeUrl:           "type.googleapis.com/envoy.api.v2.route.VirtualHost",
			SystemVersionInfo: "test",
			Resources:         updates,
			RemovedResources:  []string{},
			Nonce:             fmt.Sprintf("%d", time.Now().UnixNano()),
		})
	}

	return err
}
