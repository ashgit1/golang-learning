package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("------------------------------------------------------------------")
	fmt.Println("Welcome to Class on Slices")
	var fruitList = []string{"Apple", "Banana", "Tomato"}
	fmt.Printf("Type of fruit is %T \n", fruitList)
	fmt.Println("Original fruit list is", fruitList)

	// Dynamically add more fruits to the Slice
	fruitList = append(fruitList, "Mango", "Peach")
	fmt.Println("New fruit list is", fruitList)

	// how we can get subset in slices
	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)
	fmt.Println("------------------------------------------------------------------")

	// dynamic allocation for integers slices example
	highScores := make([]int, 4)
	highScores[0] = 222
	highScores[1] = 777
	highScores[2] = 333
	highScores[3] = 444
	//highScores[4] = 555 // at this point we can't do this
	fmt.Println("Scores are ", highScores)

	// we can increase the dynamic allocation of memory
	highScores = append(highScores, 555, 111, 666)
	fmt.Println("New Scores are ", highScores)

	// some methods for slices
	fmt.Println("Is highscores sorted? : ", sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println("Sorted highscores : ", highScores)
	fmt.Println("Is highscores sorted? : ", sort.IntsAreSorted(highScores))
	fmt.Println("------------------------------------------------------------------")

	// how to remove a vaue from slices based on index [important]
	var courses = []string{"Java", "react", "swift", "javascript", "python"}
	fmt.Println("Initial course List: ", courses)
	// remove the course swift from this list
	var index int = 2
	courses = append(courses[0:index], courses[(index+1):]...)
	fmt.Println("New Course List: ", courses)
	fmt.Println("------------------------------------------------------------------")
}
