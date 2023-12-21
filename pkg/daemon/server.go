package daemon

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type daemonServer struct {
	address string
}

func NewDaemonServer(ip string, port string) *daemonServer {
	return &daemonServer{
		address: fmt.Sprintf("%s:%s", ip, port),
	}
}

func (s *daemonServer) ListenAndServe() error {
	mux := http.NewServeMux()
	mux.Handle(NewSecretServiceHandler())
	mux.Handle(NewControlServiceHandler(func() {
		// TODO: Improve shutdown of daemon. The http package provides a context to shut it down.
		os.Exit(0)
	}))

	return http.ListenAndServe(s.address, h2c.NewHandler(mux, &http2.Server{}))
}
