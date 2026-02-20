package entities

import (
	"time"

	"github.com/fiap-grupo95/os-service-api/internal/domain/valueobject"
)

type User struct {
	ID        string
	Email     string
	Password  string
	UserType  valueobject.UserType
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Customer  *Customer
}
