package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Pizza App")

	fmt.Println("Please rate our Pizza between 1 and 5")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the rating, ", input)

	// Add 1 to your rating
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println("Error is ", err)
	} else {
		fmt.Println("Added 1 to your rating: ", (numRating + 1))
	}

}
