package commons

import (
	"api-rest/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetConnection() *gorm.DB {

	db, err := gorm.Open("mysql", "root:1234@/person?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println("no se pudo")
		//log.Fatal(err)
	}
	return db
}
func Migrate() {
	db := GetConnection()
	defer db.Close()
	log.Println("Migrating...")

	db.AutoMigrate(&models.Person{})
}
