package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type ResponseCreateInvoice struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Number  string `json:"number"`
}

func CreateInvoice(data []map[string]string, merged bool) (bool, error) {
	positions := []map[string]interface{}{}

	// Populate the positions slice from your data
	for _, item := range data {
		newItem := map[string]interface{}{
			"name":             item["Product"],
			"tax":              0,
			"total_price_gross": item["Price"],
			"quantity":         item["Quantity"],
		}
		positions = append(positions, newItem)
	}
	var dateTimeInvoice time.Time

	client := &http.Client{}

	dateTime, err := time.Parse("2006-01-02 15:04:05", data[0]["Sale Date"])
	if err != nil {
		dateTime, err = time.Parse("2006-01-02", data[0]["Sale Date"])
		if err != nil {
			return false, err
		}
	}
	dateOnly := dateTime.Format("2006-01-02")

	if data[0]["Invoice Number"] == "auto" || data[0]["Invoice Number"] == "" {
		data[0]["Invoice Number"] = ""
	}

	if data[0]["Invoice Date"] == "today" || data[0]["Invoice Date"] == "" {
		data[0]["Invoice Date"] = ""
	} else {
		dateTimeInvoice, err = time.Parse("2006-01-02 15:04:05", data[0]["Invoice Date"])
		if err != nil {
			dateTimeInvoice, err = time.Parse("2006-01-02", data[0]["Invoice Date"])
			if err != nil {
				return false, err
			}
		}
		data[0]["Invoice Date"] = dateTimeInvoice.Format("2006-01-02")
	}



	invoiceData := map[string]interface{}{
		"api_token": API_KEY,
		"invoice": map[string]interface{}{
			"kind":              data[0]["Invoice Type"],
			"number":            data[0]["Invoice Number"], 
			"status":            "paid",
			"currency":          data[0]["Currency"],
			"exchange_currency": "PLN",
			"sell_date":         dateOnly,
			"issue_date":        data[0]["Invoice Date"],
			"place":             City,
			"payment_type":      "transfer",
			"payment_to_kind":   "off",
			"client_id":         ClientId,
			"description":       data[0]["Additional Notes"],
			"buyer_override":    true,
			"buyer_tax_no":      data[0]["VATID"],
			"positions":         positions, // Add the positions slice here
		},
	}
	invoiceDataJSON, err := json.Marshal(invoiceData)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return false, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("https://%s.fakturownia.pl/invoices.json", Domain), strings.NewReader(string(invoiceDataJSON)))
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	var responseData ResponseCreateInvoice
	err = json.Unmarshal(bodyText, &responseData)
	if err != nil {
		err := errors.New(string(string(bodyText)))
		return false, err
	}

	if resp.StatusCode == 201 {
		fmt.Println(responseData)
		if responseData.Number != "" {
			err := errors.New(string(responseData.Number))
			return true, err
		} else {
			err := errors.New(string(responseData.Message))
			return true, err
		}
	} else {

		err := errors.New(string(responseData.Message))
		return false, err
	}
}
