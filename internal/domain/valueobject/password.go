package valueobject

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/fiap-grupo95/os-service-api/pkg/utils/password"
	"strings"

	"golang.org/x/crypto/argon2"
)

type Password string // representa sempre uma senha *hashada*

func NewPassword(p string) (Password, error) {
	if err := isValid(p); err != nil {
		return "", err
	}

	hashed, err := pkg.HashPassword(p)
	if err != nil {
		return "", err
	}

	return Password(hashed), nil
}

func (p Password) String() string {
	return string(p)
}

func (p Password) Verify(plain string) bool {
	parts := strings.Split(string(p), ".")
	if len(parts) != 2 {
		return false
	}

	salt, err := base64.StdEncoding.DecodeString(parts[0])
	if err != nil {
		fmt.Println("Erro salt:", err)
		return false
	}
	expectedHash, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		fmt.Println("Erro hash:", err)
		return false
	}

	hash := argon2.IDKey([]byte(plain), salt, 1, 64*1024, 4, 32)
	return subtleCompare(expectedHash, hash)
}

func subtleCompare(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	var result byte
	for i := range a {
		result |= a[i] ^ b[i]
	}
	return result == 0
}

func isValid(s string) error {
	if len(s) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	hasUpper := false
	hasLower := false
	hasDigit := false

	for _, char := range s {
		switch {
		case char >= 'A' && char <= 'Z':
			hasUpper = true
		case char >= 'a' && char <= 'z':
			hasLower = true
		case char >= '0' && char <= '9':
			hasDigit = true
		}
	}

	if !(hasUpper && hasLower && hasDigit) {
		return errors.New("password must contain at least one uppercase letter, one lowercase letter, and one digit")
	}
	return nil
}
