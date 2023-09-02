package service_hit

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/papannn/coda-assignment/api-gateway/config"
	"github.com/papannn/coda-assignment/discovery-service/api"
	"github.com/papannn/coda-assignment/lib/parser"
	"io"
	"net/http"
)

type IServiceHit interface {
	Post(namespace string, requestURI string, body io.ReadCloser) (*http.Response, error)
}

type Impl struct {
	Config config.Config
}

func (impl *Impl) Post(namespace string, requestURI string, body io.ReadCloser) (*http.Response, error) {
	lookupResp, err := impl.lookup(namespace)
	if err != nil {
		return nil, err
	}
	URL := fmt.Sprintf("http://%s:%s%s", lookupResp.IP, lookupResp.Port, requestURI)
	return http.Post(URL, "application/json", body)
}

func (impl *Impl) lookup(namespace string) (*api.LookupResponse, error) {
	payload := map[string]string{
		"namespace": namespace,
	}

	payloadByte, _ := json.Marshal(payload)
	discoveryServiceURL := fmt.Sprintf("%s/api/lookup", impl.Config.DiscoveryServiceBaseUrl)

	client := http.Client{}
	req, err := http.NewRequest(http.MethodPost, discoveryServiceURL, bytes.NewBuffer(payloadByte))
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("discovery service is not available right now")
	}

	respLookup := api.LookupResponse{}
	err = parser.ParseJSONBody(resp.Body, &respLookup)
	if err != nil {
		return nil, err
	}

	return &respLookup, nil
}
