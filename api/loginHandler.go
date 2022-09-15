package api

import (
	"crypto/sha256"
	"encoding/base64"

	"gitea.delsystems.academy/academy/hr-service-go/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const (
	UserDriver = "mysql"
)

//Login handler for credential validation and returning token
func (a API) LoginHandler(c *gin.Context) {
	var (
		credentials model.Login
		queries     = make(model.Queries)
	)
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.AbortWithStatus(401)
		return
	}
	queries["email"] = []string{credentials.Email}
	user, err := a.Stg[UserDriver].UserRead(queries)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}

	passwordHash := sha256.Sum256([]byte(credentials.Password))

	if err := bcrypt.CompareHashAndPassword(passwordHash[:], user[0].Password[:]); err != nil {
		c.AbortWithStatus(401)
		return
	}

	data := credentials.Email + ":" + string(passwordHash[:])
	token := base64.StdEncoding.EncodeToString([]byte(data))

	c.SetCookie("token", token, 3600, "", "", false, true)

}
