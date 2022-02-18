package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to the  variable input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating for the Pizza")

	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating ", input)
	fmt.Printf("thanks for rating: %v", input)

}
