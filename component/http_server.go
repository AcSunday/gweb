package component

import (
	"fmt"
	"go_project_demo/library/clean"
	"go_project_demo/library/http"
)

var HttpServer *http.Server

func SetupHttpServer(port int) (err error) {
	HttpServer = http.NewServer(&http.ServerConf{
		Addr: fmt.Sprintf(":%d", port),
	})

	clean.Push(HttpServer)

	return
}
