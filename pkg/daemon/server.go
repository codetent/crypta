package daemon

import (
	"fmt"
	"log"
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
	mux.Handle(NewDaemonServiceHandler())

	fo, _ := os.OpenFile("/workspaces/crypta/daemon.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer fo.Close()

	log.SetOutput(fo)

	log.Println("Starting the daemon..")

	return http.ListenAndServe(s.address, h2c.NewHandler(mux, &http2.Server{}))
}
