package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
)

const (
	defaultContentType = "application/json"
)

type Client struct {
	*http.Client
}

type ClientConf struct {
	DialTimeoutSecond     int // 连接超时
	DialKeepAliveSecond   int // 开启长连接
	MaxIdleConns          int // 最大空闲连接数
	MaxIdleConnsPerHost   int // HOST最大空闲连接数
	IdleConnTimeoutSecond int // 空闲连接超时
}

func NewClient(cf *ClientConf) (cli *Client) {
	cli = &Client{
		Client: &http.Client{
			Transport: &http.Transport{
				DisableKeepAlives:   cf.DialKeepAliveSecond < 0,
				MaxIdleConns:        cf.MaxIdleConns,
				MaxIdleConnsPerHost: cf.MaxIdleConnsPerHost,
				IdleConnTimeout:     time.Duration(cf.IdleConnTimeoutSecond) * time.Second,
				Proxy:               http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   time.Duration(cf.DialTimeoutSecond) * time.Second,
					KeepAlive: time.Duration(cf.DialKeepAliveSecond) * time.Second,
				}).DialContext,
			},
		},
	}
	return
}

func (cli *Client) Get(baseURL string, query url.Values, respst interface{}) (err error) {
	return cli.Req("GET", baseURL, query, nil, respst)
}

func (cli *Client) Post(baseURL string, reqdata, respst interface{}) (err error) {
	return cli.Req("POST", baseURL, nil, reqdata, respst)
}

func (cli *Client) Req(method string, baseURL string, query url.Values, reqdata, respst interface{}) (err error) {
	u, err := makeQueryUrl(baseURL, query)
	if err != nil {
		return
	}

	body, err := makeRequestBody(reqdata)
	if err != nil {
		return
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		err = fmt.Errorf("new request: %w", err)
		return
	}

	req.Header.Add("Content-Type", defaultContentType)

	resp, err := cli.Client.Do(req)
	if err != nil || resp == nil {
		err = fmt.Errorf("response is nil or do request: %w", err)
		return
	}

	defer resp.Body.Close()

	if err = scanresp(resp, respst); err != nil {
		err = fmt.Errorf("scan response: %w", err)
		return
	}

	return
}

func (cli *Client) Close() (err error) {
	cli.Client.CloseIdleConnections()
	return
}

func makeQueryUrl(baseURL string, query url.Values) (u *url.URL, err error) {
	u, err = url.Parse(baseURL)
	if err != nil {
		err = fmt.Errorf("url parse: %w", err)
		return
	}
	if len(query) > 0 {
		q := u.Query()
		for k, v := range query {
			q.Del(k)
			for _, vl := range v {
				q.Add(k, vl)
			}
		}
		u.RawQuery = q.Encode()
	}
	return
}

func makeRequestBody(reqdata interface{}) (body io.Reader, err error) {
	if reqdata != nil {
		data, merr := json.Marshal(reqdata)
		if merr != nil {
			err = fmt.Errorf("json marshal: %w", err)
			return
		}
		body = bytes.NewReader(data)
	}
	return
}

func scanresp(resp *http.Response, st interface{}) (err error) {
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("response code %d", resp.StatusCode)
		return
	}

	if st == nil {
		st = &(map[string]interface{}{})
	}

	if err = json.NewDecoder(resp.Body).Decode(st); err != nil {
		err = fmt.Errorf("json decode: %w", err)
		return
	}

	return
}
