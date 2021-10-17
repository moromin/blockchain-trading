package api

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	bitFlyerURL    = "https://api.bitflyer.com/v1/"
	cryptoWatchURL = "https://api.cryptowat.ch/markets/"
	cwSdkVersion   = "2.0.0-beta.6"
)

type APIClient interface {
	DoRequest(method, urlPath string, query map[string]string, data []byte) (body []byte, err error)
}

type apiClient struct {
	key        string
	secret     string
	httpClient *http.Client
	baseURL    string
}

// NewAPIClient
func NewAPIClient(key, secret, clientType string) APIClient {
	ac := &apiClient{
		key:        key,
		secret:     secret,
		httpClient: &http.Client{},
	}
	if strings.EqualFold(clientType, "bitflyer") {
		ac.baseURL = bitFlyerURL
	} else if strings.EqualFold(clientType, "cryptowatch") {
		ac.baseURL = cryptoWatchURL
	}
	return ac
}

// header returns a different map depending on baseURL.
func (api apiClient) header(method, endpoint string, body []byte) map[string]string {
	if strings.EqualFold(api.baseURL, bitFlyerURL) {
		timestamp := strconv.FormatInt(time.Now().Unix(), 10)
		message := timestamp + method + endpoint + string(body)

		mac := hmac.New(sha256.New, []byte(api.secret))
		mac.Write([]byte(message))
		sign := hex.EncodeToString(mac.Sum(nil))
		return map[string]string{
			"ACCESS-KEY":       api.key,
			"ACCESS-TIMESTAMP": timestamp,
			"ACCESS-SIGN":      sign,
			"Content-Type":     "application/json",
		}
	} else if strings.EqualFold(api.baseURL, cryptoWatchURL) {
		return map[string]string{
			"X-CW-API-Key": api.key,
			"User-Agent":   fmt.Sprintf("cw-sdk-go@%s", cwSdkVersion),
		}
	}
	return map[string]string{}
}

func (api *apiClient) DoRequest(method, urlPath string, query map[string]string, data []byte) (body []byte, err error) {
	baseURL, err := url.Parse(api.baseURL)
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

	for key, value := range api.header(method, req.URL.RequestURI(), data) {
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
