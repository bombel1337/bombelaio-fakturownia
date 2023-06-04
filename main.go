package main

import (
	"bombelaio-fakturownia/utils"
	"fmt"
	"time"
	"bufio"
	"os"
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
)
const asciiLogo string =`___  ____  __  ______  ______     
/ _ )/ __ \/  |/  / _ )/ __/ /     
/ _  / /_/ / /|_/ / _  / _// /__    
/____/\____/_/  /_/____/___/____/`


func main() {

	utils.Logger = logrus.New()
	utils.Logger.Formatter = &utils.CustomFormatter{}
    utils.Logger.SetOutput(colorable.NewColorableStdout())
	utils.Log(utils.Logger, logrus.WarnLevel,fmt.Sprintf("\n\n %s \n\n", asciiLogo))

	filepathJSON := "config.json"
	err := utils.ReadJSON(filepathJSON)
    if err != nil {
		utils.Log(utils.Logger, logrus.ErrorLevel, fmt.Sprintf("Error reading csv: %v", err))
        return
    }



	filepathCSV := "invoices.csv"
    dataCSV, err := utils.ReadCSV(filepathCSV)
    if err != nil {
		utils.Log(utils.Logger, logrus.ErrorLevel, fmt.Sprintf("Error reading csv: %v", err))
        return
    }

	done := make(chan bool)

	if len(dataCSV) == 0 {
		utils.Log(utils.Logger, logrus.ErrorLevel, fmt.Sprintf("Invoices.csv file is empty."))

	} else {
		for  index, value := range dataCSV {

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
	
		for i := 0; i < len(dataCSV); i++ {
			<-done
		}

	}


	utils.Log(utils.Logger, logrus.WarnLevel, fmt.Sprintf("Press Enter to close!"))
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}