package validators

import (
	"testing"
)

func TestCpfIsValid(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "CPF válido",
			input:   "52998224725", // CPF válido real
			wantErr: false,
		},
		{
			name:    "Formato inválido (menos dígitos)",
			input:   "1234567890",
			wantErr: true,
			errMsg:  "invalid CPF format",
		},
		{
			name:    "Formato inválido (caracteres não numéricos)",
			input:   "52998224abc",
			wantErr: true,
			errMsg:  "invalid CPF format",
		},
		{
			name:    "Todos dígitos iguais",
			input:   "11111111111",
			wantErr: true,
			errMsg:  "CPF cannot have all characters the same",
		},
		{
			name:    "Dígito verificador 1 incorreto",
			input:   "52998224735", // alterado último dígito antes do segundo verificador
			wantErr: true,
			errMsg:  "CPF validation digit 1 does not match",
		},
		{
			name:    "Dígito verificador 2 incorreto",
			input:   "52998224726", // alterado último dígito
			wantErr: true,
			errMsg:  "CPF validation digit 2 does not match",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CpfIsValid(tt.input)

			if tt.wantErr && err == nil {
				t.Errorf("esperava erro, mas retornou nil")
			}

			if !tt.wantErr && err != nil {
				t.Errorf("não esperava erro, mas retornou: %v", err)
			}

			if tt.wantErr && err != nil && err.Error() != tt.errMsg {
				t.Errorf("mensagem de erro esperada: '%s', mas retornou: '%s'", tt.errMsg, err.Error())
			}
		})
	}
}
