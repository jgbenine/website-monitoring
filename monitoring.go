package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const numberOfMonitoring = 2
const delayOfMonitoring = 10

func main(){
	introduction()
	registerLog("site-falso", false)

	for {
		menuOptions()
	
		optionDefined := optionsSet()
	
		switch optionDefined {
		case 1:
			initialMonitoring()
		case 2:
			fmt.Println("View logs...")
		
		case 0:
			fmt.Println("Exiting program...")
			os.Exit(0)
		default:
			fmt.Println("Unknown command option")
			os.Exit(-1)
		}
	}
}

func introduction(){
	version := 1.1
	fmt.Println("Hello.")
	fmt.Println("This program is in version", version)
}

func menuOptions(){
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - View logs")
	fmt.Println("0 - Exit program")
}

func optionsSet() int{
	var commandSet int
	fmt.Scan(&commandSet)
	fmt.Println("Chosen command was", commandSet)

	return commandSet
}


func initialMonitoring(){
	fmt.Println("Monitoring...")

	websites := getWebSitesFile()


	for i:= 0; i < numberOfMonitoring; i++{
		for i, website:= range	websites {
			fmt.Println("Testing website", i, ":", website)
			verifyWebSite(website)	
		}
		time.Sleep(delayOfMonitoring * time.Second)
	}
}


func verifyWebSite(website string){
	response, err := http.Get(website)

	if err != nil{
		fmt.Println("Error detected in verify website:", err)
	}

	if response.StatusCode == 200 {
		fmt.Println("Website" , website, "loaded successfully")
		registerLog(website, true)
	}else{
		fmt.Println("Website", website, "having problems", response.StatusCode)
		registerLog(website, false)
	}
}


func getWebSitesFile() []string{
	var sites []string

	file, err := os.Open("websites.txt")

	if err != nil {
		fmt.Println("Error detected in getWebSitesFile:" , err)
	}
	
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF{
			break
		}
		line = strings.TrimSpace(line)
		sites = append(sites, line)
	}	

	file.Close()
	return sites
}


func registerLog(website string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil{
		fmt.Println("Error when register log", err)
	}

	file.WriteString(website + "- online:" + strconv.FormatBool(status) + "\n")

	file.Close()

}

