package main

import (
	"learn-gin/auth"
	"learn-gin/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	// migrate
	user.Migrate(db)
	// router
	user.InitialRouter(r, db)
	auth.InitialRouter(r, db)
	r.Run(":8888")
}
