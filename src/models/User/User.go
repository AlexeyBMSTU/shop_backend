package User

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"primaryKey" json:"user_uid"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
	Email    *string   `json:"email,omitempty"`
}

type UserStore interface {
	GetUserByName(name string) (User, error)
}

type GetUserByNameType func(name string) (User, error)
