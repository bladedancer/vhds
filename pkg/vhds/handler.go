package vhds

import (
	"encoding/json"
	"fmt"
	"sync/atomic"
	"time"

	api "github.com/envoyproxy/go-control-plane/envoy/api/v2"
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
					Resources:         []*api.Resource{},
					TypeUrl:           "type.googleapis.com/envoy.api.v2.route.VirtualHost",
					RemovedResources:  []string{},
					Nonce:             fmt.Sprintf("%d", time.Now().UnixNano()), //deltaReq.ResponseNonce,
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
