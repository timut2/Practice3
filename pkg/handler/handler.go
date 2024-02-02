package handler

import ("github.com/gin-gonic/gin"
"github.com/timut2/Practice3/pkg/service"
)
type Handler struct{
	service *service.Service
}


func (h *Handler) InitRoutes() *gin.Engine {
	router:= gin.New()

	auth :=router.Group("/auth")
	{
		auth.POST("/sign-up", h.signup)
		auth.POST("/sign-in", h.signin)
	}

	api:= router.Group("/api")
	{
		lists:= api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id",h.getListsById)
			lists.PUT("/:id",h.updateItem)
			lists.DELETE("/:id",h.deleteItem)
		
		items :=lists.Group(":id/items")
		{
			items.POST("/", h.createItem)
			items.GET("/", h.getAllItems)
			items.GET("/:item_id", h.getListsById)
			items.PUT("/:item_id", h.updateItem)
			items.DELETE("/:item_id", h.deleteItem)
		}
		}

	}
	return router
}