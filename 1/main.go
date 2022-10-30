package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)


func main() {
	reader := bufio.NewReader(os.Stdin)

	words := doctor.Intro()
	fmt.Println(words)

	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')

		
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		
		
		if userInput == "quit" {
			break
		} else {
			fmt.Println(doctor.Response(userInput))
		}

	}
}