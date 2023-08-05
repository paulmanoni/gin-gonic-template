package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/paulmanoni/gin-gonic-template/pkg/config"
	"github.com/paulmanoni/gin-gonic-template/pkg/controller"
	"github.com/paulmanoni/gin-gonic-template/pkg/server"
	"go.uber.org/fx"
	"strconv"
)

var Module = fx.Module(
	"pkg",
	config.Module,
	server.Module,
	controller.Module,
	fx.Invoke(func(conf *config.Config, r *gin.Engine) {
		err := r.Run(":" + strconv.Itoa(conf.Port))
		if err != nil {
			return
		}
	}),
)
