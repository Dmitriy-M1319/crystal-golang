package api

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Dmitriy-M1319/crystal-golang/config"
	baseapp "github.com/Dmitriy-M1319/crystal-golang/internal/base-app"
	"github.com/Dmitriy-M1319/crystal-golang/internal/generator"
	"github.com/gin-gonic/gin"
)

type GeneratorHandler struct {
	service    *generator.GeneratorService
	ordService *baseapp.OrderService
	settings   *config.Settings
}

type generationBody struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type fileId struct {
	ID uint64 `uri:"id" binding:"required"`
}

func NewGeneratorHandler(s *generator.GeneratorService, o *baseapp.OrderService) *GeneratorHandler {
	return &GeneratorHandler{service: s, ordService: o, settings: config.GetSettings()}
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

	var body generationBody
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("%v", body)
	from, err := time.Parse("2006-01-02", body.From)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	to, err := time.Parse("2006-01-02", body.To)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := h.service.GenerateNewReport(h.settings.Storage, from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.File(file)
}

func (h *GeneratorHandler) GetFileById(c *gin.Context) {
	var id fileId
	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := h.service.GetFileById(id.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.File(h.settings.Storage + "/" + file)
}

func (h *GeneratorHandler) GetFilesListByPeriod(c *gin.Context) {

	from_str, ok := c.GetQuery("from")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": `"from" field not found`})
		return
	}

	to_str, ok := c.GetQuery("to")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": `"to" field not found`})
		return
	}

	from, err := time.Parse("2006-01-02", from_str)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	to, err := time.Parse("2006-01-02", to_str)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	files, err := h.service.GetFileListByPeriod(from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"files": files,
	})
}

func (h *GeneratorHandler) DeleteFileById(c *gin.Context) {
	var id fileId
	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.DeleteFileById(id.ID, h.settings.Storage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
