package models

import "github.com/jinzhu/gorm"

type Person struct {
	gorm.Model

	ID int64 `json:"id"`
	Name string `json:"name"`
	Age int64 `json:"age"`
}