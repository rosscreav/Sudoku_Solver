package main

import (
	"fmt"
	io "example.com/sudoku/fileio"
	"time"
	"os"
)

var row int
var col int

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0{
		argsWithoutProg = append(argsWithoutProg, "fileio/TestCases/puzzles0_kaggle")
	}	
	for _,filename := range argsWithoutProg{
	    var testCases = io.Read_file(filename)

	    var solvedtimes []time.Duration
	    var unsolvedtimes []time.Duration
	    Totalstart := time.Now()

	    for _,testcase := range testCases{
	    	start := time.Now()
	    	if solve(testcase){
	    		solvedtimes = append(solvedtimes,time.Since(start))
	    	} else{
	    		unsolvedtimes = append(unsolvedtimes,time.Since(start))
	   		}
	    }
	    elapsed := time.Since(Totalstart)
	    fmt.Printf("Testing %s",filename)
		fmt.Printf("\nTested %d cases in %s\n",len(solvedtimes),elapsed)
	    fmt.Printf("Solved: %d\nUnsolved: %d\n",len(solvedtimes),len(unsolvedtimes))
	    //Solved Statistics
	    fmt.Printf("\nSolved Case statistics\n")
	    var min,max,mean = statistics(solvedtimes)
	    fmt.Printf("Minimum Time: %s\nMaximum Time: %s\nMean Time (μ): %s\n",min,max,mean)
	    //Unsolved Statistics
	    fmt.Printf("\nUnsolved Case statistics\n")
	    min,max,mean = statistics(unsolvedtimes)
	    fmt.Printf("Minimum Time: %s\nMaximum Time: %s\nMean Time (μ): %s\n",min,max,mean)
	    fmt.Printf("\n")
   	}
}

func solve(input [9][9]int) bool{
	row = 0
	col = 0
	var current [2]int
	//If there is no empty spaces return with exit code true
	if !find_next_empty(input){
		//print(input)
		return true
	}
	//Start checking numbers
	for num:=1; num<10; num++{
		//fmt.Printf("Checking %d at [%d,%d]\n",num,row,col)
		if check_location_safety(input, num){
			//fmt.Println("safe num")
			input[row][col] = num
			//Store the current position to avoid the recursion modifying variables
			current[0] = row
			current[1] = col
			//Recursive check return
			if solve(input){
				return true	
			}
			//fmt.Println("failed")
			//If it has failed
			row = current[0]
			col = current[1]
			input[row][col] = 0
		}
	}
	
	//Backtrack or return unsolvable
	return false

}

//Find next empty location
func find_next_empty(array [9][9]int) bool {
	for i := 0; i < 9; i++{
		for j := 0; j < 9; j++{
			if array[i][j] == 0 {
				row = i
				col = j 
				return true
			}
		}
	}
	return false
}

func check_location_safety(array [9][9]int, num int) bool{
	for i := 0; i<9; i++{
		if array[row][i] == num || array[i][col] == num{
			return false
		}
	}
	var box_start_row = row - row%3
	var box_start_col = col - col%3
	for i := 0; i < 3; i++{
		for j := 0; j < 3; j++{
			if array[i+box_start_row][box_start_col+j] == num {
				return false
			}
		}
	}
	return true
}




//Prints out the array of numbers with formatting
func print(array [9][9]int){
	fmt.Printf("\n")
	for index, element := range array {
    	for subindex, subelement := range element {
    		if subindex == 3 || subindex == 6{
    			fmt.Print(" | ")
    		}
    		fmt.Printf("%d ",subelement)
    	}
    	fmt.Printf("\n")
    	if index == 2 || index == 5{
    		fmt.Print("------  ------   ------\n")
    	}
    	
	}
	fmt.Printf("\n")
}

//Gets min max mean of a slice
func statistics(array []time.Duration)(time.Duration,time.Duration,time.Duration){
	//If the array is empty return null
	if len(array) == 0{
		return time.Duration(0),time.Duration(0),time.Duration(0)
	}
    var max time.Duration = array[0]
    var min time.Duration = array[0]
    var total time.Duration
    for _, value := range array {
    	total= total+value
        if max < value {
            max = value
        }
        if min > value {
            min = value
        }
    }
    return min, max, total/time.Duration(len(array))

}





