// Curtis Helsel - 2018
//
// Commandline visualization of the first come first scheduling algorithm


package main

import (
	"fmt"
	"time"
	"os"
	"os/exec"
)

const limit int = 15

type fcfsStruct struct{
	curProc int
	process [][]int
	queue []int	
	time int
	timeLine []int
}

// Function populates the arrival time and burst cycle for the processes
func createProcess(fcfs *fcfsStruct){

	fcfs.process = make([][]int, 5)

	for i := range fcfs.process{
		fcfs.process[i] = make([]int, 3)
	}	

	fcfs.process = [][]int{
		{1, 5, 5}, 
		{2, 3, 3}, 
		{4, 4, 4}, 
		{5, 1, 1}, 
		{7, 2, 2},
	}	
}
// Prints process list
func getProcessList(process [][]int){

	fmt.Println("Process\tArrival\tBurst")

	for i := 0; i < len(process); i++{
		fmt.Printf("P%d:\t%4d\t%3d\n", i+1, process[i][0], process[i][1])
	}

	fmt.Println()
}

// Prints status of the algorithm
func getStatus(status int, process int){
	
	fmt.Println()
	fmt.Printf("Status: ")

	switch status{
	case 1:
		fmt.Printf("P%d has arrived, adding to queue.\n", process + 1)
	case 2:
		fmt.Printf("P%d has finished. Selecting next process in queue.\n", process + 1)
	case 3:
		fmt.Printf("P%d has finished.\n", process + 1) 
	case 4:
		fmt.Printf("P%d is being worked.\n", process + 1)
	case 5:
		fmt.Println("Visualization complete.")
	default:
	}
}

// Prints the contents of the ready queue
func getQueue(queue  []int){

	if len(queue) == 0 {
		fmt.Println()
		fmt.Println("Ready Queue: Currently empty.")
		fmt.Println()
	}else{

		fmt.Printf("             -") 
		for i:= 0; i < len(queue); i++{
			fmt.Printf("-----")
		}
		fmt.Printf("\nReady Queue: |")
		for i := 0; i < len(queue); i++{
			fmt.Printf(" P%d |", queue[i] + 1)
		}

		fmt.Println()
		fmt.Printf("             -") 
		for i:= 0; i < len(queue); i++{
			fmt.Printf("-----")
		}
		fmt.Println()
	}
	
	fmt.Println()
}

// Prints the timeline of the algorithm
func getTime(time int, timeline []int){
	
	fmt.Println("Timeline:")

	for i := 0; i < time; i++ {

		fmt.Printf("      ")
	}

	if time < limit {
		fmt.Println("   *")
	} else{
		fmt.Println()
	}	

	for i := 0; i < limit; i++ {
		fmt.Printf("------")
	}

	fmt.Println("-")

	for i := 0; i < limit; i++ {
	
		fmt.Printf("|")
		fmt.Printf("%3d  ", (i + 1) % 10)
	}	

	fmt.Println("|")

	for i := 0; i < limit; i++ {
		fmt.Printf("------")
	}

	fmt.Println("-")

	fmt.Printf("|")

	for i := 0; i < len(timeline); i++ {
	
		fmt.Printf("%3d  |", timeline[i])
	}	

	for i := 0; i < (limit - len(timeline)); i++ {
		fmt.Printf("     |")
	}

	fmt.Println()

	for i := 0; i < limit; i++ {
		fmt.Printf("------")
	}

	fmt.Println("-")
}

// Prints current screen in algorithm 
func printData(fcfs *fcfsStruct, status int){

	// clear ther terminal
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	clear.Run()

	fmt.Println()

	fmt.Println("First Come First Served Algorithm")
	fmt.Println()
	
	getProcessList(fcfs.process)
	getQueue(fcfs.queue)
	getTime(fcfs.time, fcfs.timeLine)
	getStatus(status, fcfs.curProc)
	
	time.Sleep(time.Second * 2)
}

func dispatcher(fcfs *fcfsStruct){
	
	var cur int = -1

	for i := 1; i <= limit; i++{

		// Checks for arrival of processes
		for j := 0; j < len(fcfs.process); j++{
			
			if fcfs.process[j][0] == i {
				temp := fcfs.curProc
				fcfs.curProc = j
				printData(fcfs, 1)
				fcfs.queue = append(fcfs.queue, fcfs.curProc)
				fcfs.curProc = temp
			}
		}

		// Checks for process finishing
		if cur > -1  && fcfs.process[cur][2] > 0 {

			fcfs.process[cur][2]--

			if fcfs.process[cur][2] == 0 { 
				cur = -1	

				if len(fcfs.queue) > 0 { 
					printData(fcfs,2)
				} else{
					printData(fcfs,3)
				}
			}
		}

		// Select process
		if len(fcfs.queue) > 0 && cur == -1 {
			cur = fcfs.queue[0]
			fcfs.curProc = cur
			fcfs.queue = fcfs.queue[1:]
		}

		// Adds current process to timeline
		if cur > -1 {
			fcfs.timeLine = append(fcfs.timeLine, cur + 1)
		} 
		
		printData(fcfs,4)

		fcfs.time = i 
	}

	printData(fcfs,5)
}


func main() {

	fcfs := fcfsStruct{}

	createProcess(&fcfs)

	dispatcher(&fcfs)

}
