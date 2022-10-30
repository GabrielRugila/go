package main

import (
	"bufio"
	"fmt"
	"os"
	"math/rand"
	"time"
)

func main() {
	//seed random number generator
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Choose the number of sides for your dice")
	rollDice()

}

func rollDice(mxSides int) {
	var diceRoll = rand.Intn(mxSides) - 1 
}