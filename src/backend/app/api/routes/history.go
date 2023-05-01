package routes

import (
	"backend/app/api/controllers"
	"backend/app/server"
)

type HistoryRoutes struct {
	requestHandler    server.RequestHandler
	historyController controllers.HistoryController
}
