package modules

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"
	"strconv"
	"time"
)

func LearningExcel() {
	CreateExcel()
	ReadingExcel()
}

type Student struct {
	Name   string
	Gender string
	Age    int
}

var DummyData = []Student{
	{"Jon", "Male", 18},
	{"David", "Female", 20},
	{"Jack", "Female", 20},
	{"Robert", "Male", 20},
}

func CreateExcel() {
	var (
		xlsx     = excelize.NewFile()
		filePath = "./assets/excel/" + time.Now().Format("01-02-2006") + ".xlsx"
	)

	CreateBasicSheet(xlsx)

	CreateBasicSheetWithMergedCell(xlsx)

	err := xlsx.SaveAs(filePath)
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}
}

func CreateBasicSheet(xlsx *excelize.File) {
	var (
		sheetName        = "student_data"
		defaultSheetName = xlsx.GetSheetName(1)
	)

	xlsx.SetSheetName(defaultSheetName, sheetName)

	// let's make the excel header first
	xlsx.SetCellValue(sheetName, "A1", "Name")
	xlsx.SetCellValue(sheetName, "B1", "Gender")
	xlsx.SetCellValue(sheetName, "C1", "Age")

	err := xlsx.AutoFilter(sheetName, "A1", "C1", "")
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}

	// let's fill in the data
	for indeks, data := range DummyData {
		xlsx.SetCellValue(sheetName, fmt.Sprintf("A%d", indeks+2), data.Name)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("B%d", indeks+2), data.Gender)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("C%d", indeks+2), data.Age)
	}
}

func CreateBasicSheetWithMergedCell(xlsx *excelize.File) {
	var (
		sheetName = "student_data_style"
	)

	xlsx.NewSheet(sheetName)

	xlsx.SetCellValue(sheetName, "A1", "Student")
	xlsx.MergeCell(sheetName, "A1", "C1")

	// set style
	style, err := xlsx.NewStyle(`{
			"font": {
				"bold": true,
				"size": 36
			},
			"fill": {
				"type": "pattern",
				"color": ["#E0EBF5"],
				"pattern": 1
			}
		}`)
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}

	xlsx.SetCellStyle(sheetName, "A1", "C1", style)

	xlsx.SetCellValue(sheetName, "A2", "Name")
	xlsx.SetCellValue(sheetName, "B2", "Gender")
	xlsx.SetCellValue(sheetName, "C2", "Age")
	for indeks, data := range DummyData {
		xlsx.SetCellValue(sheetName, fmt.Sprintf("A%d", indeks+3), data.Name)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("B%d", indeks+3), data.Gender)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("C%d", indeks+3), data.Age)
	}
}

func ReadingExcel() {
	var (
		sheetName = "student_data"
		listData  []Student
	)

	xlsx, err := excelize.OpenFile("./assets/excel/07-11-2024.xlsx")
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}

	for i := 2; i <= 5; i++ {
		student := Student{
			Name:   xlsx.GetCellValue(sheetName, fmt.Sprintf("A%d", i)),
			Gender: xlsx.GetCellValue(sheetName, fmt.Sprintf("B%d", i)),
		}

		student.Age, err = strconv.Atoi(xlsx.GetCellValue(sheetName, fmt.Sprintf("C%d", i)))
		if err != nil {
			log.Fatal(err.Error())
		}

		listData = append(listData, student)
	}

	fmt.Println(listData)
}
