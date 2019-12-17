package central

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

var log logrus.FieldLogger
var config *Central

// Init initializes the base resources.
func Init(logger *logrus.Logger, c *Central) {
	log = logger.WithField("package", "central")
	config = c
	log.Infof("Central config: %+v", config)

	var protocol string
	if config.SyncTLS {
		protocol = "https"
	} else {
		protocol = "http"
	}

	// Derived/static config
	config.timestampHeader = "x-axway-configs-ts"
	config.syncURL = fmt.Sprintf("%s://%s:%d/api/v1/sync", protocol, config.SyncHost, config.SyncPort)
	config.client = initClient(config.SyncTLS, config.SyncTimeout)
}

func initClient(tlsEnabled bool, timeout uint32) (client *http.Client) {
	tr := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: time.Duration(timeout) * time.Second,
		}).Dial,
	}

	if tlsEnabled {
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		tr.TLSHandshakeTimeout = time.Duration(timeout) * time.Second
	}

	client = &http.Client{
		Transport: tr,
		Timeout:   0, // Patience is a virtue, don't timeout active connection
	}
	return client
}
