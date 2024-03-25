package handler

func (h *Handler) routes() {
	// Healthz segment
	h.router.GET("/healthz", h.healthCheck)

	// Movies segment
	movieGroup := h.router.Group("/Movies")
	movieGroup.GET("", h.listMovie)
	movieGroup.GET("/:id", h.detailMovie)
	movieGroup.POST("", h.addNewMovie)
	movieGroup.PATCH("/:id", h.updateMovie)
	movieGroup.DELETE("/:id", h.deleteMovie)
}
