package xds

import (
	"context"

	v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
)

// Calls The XDS callbacks.
type Calls struct{}

// OnStreamOpen is called once an xDS stream is open with a stream ID and the type URL (or "" for ADS).
// Returning an error will end processing and close the stream. OnStreamClosed will still be called.
func (c Calls) OnStreamOpen(context.Context, int64, string) error {
	log.Info("OnStreamOpen")
	return nil
}

// OnStreamClosed is called immediately prior to closing an xDS stream with a stream ID.
func (c Calls) OnStreamClosed(int64) {
	log.Info("OnStreamClosed")
}

// OnStreamRequest is called once a request is received on a stream.
// Returning an error will end processing and close the stream. OnStreamClosed will still be called.
func (c Calls) OnStreamRequest(id int64, req *v2.DiscoveryRequest) error {
	log.Infof("OnStreamRequest id:%d %v", id, req.ResourceNames)
	return nil
}

// OnStreamResponse is called immediately prior to sending a response on a stream.
func (c Calls) OnStreamResponse(id int64, req *v2.DiscoveryRequest, res *v2.DiscoveryResponse) {
	log.Infof("OnStreamResponse id:%d %v", id, req.ResourceNames)
}

// OnFetchRequest is called for each Fetch request. Returning an error will end processing of the
// request and respond with an error.
func (c Calls) OnFetchRequest(context.Context, *v2.DiscoveryRequest) error {
	log.Info("OnFetchRequest")
	return nil
}

// OnFetchResponse is called immediately prior to sending a response.
func (c Calls) OnFetchResponse(*v2.DiscoveryRequest, *v2.DiscoveryResponse) {
	log.Info("OnFetchRequest")
}
