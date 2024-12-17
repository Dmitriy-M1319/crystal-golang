package generator

import (
	"github.com/gin-gonic/gin"
)

type GeneratorRouter struct {
	handlers *GeneratorHandler
	router   *gin.Engine
}

func NewGeneratorRouter(h *GeneratorHandler, e *gin.Engine) *GeneratorRouter {
	r := GeneratorRouter{handlers: h, router: gin.Default()}
	r.router.POST("/dummy", r.handlers.GetDummyFile)
	return &r
}

func (r *GeneratorRouter) Run(addr string) {
	r.router.Run(addr)
}
