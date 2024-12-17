package generator

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

type GeneratorService struct {
}

func NewGeneratorService() *GeneratorService {
	return &GeneratorService{}
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
