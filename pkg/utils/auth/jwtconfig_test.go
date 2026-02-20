package auth

import (
	"os"
	"testing"
	"time"
)

const secretKeyErrMsg = "esperado SecretKey = %q, obtido %q"

func TestLoadJWTConfigDefaults(t *testing.T) {
	// Garantir que variáveis não estão setadas
	os.Unsetenv("JWT_SECRET")
	os.Unsetenv("JWT_TTL")

	cfg := LoadJWTConfig()

	if cfg.SecretKey != "default_secret_key" {
		t.Errorf(secretKeyErrMsg, "default_secret_key", cfg.SecretKey)
	}

	if cfg.ExpirationTTL != time.Hour*24 {
		t.Errorf("esperado ExpirationTTL = %v, obtido %v", time.Hour*24, cfg.ExpirationTTL)
	}
}

func TestLoadJWTConfigFromEnv(t *testing.T) {
	os.Setenv("JWT_SECRET", "my_secret")
	os.Setenv("JWT_TTL", "2h")
	defer os.Unsetenv("JWT_SECRET")
	defer os.Unsetenv("JWT_TTL")

	cfg := LoadJWTConfig()

	if cfg.SecretKey != "my_secret" {
		t.Errorf(secretKeyErrMsg, "my_secret", cfg.SecretKey)
	}

	if cfg.ExpirationTTL != 2*time.Hour {
		t.Errorf("esperado ExpirationTTL = %v, obtido %v", 2*time.Hour, cfg.ExpirationTTL)
	}
}

func TestLoadJWTConfigInvalidDuration(t *testing.T) {
	os.Setenv("JWT_SECRET", "valid_secret")
	os.Setenv("JWT_TTL", "invalid_duration")
	defer os.Unsetenv("JWT_SECRET")
	defer os.Unsetenv("JWT_TTL")

	cfg := LoadJWTConfig()

	if cfg.SecretKey != "valid_secret" {
		t.Errorf(secretKeyErrMsg, "valid_secret", cfg.SecretKey)
	}

	// Deve cair para o valor padrão
	if cfg.ExpirationTTL != time.Hour*24 {
		t.Errorf("esperado ExpirationTTL padrão = %v, obtido %v", time.Hour*24, cfg.ExpirationTTL)
	}
}
