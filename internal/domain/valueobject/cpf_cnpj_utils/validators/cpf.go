package validators

import (
	"errors"
	"regexp"
	"strconv"
)

var (
	regexCPFPattern = regexp.MustCompile("^[0-9]{11}$")
)

func CpfIsValid(c string) error {
	if !regexCPFPattern.MatchString(c) {
		return errors.New("invalid CPF format")
	}

	charsRepeated := true
	for i := 1; i < len(c); i++ {
		if c[0] != c[i] {
			charsRepeated = false
			break
		}
	}
	if charsRepeated {
		return errors.New("CPF cannot have all characters the same")
	}

	i, sum := 10, 0
	for index := 0; index < len(c)-2; index++ {
		pos, _ := strconv.Atoi(string(c[index]))
		sum += pos * i
		i--
	}

	prod := sum * 10
	mod := prod % 11

	if mod == 10 {
		mod = 0
	}

	digit1, _ := strconv.Atoi(string(c[9]))
	if mod != digit1 {
		return errors.New("CPF validation digit 1 does not match")
	}

	i, sum = 11, 0
	for index := 0; index < len(c)-1; index++ {
		pos, _ := strconv.Atoi(string(c[index]))
		sum += pos * i
		i--
	}

	prod = sum * 10
	mod = prod % 11

	if mod == 10 {
		mod = 0
	}

	digit2, _ := strconv.Atoi(string(c[10]))
	if mod != digit2 {
		return errors.New("CPF validation digit 2 does not match")
	}

	return nil
}
