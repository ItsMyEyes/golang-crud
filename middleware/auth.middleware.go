package middleware

import (
	"crud_v2/app/jwt"
	"crud_v2/app/redis"
	"crud_v2/entity"
	"crud_v2/repository"
	. "crud_v2/utils"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		errorReponse(c)
		c.Abort()
		return
	}
	token := strings.Replace(header, "Bearer ", "", -1)
	checkAuth, _, claims := jwt.Verify(token)

	if !checkAuth && claims == nil {
		errorReponse(c)
		c.Abort()
		return
	}

	if !repository.ValidUser(claims.UserId) {
		errorReponse(c)
		c.Abort()
		return
	}

	var user entity.User
	resUser := redis.GetKey(claims.AuthId)
	errz := json.Unmarshal([]byte(resUser.Val()), &user)
	if errz != nil {
		log.Println(errz)
		errorReponse(c)
		c.Abort()
		return
	}
	c.Set("AuthId", claims.AuthId)
	c.Set("user", user)

	// Pass on to the next-in-chain
	c.Next()
}

func errorReponse(c *gin.Context) {
	log.Println("AuthMiddleware error", c.Request.URL.Path)
	res := map[string]interface{}{
		"error": "You need to login first",
	}
	ResponseJson(c, http.StatusUnauthorized, res)
	return
}
