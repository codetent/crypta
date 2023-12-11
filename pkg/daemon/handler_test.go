package daemon

import (
	"context"
	"reflect"
	"testing"

	connect "connectrpc.com/connect"
	secretv1 "github.com/codetent/crypta/gen/secret/v1"
	m_daemon "github.com/codetent/crypta/mocks/github.com/codetent/crypta/pkg/daemon"
)

func Test_secretServiceServer_SetSecret(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[secretv1.SetSecretRequest]
	}
	tests := []struct {
		name      string
		args      args
		want      *connect.Response[secretv1.SetSecretResponse]
		wantCalls func(m *m_daemon.MockSecretStore)
		wantErr   bool
	}{
		{
			name: "Set value by request",
			args: args{
				ctx: context.Background(),
				req: &connect.Request[secretv1.SetSecretRequest]{
					Msg: &secretv1.SetSecretRequest{
						Name:  "foo",
						Value: "bar",
					},
				},
			},
			want: &connect.Response[secretv1.SetSecretResponse]{
				Msg: &secretv1.SetSecretResponse{},
			},
			wantCalls: func(m *m_daemon.MockSecretStore) {
				m.EXPECT().SetSecret("foo", "bar").Return()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := m_daemon.NewMockSecretStore(t)
			s := &secretServiceServer{
				store: m,
			}

			tt.wantCalls(m)

			got, err := s.SetSecret(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("secretServiceServer.SetSecret() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("secretServiceServer.SetSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_secretServiceServer_GetSecret(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[secretv1.GetSecretRequest]
	}
	tests := []struct {
		name      string
		args      args
		want      *connect.Response[secretv1.GetSecretResponse]
		wantCalls func(m *m_daemon.MockSecretStore)
		wantErr   bool
	}{
		{
			name: "Get existing value by request",
			args: args{
				ctx: context.Background(),
				req: &connect.Request[secretv1.GetSecretRequest]{
					Msg: &secretv1.GetSecretRequest{
						Name: "foo",
					},
				},
			},
			want: &connect.Response[secretv1.GetSecretResponse]{
				Msg: &secretv1.GetSecretResponse{
					Value:  "bar",
					Exists: true,
				},
			},
			wantCalls: func(m *m_daemon.MockSecretStore) {
				m.EXPECT().GetSecret("foo").Return("bar", true)
			},
			wantErr: false,
		},
		{
			name: "Get missing value by request",
			args: args{
				ctx: context.Background(),
				req: &connect.Request[secretv1.GetSecretRequest]{
					Msg: &secretv1.GetSecretRequest{
						Name: "foo",
					},
				},
			},
			want: &connect.Response[secretv1.GetSecretResponse]{
				Msg: &secretv1.GetSecretResponse{
					Value:  "",
					Exists: false,
				},
			},
			wantCalls: func(m *m_daemon.MockSecretStore) {
				m.EXPECT().GetSecret("foo").Return("", false)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := m_daemon.NewMockSecretStore(t)
			s := &secretServiceServer{
				store: m,
			}

			tt.wantCalls(m)

			got, err := s.GetSecret(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("secretServiceServer.GetSecret() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("secretServiceServer.GetSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}
