package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const numberOfMonitoring = 2
const delayOfMonitoring = 5

func main(){
	introduction()

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

	websites := []string{
		"https://www.google.com/",
		"https://www.microsoft.com/",
		"https://vercel.com",
		"https://github.com/"}


	for i :=0; i < numberOfMonitoring; i++{
		for i, website:= range	websites {
			fmt.Println("Testing website", i, ":", website)
			verifyWebSite(website)	
		}
		time.Sleep(delayOfMonitoring * time.Second)
	}
}


func verifyWebSite(website string){
	response, _ := http.Get(website)
	// fmt.Println(response)
	if response.StatusCode == 200 {
		fmt.Println("Website" , website, "loaded successfully")
	}else{
		fmt.Println("Website", website, "having problems", response.StatusCode)
	}
}


