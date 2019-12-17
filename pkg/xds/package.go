package xds

import (
	"github.com/bladedancer/xdsing/pkg/base"
	"github.com/sirupsen/logrus"
)

var log logrus.FieldLogger
var config *base.Config

// Init initializes the XDS package.
func Init(logger *logrus.Logger) {
	log = logger.WithField("package", "xds")
	config = base.GetConfig()
}
