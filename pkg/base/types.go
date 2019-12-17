package base

// Config defines the configuration needed for Envoy XDS
type Config struct {
	Port           uint32
	Path           string
	CertPath       string
	Domain         string
	DNSRefreshRate int64
	RespectDNSTTL  bool
	NumShards      int
	UseProxyProto  bool

	ReadinessChan chan Readiness
	ReadinessPort uint32
}

// Readiness is the state used for readiness channel.
type Readiness interface {
	IsReady() bool
	GetMessage() string
}
