package daemon

type SecretStore struct {
	secrets map[string]string
}

func NewSecretStore() *SecretStore {
	return &SecretStore{
		secrets: map[string]string{},
	}
}

func (s *SecretStore) SetSecret(name string, value string) {
	s.secrets[name] = value
}

func (s *SecretStore) GetSecret(name string) (string, bool) {
	value, exists := s.secrets[name]
	return value, exists
}
