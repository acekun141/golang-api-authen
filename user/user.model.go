package user

import (
	"log"

	"gorm.io/gorm"
)

//User orm for user table
type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
}

//Migrate migrate user model
func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err.Error())
	}
}
