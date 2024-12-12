package rest

import (
	"base-gin/server"

	"github.com/gin-gonic/gin"
)

type RestHandler interface {
	init(*server.Handler)
	Route(*gin.Engine)
}

var handlers []RestHandler = []RestHandler{
	&AccountHandler{},
	&PersonHandler{},
}

func SetupRestHandlers(app *gin.Engine) {
	handler := server.GetHandler()

	for _, v := range handlers {
		v.init(handler)
		v.Route(app)
	}
}
