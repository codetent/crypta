package daemon

import (
	"os"
	"strings"
)

type SecretStore interface {
	SetSecret(string, string)
	GetSecret(string) (string, bool)
}

func PopulateStore(store SecretStore) {
	const prefix string = "CRYPTA_SECRET_"

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)

		if strings.HasPrefix(pair[0], prefix) {
			key := pair[0][len(prefix):]

			if len(key) == 0 || len(pair[1]) == 0 {
				continue
			}

			store.SetSecret(key, pair[1])
		}
	}
}

type LocalSecretStore struct {
	secrets map[string]string
}

func NewLocalSecretStore() SecretStore {
	return &LocalSecretStore{
		secrets: map[string]string{},
	}
}

func (s *LocalSecretStore) SetSecret(name string, value string) {
	s.secrets[name] = value
}

func (s *LocalSecretStore) GetSecret(name string) (string, bool) {
	value, exists := s.secrets[name]
	return value, exists
}
