package xds

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
)

// ShardRouter Router Struct for service
type ShardRouter struct {
}

// MakeShardRouter Return a new Router Server
func MakeShardRouter() *ShardRouter {
	return &ShardRouter{}
}

//Run Start the service
func (t *ShardRouter) Run() {

	http.HandleFunc("/shard",
		func(w http.ResponseWriter, r *http.Request) {
			buf := new(bytes.Buffer)
			buf.ReadFrom(r.Body)
			log.Infof("Received %s", buf.String())
			toks := strings.Split(buf.String(), ":")
			host := toks[0]
			fmt.Fprintf(w, "%s", getShard(host))

		},
	)
	go http.ListenAndServe(":12001", nil)
	log.Info("Simple Service Started")
}

// getShard get the shard hosting the runtime group
func getShard(host string) string {
	log.Debugf("Getting shard for %s", host)
	depKey := extractDeploymentKey(host)
	shard := deploymentManager.GetShardName(depKey)
	s := deploymentManager.shards[shard]
	if s == nil {
		// Edge case for unknown tennat request
		// Should not occur, but will result in 404
		log.Warnf("shard should not be nil  %s", shard)
		return ""
	}
	log.Infof("Host: %s, Shard %s", host, shard)
	return shard
}

//extractDeploymentKey the deployment key is encoded in the url.
func extractDeploymentKey(host string) string {
	// Key is the vhost-instancename which is the machine name in the host.
	return strings.Split(host, ".")[0]
}
