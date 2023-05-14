package utils

import (
    "encoding/csv"
    "os"
)

func ReadCSV(filepath string) (map[int]map[string]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	results := make(map[int]map[string]string)

	headers := rows[0]

	for i, row := range rows[1:] {
		rowData := make(map[string]string)

		for j, column := range row {
			rowData[headers[j]] = column
		}

		results[i+1] = rowData
	}

	return results, nil
}