package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("enter the time")

	today := time.Now()

	// fmt.Println(today)
	fmt.Println(today.Format("02-01-2006 11:30:40 Monday"))
}
