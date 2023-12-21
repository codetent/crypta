package daemon

import (
	"context"
	"net/http"
	"time"

	"connectrpc.com/connect"
	controlv1 "github.com/codetent/crypta/gen/control/v1"
	"github.com/codetent/crypta/gen/control/v1/controlv1connect"
)

type shutdowner func()

type controlServiceServer struct {
	controlv1connect.UnimplementedControlServiceHandler
	shutdown shutdowner
}

func NewControlServiceHandler(shut shutdowner) (string, http.Handler) {
	return controlv1connect.NewControlServiceHandler(&controlServiceServer{shutdown: shut})
}

func (s *controlServiceServer) Shutdown(
	ctx context.Context,
	req *connect.Request[controlv1.ShutdownRequest],
) (*connect.Response[controlv1.ShutdownResponse], error) {
	// wait so the response can be returned, and the server is shutdown afterwards
	go time.AfterFunc(1*time.Second, s.shutdown)

	return connect.NewResponse(&controlv1.ShutdownResponse{}), nil
}
