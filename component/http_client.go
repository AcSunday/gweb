package component

import (
	"fmt"
	"go_project_demo/library/clean"
	"go_project_demo/library/http"
)

var HttpClient *http.ClientContainer

type httpClientConfig struct {
	DialTimeoutSecond     int `env:"HTTP_CLIENT_DIAL_TIMEOUT_SECOND"`
	DialKeepAliveSecond   int `env:"HTTP_CLIENT_DIAL_KEEP_ALIVE_SECOND"`
	MaxIdleConns          int `env:"HTTP_CLIENT_MAX_IDLE_CONNS"`
	MaxIdleConnsPerHost   int `env:"HTTP_CLIENT_MAX_IDLE_CONNS_PER_HOST"`
	IdleConnTimeoutSecond int `env:"HTTP_CLIENT_MAX_CONN_TIMEOUT_SECOND"`
}

func SetupHttpClient() (err error) {
	HttpClient, err = http.NewClientContainer(getHttpClientConf)
	if err != nil {
		err = fmt.Errorf("http.NewClientContainer: %w", err)
		return
	}

	clean.Push(HttpClient)

	return
}

func getHttpClientConf() (cf *http.ClientConf, err error) {
	cfg := &httpClientConfig{}

	if err = Conf.Scan(cfg, "env"); err != nil {
		err = fmt.Errorf("Conf.Scan: %w", err)
		return
	}

	cf = &http.ClientConf{
		DialTimeoutSecond:     cfg.DialTimeoutSecond,
		DialKeepAliveSecond:   cfg.DialKeepAliveSecond,
		MaxIdleConns:          cfg.MaxIdleConns,
		MaxIdleConnsPerHost:   cfg.MaxIdleConnsPerHost,
		IdleConnTimeoutSecond: cfg.IdleConnTimeoutSecond,
	}

	return
}
