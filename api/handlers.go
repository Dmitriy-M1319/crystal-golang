package api

import (
	"net/http"
	"os"

	"github.com/Dmitriy-M1319/crystal-golang/config"
	"github.com/Dmitriy-M1319/crystal-golang/internal/generator"
	"github.com/gin-gonic/gin"
)

type GeneratorHandler struct {
	service  *generator.GeneratorService
	settings *config.Settings
}

func NewGeneratorHandler(s *generator.GeneratorService) *GeneratorHandler {
	return &GeneratorHandler{service: s, settings: config.GetSettings()}
}

func (h *GeneratorHandler) GetDummyFile(c *gin.Context) {
	_, err := os.Stat(h.settings.Storage)
	if os.IsNotExist(err) {
		err = os.Mkdir(h.settings.Storage, 0777)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	filename := h.settings.Storage + "/file1.xlsx"
	err = h.service.CreateDummyFile(filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.File(filename)
}

func (h *GeneratorHandler) GenerateReport(c *gin.Context) {
	_, err := os.Stat(h.settings.Storage)
	if os.IsNotExist(err) {
		err = os.Mkdir(h.settings.Storage, 0777)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	file, err := h.service.GenerateNewReport(h.settings.Storage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.File(file)
}

func (h *GeneratorHandler) GetFileById(c *gin.Context) {
	file, err := h.service.GetFileById(c.GetUint64("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.File(h.settings.Storage + "/" + file)
}

func (h *GeneratorHandler) GetFilesListByPeriod(c *gin.Context) {
	files, err := h.service.GetFileListByPeriod(c.GetTime("from"), c.GetTime("to"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"files": files,
	})
}

func (h *GeneratorHandler) DeleteFileById(c *gin.Context) {
	err := h.service.DeleteFileById(c.GetUint64("id"), h.settings.Storage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
