package daemon

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	secretv1 "github.com/codetent/crypta/gen/secret/v1"
	"github.com/codetent/crypta/gen/secret/v1/secretv1connect"
)

var (
	ErrSecretNotExists = errors.New("secret does not exist")
)

type daemonClient struct {
	client secretv1connect.SecretServiceClient
}

func NewDaemonClient(ip string, port string) *daemonClient {
	return &daemonClient{
		client: secretv1connect.NewSecretServiceClient(
			http.DefaultClient,
			fmt.Sprintf("http://%s:%s", ip, port),
		),
	}
}

func (c *daemonClient) SetSecret(ctx context.Context, name string, value string) error {
	_, err := c.client.SetSecret(
		ctx,
		connect.NewRequest(&secretv1.SetSecretRequest{
			Name:  name,
			Value: value,
		}),
	)
	return err
}

func (c *daemonClient) GetSecret(ctx context.Context, name string) (string, error) {
	res, err := c.client.GetSecret(
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
