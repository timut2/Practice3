package handler

import "github.com/gin-gonic/gin"
type handler struct{}


func (h *Handler) InitRoutes() *gin.Engine {
	router:= gin.New()

	auth :=router.Group("/auth")
	{
		auth.post("/sign-up", h.signup)
		auth.post("/sign-in", h.singin)
	}

	api:= router.Group("/api")
	{
		lists:= api.Group("/lists")
		{
			lists.Post("/", h.createList)
			lists.Get("/", h.getAllLists)
			lists.Get("/:id",h.getListsById)
			lists.Put("/:id",h.updateItem)
			lists.Delete("/:id",h.deleteItem)
		
		items :=lists.Group(":id/items")
		{
			items.Post("/", h.createItem)
			items.Get("/", h.getAllItems)
			items.Get("/:item_id", h.getListsById)
			items.Put("/:item_id", h.updateItem)
			items.Delete("/:item_id", h.deleteItem)
		}
		}

	}
	return router
}