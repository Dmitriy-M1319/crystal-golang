package generator

import (
	"fmt"
	"os"
	"time"

	baseapp "github.com/Dmitriy-M1319/crystal-golang/internal/base-app"
	"github.com/xuri/excelize/v2"
)

type GeneratorService struct {
	repository   IFileRepository
	orderService *baseapp.OrderService
}

func NewGeneratorService(r IFileRepository, o *baseapp.OrderService) *GeneratorService {
	return &GeneratorService{repository: r, orderService: o}
}

func (g *GeneratorService) CreateDummyFile(name string) error {
	file := excelize.NewFile()
	defer file.Close()

	style, _ := file.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
		Border: []excelize.Border{
			{Type: "left", Color: "00FF0000", Style: 1},
			{Type: "right", Color: "00FF0000", Style: 1},
			{Type: "top", Color: "00FF0000", Style: 1},
			{Type: "bottom", Color: "00FF0000", Style: 1},
		},
	})

	file.MergeCell("Sheet1", "A1", "AB1")
	file.SetCellValue("Sheet1", "A1", "ОТЧЕТ ПО ПРОДАЖАМ")

	file.SetCellValue("Sheet1", "A2", "Начало:")
	file.MergeCell("Sheet1", "B2", "C2")
	file.SetCellValue("Sheet1", "B2", time.Now().Format("02 January 2006"))
	file.SetCellValue("Sheet1", "D2", "Конец:")
	file.MergeCell("Sheet1", "E2", "F2")
	file.SetCellValue("Sheet1", "E2", time.Now().Format("02 January 2006"))
	file.MergeCell("Sheet1", "G2", "AB2")

	file.SetCellValue("Sheet1", "A3", "№ заказа")
	file.MergeCell("Sheet1", "B3", "C3")
	file.SetCellValue("Sheet1", "B3", "ФИО клиента")
	file.MergeCell("Sheet1", "D3", "E3")
	file.SetCellValue("Sheet1", "D3", "Номер телефона")
	file.MergeCell("Sheet1", "F3", "G3")
	file.SetCellValue("Sheet1", "F3", "Тип получения")
	file.MergeCell("Sheet1", "H3", "K3")
	file.SetCellValue("Sheet1", "H3", "Адрес доставки (если имеется)")
	file.MergeCell("Sheet1", "L3", "O3")
	file.SetCellValue("Sheet1", "L3", "Создан")
	file.MergeCell("Sheet1", "P3", "S3")
	file.SetCellValue("Sheet1", "P3", "Обновлен")
	file.MergeCell("Sheet1", "T3", "AA3")
	file.SetCellValue("Sheet1", "T3", "Товары")
	file.SetCellValue("Sheet1", "AB3", "Итого")

	// Колонки для товаров
	file.MergeCell("Sheet1", "T4", "U4")
	file.SetCellValue("Sheet1", "T4", "Название")
	file.MergeCell("Sheet1", "V4", "W4")
	file.SetCellValue("Sheet1", "V4", "Производитель")
	file.MergeCell("Sheet1", "X4", "Y4")
	file.SetCellValue("Sheet1", "X4", "Розничная цена")
	file.MergeCell("Sheet1", "Z4", "AA4")
	file.SetCellValue("Sheet1", "Z4", "Количество в заказе")

	file.SetCellStyle("Sheet1", "A1", "AB4", style)

	err := file.SaveAs(name)
	return err
}

func insertOrder(file *excelize.File, order *baseapp.OrderWithProducts, leftIdx int32, isFirst bool) int32 {
	productsCount := len(order.Products)
	if !isFirst {
		productsCount -= 1
	}

	// Заполняем ячейки перед списком товаров
	file.MergeCell("Sheet1", fmt.Sprintf("A%d", leftIdx), fmt.Sprintf("A%d", leftIdx+int32(productsCount)))
	file.SetCellUint("Sheet1", fmt.Sprintf("A%d", leftIdx), order.Order.ID)
	file.MergeCell("Sheet1", fmt.Sprintf("B%d", leftIdx), fmt.Sprintf("C%d", leftIdx+int32(productsCount)))
	file.SetCellValue("Sheet1", fmt.Sprintf("B%d", leftIdx), order.User.Surname+" "+order.User.Name)
	file.MergeCell("Sheet1", fmt.Sprintf("D%d", leftIdx), fmt.Sprintf("E%d", leftIdx+int32(productsCount)))
	file.SetCellValue("Sheet1", fmt.Sprintf("D%d", leftIdx), order.User.Phone)
	file.MergeCell("Sheet1", fmt.Sprintf("F%d", leftIdx), fmt.Sprintf("G%d", leftIdx+int32(productsCount)))
	var t string
	if order.Order.IsDelivery {
		t = "Доставка"
	} else {
		t = "Самовывоз"
	}

	file.SetCellValue("Sheet1", fmt.Sprintf("F%d", leftIdx), t)
	file.MergeCell("Sheet1", fmt.Sprintf("H%d", leftIdx), fmt.Sprintf("K%d", leftIdx+int32(productsCount)))
	file.SetCellValue("Sheet1", fmt.Sprintf("H%d", leftIdx), order.Order.Address)
	file.MergeCell("Sheet1", fmt.Sprintf("L%d", leftIdx), fmt.Sprintf("O%d", leftIdx+int32(productsCount)))
	file.SetCellValue("Sheet1", fmt.Sprintf("L%d", leftIdx), order.Order.CreatedAt.Format("02 January 2006 15:04"))
	file.MergeCell("Sheet1", fmt.Sprintf("P%d", leftIdx), fmt.Sprintf("S%d", leftIdx+int32(productsCount)))
	file.SetCellValue("Sheet1", fmt.Sprintf("P%d", leftIdx), order.Order.UpdatedAt.Format("02 January 2006 15:04"))

	file.MergeCell("Sheet1", fmt.Sprintf("AB%d", leftIdx), fmt.Sprintf("AB%d", leftIdx+int32(productsCount)))
	file.SetCellFloat("Sheet1", fmt.Sprintf("AB%d", leftIdx), order.Order.TotalPrice, 2, 32)

	// Теперь заполняем список товаров
	for i, product := range order.Products {
		line := leftIdx + int32(i)
		if isFirst {
			line += 1
		}
		file.MergeCell("Sheet1", fmt.Sprintf("T%d", line), fmt.Sprintf("U%d", line))
		file.SetCellValue("Sheet1", fmt.Sprintf("T%d", line), product.Name)
		file.MergeCell("Sheet1", fmt.Sprintf("V%d", line), fmt.Sprintf("W%d", line))
		file.SetCellValue("Sheet1", fmt.Sprintf("V%d", line), product.Company)
		file.MergeCell("Sheet1", fmt.Sprintf("X%d", line), fmt.Sprintf("Y%d", line))
		file.SetCellFloat("Sheet1", fmt.Sprintf("X%d", line), product.ClientPrice, 2, 32)
		file.MergeCell("Sheet1", fmt.Sprintf("Z%d", line), fmt.Sprintf("AA%d", line))
		file.SetCellInt("Sheet1", fmt.Sprintf("Z%d", line), int(order.Counts[i]))

	}

	result := leftIdx + int32(productsCount) + 1

	return result
}

