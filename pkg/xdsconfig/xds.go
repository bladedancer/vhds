package xdsconfig

import (
	"github.com/envoyproxy/go-control-plane/pkg/cache"
)

// XDS The xds resources
type XDS struct {
	LDS []cache.Resource
	CDS []cache.Resource
	RDS []cache.Resource
}

// Add an XDS to this xds.
func (xds *XDS) Add(other *XDS) {
	xds.LDS = append(xds.LDS, other.LDS...)
	xds.CDS = append(xds.CDS, other.CDS...)
	xds.RDS = append(xds.RDS, other.RDS...)
}
