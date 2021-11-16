package infra

import (
	"blockchain-trading/config"
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

const (
	BitFlyerURL       = "https://api.bitflyer.com/v1/"
	realtimeAPIScheme = "wss"
	realtimeAPIHost   = "ws.lightstream.bitflyer.com"
	realtimeAPIPath   = "/json-rpc"
	jsonRPCVersion    = "2.0"
	CryptoWatchURL    = "https://api.cryptowat.ch/markets/"
	CwSdkVersion      = "2.0.0-beta.6"
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

func NewAPIClient(target Target) APIClient {
	ac := &apiClient{
		httpClient: &http.Client{},
		target:     target,
	}
	return ac
}

func (api *apiClient) DoRequest(method, urlPath string, query, header map[string]string, data []byte) ([]byte, error) {
	baseURL, err := url.Parse(api.target.BaseURL)
	if err != nil {
		return nil, errors.Wrap(err, "Parse baseURL")
	}
	apiURL, err := url.Parse(urlPath)
	if err != nil {
		return nil, errors.Wrap(err, "Parse urlPath")
	}
	endpoint := baseURL.ResolveReference(apiURL).String()
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, errors.Wrap(err, "New request")
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
		return nil, errors.Wrap(err, "HTTP client request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Read body")
	}
	return body, nil
}

func getBitFlyerPrivateHeader(method, urlPath string, body []byte) map[string]string {
	u, err := url.Parse(BitFlyerURL + urlPath)
	if err != nil {
		return nil
	}
	endpoint := u.RequestURI()
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	message := timestamp + method + endpoint + string(body)

	mac := hmac.New(sha256.New, []byte(config.Env.BfSecret))
	mac.Write([]byte(message))
	sign := hex.EncodeToString(mac.Sum(nil))
	return map[string]string{
		"ACCESS-KEY":       config.Env.BfKey,
		"ACCESS-TIMESTAMP": timestamp,
		"ACCESS-SIGN":      sign,
	}
}
