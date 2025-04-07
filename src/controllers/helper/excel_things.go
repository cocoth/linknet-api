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
		"NAMA CUSTOMER",
		"ALAMAT",
		"NODE FDT",
		"TANGGAL SURVEY",
		"STATUS",
		"REMARK",
		"SURVEYOR",
	}

	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, header)
	}

	row := 2
	f.SetCellValue(sheet, "A"+strconv.Itoa(row), data.FormNumber)
	f.SetCellValue(sheet, "B"+strconv.Itoa(row), data.CustomerName)
	f.SetCellValue(sheet, "C"+strconv.Itoa(row), data.Address)
	f.SetCellValue(sheet, "D"+strconv.Itoa(row), data.NodeFDT)
	f.SetCellValue(sheet, "E"+strconv.Itoa(row), data.SurveyDate.Format("2006-01-02")) // Format tanggal
	f.SetCellValue(sheet, "F"+strconv.Itoa(row), data.Status)
	f.SetCellValue(sheet, "G"+strconv.Itoa(row), data.Remark)
	f.SetCellValue(sheet, "H"+strconv.Itoa(row), joinNames(data.Surveyors)) // Corrected field name

	// Simpan file ke buffer
	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil

}
