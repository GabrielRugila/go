package main

import (
	"myapp/packageone"
)

	var myVar = "My variable"

func main() {
	blockVar := "Block variable"

	packageone.PrintMe(myVar, blockVar)
}

