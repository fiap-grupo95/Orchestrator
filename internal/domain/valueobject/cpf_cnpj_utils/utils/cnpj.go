package utils

import "regexp"

func MaskCNPJ(c string) string {
	re := regexp.MustCompile(`^(\d{2})(\d{3})(\d{3})(\d{4})(\d{2})$`)
	return re.ReplaceAllString(c, "$1.$2.$3/$4-$5")
}
