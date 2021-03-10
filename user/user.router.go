package user

import (
	"learn-gin/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Router struct {
	engine  *gin.Engine
	service Service
}

func (r *Router) SignUp(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	if user.Username == "" || user.Password == "" || user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please complete all fields"})
		return
	}
	if err := r.service.CreateUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func (r *Router) GetCurrentUser(c *gin.Context) {
	userID, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"userID": userID})
}

func InitialRouter(engine *gin.Engine, db *gorm.DB) {
	router := Router{
		engine:  engine,
		service: Service{DB: db},
	}
	router.engine.POST("/user", router.SignUp)
	router.engine.GET("/user", middleware.Authen, router.GetCurrentUser)
}
