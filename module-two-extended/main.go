package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Enter the Pizza rating from 1 to 5:"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println(reader)

	// input, _ := reader.ReadString('\n')
	// fmt.Println("thanks for rating")
	// fmt.Printf(input)

	// numRating, err := strconv.ParseFloat(input, 64)

	// if err != nil {
	// 	fmt.Println("error occred", err)
	// } else {
	// 	fmt.Println("rating:", numRating)
	// }

}
