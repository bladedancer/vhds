package vhds

import (
	"encoding/json"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/bladedancer/vhds/pkg/central"
	"github.com/bladedancer/vhds/pkg/xdsconfig"
	api "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/golang/protobuf/ptypes"
	any "github.com/golang/protobuf/ptypes/any"
)

// AxwayVirtualHostDiscoveryServiceServer implements the Stream Access Logs endpoint.
type AxwayVirtualHostDiscoveryServiceServer struct {
}

// DeltaVirtualHosts serves the virtual host delta.
func (*AxwayVirtualHostDiscoveryServiceServer) DeltaVirtualHosts(srv api.VirtualHostDiscoveryService_DeltaVirtualHostsServer) (err error) {
	reqCh := make(chan *api.DeltaDiscoveryRequest)
	reqStop := int32(0)

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

	for {
		done := false
		select {
		case <-srv.Context().Done():
			done = true
		case req := <-reqCh:
			reqStr, err := json.Marshal(req)
			log.Infof("Msg: %s", reqStr)
			if err != nil {
				break
			}

			if req.GetResponseNonce() == "" {
				// Initial response
				err = srv.Send(&api.DeltaDiscoveryResponse{
					SystemVersionInfo: "test",
					Resources: []*api.Resource{
						&api.Resource{
							Name:     "foo",
							Version:  "1",
							Aliases:  []string{"Foo"},
							Resource: getFakeVHost(1),
						},
					},
					TypeUrl:          "type.googleapis.com/envoy.api.v2.route.VirtualHost",
					RemovedResources: []string{},
					Nonce:            fmt.Sprintf("%d", time.Now().UnixNano()), //deltaReq.ResponseNonce,
				})
			} else {
				// ACK/NACK?
				log.Info("ACK/NACK?")
				// on ack we'd wait for something to do, on nack we'd need retry/error logic
			}
		}

		if done {
			break
		}
	}

	atomic.StoreInt32(&reqStop, 1)
	return err
}

func getFakeVHost(i int) *any.Any {
	rtg := xdsconfig.MakeRuntimeGroup(fakeGetListener(i))
	vh, _ := ptypes.MarshalAny(rtg.VirtualHosts[0])
	return vh
}

func fakeGetListener(i int) *central.Listener {
	listener := &central.Listener{
		ID:             fmt.Sprintf("id-%d", i),
		Activated:      true,
		Name:           fmt.Sprintf("id-%d", i),
		Protocol:       "http",
		BindAddress:    "0.0.0.0",
		Port:           "80",
		VirtualHosts:   []string{"test", "prod"},
		RuntimeGroupID: fmt.Sprintf("%d", i),
		Metadata:       nil,
		InstanceName:   fmt.Sprintf("inst-%d", i%2),
		TenantID:       fmt.Sprintf("ten-%d", i%2),
	}

	return listener
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
