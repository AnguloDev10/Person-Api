package commons

import (
	"log"

	"github.com/jinzhu/gorm"
)

func GetConnection() *gorm.DB {

	db, err := gorm.Open("mysql", "root:1234@/person?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
