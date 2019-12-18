package accesslog

import (
	"github.com/bladedancer/vhds/pkg/base"
	"github.com/sirupsen/logrus"
)

var log logrus.FieldLogger
var config *base.Config

// Init initializes the Access Server
func Init(logger *logrus.Logger) {
	log = logger.WithField("package", "accesslog")
	config = base.GetConfig()
}
