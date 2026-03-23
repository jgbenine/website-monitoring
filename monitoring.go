package main

import (
	"fmt"
	"os"
)


func main(){

	introduction()
	menuOptions()
	
	optionDefined := optionsSet()
	
	switch optionDefined {
	case 1:
		fmt.Println("Monitoring...")
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



