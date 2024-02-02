package store

type SecretStore interface {
	SetSecret(string, string)
	GetSecret(string) (string, bool)
}

type LocalSecretStore struct {
	secrets map[string]string
}

func NewLocalSecretStore(opts ...SecretStoreOption) SecretStore {
	store := &LocalSecretStore{
		secrets: map[string]string{},
	}

	for _, opt := range opts {
		opt(store)
	}

	return store
}

func (s *LocalSecretStore) SetSecret(name string, value string) {
	s.secrets[name] = value
}

func (s *LocalSecretStore) GetSecret(name string) (string, bool) {
	value, exists := s.secrets[name]
	return value, exists
}
