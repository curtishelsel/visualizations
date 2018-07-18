package main

import (
	"fmt"
	"math/rand"
	"time"
)

func createProcess(process [][]int){

	for i:= 0; i < len(process); i++{
		process[i][0] = rand.Intn(10 - 1) 
		process[i][1] = rand.Intn(10 - 1) + 1
		//printStatus((i+1) % 4)
	}
	
}

func printStatus(code int){
	
	switch code{
	case 1:
		fmt.Println("Generating processes.")
		time.Sleep(time.Second)
	case 2:
		fmt.Println("Generating processes..")
		time.Sleep(time.Second)
	case 3:
		fmt.Println("Generating processes...")
		time.Sleep(time.Second)
	default:
	}

}

func fcfs(process [][]int){
	
	queue := []int{}
	var cur int = -1

	fmt.Println(process)

	for i := 0; i < 100; i++{
		for j := 0; j < len(process); j++{
			
			if process[j][0] == i {
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
			}
		}

		if len(queue) > 0 && cur == -1 {
			cur = queue[0]
			queue = queue[1:]
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
	
	fcfs(process)


}
