package main

import (
	"fmt"
	"math/rand"
	"time"
)

type fcfsStruct struct{
	curProc int
	process [][]int
	queue []int	
}

func createProcess(fcfs *fcfsStruct){

	fcfs.process = make([][]int, 5)
	for i := range fcfs.process{
		fcfs.process[i] = make([]int, 2)
	}	
	
	rand.Seed(time.Now().Unix())

	printData(fcfs, 4)

	for i:= 0; i < len(fcfs.process); i++{
		fcfs.process[i][0] = rand.Intn(9) + 1
		fcfs.process[i][1] = rand.Intn(9) + 1

		for j := 0; j < 3; j++ {
			printData(fcfs, (j) % 3)
		}
	}
	
	printData(fcfs,3)
	
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


func getStatus(status int, process int){
	
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
		fmt.Println("Starting Program.")
	case 5:
		fmt.Println("Starting Algorithm.")
	case 6:
		fmt.Printf("P%d has arrived, adding to queue.", process + 1)
		fmt.Println()
	default:
	}

	fmt.Println()

}

func getQueue(queue  []int){
	if len(queue) == 0 {
		fmt.Println("Queue: Currently empty")
	}else{
		fmt.Printf("Queue: ")
		for i := 0; i < len(queue); i++{
			fmt.Printf("%d ", queue[i] + 1)
		}
		fmt.Println()
	}
	
}

func printData(fcfs *fcfsStruct, status int){

	for i := 0; i < 35; i++{
		fmt.Printf("-")
	}
	fmt.Println()
	getStatus(status, fcfs.curProc)
	getProcessList(fcfs.process)
	getQueue(fcfs.queue)
	time.Sleep(time.Second)
}

func fcfsAlgo(fcfs *fcfsStruct){
	
	var cur int = -1

	printData(fcfs, 5)

	limit := len(fcfs.process)

	for i := 0; i < 100; i++{
		for j := 0; j < len(fcfs.process); j++{
			
			if fcfs.process[j][0] == i {
				fcfs.curProc = j
				printData(fcfs, 6)
				fcfs.queue = append(fcfs.queue, fcfs.curProc)
				printData(fcfs, 6)
			}
		}

		if cur > -1  && fcfs.process[cur][1] > 0 {
			fcfs.process[cur][1]--

			if fcfs.process[cur][1] == 0 { 
				cur = -1	
				limit--
			}
		}

		if len(fcfs.queue) > 0 && cur == -1 {
			cur = fcfs.queue[0]
			fcfs.queue = fcfs.queue[1:]
		}
		
		if limit == 0{
			break
		}

	}
	
}


func main() {

	fcfs := fcfsStruct{}

//	printStatus(1)

	createProcess(&fcfs)
	
	fcfsAlgo(&fcfs)


}
