package infrastructure

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/pkg/errors"
)

type APIClient struct {
	HTTPClient *http.Client
	Target     Target
}

type Target struct {
	BaseURL *url.URL
	Header  map[string]string
}

func (api *APIClient) DoRequest(method, urlPath string, query, header map[string]string, data []byte) ([]byte, error) {
	copyBaseURL := *api.Target.BaseURL
	copyBaseURL.Path = path.Join(urlPath)
	endpoint := copyBaseURL.String()
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, errors.Wrap(err, "New request")
	}
	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	reqHeader := api.Target.Header
	for key, value := range header {
		reqHeader[key] = value
	}
	for key, value := range reqHeader {
		req.Header.Add(key, value)
	}
	resp, err := api.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "HTTP client request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Read body")
	}
	return body, nil
}
