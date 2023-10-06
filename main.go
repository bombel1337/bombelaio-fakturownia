package main

import (
	"bombelaio-fakturownia/utils"
	"bufio"
	"fmt"
	// "log"
	"os"
	"syscall"
	"unsafe"

	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
)

const (
	asciiLogo string = `___  ____  __  ______  ______     
/ _ )/ __ \/  |/  / _ )/ __/ /     
/ _  / /_/ / /|_/ / _  / _// /__    
/____/\____/_/  /_/____/___/____/`

	STDOUT = -11
)

var (
	modkernel32              = syscall.NewLazyDLL("kernel32.dll")
	procSetConsoleTitleW     = modkernel32.NewProc("SetConsoleTitleW")
	invoicesCreated      int = 0
	productsInInvoices int = 0
) 

func setConsoleTitle(title string) {
	utf16Title := syscall.StringToUTF16(title)
	procSetConsoleTitleW.Call(uintptr(unsafe.Pointer(&utf16Title[0])))
}

func main() {

	utils.Logger = logrus.New()
	utils.Logger.Formatter = &utils.CustomFormatter{}
	utils.Logger.SetOutput(colorable.NewColorableStdout())
	utils.Log(utils.Logger, logrus.WarnLevel, fmt.Sprintf("\n\n %s \n\n", asciiLogo))

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

	if len(dataCSV) == 0 {
		utils.Log(utils.Logger, logrus.ErrorLevel, fmt.Sprintf("Invoices.csv file is empty."))

	} else {
		var mergedValues []map[string]string
		var productsInLastInvoice int = 0;
		for index, value := range dataCSV {
			indexInvoice := fmt.Sprintf("%03d", index + 1)
			
			if (index == len(dataCSV)-1  || dataCSV[index+1]["Merged Invoice"] == "no" || dataCSV[index+1]["Invoice Number"] != "") && value["Merged Invoice"] == "yes" {
				productsInInvoices++;
				productsInLastInvoice++;
				mergedValues = append(mergedValues, value)

				created, err := utils.CreateInvoice(mergedValues, true)
				if(created) {
					invoicesCreated++;
					setConsoleTitle(fmt.Sprintf("bombel invoice maker, Invoices created %x out of %v. Total products: %v", invoicesCreated, len(dataCSV), productsInInvoices))
					utils.Log(utils.Logger, logrus.InfoLevel, fmt.Sprintf("[%v] Invoice number: %v, has been created, total products in invoice: %v!", indexInvoice, err, productsInLastInvoice))
				} else {
					utils.Log(utils.Logger, logrus.ErrorLevel, fmt.Sprintf("[%v] Error making invoice: %v", indexInvoice, err))
				}
				mergedValues = nil
				productsInLastInvoice = 0;
			} else if value["Merged Invoice"] == "yes" {
				productsInLastInvoice++;
				productsInInvoices++;
				mergedValues = append(mergedValues, value)
			} else if (value["Merged Invoice"] == "no") {
				productsInInvoices++;
				mergedValues = append(mergedValues, value)
				created, err := utils.CreateInvoice(mergedValues, false)
				if(created){
					invoicesCreated++;
					setConsoleTitle(fmt.Sprintf("bombel invoice maker, Invoices created %x out of %v. Total products: %v", invoicesCreated, len(dataCSV), productsInInvoices))
					utils.Log(utils.Logger, logrus.InfoLevel, fmt.Sprintf("[%v] Invoice number: %v, has been created!", indexInvoice, err))
				} else {
					utils.Log(utils.Logger, logrus.ErrorLevel, fmt.Sprintf("[%v] Error making invoice: %v", indexInvoice, err))

				}
				mergedValues = nil
			}


		}

	}
	utils.Log(utils.Logger, logrus.WarnLevel, fmt.Sprintf("Press Enter to close!"))
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}
