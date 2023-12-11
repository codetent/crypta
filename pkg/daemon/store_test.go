package daemon

import (
	"reflect"
	"testing"
)

func TestSecretStore_SetSecret(t *testing.T) {
	type fields struct {
		secrets map[string]string
	}
	type args struct {
		name  string
		value string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantFields fields
	}{
		{
			name: "Set value",
			fields: fields{
				secrets: map[string]string{},
			},
			args: args{
				name:  "foo",
				value: "bar",
			},
			wantFields: fields{
				secrets: map[string]string{
					"foo": "bar",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &LocalSecretStore{
				secrets: tt.fields.secrets,
			}
			s.SetSecret(tt.args.name, tt.args.value)

			if !reflect.DeepEqual(s.secrets, tt.wantFields.secrets) {
				t.Errorf("SecretStore.secrets got = %v, want %v", s.secrets, tt.wantFields.secrets)
			}
		})
	}
}

func TestSecretStore_GetSecret(t *testing.T) {
	type fields struct {
		secrets map[string]string
	}
	type args struct {
		name string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		want       string
		wantExists bool
	}{
		{
			name: "Get existing secret",
			fields: fields{
				secrets: map[string]string{
					"foo": "bar",
				},
			},
			args: args{
				name: "foo",
			},
			want:       "bar",
			wantExists: true,
		},
		{
			name: "Get missing secret",
			fields: fields{
				secrets: map[string]string{},
			},
			args: args{
				name: "foo",
			},
			want:       "",
			wantExists: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &LocalSecretStore{
				secrets: tt.fields.secrets,
			}
			got, gotExists := s.GetSecret(tt.args.name)
			if got != tt.want {
				t.Errorf("SecretStore.GetSecret() got = %v, want %v", got, tt.want)
			}
			if gotExists != tt.wantExists {
				t.Errorf("SecretStore.GetSecret() gotExists = %v, want %v", gotExists, tt.wantExists)
			}
		})
	}
}

func TestNewSecretStore(t *testing.T) {
	tests := []struct {
		name string
		want SecretStore
	}{
		{
			name: "Create empty secret store",
			want: &LocalSecretStore{
				secrets: map[string]string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLocalSecretStore(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSecretStore() = %v, want %v", got, tt.want)
			}
		})
	}
}
