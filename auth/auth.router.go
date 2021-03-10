package auth

import (
	"learn-gin/user"
	"learn-gin/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Router struct {
	engine *gin.Engine
	db     *gorm.DB
}

func (r *Router) Login(c *gin.Context) {
	var userFromBody user.User
	if err := c.ShouldBindJSON(&userFromBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	userService := user.Service{DB: r.db}
	userLogin, err := userService.GetUserLogin(userFromBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	token, err := util.CreateToken(userLogin.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func InitialRouter(engine *gin.Engine, db *gorm.DB) {
	router := Router{
		engine: engine,
		db:     db,
	}
	router.engine.POST("/auth/login", router.Login)
}
