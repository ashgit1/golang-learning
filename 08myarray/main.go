package main

import "fmt"

func main() {

	fmt.Println("Welcome to the class on arrays")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[2] = "Orange"
	fmt.Println("List of Fruits: ", fruitList)
	fmt.Println("Length of Fruitlist: ", len(fruitList))

	var vegList = [3]string{"Potato", "Tomato", "Mushroom"}
	fmt.Println("List of Veges: ", vegList)
	fmt.Println("Length of Veges: ", len(vegList))

}
