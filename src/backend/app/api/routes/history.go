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
	api := h.requestHandler.Gin.Group("history")
	{
		api.POST("", h.historyController.GetHistories)
		api.POST(":sessionId", h.historyController.GetHistory)
	}
}

func NewHistoryRoutes(requestHandler *server.RequestHandler,
	historyController controllers.HistoryController) HistoryRoutes {
	return HistoryRoutes{
		requestHandler:    *requestHandler,
		historyController: historyController,
	}
}
