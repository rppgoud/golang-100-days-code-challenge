package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Enter the Pizza rating from 1 to 5:"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the  value")

	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("error occred", err)
	} else {
		fmt.Printf("rating: %T", numRating)
	}

}
