package generator

import "github.com/xuri/excelize/v2"

type GeneratorService struct {
}

func NewGeneratorService() *GeneratorService {
	return &GeneratorService{}
}

func (g *GeneratorService) CreateDummyFile(name string) error {
	file := excelize.NewFile()
	defer file.Close()
	err := file.SaveAs(name)
	return err
}
