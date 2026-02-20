package valueobject

import "regexp"

type Email string

func ParseEmail(value string) Email {
	return Email(value)
}

func (e Email) String() string {
	return string(e)
}

func (e Email) IsValidFormat() bool {
	// A simple regex for email validation
	// Note: This is a basic check and may not cover all edge cases.
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(emailRegex).MatchString(string(e))
}
