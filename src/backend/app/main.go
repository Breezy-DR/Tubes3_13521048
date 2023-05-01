package main

import (
	"backend/app/modules"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		modules.Module,
	).Run()
}
