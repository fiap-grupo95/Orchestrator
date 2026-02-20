package valueobject

import "regexp"

type Plate string

var (
	plateRegex = regexp.MustCompile(`^[A-Z]{3}(?:[0-9]{4}|[0-9][A-Z][0-9]{2})$`)
)

func ParsePlate(value string) Plate {
	return Plate(value)
}

func (v Plate) IsValidFormat() bool {
	return plateRegex.MatchString(string(v))
}

func (v Plate) String() string {
	return string(v)
}
