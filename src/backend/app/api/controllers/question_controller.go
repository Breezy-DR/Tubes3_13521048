package controllers

import (
	"backend/app/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type QuestionController struct {
	questionService services.QuestionService
	sessionService  services.SessionService
}

type QuestionRequestBody struct {
	Question  string `json:"question" binding:"required"`
	SearchAlg string `json:"search_algorithm" binding:"required"`
	SessionID string `json:"session_id"`
}

func (q QuestionController) GetAnswer(ctx *gin.Context) {
	var reqBody QuestionRequestBody
	if err := ctx.BindJSON(&reqBody); err != nil {
		fmt.Println("errrrrrrr")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request format",
		})
		return
	}

	var sessionId string
	if sessionId = reqBody.SessionID; sessionId == "" {
		var sessionName string
		if len(reqBody.Question) > 15 {
			sessionName = reqBody.Question[:15]
		} else {
			sessionName = reqBody.Question
		}
		sessionId = q.sessionService.GetNewSession(sessionName)
	}

	resp, found := q.questionService.GetAnswer(reqBody.Question, sessionId, reqBody.SearchAlg)
	ctx.JSON(http.StatusOK, gin.H{
		"answer_found": found,
		"response":     resp,
		"session_id":   sessionId,
	})
}

func NewQuestionController(questionService services.QuestionService,
	sessionService services.SessionService) QuestionController {
	return QuestionController{
		questionService: questionService,
		sessionService:  sessionService,
	}
}
