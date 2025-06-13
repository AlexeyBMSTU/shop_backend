package User

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"primaryKey" json:"user_uid"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
}
