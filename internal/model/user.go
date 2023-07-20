package model

type User struct {
	Base
	UserId   string `gorm:"unique;not null"`
	Username string `gorm:"unique;not null"`
	Nickname string `gorm:"not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"not null"`
}

func (u *User) TableName() string {
	return "users"
}
