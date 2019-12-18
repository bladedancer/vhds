package xds

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/bladedancer/vhds/pkg/accesslog"
	"github.com/bladedancer/vhds/pkg/vhds"
	api "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	als "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v2"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v2"
	"github.com/envoyproxy/go-control-plane/pkg/cache"
	xds "github.com/envoyproxy/go-control-plane/pkg/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Run entry point for Envoy XDS command line.
func Run() error {

	callbacks := Calls{}
	snapshotCache := cache.NewSnapshotCache(false, cache.IDHash{}, nil)
	server := xds.NewServer(context.Background(), snapshotCache, callbacks)
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", config.Port))
	if err != nil {
		log.Fatal(err)
	}

	discovery.RegisterAggregatedDiscoveryServiceServer(grpcServer, server)
	api.RegisterEndpointDiscoveryServiceServer(grpcServer, server)
	api.RegisterClusterDiscoveryServiceServer(grpcServer, server)
	api.RegisterRouteDiscoveryServiceServer(grpcServer, server)
	api.RegisterListenerDiscoveryServiceServer(grpcServer, server)

	// Virtual Host Discovery Service
	api.RegisterVirtualHostDiscoveryServiceServer(grpcServer, &vhds.AxwayVirtualHostDiscoveryServiceServer{})

	// Configure the Access Log server.
	als.RegisterAccessLogServiceServer(grpcServer, &accesslog.Handler{})

	watch(snapshotCache)

	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	log.Infof("Listening on %d", config.Port)
	sig := <-gracefulStop
	log.Debugf("Got signal: %s", sig)
	grpcServer.GracefulStop()
	log.Info("Shutdown")
	return nil
}
