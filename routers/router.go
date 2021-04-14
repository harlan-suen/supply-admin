package routers

import (
	"github.com/gin-gonic/gin"
	"supply-admin/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("api/auth/login", api.GetAuth)
	r.POST("api/auth/register", api.Register)

	guardApi := r.Group("guard")
	{
		guardApi.GET("user", nil)
	}

	return r
}
