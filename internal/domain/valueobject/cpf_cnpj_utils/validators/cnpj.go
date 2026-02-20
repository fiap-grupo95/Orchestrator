package validators

import (
	"errors"
	"fmt"
	"regexp"
)

const (
	sizeCNPJWithoutDV       = 12
	baseValue               = int('0')
	cnpjZerosOnly           = "00000000000000"
	operationAndErrorFormat = "%s: %w"
)

var (
	regexCNPJWithoutDV   = regexp.MustCompile("^[0-9]{12}$")
	regexCNPJ            = regexp.MustCompile("^[0-9]{14}$")
	regexCharsNotAllowed = regexp.MustCompile("[^0-9./-]")
	regexNonNumericChars = regexp.MustCompile("[^0-9]")
	weightDV             = []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
)

func CnpjIsValid(c string) error {
	if regexCharsNotAllowed.MatchString(c) {
		return errors.New("CNPJ contains invalid characters")
	}

	cnpjWithoutMask := clean(c)
	if !regexCNPJ.MatchString(cnpjWithoutMask) || cnpjWithoutMask == cnpjZerosOnly {
		return errors.New("CNPJ has an invalid pattern")
	}

	calculatedCNPJDV, err := calculateAndValidateCNPJDV(cnpjWithoutMask[:sizeCNPJWithoutDV])
	if err != nil {
		return err
	}

	informedCNPJDV := cnpjWithoutMask[sizeCNPJWithoutDV:]
	if informedCNPJDV != calculatedCNPJDV {
		return errors.New("CNPJ validation digits do not match")
	}

	return nil
}

func calculateAndValidateCNPJDV(c string) (string, error) {
	const op = "validator.calculateAndValidateCNPJDV"

	if regexCharsNotAllowed.MatchString(c) {
		return "", fmt.Errorf(
			operationAndErrorFormat,
			op,
			errors.New("CNPJ contains invalid characters"),
		)
	}

	CNPJWithoutMask := clean(c)
	if !regexCNPJWithoutDV.MatchString(CNPJWithoutMask) || CNPJWithoutMask == cnpjZerosOnly[:sizeCNPJWithoutDV] {
		return "", fmt.Errorf(
			operationAndErrorFormat,
			op,
			errors.New("CNPJ has an invalid pattern"),
		)
	}

	sumDV1 := 0
	sumDV2 := 0
	for i := 0; i < sizeCNPJWithoutDV; i++ {
		asciiDigit := int(CNPJWithoutMask[i]) - baseValue

		sumDV1 += asciiDigit * weightDV[i+1]
		sumDV2 += asciiDigit * weightDV[i]
	}

	var dv1 int
	if sumDV1%11 < 2 {
		dv1 = 0
	} else {
		dv1 = 11 - (sumDV1 % 11)
	}

	sumDV2 += dv1 * weightDV[sizeCNPJWithoutDV]

	var dv2 int
	if sumDV2%11 < 2 {
		dv2 = 0
	} else {
		dv2 = 11 - (sumDV2 % 11)
	}

	return fmt.Sprintf("%d%d", dv1, dv2), nil
}

func clean(c string) string {
	return regexNonNumericChars.ReplaceAllString(c, "")
}
