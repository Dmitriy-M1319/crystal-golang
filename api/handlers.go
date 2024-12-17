package api

import (
	"net/http"
	"os"

	"github.com/Dmitriy-M1319/crystal-golang/config"
	"github.com/Dmitriy-M1319/crystal-golang/internal/generator"
	"github.com/gin-gonic/gin"
)

type GeneratorHandler struct {
	service *generator.GeneratorService
	env     string
}

func NewGeneratorHandler(s *generator.GeneratorService, envFile string) *GeneratorHandler {
	return &GeneratorHandler{service: s, env: envFile}
}

func (h *GeneratorHandler) GetDummyFile(c *gin.Context) {
	settings, _ := config.GetSettings(h.env)
	_, err := os.Stat(settings.Storage)
	if os.IsNotExist(err) {
		err = os.Mkdir(settings.Storage, 0777)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	filename := settings.Storage + "/file1.xlsx"
	err = h.service.CreateDummyFile(filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.File(filename)
}
