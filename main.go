package main

//Imports
import (
	//Print outs
	"fmt"
	//Imported file to load test cases /fileio/fileio.go
	io "example.com/sudoku/fileio"
	//Used for getting time statisitics
	"time"
	//Used to handle cmd line arguments
	"os"
	//Progress bar for command line
	"github.com/cheggaaa/pb/v3"
)

//Holds the current row and column being looked at
var row int
var col int

//Main function called on startup
func main() {
	//Read the command line arguments ignoring the name of the file
	argsWithoutProg := os.Args[1:]
	//If there is no other arguments then default 
	if len(argsWithoutProg) == 0{
		//Add puzzles0_kaggle to the test set
		argsWithoutProg = append(argsWithoutProg, "fileio/TestCases/puzzles0_kaggle")
	}	

	//Loop over the test files
	for _,filename := range argsWithoutProg{
		//Read in the file into an slice of test cases
	    var testCases = io.Read_file(filename)
	    //Create a new progress bar (amount of test cases)
	    bar := pb.StartNew(len(testCases))
	    //Declare solvedtimes and unsolved times (used to get average min and max for results)
	    var solvedtimes []time.Duration
	    var unsolvedtimes []time.Duration
	    //Total time declared before loop start (all test cases in a file time)
	    Totalstart := time.Now()

	    //Loop over each case within the test set/file
	    for _,testcase := range testCases{
	    	//Starting time of the solve
	    	start := time.Now()
	    	//If the case is solved
	    	if solve(testcase){
	    		//Then record the time it took to solve in solvedtimes
	    		solvedtimes = append(solvedtimes,time.Since(start))
	    	} else{
	    		//If not solved then record the time in unsolved times
	    		unsolvedtimes = append(unsolvedtimes,time.Since(start))
	   		}
	   		//Increment the progress bar by one case
	   		bar.Increment()
	    }
	    //Finish the progress bar
	    bar.Finish()
	    //Record the time taken to finish all test cases
	    elapsed := time.Since(Totalstart)

	    //Statistic print out
	    //Basic details on the test cases
	    fmt.Printf("Testing %s",filename)
		fmt.Printf("\nTested %d cases in %s\n",len(solvedtimes),elapsed)
	    fmt.Printf("Solved: %d\nUnsolved: %d\n",len(solvedtimes),len(unsolvedtimes))
	    //Solved Statistics
	    fmt.Printf("\nSolved Case statistics\n")
	    //Get the max min and mean of the solved times
	    var min,max,mean = statistics(solvedtimes)
	    fmt.Printf("Minimum Time: %s\nMaximum Time: %s\nMean Time (μ): %s\n",min,max,mean)
	    //Unsolved Statistics
	    fmt.Printf("\nUnsolved Case statistics\n")
	    //Get the max min and mean of the unsolved times
	    min,max,mean = statistics(unsolvedtimes)
	    fmt.Printf("Minimum Time: %s\nMaximum Time: %s\nMean Time (μ): %s\n",min,max,mean)
	    fmt.Printf("\n")
   	}
}

//Recursive function to solve the puzzle
func solve(input [9][9]int) bool{
	//Default row and column
	row = 0
	col = 0
	//Variable used to store the corrindinates within each instance of the function
	var current [2]int
	//Checks for empty spaces and sets row and col pointer if there is one
	if !find_next_empty(input){
		//Solve is complete
		return true
	}
	//Start checking numbers from 1-9
	for num:=1; num<10; num++{
		//If the number can work in that location
		if check_location_safety(input, num){
			//Set the space to the safe number
			input[row][col] = num
			//Store the current position to avoid the recursion modifying variables
			current[0] = row
			current[1] = col
			//Recursively check the next position and so on till they all return
			if solve(input){
				//If it is fully solved return true
				return true	
			}
			//If the num fails then reset that position to zero
			row = current[0]
			col = current[1]
			input[row][col] = 0
			//Try a new number
		}
	}
	//Backtrack or return unsolvable
	return false
}

//Find next empty location (returns true if one is found)
func find_next_empty(array [9][9]int) bool {
	//Loop counting across the row
	for i := 0; i < 9; i++{
		//Loop counting across the column
		for j := 0; j < 9; j++{
			//If the poistion is 0
			if array[i][j] == 0 {
				//Set the row and col to that address
				row = i
				col = j 
				//Position found
				return true
			}
		}
	}
	//No position was found
	return false
}

//Check if a number is safe to put in a location
func check_location_safety(array [9][9]int, num int) bool{
	//Check if the number is used in the row or col
	for i := 0; i<9; i++{
		//If in row or in col
		if array[row][i] == num || array[i][col] == num{
			//Is used
			return false
		}
	}
	//Check if the number is used in the box
	//Returns the top left cordinate for the box
	var box_start_row = row - row%3
	var box_start_col = col - col%3
	//Loop through the row of the box
	for i := 0; i < 3; i++{
		//Loop through the column of the box
		for j := 0; j < 3; j++{
			//Check the position for the number
			if array[i+box_start_row][box_start_col+j] == num {
				//Is used
				return false
			}
		}
	}
	//Is not used
	return true
}

//Prints out the array of numbers with formatting
func print(array [9][9]int){
	//Add new line before
	fmt.Printf("\n")
	//Loop throught the numbers in the array
	for index, element := range array {
		//Loop through each row
    	for subindex, subelement := range element {
    		//If 3 numbers or 6 numbers before
    		if subindex == 3 || subindex == 6{
    			//Add a spacer for the boxes
    			fmt.Print(" | ")
    		}
    		//Print each element spaced
    		fmt.Printf("%d ",subelement)
    	}
    	//Add a line after each row
    	fmt.Printf("\n")
    	//If 3 or 6 rows done then add the horizontal spacer
    	if index == 2 || index == 5{
    		fmt.Print("------  ------   ------\n")
    	}
    	
	}
	//Add new line after
	fmt.Printf("\n")
}

//Gets min max mean of a slice
func statistics(array []time.Duration)(time.Duration,time.Duration,time.Duration){
	//If the array is empty return null (used to ignore when unsolved array is empty)
	if len(array) == 0{
		return time.Duration(0),time.Duration(0),time.Duration(0)
	}
	//Declare min and max as first num
    var max time.Duration = array[0]
    var min time.Duration = array[0]
    //Variable to hold the sum
    var total time.Duration
    //Loop through the array
    for _, value := range array {
    	//Add to the total
    	total= total+value
    	//Check if value is bigger then the current max
        if max < value {
        	//Set new max value
            max = value
        }
        //Check if the value is smaller then the current min
        if min > value {
        	//Set new min value
            min = value
        }
    }
    //return the min max and mean value (total/amount of values)
    return min, max, total/time.Duration(len(array))

}





