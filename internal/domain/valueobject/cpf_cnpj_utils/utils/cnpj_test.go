package utils

import "testing"

func TestMaskCNPJ(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"12345678000195", "12.345.678/0001-95"},
		{"00000000000000", "00.000.000/0000-00"},
		{"98765432000199", "98.765.432/0001-99"},
		// Caso inválido: menos de 14 dígitos (não casa com regex, deve retornar igual)
		{"1234567890123", "1234567890123"},
		// Caso inválido: mais de 14 dígitos (não casa com regex, deve retornar igual)
		{"123456789012345", "123456789012345"},
	}

	for _, test := range tests {
		got := MaskCNPJ(test.input)
		if got != test.expected {
			t.Errorf("MaskCNPJ(%q) = %q; want %q", test.input, got, test.expected)
		}
	}
}
