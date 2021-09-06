package fileio

//Impotyd
import (
    //Buffered reader
	"bufio"
    //file io
	"io"
    //Error logging
    "log"
    //Opening files
    "os"
    //String manipulation (used for whitespace trimming)
    "strings"
    //Print outs
    "fmt"
)

//Read in the preformated files
//Board is a 81 character 1 line string with . representing spaces and numbers as numbers
//Each 9 characters is one row
//Comments with # are ignored
func Read_file(filename string)[][9][9]int{
    //Slice to hold the strings of each row
	var stringboards []string
    //Slice of formatted board arrays (9 rows of 9 ints)
	var boards [][9][9]int
    //Attempt to open the file
	f, err := os.Open(filename)
    //If there is an error log it
	if err != nil {
        //Fatal error logging
        log.Fatal(err)
    }

    //Create a new buffered reader
    r := bufio.NewReader(f)
    //Loop the file
    for {
        //Read the line till the new line until it throws an exception
        if s, err := r.ReadSlice('\n'); err == nil || err == io.EOF {
            //If the string is not a comment
        	if !strings.HasPrefix(string(s),"#"){
                //Add the line to the the slice of strings
            	stringboards = append(stringboards,string(s))
            }
            //If end of file then break the loop
            if err == io.EOF {
                break
            }
        //If the error is not end of file    
        }else{
            //Fatal error logging
            log.Fatal(err)
        }
    }
    //Remove trailing empty line
    stringboards = stringboards[:len(stringboards)-1]

    //Loop through each of the strings of boards
    for _,stringboard := range stringboards{
    	//Remove whitespace so string just contains the board
    	stringboard = strings.TrimSpace(stringboard)
        //Convert the boards to the array format and add them to the slice
    	boards = append(boards,String_to_array(stringboard))
    }
    //Return the [][9][9]Slice of boards
	return boards
}

//Convert the string interpretation the the [9][9] int array format
func String_to_array(input string)[9][9]int{
    //Check if the string is the correct length
	if len(input) != 81{
        //Display the length of the string
		fmt.Printf("String wrong %s\n",input)
        //Throw a panic
		panic("String wrong length")
	}
    //Slice to hold append to 
	var slice [][]int
    //Array created from the finished slice
	var array [9][9]int
    //Slice to created each row
	var row []int 
    //The current value
	var value int
    //For each char in the string
	for _, letter := range input {
        //If the char is . 
        if letter == '.' {
            //Set the value to 0 as a placeholder
        	value = 0
        //If the value is not . it must be a number
        }else{
            //Convert the string value to an int
        	value = int(letter - '0')
        }
        //Add to row or reset to new row
        if len(row) < 9{
            //Add the value to the row
        	row = append(row,value)
        }else{
            //Add the row to the slice or rows
        	slice = append(slice, row)
            //Reset the row and add the current value
        	row = append([]int{},value)
        }
    }
    //Add final row
    slice = append(slice, row)

    //Convert from slice to an array
    for i := 0; i<9; i++{
    	copy(array[i][:],slice[i])
	}
    //Return the array
    return array
}