func (g *GeneratorService) CreateReport(name string, from, to time.Time) error {
	orders, err := g.orderService.GetOrdersInfo(from, to)
	if err != nil {
		return err
	}

	file := excelize.NewFile()
	defer file.Close()

	style, _ := file.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
		Border: []excelize.Border{
			{Type: "left", Color: "00FF0000", Style: 1},
			{Type: "right", Color: "00FF0000", Style: 1},
			{Type: "top", Color: "00FF0000", Style: 1},
			{Type: "bottom", Color: "00FF0000", Style: 1},
		},
	})

	// Подготовка заголовка
	file.MergeCell("Sheet1", "A1", "AB1")
	file.SetCellValue("Sheet1", "A1", "ОТЧЕТ ПО ПРОДАЖАМ")

	file.SetCellValue("Sheet1", "A2", "Начало:")
	file.MergeCell("Sheet1", "B2", "C2")
	file.SetCellValue("Sheet1", "B2", from.Format("02 January 2006"))
	file.SetCellValue("Sheet1", "D2", "Конец:")
	file.MergeCell("Sheet1", "E2", "F2")
	file.SetCellValue("Sheet1", "E2", to.Format("02 January 2006"))
	file.MergeCell("Sheet1", "G2", "AB2")

	file.SetCellValue("Sheet1", "A3", "№ заказа")
	file.MergeCell("Sheet1", "B3", "C3")
	file.SetCellValue("Sheet1", "B3", "ФИО клиента")
	file.MergeCell("Sheet1", "D3", "E3")
	file.SetCellValue("Sheet1", "D3", "Номер телефона")
	file.MergeCell("Sheet1", "F3", "G3")
	file.SetCellValue("Sheet1", "F3", "Тип получения")
	file.MergeCell("Sheet1", "H3", "K3")
	file.SetCellValue("Sheet1", "H3", "Адрес доставки (если имеется)")
	file.MergeCell("Sheet1", "L3", "O3")
	file.SetCellValue("Sheet1", "L3", "Создан")
	file.MergeCell("Sheet1", "P3", "S3")
	file.SetCellValue("Sheet1", "P3", "Обновлен")
	file.MergeCell("Sheet1", "T3", "AA3")
	file.SetCellValue("Sheet1", "T3", "Товары")
	file.SetCellValue("Sheet1", "AB3", "Итого")

	// Колонки для товаров
	file.MergeCell("Sheet1", "T4", "U4")
	file.SetCellValue("Sheet1", "T4", "Название")
	file.MergeCell("Sheet1", "V4", "W4")
	file.SetCellValue("Sheet1", "V4", "Производитель")
	file.MergeCell("Sheet1", "X4", "Y4")
	file.SetCellValue("Sheet1", "X4", "Розничная цена")
	file.MergeCell("Sheet1", "Z4", "AA4")
	file.SetCellValue("Sheet1", "Z4", "Количество в заказе")

	// Теперь добавляем заказы
	var leftIdx = 4
	isFirst := true
	for _, order := range orders {
		leftIdx = int(insertOrder(file, &order, int32(leftIdx), isFirst))
		isFirst = false
	}

	file.SetCellStyle("Sheet1", "A1", fmt.Sprintf("AB%d", leftIdx+1-len(orders[len(orders)-1].Products)), style)
	err = file.SaveAs(name)
	return err
}

func (g *GeneratorService) GenerateNewReport(path string, from, to time.Time) (string, error) {
	now := time.Now()
	filename := "Report_" + now.Format("2006-01-02") + ".xlsx"
	xlsxFile := XlsxFile{CreatedAt: now, Filename: filename}

	// TODO: Сделать миграцию к бд
	err := g.repository.InsertFile(&xlsxFile)
	if err != nil {
		return "", err
	}

	err = g.CreateReport(path+"/"+filename, from, to)
	if err != nil {
		_ = g.repository.DeleteFile(xlsxFile.ID)
		return "", err
	}

	return path + "/" + filename, nil
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
