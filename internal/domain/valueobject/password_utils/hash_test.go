package pkg

import (
	"encoding/base64"
	"errors"
	"strings"
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "minhaSenha123"

	hashed, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword retornou erro inesperado: %v", err)
	}

	// Verifica se contém exatamente um ponto separando salt e hash
	parts := strings.Split(hashed, ".")
	if len(parts) != 2 {
		t.Errorf("HashPassword deve retornar string com formato salt.hash, mas retornou: %s", hashed)
	}

	// Verifica se salt e hash são base64 válidos (tentativa de decode)
	if _, err := decodeBase64(parts[0]); err != nil {
		t.Errorf("Salt não é base64 válido: %v", err)
	}
	if _, err := decodeBase64(parts[1]); err != nil {
		t.Errorf("Hash não é base64 válido: %v", err)
	}

	// Opcional: testar que o hash para duas senhas iguais é diferente (por causa do salt)
	hashed2, _ := HashPassword(password)
	if hashed == hashed2 {
		t.Errorf("HashPassword retornou hash idêntico para duas chamadas, deveria variar por causa do salt")
	}
}

// decodeBase64 tenta decodificar base64 e retorna erro se inválido
func decodeBase64(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

func TestHashPasswordSaltGenerationError(t *testing.T) {
	// Mockar generateSaltFunc para sempre retornar erro
	generateSaltFunc = func(length int) ([]byte, error) {
		return nil, errors.New("salt error")
	}
	defer func() { generateSaltFunc = generateSalt }() // resetar após teste

	_, err := HashPassword("teste")
	if err == nil || !strings.Contains(err.Error(), "failed to generate salt") {
		t.Errorf("Esperava erro na geração de salt, mas não recebeu erro correto")
	}
}
