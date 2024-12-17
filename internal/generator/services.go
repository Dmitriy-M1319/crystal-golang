package generator

import (
	"fmt"
	"os"
	"time"

	"github.com/xuri/excelize/v2"
)

type GeneratorService struct {
	repository IFileRepository
}

func NewGeneratorService(r IFileRepository) *GeneratorService {
	return &GeneratorService{repository: r}
}

func (g *GeneratorService) CreateDummyFile(name string) error {
	file := excelize.NewFile()
	defer file.Close()

	data := [][]interface{}{
		{1, "John", 30},
		{2, "Alex", 20},
		{3, "Bob", 40},
	}

	for i, row := range data {
		dataRow := i + 2
		for j, col := range row {
			file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+j)), dataRow), col)
		}
	}

	err := file.SaveAs(name)
	return err
}

func (g *GeneratorService) GenerateNewReport(path string) (string, error) {
	now := time.Now()
	filename := "Report_" + now.Format("2006-01-02_15:04:05")
	xlsxFile := XlsxFile{CreatedAt: now, Filename: filename}

	// Логика заполнения файла
	err := g.repository.InsertFile(&xlsxFile)
	if err != nil {
		return "", nil
	}

	err = g.CreateDummyFile(path + "/" + filename)
	if err != nil {
		_ = g.repository.DeleteFile(xlsxFile.ID)
		return "", err
	}

	return filename, nil
}

func (g *GeneratorService) GetFileById(id uint64) (string, error) {
	file, err := g.repository.GetFileById(id)
	if err != nil {
		return "", err
	} else {
		return file.Filename, nil
	}
}

func (g *GeneratorService) GetFileListByPeriod(from time.Time, to time.Time) ([]string, error) {
	files, err := g.repository.GetFilesForPeriod(from, to)
	if err != nil {
		return nil, err
	} else {
		result := make([]string, 0)
		for _, file := range files {
			result = append(result, file.Filename)
		}
		return result, nil
	}
}

func (g *GeneratorService) DeleteFileById(id uint64, path string) error {
	filename, err := g.GetFileById(id)
	if err != nil {
		return err
	}
	err = os.Remove(path + "/" + filename)
	if err != nil {
		return err
	}

	err = g.repository.DeleteFile(id)
	if err != nil {
		return err
	}

	return nil
}
