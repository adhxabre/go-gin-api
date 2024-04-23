package models

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	Name      string    `json:"name" form:"name" gorm:"type:varchar(100)"`
	Email     string    `json:"email" form:"email" gorm:"varchar(255)"`
	Password  string    `json:"password" form:"password" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
