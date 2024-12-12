package rest

import (
	"base-gin/server"

	"github.com/gin-gonic/gin"
)

type RestHandler interface {
	init(*server.Handler)
	Route(*gin.Engine)
}

func SetupRestHandlers(app *gin.Engine) {
	handler := server.GetHandler()

	for _, v := range handlers {
		v.init(handler)
		v.Route(app)
	}
}
