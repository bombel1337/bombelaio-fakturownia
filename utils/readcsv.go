package utils

import (
	"encoding/csv"
	"os"
)

func ReadCSV(filepath string) ([]map[string]string, error) {
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

	results := make([]map[string]string, 0)

	headers := rows[0]

	for _, row := range rows[1:] {
		rowData := make(map[string]string)

		for j, column := range row {
			rowData[headers[j]] = column
		}

		results = append(results, rowData)
	}

	return results, nil
}
