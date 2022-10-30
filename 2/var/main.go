package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const pressEnter = "press ENTER when ready"

func main() {
	reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().UnixNano())
	firstNum := rand.Intn(8) +2
	secondNum := rand.Intn(8) +2
	subtraction := rand.Intn(8) +2
	answer := firstNum * secondNum - subtraction

	//display intro
	fmt.Println("\tGuess The Number")
	fmt.Println("------------------\n")

	fmt.Println("Choose a Number between 1 and 10,", pressEnter)
	reader.ReadString('\n')

	// take through game
	fmt.Println("Multiply your number by", firstNum, pressEnter)
	reader.ReadString('\n')

	fmt.Println("Now multiply by", secondNum, pressEnter)
	reader.ReadString('\n')

	fmt.Println("Now divide the result by the number you originally thought of", pressEnter)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, pressEnter)
	reader.ReadString('\n')

	// give answer
	fmt.Println("The answer is", answer)
}
