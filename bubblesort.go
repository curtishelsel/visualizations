// Curtis Helsel - July 21, 2018
// Visualization of Bubble Sort

package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
	"os/exec"
)	

// Number of indices to sort
const sortNum int = 30

// Function draws the bubble sort in action
func draw(a []int, status int){

	// clears the terminal screen
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	clear.Run()

	fmt.Println("Bubble Sort Alogorithm")
	fmt.Printf("Current Array Order: ")
	fmt.Println(a)

	for j := sortNum; j > 0; j-- {
		for i := 0; i < sortNum; i++{
			if a[i] >= j {
				fmt.Printf("*")	
			}else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	// Status line for sort
	switch status{
	case 0:
		fmt.Println("Status: Sorting")
	case 1:
		fmt.Println("Status: Sorting.")
	case 2:
		fmt.Println("Status: Sorting..")
	case 3:
		fmt.Println("Status: Sorting...")
	default:
		fmt.Println("Visualization complete!")
	}	

	time.Sleep(time.Millisecond * 200)
}

// Creates a shuffled slice of integers 1-sortNum
func shuffle()([]int){

	a := make([]int, sortNum)
	
	for i := 0; i < sortNum; i++ {
		a[i] = i+1
	}
	
	rand.Seed(time.Now().UnixNano())	

	dest := make([]int, sortNum)
	perm := rand.Perm(sortNum)
	for i, v := range perm {
		dest[v] = a[i]
	}
	
	return dest
}

// Bubble Sort
func bubble(a []int) {

	for i := 0; i < sortNum - 1; i++{
		for j := 0; j < sortNum-i-1; j++ {
			if a[j] > a[j+1] {
				temp := a[j]
				a[j] = a[j+1]
				a[j+1] = temp
				draw(a, j % 4)
			}
		}
	}
	draw(a,4)
}

func main(){

	arr := shuffle()	

	bubble(arr)
}
