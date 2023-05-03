package modules

import (
	"backend/app/api/controllers"
	"backend/app/api/routes"
	"backend/app/gorms"
	"backend/app/services"
	//"backend/app/api/routes"
	"backend/app/server"
	"context"
	"go.uber.org/fx"
)

var Module = fx.Options(
	server.Module,
	services.Module,
	controllers.Module,
	routes.Module,
	gorms.Module,
	fx.Invoke(registerHooks),
)

func registerHooks(lifecycle fx.Lifecycle, h *server.RequestHandler, r routes.Routes) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go func() {
					r.Setup()
					err := h.Gin.Run(":8080")
					if err != nil {
					}
				}()
				return nil
			},
			OnStop: func(context.Context) error {
				return nil
			},
		},
	)
}
