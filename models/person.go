package models

type Person struct {
	ID        int64  `json:"id" gorm:"primary_key;auto_increment"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Address   string `json:"address"`
	Cellphone string `json:"cellphone"`
}
