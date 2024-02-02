package daemon

import (
	"context"
	"net/http"

	connect "connectrpc.com/connect"
	secretv1 "github.com/codetent/crypta/gen/secret/v1"
	"github.com/codetent/crypta/gen/secret/v1/secretv1connect"
	"github.com/codetent/crypta/pkg/store"
)

type secretServiceServer struct {
	secretv1connect.UnimplementedSecretServiceHandler
	store store.SecretStore
}

func NewSecretServiceHandler(store store.SecretStore) (string, http.Handler) {
	return secretv1connect.NewSecretServiceHandler(&secretServiceServer{
		store: store,
	})
}

func (s *secretServiceServer) SetSecret(
	ctx context.Context,
	req *connect.Request[secretv1.SetSecretRequest],
) (*connect.Response[secretv1.SetSecretResponse], error) {
	name := req.Msg.GetName()
	value := req.Msg.GetValue()

	s.store.SetSecret(name, value)

	return connect.NewResponse(&secretv1.SetSecretResponse{}), nil
}

func (s *secretServiceServer) GetSecret(
	ctx context.Context,
	req *connect.Request[secretv1.GetSecretRequest],
) (*connect.Response[secretv1.GetSecretResponse], error) {
	name := req.Msg.GetName()

	value, exists := s.store.GetSecret(name)

	return connect.NewResponse(&secretv1.GetSecretResponse{
		Value:  value,
		Exists: exists,
	}), nil
}
