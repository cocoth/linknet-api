package helper

import (
	"bytes"
	"strconv"

	"github.com/cocoth/linknet-api/src/http/response"
	"github.com/xuri/excelize/v2"
)

func joinNames(names []string) string {
	result := ""
	for i, name := range names {
		if i > 0 {
			result += ", "
		}
		result += name
	}
	return result
}

func GenerateSurveyExcel(data response.SurveyReportView) ([]byte, error) {
	f := excelize.NewFile()
	sheet := f.GetSheetName(0)

	headers := []string{
		"FORM NUMBER",
		"QUESTOR NAME",
		"FAT",
		"NAMA CUSTOMER",
		"ALAMAT",
		"NODE FDT",
		"TANGGAL SURVEY",
		"STATUS",
		"REMARK",
		"SURVEYOR",
	}

	// Define bold style for headers
	headerStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
	})
	if err != nil {
		return nil, err
	}

	// Set headers and apply bold style
	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, header)
		f.SetCellStyle(sheet, cell, cell, headerStyle)
	}

	row := 2
	f.SetCellValue(sheet, "A"+strconv.Itoa(row), data.FormNumber)
	f.SetCellValue(sheet, "B"+strconv.Itoa(row), data.QuestorName) // Corrected field name
	f.SetCellValue(sheet, "C"+strconv.Itoa(row), data.Fat)         // Corrected field name
	f.SetCellValue(sheet, "D"+strconv.Itoa(row), data.CustomerName)
	f.SetCellValue(sheet, "E"+strconv.Itoa(row), data.Address)
	f.SetCellValue(sheet, "F"+strconv.Itoa(row), data.NodeFDT)
	f.SetCellValue(sheet, "G"+strconv.Itoa(row), data.SurveyDate.Format("02/01/2006 15:04:05")) // Format tanggal
	f.SetCellValue(sheet, "H"+strconv.Itoa(row), data.Status)
	f.SetCellValue(sheet, "I"+strconv.Itoa(row), data.Remark)
	f.SetCellValue(sheet, "J"+strconv.Itoa(row), joinNames(data.Surveyors)) // Corrected field name

	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil

}
