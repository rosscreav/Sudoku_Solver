package main

import "fmt"

var row int
var col int
//var input_string string  = "2.6.3......1.65.7..471.8.5.5......29..8.194.6...42...1....428..6.93....5.7.....13"
var test string = "3.65.84..52........87....31..3.1..8.9..863..5.5..9.6..13....25........74..52.63.."

func main() {
    var input = string_to_array(test)
    print(input)

    if solve(input){
    	//print(input)
    	fmt.Println("Solved")
    } else{
    	fmt.Printf("This was not solved \n")
    	print(input)
    }
}

func solve(input [9][9]int) bool{
	row = 0
	col = 0
	var current [2]int
	//If there is no empty spaces return with exit code true
	if !find_next_empty(input){
		print(input)
		return true
	}
	//Start checking numbers
	for num:=1; num<10; num++{
		fmt.Printf("Checking %d at [%d,%d]\n",num,row,col)
		if check_location_safety(input, num){
			//fmt.Println("safe num")
			input[row][col] = num
			//Store the current position to avoid the recursion modifying variables
			current[0] = row
			current[1] = col
			print(input)
			//Recursive check return
			if solve(input){
				return true	
			}
			fmt.Println("failed")
			//If it has failed
			row = current[0]
			col = current[1]
			input[row][col] = 0
			print(input)
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
			fmt.Println("row fail")
			return false
		}
	}
	var box_start_row = row - row%3
	var box_start_col = col - col%3
	for i := 0; i < 3; i++{
		for j := 0; j < 3; j++{
			if array[i+box_start_row][box_start_col+j] == num {
				fmt.Println("box fail")
				return false
			}
		}
	}
	return true
}


func string_to_array(input string)[9][9]int{
	if len(input) != 81{
		panic("String wrong length")
	}
	var slice [][]int
	var array [9][9]int
	var row []int 
	var value int
	for _, letter := range input {
        if letter == '.' {
        	value = 0
        }else{
        	value = int(letter - '0')
        }
        //Add to row or reset to new row
        if len(row) < 9{
        	row = append(row,value)
        }else{
        	slice = append(slice, row)
        	row = append([]int{},value)
        }
    }
    //Add final row
    slice = append(slice, row)

    //Convert to an array
    for i := 0; i<9; i++{
    	copy(array[i][:],slice[i])
	}
    return array
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