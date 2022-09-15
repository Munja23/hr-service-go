package routes

import (
	"gitea.delsystems.academy/academy/hr-service-go/api"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(apiConn api.API, router *gin.Engine) {
	v1 := router.Group("/v1")

	organisation := v1.Group("/org")

	{
		OrganisationHandler(apiConn, organisation)
	}
}

func OrganisationHandler(apiConn api.API, r *gin.RouterGroup) {
	r.GET("", apiConn.ReadOrganisations)

	r.GET("/:code", apiConn.ReadOrganisations)

	r.POST("", apiConn.CreateOrganisation)
}

func UserHandler(apiConn api.API, r *gin.RouterGroup) {
	r.GET("", apiConn.ReadUsers)

	r.GET("/:email", apiConn.ReadUsers)

	r.POST("", apiConn.CreateUser)
}

func LoginHandler(apiConn api.API, r *gin.RouterGroup) {
	r.POST("/login", apiConn.LoginHandler)

	r.DELETE("/logout", apiConn.LogoutHandler)
}
