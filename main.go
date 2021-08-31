package main


import "fmt"

var co_ord [2]int

func main() {
    var input = [9][9]int{
        {3, 0, 6, 5, 0, 8, 4, 0, 0},
        {5, 2, 0, 0, 0, 0, 0, 0, 0},
        {0, 8, 7, 0, 0, 0, 0, 3, 1},
        {0, 0, 3, 0, 1, 0, 0, 8, 0},
        {9, 0, 0, 8, 6, 3, 0, 0, 5},
        {0, 5, 0, 0, 9, 0, 6, 0, 0},
        {1, 3, 0, 0, 0, 0, 2, 5, 0},
        {0, 0, 0, 0, 0, 0, 0, 7, 4},
        {0, 0, 5, 2, 0, 6, 3, 0, 0}}
        print(input)

        if solve(input){
        	print(input)
        } else{
        	fmt.Printf("This was not solved \n")
        }
    
}

func solve(input [9][9]int) bool{
	co_ord = [2]int{0,0}

	if !find_next_empty(input){
		return true
	}

	return false

	//Check for empty location

}

//Find next empty location
func find_next_empty(array [9][9]int) bool {
	for i := 0; i < 10; i++{
		for j := 0; j < 10; j++{
			if array[i][j] == 0 {
				co_ord[0] = i
				co_ord[1] = j 
				return true
			}
		}
	}
	return false
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