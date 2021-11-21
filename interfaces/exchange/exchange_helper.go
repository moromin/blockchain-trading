package exchange

import (
	"blockchain-trading/config"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"strconv"
	"time"
)

const (
	BitFlyerURL       = "https://api.bitflyer.com/"
	realtimeAPIScheme = "wss"
	realtimeAPIHost   = "ws.lightstream.bitflyer.com"
	realtimeAPIPath   = "/json-rpc"
	jsonRPCVersion    = "2.0"
	CryptoWatchURL    = "https://api.cryptowat.ch/markets/"
	CwSdkVersion      = "2.0.0-beta.6"
)

type Endpoint string

const (
	GetBalance Endpoint = "v1/me/getbalance"
	GetTicker  Endpoint = "v1/ticker"
	SendOrder  Endpoint = "v1/me/sendchildorder"
)

func (e Endpoint) String() string {
	return string(e)
}

func (e Endpoint) Method() string {
	switch e {
	case GetBalance, GetTicker:
		return "GET"
	case SendOrder:
		return "POST"
	default:
		return ""
	}
}

func (e Endpoint) Header(body []byte) map[string]string {
	switch e {
	case GetBalance, SendOrder:
		return getBitFlyerPrivateHeader(e.Method(), e.String(), body)
	default:
		return nil
	}
}

func getBitFlyerPrivateHeader(method, urlPath string, body []byte) map[string]string {
	u, err := url.Parse(BitFlyerURL + urlPath)
	if err != nil {
		panic(err)
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
