package base

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ReadinessResponse is the response for the readiness probe
type ReadinessResponse struct {
	Ready   bool   `json:"ready"`
	Message string `json:"message"`
}

func initReadinessProbe() chan Readiness {
	readinessChan := make(chan Readiness)

	http.HandleFunc("/readiness", readinessHandler(readinessChan))
	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%d", config.ReadinessPort), nil)
		if err != nil {
			panic(err)
		}
	}()
	log.Infof("Healthcheck started on port %d", config.ReadinessPort)

	return readinessChan
}

func readinessHandler(readinessChan chan Readiness) func(http.ResponseWriter, *http.Request) {
	response := &ReadinessResponse{
		Ready:   false,
		Message: "Initializing",
	}

	go func() {
		for readiness := range readinessChan {
			response = &ReadinessResponse{
				Ready:   readiness.IsReady(),
				Message: readiness.GetMessage(),
			}
		}
	}()

	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Only GET is allowed", http.StatusMethodNotAllowed)
			return
		}
		status := http.StatusOK
		// TODO: REINSTATE FOR NOW ALWAYS READY
		// if !response.Ready {
		// 	status = http.StatusServiceUnavailable
		// }
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(response)
	}

	return handler
}
