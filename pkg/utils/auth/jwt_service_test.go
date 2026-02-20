package auth

import (
	"encoding/base64"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestJWTServiceInvalidSigningMethod(t *testing.T) {
	cfg := &JWTConfig{
		SecretKey:     "test_secret",
		ExpirationTTL: time.Hour,
	}
	service := NewJWTService(cfg)

	header := `{"alg":"RS256","typ":"JWT"}`
	payload := fmt.Sprintf(`{"sub":"user123","exp":%d,"iat":%d}`,
		time.Now().Add(time.Hour).Unix(),
		time.Now().Unix())

	segment := func(s string) string {
		return base64.RawURLEncoding.EncodeToString([]byte(s))
	}

	fakeToken := segment(header) + "." + segment(payload) + "." + segment("fake_signature")

	_, err := service.ValidateToken(fakeToken)
	if err == nil {
		t.Fatalf("esperado erro, mas veio nil")
	}

	// Aceita qualquer erro que prove que a verificação de método foi acionada
	if !strings.Contains(err.Error(), "invalid") {
		t.Errorf("erro inesperado: %v", err)
	}
}

func TestJWTService(t *testing.T) {
	cfg := &JWTConfig{
		SecretKey:     "test_secret",
		ExpirationTTL: time.Hour,
	}
	service := NewJWTService(cfg)

	// Teste: Gerar e validar token
	subject := "user123"
	tokenStr, err := service.GenerateToken(subject)
	if err != nil {
		t.Fatalf("erro ao gerar token: %v", err)
	}

	token, err := service.ValidateToken(tokenStr)
	if err != nil {
		t.Fatalf("erro ao validar token: %v", err)
	}

	if !token.Valid {
		t.Errorf("token gerado deveria ser válido")
	}

	// Extrair claims e verificar subject
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		t.Fatalf("falha ao converter claims para MapClaims")
	}

	if claims["sub"] != subject {
		t.Errorf("subject incorreto, esperado %q, obtido %q", subject, claims["sub"])
	}

	// Teste: Token inválido
	invalidToken := tokenStr + "abc"
	_, err = service.ValidateToken(invalidToken)
	if err == nil {
		t.Errorf("token inválido deveria gerar erro")
	}
}

func TestJWTServiceExpiredToken(t *testing.T) {
	cfg := &JWTConfig{
		SecretKey:     "test_secret",
		ExpirationTTL: -time.Second, // já expirado
	}
	service := NewJWTService(cfg)

	tokenStr, err := service.GenerateToken("expired_user")
	if err != nil {
		t.Fatalf("erro ao gerar token: %v", err)
	}

	token, err := service.ValidateToken(tokenStr)
	if err == nil && token.Valid {
		t.Errorf("token expirado deveria ser inválido")
	}
}
