package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("hello welcome to the slices")

	var FruitsList = []string{"app", "banan", "guva"}
	fmt.Println("FrutisList", FruitsList)

	fmt.Println("FruitsListadded", append(FruitsList, "pavi"))
	// fmt.Println("FruitsList", FruitsList)
	FruitsList = append(FruitsList[1:3])
	fmt.Println(FruitsList)

	numberList := make([]int, 4)
	numberList[0] = 952
	numberList[1] = 261
	numberList[2] = 821
	numberList[3] = 122
	//numberList[4] = 952

	fmt.Println("the values of number are", numberList)

	fmt.Println("the values of number are", append(numberList, 232, 992, 421))
	fmt.Println("the sorted values", numberList)

	sort.Ints(numberList)

	fmt.Println("hello", numberList)

}
