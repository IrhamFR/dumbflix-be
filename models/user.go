package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" gorm:"type: varchar(255)"`
	FullName  string    `json:"fullname" gorm:"type: varchar(255)"`
	Email     string    `json:"email" gorm:"type: varchar(255)"`
	Password  string    `json:"password" gorm:"type: varchar(255)"`
	Gender    string    `json:"gender" gorm:"type: varchar(255)"`
	Phone     string    `json:"phone" gorm:"type: varchar(20)"`
	Address   string    `json:"address" gorm:"type: text"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Subscribe bool      `json:"subscribe"`
}

type UsersProfileResponse struct {
	ID        int    `json:"id"`
	FullName  string `json:"fullname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Gender    string `json:"gender"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Subscribe bool   `json:"subscribe"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
