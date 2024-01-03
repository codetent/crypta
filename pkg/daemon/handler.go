package daemon

import (
	"context"
	"log"
	"net/http"
	"os"

	connect "connectrpc.com/connect"
	daemonv1 "github.com/codetent/crypta/gen/daemon/v1"
	"github.com/codetent/crypta/gen/daemon/v1/daemonv1connect"
)

type daemonServiceServer struct {
	daemonv1connect.UnimplementedDaemonServiceHandler
	store SecretStore
}

func NewDaemonServiceHandler() (string, http.Handler) {
	return daemonv1connect.NewDaemonServiceHandler(&daemonServiceServer{
		store: NewLocalSecretStore(),
	})
}

func (s *daemonServiceServer) SetSecret(
	ctx context.Context,
	req *connect.Request[daemonv1.SetSecretRequest],
) (*connect.Response[daemonv1.SetSecretResponse], error) {
	name := req.Msg.GetName()
	value := req.Msg.GetValue()

	s.store.SetSecret(name, value)

	return connect.NewResponse(&daemonv1.SetSecretResponse{}), nil
}

func (s *daemonServiceServer) GetSecret(
	ctx context.Context,
	req *connect.Request[daemonv1.GetSecretRequest],
) (*connect.Response[daemonv1.GetSecretResponse], error) {
	name := req.Msg.GetName()

	value, exists := s.store.GetSecret(name)

	return connect.NewResponse(&daemonv1.GetSecretResponse{
		Value:  value,
		Exists: exists,
	}), nil
}

func (s *daemonServiceServer) GetProcessId(
	ctx context.Context,
	req *connect.Request[daemonv1.GetProcessIdRequest],
) (*connect.Response[daemonv1.GetProcessIdResponse], error) {
	log.Println("GetProcessId called - exiting")

	os.Exit(0)

	return connect.NewResponse(&daemonv1.GetProcessIdResponse{
		Pid: int32(os.Getpid()),
	}), nil
}
