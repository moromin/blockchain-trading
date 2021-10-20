package api

import (
	"blockchain-trading/config"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"strconv"
	"time"
)

func GetBitFlyerPrivateHeader(method, urlPath string, body []byte) map[string]string {
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
