package daemon

import (
	"context"
	"testing"
	"time"

	"connectrpc.com/connect"
	secretv1 "github.com/codetent/crypta/gen/secret/v1"
	m_secretv1connect "github.com/codetent/crypta/mocks/github.com/codetent/crypta/gen/secret/v1/secretv1connect"
)

func Test_daemonClient_SetSecret(t *testing.T) {
	type args struct {
		ctx   context.Context
		name  string
		value string
	}
	tests := []struct {
		name      string
		args      args
		wantCalls func(m *m_secretv1connect.MockSecretServiceClient)
		wantErr   bool
	}{
		{
			name: "Send SetSecretRequest",
			args: args{
				ctx:   context.Background(),
				name:  "foo",
				value: "bar",
			},
			wantCalls: func(m *m_secretv1connect.MockSecretServiceClient) {
				m.EXPECT().SetSecret(
					context.Background(),
					connect.NewRequest(&secretv1.SetSecretRequest{
						Name:  "foo",
						Value: "bar",
					}),
				).Return(
					&connect.Response[secretv1.SetSecretResponse]{},
					nil,
				)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := m_secretv1connect.NewMockSecretServiceClient(t)
			c := &daemonClient{
				client: m,
			}

			tt.wantCalls(m)

			if err := c.SetSecret(tt.args.ctx, tt.args.name, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("daemonClient.SetSecret() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_daemonClient_GetSecret(t *testing.T) {
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name      string
		args      args
		want      string
		wantCalls func(m *m_secretv1connect.MockSecretServiceClient)
		wantErr   bool
	}{
		{
			name: "Send GetSecretRequest and return existing value",
			args: args{
				ctx:  context.Background(),
				name: "foo",
			},
			want: "bar",
			wantCalls: func(m *m_secretv1connect.MockSecretServiceClient) {
				m.EXPECT().GetSecret(
					context.Background(),
					connect.NewRequest(&secretv1.GetSecretRequest{
						Name: "foo",
					}),
				).Return(
					&connect.Response[secretv1.GetSecretResponse]{
						Msg: &secretv1.GetSecretResponse{
							Value:  "bar",
							Exists: true,
						},
					},
					nil,
				)
			},
			wantErr: false,
		},
		{
			name: "Send GetSecretRequest and fail when name not found",
			args: args{
				ctx:  context.Background(),
				name: "foo",
			},
			want: "",
			wantCalls: func(m *m_secretv1connect.MockSecretServiceClient) {
				m.EXPECT().GetSecret(
					context.Background(),
					connect.NewRequest(&secretv1.GetSecretRequest{
						Name: "foo",
					}),
				).Return(
					&connect.Response[secretv1.GetSecretResponse]{
						Msg: &secretv1.GetSecretResponse{
							Value:  "",
							Exists: false,
						},
					},
					nil,
				)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := m_secretv1connect.NewMockSecretServiceClient(t)
			c := &daemonClient{
				client: m,
			}

			tt.wantCalls(m)

			got, err := c.GetSecret(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("daemonClient.GetSecret() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("daemonClient.GetSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateRetryTimeout(t *testing.T) {
	tests := []struct {
		name       string
		maxTimeout string
		want       time.Duration
	}{
		{
			name:       "Calculates retry timeout based on default maximum timeout because the env var is not set",
			maxTimeout: "",
			want:       1 * time.Second,
		},
		{
			name:       "Calculates retry timeout based on set maximum timeout because the env var is set",
			maxTimeout: "5",
			want:       1 * time.Second,
		},
		{
			name:       "Allows to calculate retry timeout based on env var value, even if float value is given",
			maxTimeout: "0.1",
			want:       20 * time.Millisecond,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("CRYPTA_TIMEOUT", tt.maxTimeout)

			if got := calculateRetryTimeout(); got != tt.want {
				t.Errorf("calculateRetryTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}
