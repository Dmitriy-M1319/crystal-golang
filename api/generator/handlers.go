package generator

import (
	"net/http"

	"github.com/Dmitriy-M1319/crystal-golang/internal/generator"
	"github.com/gin-gonic/gin"
)

type GeneratorHandler struct {
	service *generator.GeneratorService
}

func NewGeneratorHandler(s *generator.GeneratorService) *GeneratorHandler {
	return &GeneratorHandler{service: s}
}

func (h *GeneratorHandler) GetDummyFile(c *gin.Context) {
	filename := "file1.xlsx"
	err := h.service.CreateDummyFile(filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.File(filename)
}
