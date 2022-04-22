package handler

import (
	"crud_v2/app/database"
	"crud_v2/app/jwt"
	"crud_v2/app/redis"
	"crud_v2/entity"
	"fmt"
	"net/http"

	. "crud_v2/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func HomePage(c *gin.Context) {
	user := c.MustGet("user").(entity.User)
	welcomeMessage := fmt.Sprintf("Welcome %s", user.Username)
	ResponseJson(c, http.StatusOK, gin.H{"message": welcomeMessage})
}

func CreateUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBind(&user); err != nil {
		ResponseJson(c, http.StatusBadRequest, gin.H{"error": GetErrorString(err)})
		return
	}

	if user.PlainPassword != user.Password {
		ResponseJson(c, http.StatusBadRequest, gin.H{"error": "Password not match"})
		return
	}

	data, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(data)
	user.PlainPassword = "secret"
	if err := database.Connector.Create(&user).Error; err != nil {
		ResponseJson(c, http.StatusBadRequest, gin.H{"error": GetErrorString(err)})
		return
	}

	ResponseJson(c, http.StatusOK, user)
}

func LoginHandler(c *gin.Context) {
	var login login
	if err := c.ShouldBind(&login); err != nil {
		ResponseJson(c, http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user entity.User
	if err := database.Connector.Where("username = ?", login.Username).First(&user).Error; err != nil {
		ResponseJson(c, http.StatusBadRequest, gin.H{"error": GetErrorString(err)})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		ResponseJson(c, http.StatusBadRequest, gin.H{"error": GetErrorString(err)})
		return
	}

	token, err := jwt.MakeJWT(&user)
	if err != nil {
		ResponseJson(c, http.StatusBadRequest, gin.H{"error": GetErrorString(err)})
		return
	}

	resultData := map[string]interface{}{
		"token": token,
	}

	ResponseJson(c, http.StatusOK, resultData)
}

func LogoutHandler(c *gin.Context) {
	auid, _ := c.MustGet("AuthId").(string)
	redis.RemoveKey(auid)
	redis.RemoveKey(fmt.Sprintf("user-%s", auid))
	ResponseJson(c, http.StatusOK, gin.H{"message": "Logout success"})
}
