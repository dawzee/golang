package main

import (
   "github.com/360EntSecGroup-Skylar/excelize"
)

func main() {

	f := excelize.NewFile()
	f.SetCellValue("Sheet1", "A1", "passionfruit")
	if err := f.SaveAs("./Book1.xlsx"); err != nil {
		panic(err)
	}
}
