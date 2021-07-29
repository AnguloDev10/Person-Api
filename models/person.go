package models

import "github.com/jinzhu/gorm"

type Person struct {
	gorm.Model
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Address   string `json:"address"`
	Cellphone string `json:"cellphone"`
}
