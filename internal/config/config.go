package config

import (
	"fmt"
	"os"
	"path/filepath"
)

const appDataName = "DoubaoSpace"

// EnvOr returns the environment value when it is set, otherwise fallback.
func EnvOr(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

// ResolveDataDir keeps mutable data outside the executable and application bundle.
// Development uses ./data for easy inspection; production uses the OS user config dir.
func ResolveDataDir(explicit string, dev bool) (string, error) {
	if explicit != "" {
		return absolutePath(explicit)
	}
	if dev {
		return absolutePath("data")
	}

	base, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("get user config directory: %w", err)
	}
	return filepath.Join(base, appDataName), nil
}

// PrepareDataDir creates only directories that are part of the runtime contract.
func PrepareDataDir(dataDir string) error {
	for _, dir := range []string{
		dataDir,
		filepath.Join(dataDir, "backups"),
		filepath.Join(dataDir, "uploads"),
	} {
		if err := os.MkdirAll(dir, 0o700); err != nil {
			return fmt.Errorf("create %s: %w", dir, err)
		}
	}
	return nil
}

func absolutePath(value string) (string, error) {
	path, err := filepath.Abs(value)
	if err != nil {
		return "", fmt.Errorf("resolve path %q: %w", value, err)
	}
	return filepath.Clean(path), nil
}
