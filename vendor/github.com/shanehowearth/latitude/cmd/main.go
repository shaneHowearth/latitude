package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	latitude "github.com/shanehowearth/latitude/pkg"
)

func main() {
	fmt.Println("Latitude max sharetrading profit calculator")
	input := os.Args[1:]

	// allow comma separated or space separated input
	inputRegExp, _ := regexp.Compile(",| ")
	strVals := inputRegExp.Split(strings.Join(input, ","), -1)

	// Convert the input to ints
	vals := make([]int, len(strVals))
	for i := range vals {
		tmp, err := strconv.Atoi(strVals[i])
		if err != nil {
			log.Fatalf("Your input contained a non number")
		}
		vals[i] = tmp
	}

	// Nice wee output
	fmt.Printf("Calculating for: %v\n", vals)
	minIdx, maxIdx, profit, err := latitude.GetMaxProfit(vals)
	if err != nil {
		log.Fatalf("Unable to process request with error: %v", err)
	}
	fmt.Println("Buy at: ", vals[minIdx])
	fmt.Println("Sell at: ", vals[maxIdx])
	fmt.Println("Profit: ", profit)
}
