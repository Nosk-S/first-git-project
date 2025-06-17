package route

import (
	"gin-project/internal/delivery/http/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(router *gin.Engine, publicHandler *handler.PublicHandler, cardsHandler *handler.CardsHandler) {
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
	api := public.Group("/api")
	{
		api.GET("/cards", cardsHandler.GetCards)
		api.GET("/card", cardsHandler.GetCard)
		api.POST("/research", cardsHandler.CardWhere)
		api.Use(gin.BasicAuth(gin.Accounts{
			"admin": "admin",
		}))

		api.POST("/card", cardsHandler.InsertCard)
		api.DELETE("/card", cardsHandler.DeleteCard)
		api.PATCH("/card", cardsHandler.UpdateCard)

	}
}
