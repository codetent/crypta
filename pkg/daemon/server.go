package daemon

import (
	"fmt"
	"net/http"

	"github.com/codetent/crypta/pkg/store"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type daemonServer struct {
	store store.SecretStore
}

func NewDaemonServer(store store.SecretStore) *daemonServer {
	return &daemonServer{
		store: store,
	}
}

func (s *daemonServer) ListenAndServe(ip string, port string) error {
	mux := http.NewServeMux()
	mux.Handle(NewSecretServiceHandler(s.store))

	return http.ListenAndServe(
		fmt.Sprintf("%s:%s", ip, port),
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
