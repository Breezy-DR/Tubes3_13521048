package routes

import (
	"backend/app/api/controllers"
	"backend/app/server"
)

type QuestionRoutes struct {
	requestHandler     server.RequestHandler
	questionController controllers.QuestionController
}

func (q QuestionRoutes) Setup() {
	api := q.requestHandler.Gin.Group("question")
	{
		api.POST("", q.questionController.GetAnswer)
	}
}

func NewQuestionRoutes(requestHandler *server.RequestHandler,
	questionController controllers.QuestionController) QuestionRoutes {
	return QuestionRoutes{
		requestHandler:     *requestHandler,
		questionController: questionController,
	}
}
