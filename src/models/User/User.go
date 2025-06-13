package User

type User struct {
	ID       uint64 `gorm:"primaryKey" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
