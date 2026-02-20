package entities

import (
	"github.com/daniloAleite/orchestrator/internal/domain/valueobject"
	"time"
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
