package sio2

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func ReadExcel(fileName string) (data [][]string, err error) {
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Get value from cell by given worksheet name and cell reference.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return rows, nil
}
