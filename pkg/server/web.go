package server

import (
	"github.com/gin-gonic/gin"
	"github.com/paulmanoni/gin-gonic-template/pkg/controller"
	"go.uber.org/fx"
)

type WebServerParams struct {
	fx.In
	Main controller.MainController
}

func NewWebServer(Param WebServerParams) *gin.Engine {
	r := gin.Default()

	r.GET("/", Param.Main.Home)

	return r
}
