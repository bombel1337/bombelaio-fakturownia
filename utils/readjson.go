package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ConstantConfigData struct {
	Domain   string `json:"domain"`
	ClientId int    `json:"client_id"`
	API_KEY  string `json:"api_key"`
	City     string `json:"email"`
}

var (
	Domain   string
	ClientId int
	API_KEY  string
	City     string
)

func ReadJSON(filepath string) error {
	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return err
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return err
	}

	var dataJSON ConstantConfigData
	err = json.Unmarshal(content, &dataJSON)
	if err != nil {
		fmt.Println("Error parsing JSON data:", err)
		return err
	}
	Domain = dataJSON.Domain
	ClientId = dataJSON.ClientId
	City = dataJSON.City
	API_KEY = dataJSON.API_KEY

	return err
}
