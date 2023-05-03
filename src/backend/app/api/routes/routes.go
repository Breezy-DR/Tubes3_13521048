package routes

import (
	"go.uber.org/fx"
)

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(questionRoutes QuestionRoutes,
	historyRoutes HistoryRoutes) Routes {
	return Routes{
		questionRoutes,
		historyRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}

var Module = fx.Options(
	fx.Provide(NewQuestionRoutes),
	fx.Provide(NewRoutes),
	fx.Provide(NewHistoryRoutes),
)
