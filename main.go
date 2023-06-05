package main

import (
	"bombelaio-fakturownia/utils"
	"fmt"
	"time"
	"bufio"
	"os"
	"syscall"
	"unsafe"
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
)

const (
	asciiLogo string =`___  ____  __  ______  ______     
/ _ )/ __ \/  |/  / _ )/ __/ /     
/ _  / /_/ / /|_/ / _  / _// /__    
/____/\____/_/  /_/____/___/____/`

	STDOUT = -11
)

var (
	modkernel32 = syscall.NewLazyDLL("kernel32.dll")
	procSetConsoleTitleW = modkernel32.NewProc("SetConsoleTitleW")
	invoicesCreated int = 0
)

func setConsoleTitle(title string) {
	// Call the Windows API function to set the console title
	utf16Title := syscall.StringToUTF16(title)
	procSetConsoleTitleW.Call(uintptr(unsafe.Pointer(&utf16Title[0])))
}


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


	if len(dataCSV) == 0 {
		utils.Log(utils.Logger, logrus.ErrorLevel, fmt.Sprintf("Invoices.csv file is empty."))

	} else {
		for  index, value := range dataCSV {

			// go func(i int, value map[string]string) {
				indexInvoice := fmt.Sprintf("%03d", index)
	
				createdInvoice, err := utils.CreateInvoice(value);
				if createdInvoice {
					invoicesCreated++;
					setConsoleTitle(fmt.Sprintf("bombel invoice maker, Invoices created %x out of %v.", invoicesCreated, len(dataCSV)))

					utils.Log(utils.Logger, logrus.InfoLevel, fmt.Sprintf("[%v] Invoice number: %v, has been created!", indexInvoice, err) )
				} else {
					utils.Log(utils.Logger, logrus.ErrorLevel, fmt.Sprintf("[%v] Error making invoice: %v", indexInvoice, err))
				}

				// done <- true
			// }(index, value)
			time.Sleep(1 * time.Second)
	
		}
	
		// for i := 0; i < len(dataCSV); i++ {
		// 	<-done
		// }

	}


	utils.Log(utils.Logger, logrus.WarnLevel, fmt.Sprintf("Press Enter to close!"))
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}