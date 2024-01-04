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
		name, value, _ := strings.Cut(e, "=")

		// extract "<KEY>" from "<PREFIX><KEY>"
		key, found := strings.CutPrefix(name, prefix)

		if found {
			if len(key) == 0 || len(value) == 0 {
				continue
			}

			store.SetSecret(key, value)
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
