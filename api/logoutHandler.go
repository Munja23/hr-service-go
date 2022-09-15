package api

import "github.com/gin-gonic/gin"

//Logout handler
func (a API) LogoutHandler(c *gin.Context) {

	c.SetCookie("token", "", -1, "", "", false, true)

}
