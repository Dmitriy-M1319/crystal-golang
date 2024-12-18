package api

import (
	"github.com/gin-gonic/gin"
)

type GeneratorRouter struct {
	handlers *GeneratorHandler
	router   *gin.Engine
}

func NewGeneratorRouter(h *GeneratorHandler) *GeneratorRouter {
	r := GeneratorRouter{handlers: h, router: gin.Default()}
	r.router.POST("/report", r.handlers.GenerateReport)
	r.router.POST("/dummy", r.handlers.GetDummyFile)
	r.router.GET("/files/:id", r.handlers.GetFileById)
	r.router.GET("/files", r.handlers.GetFilesListByPeriod)
	r.router.DELETE("/files/:id", r.handlers.DeleteFileById)
	return &r
}

func (r *GeneratorRouter) Run(addr string) {
	r.router.Run(addr)
}
