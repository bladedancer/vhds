package accesslog

import als "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v2"

// Handler implements the Stream Access Logs endpoint.
type Handler struct{}

// StreamAccessLogs the sink for access logs.
func (svc *Handler) StreamAccessLogs(stream als.AccessLogService_StreamAccessLogsServer) error {
	for {
		msg, err := stream.Recv()
		if err != nil {
			panic(err)
		}
		switch entries := msg.LogEntries.(type) {
		case *als.StreamAccessLogsMessage_HttpLogs:
			for _, entry := range entries.HttpLogs.LogEntry {
				log.Infof("Access Log: %+v", entry)
			}
		}
	}
}
