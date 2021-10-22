package api

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	BitFlyerURL    = "https://api.bitflyer.com/v1/"
	CryptoWatchURL = "https://api.cryptowat.ch/markets/"
	CwSdkVersion   = "2.0.0-beta.6"
)

type APIClient interface {
	DoRequest(method, urlPath string, query, header map[string]string, data []byte) (body []byte, err error)
}

type apiClient struct {
	httpClient *http.Client
	target     Target
}

type Target struct {
	BaseURL string
	Header  map[string]string
}

// NewAPIClient
func NewAPIClient(target Target) APIClient {
	ac := &apiClient{
		httpClient: &http.Client{},
		target:     target,
	}
	return ac
}

func (api *apiClient) DoRequest(method, urlPath string, query, header map[string]string, data []byte) (body []byte, err error) {
	baseURL, err := url.Parse(api.target.BaseURL)
	if err != nil {
		return
	}
	apiURL, err := url.Parse(urlPath)
	if err != nil {
		return
	}
	endpoint := baseURL.ResolveReference(apiURL).String()
	log.Printf("action=doRequest endpoint=%s", endpoint)
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(data))
	if err != nil {
		return
	}
	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	reqHeader := api.target.Header
	for key, value := range header {
		reqHeader[key] = value
	}
	for key, value := range reqHeader {
		req.Header.Add(key, value)
	}
	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
