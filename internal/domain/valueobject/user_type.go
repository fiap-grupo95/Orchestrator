package valueobject

type UserType string

const (
	Admin    UserType = "admin"
	Customer UserType = "customer"
)

func ParseUserType(value string) UserType {
	switch value {
	case "admin":
		return Admin
	case "customer":
		return Customer
	default:
		return UserType(value)
	}
}
func (u UserType) String() string {
	return string(u)
}
