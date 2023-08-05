package main

import (
	"github.com/paulmanoni/gin-gonic-template/pkg"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(pkg.Module)
	app.Run()
}
