package store

import (
	"os"
	"path/filepath"
	"strings"
)

type SecretStoreOption func(store SecretStore)

func WithEnvPrefix(prefix string) SecretStoreOption {
	return func(store SecretStore) {
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
}

func WithLocalPath(dir string) SecretStoreOption {
	return func(store SecretStore) {
		files, err := os.ReadDir(dir)
		if err != nil {
			return
		}

		for _, f := range files {
			if !f.IsDir() {
				content, err := os.ReadFile(filepath.Join(dir, f.Name()))
				if err != nil {
					return
				}

				store.SetSecret(f.Name(), strings.TrimRight(string(content), "\r\n"))
			}
		}
	}
}
