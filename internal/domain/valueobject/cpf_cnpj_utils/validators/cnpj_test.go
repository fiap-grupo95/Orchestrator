package validators

import (
	"testing"
)

func TestCnpjIsValid(t *testing.T) {
	tests := []struct {
		name    string
		cnpj    string
		wantErr bool
	}{
		{"válido sem máscara", "11222333000181", false},
		{"válido com máscara", "11.222.333/0001-81", false},
		{"caracteres inválidos", "11.222.333/0001-8A", true},
		{"dígitos verificadores errados", "11222333000182", true},
		{"todos zeros", "00.000.000/0000-00", true},
		{"tamanho errado", "123456789", true},
		{"vazio", "", true},
		{
			name:    "caracteres especiais extras",
			cnpj:    "11@222!333#0001$81",
			wantErr: true, // esperado erro pois há caracteres inválidos
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CnpjIsValid(tt.cnpj)
			if (err != nil) != tt.wantErr {
				t.Errorf("CnpjIsValid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCalculateAndValidateCNPJDV(t *testing.T) {
	tests := []struct {
		name    string
		cnpj12  string
		wantDV  string
		wantErr bool
	}{
		{"válido sem máscara", "112223330001", "81", false},
		{"caracteres inválidos", "11A223330001", "", true},
		{"todos zeros", "000000000000", "", true},
		{"tamanho errado", "123", "", true},
		{"com caracteres especiais", "11.222.333/0001", "81", false}, // deve limpar e calcular
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dv, err := calculateAndValidateCNPJDV(tt.cnpj12)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateAndValidateCNPJDV() error = %v, wantErr %v", err, tt.wantErr)
			}
			if dv != tt.wantDV {
				t.Errorf("CalculateAndValidateCNPJDV() = %v, want %v", dv, tt.wantDV)
			}
		})
	}
}
