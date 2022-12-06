package commons

import (
	"log"

	"github.com/jinzhu/gorm"
)

func GetConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "usreprac:@/test")

	if err != nil {
		log.Fatal(err)
	}
	return db
}
