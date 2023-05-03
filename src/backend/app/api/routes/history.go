package routes

import (
	"backend/app/api/controllers"
	"backend/app/server"
)

type HistoryRoutes struct {
	requestHandler    server.RequestHandler
	historyController controllers.HistoryController
}

func (h HistoryRoutes) Setup() {
	api := h.requestHandler.Gin.Group("/history")
	{
		api.GET("/", h.historyController.GetHistories)
		api.GET("/:sessionId", h.historyController.GetHistory)
	}
}

func NewHistoryRoutes(requestHandler *server.RequestHandler,
	historyController controllers.HistoryController) HistoryRoutes {
	return HistoryRoutes{
		requestHandler:    *requestHandler,
		historyController: historyController,
	}
}
