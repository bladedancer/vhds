package xdsconfig

// Shard configuration for a single envoy node
type Shard interface {
	GetName() string
	GetXDS() *XDS
}
