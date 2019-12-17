package central

import "net/http"

// Central config
type Central struct {
	SyncTLS      bool
	SyncHost     string
	SyncPort     uint32
	SyncTimeout  uint32
	SyncInterval uint32

	ServiceUser   string
	ServiceSecret string

	timestampHeader string
	syncURL         string
	client          *http.Client
	timestamp       string
}

type _SyncReadiness struct {
	Ready   bool
	Message string
}

// IsReady is sync ready.
func (s *_SyncReadiness) IsReady() bool {
	return s.Ready
}

// GetMessage get the readiness message.
func (s *_SyncReadiness) GetMessage() string {
	return s.Message
}

// Listener returned from Sync API
type Listener struct {
	ID             string            `json:"id"`
	Activated      bool              `json:"activated"`
	Name           string            `json:"name"`
	Protocol       string            `json:"protocol"`
	BindAddress    string            `json:"bindAddress"`
	Port           string            `json:"port"`
	VirtualHosts   []string          `json:"virtualHosts"`
	RuntimeGroupID string            `json:"runtimeGroupId"`
	Metadata       map[string]string `json:"Metadata"`
	InstanceName   string            `json:"instanceName"`
	TenantID       string            `json:"tenantId"`
}
