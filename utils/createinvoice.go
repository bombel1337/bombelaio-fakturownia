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



func CreateInvoice(data map[string]string) (bool, error) {

	client := &http.Client{}

	dateTime, err := time.Parse("2006-01-02 15:04:05", data["Sale Date"])
	if err != nil {
		dateTime, err = time.Parse("2006-01-02", data["Sale Date"])
		if err != nil {
			return false, err
		}
	}
	dateOnly := dateTime.Format("2006-01-02")
	
	if data["Invoice Number"] == "auto" || data["Invoice Number"] == "" {
		data["Invoice Number"] = "null";
	} else {
		data["Invoice Number"] = `"` + data["Invoice Number"] + `"`
	}

	var invoiceData = fmt.Sprintf(`{
        "api_token": "%v",
        "invoice": {
			"kind":"vat",
			"number": %s,
			"status":"paid",
			"currency":  "EUR",
			"exchange_currency": "PLN",
			"sell_date": "%s",
			"place" : "%v",
			"payment_type" : "transfer",
			"payment_to_kind": "off",
			"client_id": %v,
			"description":"%v",
			"buyer_override": true,
			"buyer_tax_no": %q,
            "positions":[
                {"name":"%s", "tax":0, "total_price_gross":%v, "quantity":%v}
            ]
        }
    }`,API_KEY, data["Invoice Number"], dateOnly, City, ClientId, data["Additional Notes"], data["VATID"], data["Product"], data["Price"], data["Quantity"])


	req, err := http.NewRequest("POST", fmt.Sprintf("https://%s.fakturownia.pl/invoices.json", Domain), strings.NewReader(invoiceData))
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

	if (resp.StatusCode == 201){
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

