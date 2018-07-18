package main

import (
	"fmt"
	"math/rand"
	"time"
)

func createProcess(process [][]int){

	printData(process,4)

	for i:= 0; i < len(process); i++{
		process[i][0] = rand.Intn(9) + 1
		process[i][1] = rand.Intn(9) + 1

		for j := 0; j < 3; j++ {
			printData(process, (j) % 3)
		}
	}
	
	printData(process,3)
	
}

func getProcessList(process [][]int){
	

	fmt.Println("Process\tArrival\tBurst")
	for i := 0; i < len(process); i++{
		
		if process[i][0] == 0 {
			fmt.Printf("P%d: Not created\n", i+1)
		} else{

			fmt.Printf("P%d:\t%4d\t%3d\n", i+1, process[i][0], process[i][1])	
		}
	}

	fmt.Println()
}


func getStatus(status int){
	
	fmt.Printf("Status: ")

	switch status{
	case 0:
		fmt.Println("Generating processes.")
	case 1:
		fmt.Println("Generating processes..")
	case 2:
		fmt.Println("Generating processes...")
	case 3:
		fmt.Println("Processes generated!")
	case 4:
		fmt.Println("Starting Algorthm")
	default:
	}

	fmt.Println()

}

func getQueue(queue []int){
	
}

func printData(process [][]int, status int){

	for i := 0; i < 35; i++{
		fmt.Printf("-")
	}
	fmt.Println()
	getStatus(status)
	getProcessList(process)
	getQueue(queue)
	time.Sleep(time.Second)
}

func fcfs(process [][]int){
	
	queue := []int{}
	var cur int = -1

	fmt.Println(process)
	limit := len(process)

	for i := 0; i < 100; i++{
		for j := 0; j < len(process); j++{
			
			if process[j][0] == i {
				fmt.Printf("Time %d : %d arrived\n", i, j)
				queue = append(queue, j)
			}
		}
		if len(queue) > 0 {
			fmt.Println(queue)
		}
		if cur > -1  && process[cur][1] > 0 {
			process[cur][1]--

			if process[cur][1] == 0 { 
				cur = -1	
				limit--
			}
		}

		if len(queue) > 0 && cur == -1 {
			cur = queue[0]
			queue = queue[1:]
		}
		
		if limit == 0{
			fmt.Println(i)
			break
		}

	
		
	}
	
}


func main() {

	process := make([][]int, 5)
	for i := range process{
		process[i] = make([]int, 2)
	}	
	
	rand.Seed(time.Now().Unix())
	

//	printStatus(1)

	createProcess(process)
	
//	fcfs(process)


}
