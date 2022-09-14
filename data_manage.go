package main

import (
	"strings"

	"github.com/x-sanya/trains_problem/train"
	"github.com/xuri/excelize/v2"
)

func UploadData(fileName, sheetName string) ([]train.Train, error) {
	xlsx, err := excelize.OpenFile(fileName)
	if err != nil {
		return nil, err
	}
	defer xlsx.Close()

	rows, err := xlsx.GetRows(sheetName)

	if err != nil {
		return nil, err
	}

	trains := make([]train.Train, len(rows))

	for i, row := range rows {
		trains[i], err = train.NewTrain(strings.Split(row[0], ";"))
		if err != nil {
			return nil, err
		}
	}

	return trains, nil
}
