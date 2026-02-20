package valueobject

import "regexp"

type PhoneNumber string

// ParsePhoneNumber cria um parse PhoneNumber que retorne um PhoneNumber
func ParsePhoneNumber(value string) PhoneNumber {
	return PhoneNumber(value)
}

// IsValidFormat verifica se o número de telefone está no formato válido
func (p PhoneNumber) IsValidFormat() bool {
	// Um regex simples para validação de número de telefone
	// Nota: Isso é uma verificação básica e pode não cobrir todos os casos.
	phoneRegex := `^\+?[1-9]\d{1,14}$`
	return regexp.MustCompile(phoneRegex).MatchString(string(p))
}

// String retorna a representação em string do PhoneNumber
func (p PhoneNumber) String() string {
	return string(p)
}

// IsEmpty verifica se o PhoneNumber está vazio
func (p PhoneNumber) IsEmpty() bool {
	return len(p) == 0
}

// IsEqual verifica se dois PhoneNumbers são iguais
func (p PhoneNumber) IsEqual(other PhoneNumber) bool {
	return p.String() == other.String()
}

// IsNotEqual verifica se dois PhoneNumbers são diferentes
func (p PhoneNumber) IsNotEqual(other PhoneNumber) bool {
	return !p.IsEqual(other)
}

// IsSame verifica se dois PhoneNumbers são iguais
func (p PhoneNumber) IsSame(other PhoneNumber) bool {
	return p.IsEqual(other)
}

// IsValid verifica se o PhoneNumber é válido
func (p PhoneNumber) IsValid() bool {
	return p.IsValidFormat() && !p.IsEmpty()
}
