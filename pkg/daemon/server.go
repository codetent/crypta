package daemon

import (
	"fmt"
	"net/http"

	"github.com/codetent/crypta/pkg/store"
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
	store := store.NewLocalSecretStore(
		store.WithEnvPrefix("CRYPTA_SECRET_"),
		store.WithLocalPath("/var/run/secrets/crypta"),
	)

	mux := http.NewServeMux()
	mux.Handle(NewSecretServiceHandler(store))

	return http.ListenAndServe(s.address, h2c.NewHandler(mux, &http2.Server{}))
}
