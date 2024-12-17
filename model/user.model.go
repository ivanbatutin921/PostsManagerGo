package model

type Role string

const (
	ADMIN Role = "admin"
	USER  Role = "user"
)

type User struct {
	ID       int `json:"id" gorm:"type:uuid;primary_key;autoIncrement"`
	Login    string `json:"login" gorm:"unique;not null"`
	Password string `json:"-"`
	Role     Role   `json:"role"`
}
