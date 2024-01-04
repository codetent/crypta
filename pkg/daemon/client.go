package daemon

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"

	"connectrpc.com/connect"
	secretv1 "github.com/codetent/crypta/gen/secret/v1"
	"github.com/codetent/crypta/gen/secret/v1/secretv1connect"
	"github.com/hashicorp/go-retryablehttp"
)

var (
	ErrSecretNotExists = errors.New("secret does not exist")
	networkError       *net.OpError
)

type daemonClient struct {
	client secretv1connect.SecretServiceClient
}

func NewDaemonClient(ip string, port string) *daemonClient {
	c := retryablehttp.NewClient()

	// disable logging of the retries (as it is quite verbose)
	silentLogger := log.New(io.Discard, "", 0)
	c.Logger = silentLogger

	// hook the retries in order to inform the user
	c.RequestLogHook = func(l retryablehttp.Logger, r *http.Request, i int) {
		// it is important to log to Stderr here, since Stdout is read by scripts to get values etc
		// retries are counted starting from 0
		fmt.Fprintf(os.Stderr, "Daemon currently not reachable. Retry %d of %d...\n", i+1, c.RetryMax+1)
	}

	return &daemonClient{
		client: secretv1connect.NewSecretServiceClient(
			c.StandardClient(),
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
