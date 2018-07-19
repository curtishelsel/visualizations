package main

import (
	"fmt"
	"time"
)

type fcfsStruct struct{
	curProc int
	process [][]int
	queue []int	
	time int
	timeLine []int
}

func createProcess(fcfs *fcfsStruct){

	fcfs.process = make([][]int, 5)
	for i := range fcfs.process{
		fcfs.process[i] = make([]int, 3)
	}	

	fcfs.process = [][]int{
		{1, 5, 5}, 
		{2, 2, 2}, 
		{4, 4, 4}, 
		{5, 6, 6}, 
		{7, 2, 2},
	}	
		
}

func getProcessList(process [][]int){
	

	fmt.Println("Process\tArrival\tBurst Remaining")
	for i := 0; i < len(process); i++{
		
		if process[i][0] == 0 {
			fmt.Printf("P%d: Not created\n", i+1)
		} else{
			fmt.Printf("P%d:\t%4d\t%3d\t%3d\n", i+1, process[i][0], process[i][1], process[i][2])	
		}
	}

	fmt.Println()
}


func getStatus(status int, process int){
	
	fmt.Printf("Status: ")

	switch status{
	case 1:
		fmt.Printf("P%d has arrived, adding to queue.\n", process + 1)
	case 2:
		fmt.Printf("P%d has finished. Selecting next in queue.\n", process + 1)
	case 3:
		if process + 1 == 5 {
			fmt.Printf("P%d has finished.\n", process + 1) 
		} else {
			fmt.Printf("P%d has finished. Waiting for additional processes.\n", process + 1) 
		}
	case 4:
		fmt.Printf("P%d is being worked.\n", process + 1)
	case 5:
		fmt.Println("Visualization complete.")
	default:

	}

	fmt.Println()

}

func getQueue(queue  []int){
	if len(queue) == 0 {
		fmt.Println()
		fmt.Println("Queue: Currently empty")
		fmt.Println()
	}else{

		fmt.Printf("       -") 
		for i:= 0; i < len(queue); i++{
			fmt.Printf("-----")
		}
		fmt.Printf("\nQueue: |")
		for i := 0; i < len(queue); i++{
			fmt.Printf(" P%d |", queue[i] + 1)
		}

		fmt.Println()
		fmt.Printf("       -") 
		for i:= 0; i < len(queue); i++{
			fmt.Printf("-----")
		}
		fmt.Println()
	}
	
}

func getTime(time int){
	
	if time == 0 {
		fmt.Println("Time: Algorithm has not started")	
	} else{
		fmt.Printf("Time: %d\n", time)
	}
}

func getTimeLine(timeline []int){
	
	fmt.Println(timeline)
}

func printData(fcfs *fcfsStruct, status int){

	for i := 0; i < 60; i++{
		fmt.Printf("-")
	}
	fmt.Println()
	fmt.Println("First Come First Served Algorithm")
	fmt.Println()
	
	getProcessList(fcfs.process)

	getQueue(fcfs.queue)
	getTime(fcfs.time)
	getTimeLine(fcfs.timeLine)
	getStatus(status, fcfs.curProc)
	
	
	time.Sleep(time.Second * 2)
}

func fcfsAlgo(fcfs *fcfsStruct){
	
	var cur int = -1

	printData(fcfs, 5)

	limit := len(fcfs.process)

	for i := 1; i < 100; i++{
		fcfs.time = i 
		for j := 0; j < len(fcfs.process); j++{
			
			if fcfs.process[j][0] == i {
				temp := fcfs.curProc
				fcfs.curProc = j
				printData(fcfs, 1)
				fcfs.queue = append(fcfs.queue, fcfs.curProc)
				printData(fcfs, 1)
				fcfs.curProc = temp
			}
		}

		if cur > -1  && fcfs.process[cur][2] > 0 {
			fcfs.process[cur][2]--

			if fcfs.process[cur][2] == 0 { 
				cur = -1	
				limit--
				if len(fcfs.queue) > 0 { 
					printData(fcfs,2)
				} else{
					printData(fcfs,3)
				}
			}
		}

		if len(fcfs.queue) > 0 && cur == -1 {
			cur = fcfs.queue[0]
			fcfs.curProc = cur
			fcfs.queue = fcfs.queue[1:]
		}

		if cur > -1 {
			fcfs.timeLine = append(fcfs.timeLine, cur + 1)
		} 
		
		if limit == 0 {
			break
		}
		
		if cur > -1 {
			printData(fcfs,4)
		}

	}

	printData(fcfs,5)
	
}


func main() {

	fcfs := fcfsStruct{}
	createProcess(&fcfs)

	
	fcfsAlgo(&fcfs)


}
