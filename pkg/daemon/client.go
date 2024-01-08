package daemon

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"connectrpc.com/connect"
	secretv1 "github.com/codetent/crypta/gen/secret/v1"
	"github.com/codetent/crypta/gen/secret/v1/secretv1connect"
	"github.com/hashicorp/go-retryablehttp"
)

var (
	ErrSecretNotExists = errors.New("secret does not exist")
)

const (
	defaultConnectionTimeout = 5 * time.Second
	retries                  = 5
)

type daemonClient struct {
	client secretv1connect.SecretServiceClient
}

func calculateRetryTimeout() time.Duration {
	maxTimeoutStr := os.Getenv("CRYPTA_TIMEOUT")
	maxTimeout, err := strconv.ParseFloat(maxTimeoutStr, 32)

	timeout := time.Duration(maxTimeout*float64(time.Second/time.Millisecond)) * time.Millisecond

	if err != nil {
		log.Println("Using default maximum connection timeout:", defaultConnectionTimeout)
		timeout = defaultConnectionTimeout
	} else {
		log.Println("Using set maximum connection timeout:", timeout)
	}

	return timeout / retries
}

func NewDaemonClient(ip string, port string) *daemonClient {
	c := retryablehttp.NewClient()

	// disable logging of the retries (as it is quite verbose)
	silentLogger := log.New(io.Discard, "", 0)
	c.Logger = silentLogger

	retryTimeout := calculateRetryTimeout()

	// in order to calculate a maximum timeout that the connection attempt takes, a constant wait with a constant retry
	// timeout is used. The retry timeout is calculated based on the maximum timeout & the number of retries.
	c.RetryMax = retries - 1 // retries start counting with 0
	c.RetryWaitMin = retryTimeout
	c.RetryWaitMax = retryTimeout
	c.Backoff = func(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
		return max
	}

	// hook the retries in order to inform the user
	c.RequestLogHook = func(l retryablehttp.Logger, r *http.Request, i int) {
		if i > 0 {
			log.Printf("Daemon currently not reachable. Retry %d of %d...\n", i, c.RetryMax)
		}
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
		ctx,
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
