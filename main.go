package main

import (
	"bombelaio-fakturownia/utils"
	"fmt"
	"time"
	"github.com/sirupsen/logrus"
)


func main() {
	utils.Logger = logrus.New()
	utils.Logger.Formatter = &utils.CustomFormatter{}

	filepath := "example.csv"
    data, err := utils.ReadCSV(filepath)
    if err != nil {
		utils.Log(utils.Logger, logrus.ErrorLevel, fmt.Sprintf("Error reading csv: %v", err))
        return
    }

	done := make(chan bool)


    for  index, value := range data {

		go func(i int, value map[string]string) {
			indexInvoice := fmt.Sprintf("%03d", i)
			
			createdInvoice, err := utils.CreateInvoice(value);
			if createdInvoice {
				utils.Log(utils.Logger, logrus.InfoLevel, fmt.Sprintf("[%v] Invoice number: %v, has been created!", indexInvoice, err) )
			} else {
				utils.Log(utils.Logger, logrus.ErrorLevel, fmt.Sprintf("[%v] Error making invoice: %v", indexInvoice, err))
			}
            done <- true
        }(index, value)
		time.Sleep(1 * time.Second)

    }

	for i := 0; i < len(data); i++ {
        <-done
    }

}