package utils

import "testing"

func TestMaskCPF(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"12345678901", "123.456.789-01"},
		{"00000000000", "000.000.000-00"},
		{"98765432100", "987.654.321-00"},
		// Caso inválido: menos de 11 dígitos (não casa com regex, deve retornar igual)
		{"1234567890", "1234567890"},
		// Caso inválido: mais de 11 dígitos (não casa com regex, deve retornar igual)
		{"123456789012", "123456789012"},
	}

	for _, test := range tests {
		got := MaskCPF(test.input)
		if got != test.expected {
			t.Errorf("MaskCPF(%q) = %q; want %q", test.input, got, test.expected)
		}
	}
}
