package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	userRoute := router.Group("/user")
	{
		// Sign up and sign in
		userRoute.POST("/register", h.register)
		userRoute.POST("/auth", h.auth)

		// Get user info: id, name, age
		userRoute.GET("/:name", h.getUser)

		phoneRoute := userRoute.Group("/phone")
		{
			// Add phone
			phoneRoute.POST("/", h.createPhone)
			// Get phone
			//phoneRoute.GET("?q=:phone", h.getPhones)
			// Edit phone
			phoneRoute.PUT("/", h.editPhone)
			// Delete Phone
			phoneRoute.DELETE("/:phoneId", h.deletePhone)
		}
	}
	return router
}
