package route

import (
	"gin-project/internal/delivery/http/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(router *gin.Engine, publicHandler *handler.PublicHandler) {
	public := router.Group("/")
	{
		public.GET("/", publicHandler.Index)
	}
	admin := public.Group("/admin")
	{
		admin.Use(gin.BasicAuth(gin.Accounts{
			"admin": "admin",
		}))
		admin.GET("/", publicHandler.Admin)
	}
	// api := public.Group("/api")
	// {
	// 	api.GET("/cards", publicHandler.Getcards)
	// }

}
