package daemon

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	controlv1 "github.com/codetent/crypta/gen/control/v1"
	"github.com/codetent/crypta/gen/control/v1/controlv1connect"
	secretv1 "github.com/codetent/crypta/gen/secret/v1"
	"github.com/codetent/crypta/gen/secret/v1/secretv1connect"
)

var (
	ErrSecretNotExists = errors.New("secret does not exist")
)

type daemonClient struct {
	secretClient  secretv1connect.SecretServiceClient
	controlClient controlv1connect.ControlServiceClient
}

func NewDaemonClient(ip string, port string) *daemonClient {
	endpoint := fmt.Sprintf("http://%s:%s", ip, port)
	return &daemonClient{
		secretClient: secretv1connect.NewSecretServiceClient(
			http.DefaultClient,
			endpoint,
		),
		controlClient: controlv1connect.NewControlServiceClient(
			http.DefaultClient,
			endpoint,
		),
	}
}

func (c *daemonClient) SetSecret(ctx context.Context, name string, value string) error {
	_, err := c.secretClient.SetSecret(
		ctx,
		connect.NewRequest(&secretv1.SetSecretRequest{
			Name:  name,
			Value: value,
		}),
	)
	return err
}

func (c *daemonClient) GetSecret(ctx context.Context, name string) (string, error) {
	res, err := c.secretClient.GetSecret(
		context.Background(),
		connect.NewRequest(&secretv1.GetSecretRequest{
			Name: name,
		}),
	)
	if err != nil {
		return "", err
	}

	if !res.Msg.Exists {
		return "", ErrSecretNotExists
	}

	return res.Msg.Value, nil
}

func (c *daemonClient) Shutdown(ctx context.Context) error {
	_, err := c.controlClient.Shutdown(
		ctx,
		connect.NewRequest(&controlv1.ShutdownRequest{}),
	)
	return err
}
