package controllers

import (
	"backend/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HistoryController struct {
	historyService services.HistoryService
	sessionService services.SessionService
}

func (h HistoryController) GetHistories(ctx *gin.Context) {
	res := h.sessionService.GetNewestSessions()

	ctx.JSON(http.StatusOK, res)
}

func (h HistoryController) GetHistory(ctx *gin.Context) {
	sessionId := ctx.Param("sessionId")

	var exists = h.sessionService.SessionExists(sessionId)

	if !exists {
		ctx.JSON(http.StatusOK, gin.H{
			"response": "data not found",
		})
		return
	}

	res := h.historyService.GetHistory(sessionId)
	ctx.JSON(http.StatusOK, res)
}

func NewHistoryController(historyService services.HistoryService,
	sessionService services.SessionService) HistoryController {
	return HistoryController{
		historyService: historyService,
		sessionService: sessionService,
	}
}
