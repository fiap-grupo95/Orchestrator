package utils

import "regexp"

func MaskCPF(c string) string {
	re := regexp.MustCompile(`^(\d{3})(\d{3})(\d{3})(\d{2})$`)
	return re.ReplaceAllString(c, "$1.$2.$3-$4")
}
