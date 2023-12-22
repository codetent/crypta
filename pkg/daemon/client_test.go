package daemon

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	daemonv1 "github.com/codetent/crypta/gen/daemon/v1"
	m_daemonv1connect "github.com/codetent/crypta/mocks/github.com/codetent/crypta/gen/daemon/v1/daemonv1connect"
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
		wantCalls func(m *m_daemonv1connect.MockDaemonServiceClient)
		wantErr   bool
	}{
		{
			name: "Send SetSecretRequest",
			args: args{
				ctx:   context.Background(),
				name:  "foo",
				value: "bar",
			},
			wantCalls: func(m *m_daemonv1connect.MockDaemonServiceClient) {
				m.EXPECT().SetSecret(
					context.Background(),
					connect.NewRequest(&daemonv1.SetSecretRequest{
						Name:  "foo",
						Value: "bar",
					}),
				).Return(
					&connect.Response[daemonv1.SetSecretResponse]{},
					nil,
				)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := m_daemonv1connect.NewMockDaemonServiceClient(t)
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
		wantCalls func(m *m_daemonv1connect.MockDaemonServiceClient)
		wantErr   bool
	}{
		{
			name: "Send GetSecretRequest and return existing value",
			args: args{
				ctx:  context.Background(),
				name: "foo",
			},
			want: "bar",
			wantCalls: func(m *m_daemonv1connect.MockDaemonServiceClient) {
				m.EXPECT().GetSecret(
					context.Background(),
					connect.NewRequest(&daemonv1.GetSecretRequest{
						Name: "foo",
					}),
				).Return(
					&connect.Response[daemonv1.GetSecretResponse]{
						Msg: &daemonv1.GetSecretResponse{
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
			wantCalls: func(m *m_daemonv1connect.MockDaemonServiceClient) {
				m.EXPECT().GetSecret(
					context.Background(),
					connect.NewRequest(&daemonv1.GetSecretRequest{
						Name: "foo",
					}),
				).Return(
					&connect.Response[daemonv1.GetSecretResponse]{
						Msg: &daemonv1.GetSecretResponse{
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
			m := m_daemonv1connect.NewMockDaemonServiceClient(t)
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
