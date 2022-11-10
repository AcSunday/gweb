package http

import "net/url"

func (ct *ClientContainer) Req(method, baseURL string, query url.Values, reqdata, respst interface{}) (err error) {
	client := ct.MustGetClient()
	defer ct.PutClient(client)

	return client.Req(method, baseURL, query, reqdata, respst)
}
