package daemon

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	daemonv1 "github.com/codetent/crypta/gen/daemon/v1"
	"github.com/codetent/crypta/gen/daemon/v1/daemonv1connect"
)

var (
	ErrSecretNotExists = errors.New("secret does not exist")
)

type daemonClient struct {
	client daemonv1connect.DaemonServiceClient
}

func NewDaemonClient(ip string, port string) *daemonClient {
	endpoint := fmt.Sprintf("http://%s:%s", ip, port)
	return &daemonClient{
		client: daemonv1connect.NewDaemonServiceClient(
			http.DefaultClient,
			endpoint,
		),
	}
}

func (c *daemonClient) SetSecret(ctx context.Context, name string, value string) error {
	_, err := c.client.SetSecret(
		ctx,
		connect.NewRequest(&daemonv1.SetSecretRequest{
			Name:  name,
			Value: value,
		}),
	)
	return err
}

func (c *daemonClient) GetSecret(ctx context.Context, name string) (string, error) {
	res, err := c.client.GetSecret(
		context.Background(),
		connect.NewRequest(&daemonv1.GetSecretRequest{
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

func (c *daemonClient) GetProcessId(ctx context.Context) (int32, error) {
	res, err := c.client.GetProcessId(
		context.Background(),
		connect.NewRequest(&daemonv1.GetProcessIdRequest{}),
	)
	if err != nil {
		return 0, err
	}

	return res.Msg.Pid, nil
}
