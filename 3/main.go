package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()
	
	coffees := make(map[int]string)
	coffees[1] = "Cappuccino"
	coffees[2] = "Latte"
	coffees[3] = "Expresso"
	coffees[4] = "Depresso"
	coffees[5] = "Mocha"
	coffees[6] = "Machiato"

	fmt.Println("---------------------------------")
	fmt.Println("|\t\tMENU\t\t|")
	fmt.Println("---------------------------------")
	fmt.Println("|\t1 - Cappuccino\t\t|")
	fmt.Println("|\t2 - Latte\t\t|")
	fmt.Println("|\t3 - Expresso\t\t|")
	fmt.Println("|\t4 - Depresso\t\t|")
	fmt.Println("|\t5 - Mocha\t\t|")
	fmt.Println("|\t6 - Machiato\t\t|")
	fmt.Println("|\tQ - Quit\t\t|")
	fmt.Println("---------------------------------")


	for {
		char, _, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}
		
		if char == 'q' || char == 'Q'{
			break
		}

		i, _ := strconv.Atoi(string(char))
		fmt.Sprintf("You chose %s", coffees[i])

		if char == 'q' || char == 'Q'{
			break
		}

	}

	fmt.Println("Program Exited...")
}