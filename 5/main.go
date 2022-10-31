package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct{
	UserName string
	Age int
	FavouriteNumber float64
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var user User

	user.UserName = readString("What is your name?")
	user.Age = readInt("What is your age?")
	user.FavouriteNumber = readFloat("What is your favourite number?")


	fmt.Printf("Your name is %s. You are %d years old. And your favourite number is %.2f", 
	user.UserName, user.Age, user.FavouriteNumber)
}

func prompt() {
	fmt.Print("-> ")
}

func readString(s string) string{
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "" {
			fmt.Println("Please enter a value")
		} else {
			return userInput
		}
	}
}

func readInt(s string) int {
	for {
		i := readString(s)

		num, err := strconv.Atoi(i)
		if err != nil {
			fmt.Println("Please enter a whole number")
		} else {
			return num
		}
	}
}

func readFloat(s string) float64 {
	for {
		i:= readString(s)
		num, err := strconv.ParseFloat(i, 64)
		if err != nil {
			fmt.Println("Please enter a number")
		} else {
			return num
		}
	}
}