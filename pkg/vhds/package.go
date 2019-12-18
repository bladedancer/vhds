package vhds

import (
	"github.com/bladedancer/vhds/pkg/base"
	"github.com/sirupsen/logrus"
)

var log logrus.FieldLogger
var config *base.Config

// Init initializes the config and log
func Init(logger *logrus.Logger) {
	log = logger.WithField("package", "vhds")
	config = base.GetConfig()
}
