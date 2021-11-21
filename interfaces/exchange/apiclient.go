package exchange

type APIClient interface {
	DoRequest(method, urlPath string, query, header map[string]string, data []byte) (body []byte, err error)
}
