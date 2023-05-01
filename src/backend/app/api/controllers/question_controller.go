package controllers

import (
	"backend/app/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type QuestionController struct {
	questionService services.QuestionService
}

type QuestionRequestBody struct {
	Question string `json:"question" binding:"required"`
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

	ctx.JSON(http.StatusOK, gin.H{
		"data": q.questionService.GetAnswer(reqBody.Question),
	})
}

func NewQuestionController(questionService services.QuestionService) QuestionController {
	return QuestionController{
		questionService: questionService,
	}
}
