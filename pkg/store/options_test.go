package store

import (
	"os"
	"path/filepath"
	"testing"

	m_store "github.com/codetent/crypta/mocks/github.com/codetent/crypta/pkg/store"
)

func TestWithEnvPrefix(t *testing.T) {
	tests := []struct {
		name   string
		prefix string
		env    map[string]string
		expect func(store *m_store.MockSecretStore)
	}{
		{
			name:   "Does not pre-populate secret store if no fitting env variables with prefix CRYPTA_SECRET_ are set",
			prefix: "CRYPTA_SECRET_",
			env: map[string]string{
				"TEST":    "TEST",
				"CRYPTA_": "TEST",
			},
			expect: func(store *m_store.MockSecretStore) {},
		},
		{
			name:   "Does not pre-populate secret store if given key or value are empty",
			prefix: "CRYPTA_SECRET_",
			env: map[string]string{
				"CRYPTA_SECRET_":    "TEST",
				"CRYPTA_SECRET_XYZ": "",
			},
			expect: func(store *m_store.MockSecretStore) {},
		},
		{
			name:   "Pre-populates secret store with content of env variables with prefix CRYPTA_SECRET_",
			prefix: "CRYPTA_SECRET_",
			env: map[string]string{
				"TEST":                  "TEST",
				"CRYPTA_SECRET_TEST123": "ABCD",
				"CRYPTA_SECRET_XYZ":     "AFGH",
			},
			expect: func(store *m_store.MockSecretStore) {
				store.EXPECT().SetSecret("TEST123", "ABCD").Once()
				store.EXPECT().SetSecret("XYZ", "AFGH").Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := m_store.NewMockSecretStore(t)

			for k, v := range tt.env {
				t.Setenv(k, v)
			}

			tt.expect(store)

			fn := WithEnvPrefix(tt.prefix)
			fn(store)
		})
	}
}

func TestWithLocalPath(t *testing.T) {
	tests := []struct {
		name   string
		files  map[string]string
		expect func(store *m_store.MockSecretStore)
	}{
		{
			name:   "Does not populate secret store if there are no local files",
			files:  map[string]string{},
			expect: func(store *m_store.MockSecretStore) {},
		},
		{
			name: "Pre-populates secret store with content of local files if they exist",
			files: map[string]string{
				"foo":    "bar",
				"mickey": "mouse",
			},
			expect: func(store *m_store.MockSecretStore) {
				store.EXPECT().SetSecret("foo", "bar").Once()
				store.EXPECT().SetSecret("mickey", "mouse").Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := m_store.NewMockSecretStore(t)

			dir, err := os.MkdirTemp("", "crypta-test-path")
			if err != nil {
				panic(err)
			}
			defer os.RemoveAll(dir)

			for name, val := range tt.files {
				f, err := os.Create(filepath.Join(dir, name))
				if err != nil {
					panic(err)
				}
				defer f.Close()

				err = os.WriteFile(filepath.Join(dir, name), []byte(val), 0666)
				if err != nil {
					panic(err)
				}
			}

			tt.expect(store)

			fn := WithLocalPath(dir)
			fn(store)
		})
	}
}
