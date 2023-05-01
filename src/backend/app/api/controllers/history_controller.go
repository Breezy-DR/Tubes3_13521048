package controllers

import "backend/app/services"

type HistoryController struct {
	historyService services.HistoryService
}

func NewUserController(historyService services.HistoryService) HistoryController {
	return HistoryController{
		historyService: historyService,
	}
}
