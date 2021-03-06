package central

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/bladedancer/vhds/pkg/base"
)

// Watch central for changes
func Watch(ready chan base.Readiness) chan []*Listener {
	listenerUpdateChan := make(chan []*Listener)
	log.Infof("Polling every %d seconds", config.SyncInterval)
	tick := time.NewTicker(time.Duration(config.SyncInterval) * time.Second)
	go func() {
		sync(ready, listenerUpdateChan)
		for {
			select {
			case <-tick.C:
				sync(ready, listenerUpdateChan)
			}
		}
	}()

	return listenerUpdateChan
}

func sync(readyChan chan base.Readiness, listenerChan chan []*Listener) {
	//listeners, err := getListeners()
	listeners, err := fakeGetListeners()
	if err != nil {
		log.Error(err)
		readyChan <- &_SyncReadiness{Ready: false, Message: err.Error()}
		return
	}

	listenerChan <- listeners
	readyChan <- &_SyncReadiness{Ready: true}
}

var fakeID int

func fakeGetListeners() ([]*Listener, error) {
	// Faking it.
	listeners := []*Listener{}

	listeners = append(listeners, &Listener{
		ID:             fmt.Sprintf("id-%d", fakeID),
		Activated:      true,
		Name:           fmt.Sprintf("id-%d", fakeID),
		Protocol:       "http",
		BindAddress:    "0.0.0.0",
		Port:           "80",
		VirtualHosts:   []string{"test", "prod"},
		RuntimeGroupID: fmt.Sprintf("%d", fakeID),
		Metadata:       nil,
		InstanceName:   fmt.Sprintf("inst-%d", fakeID),
		TenantID:       fmt.Sprintf("ten-%d", fakeID),
	})

	fakeID = fakeID + 1

	return listeners, nil
}

func getListeners() ([]*Listener, error) {
	req, err := http.NewRequest("GET", config.syncURL, nil)
	if err != nil {
		return nil, err
	}

	// Configure the headers
	req.Header.Set("X-Axway-Instance-Id", "axway")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(config.ServiceUser, config.ServiceSecret)

	if config.timestamp != "" {
		req.Header.Set(config.timestampHeader, config.timestamp)
	}
	req.Host = config.SyncHost
	log.Debugf("Request: %+v", req)

	resp, err := config.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp != nil && resp.StatusCode != 200 {
		var responseCode = strconv.Itoa(resp.StatusCode)
		message := fmt.Sprintf("Sync API returned non 200 response. Code returned %s", responseCode)
		err = errors.New(message)
		return nil, err
	}

	timestamp := resp.Header.Get(config.timestampHeader)

	if timestamp == "" {
		err = errors.New("no response timestamp header found")
		return nil, err
	}
	config.timestamp = timestamp

	fullBody, err := ioutil.ReadAll(resp.Body)
	defer closeMe(resp.Body)

	listeners, err := parseBody(fullBody)
	if err != nil {
		return nil, err
	}

	return listeners, nil
}

func parseBody(data []byte) ([]*Listener, error) {
	log.Debugf(string(data))
	var listeners []*Listener
	err := json.Unmarshal(data, &listeners)
	return listeners, err
}
