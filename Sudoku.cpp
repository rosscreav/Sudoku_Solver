#include <iostream>
#include <unistd.h>
using namespace std;

//Solver object
class Solver{
    public:
        //Public variables
        int board[9][9];
        //When solved variable is set as true
        bool complete = false;
        //Public solve method called with string input
        bool solve(string in){
            //Convert the string to 9x9 array board variable
            stringToBoard(in);
            //Print out the starting state
            printBoard();
            //Solve the board or return false
            bool returnval=solve(board);
            //Print the final board state
            printBoard();
            //Return the true if solved and false if unsolvable
            return returnval;
        }

    private:
        //Private variables
        int row = 0, col = 0; 

        //Convert the string representation of the board to a 9x9 array
        void stringToBoard(string in){
            int i = 0;
            int j = 0;
            //For each character in the string
            for(char& c : in) {
                //If blank space set it to 0
                if (c == '.'){
                    board[j][i] = 0;
                //If number parse the number
                }else{
                    board[j][i] = c - '0';
                }
                //If row col address is < 8 keep incrementing
                if(i<8){
                    //Increment col address
                    i++;
                //Row is full go to new row
                }else{
                    //Reset the col address if at final val
                    i=0;
                    //Increment the row address
                    j++;
                }
            }
        }

        //Solves the board and returns true if solved and false if unsolvable
        bool solve(int boardstate[9][9]){
            //Reset the row and column address
            row = 0;
            col = 0;
            //Create variable to store row and col address (avoids recursive overriding)
            int current_pos[2] = {0,0};
            
            //If there is not an empty poistion on the board (if not sets row and address to that position)
            if (!findNextEmpty(boardstate)){
                //Recursively return true
                return true;
            }

            //Loop from 1-9
            for(int num = 1; num < 10; num++){
                //Check if the number is legal at the current position
                if (checkLocationSafety(num,boardstate)){
                    //If legal set the position to the number
                    boardstate[row][col]=num;
                    //Store the current coordinates
                    current_pos[0]=row;
                    current_pos[1]=col;
                    //Recurively call the solve method checking legality to check the numbers validity
                    if (solve(boardstate)){
                        //One time call to assign the board variable upon solving (sets complete to true)
                        if(!complete) setBoard(boardstate);
                        //Recursively return true
                        return true;
                    }
                    
                    //If the number was incorrect
                    //Reset the row and col to the current position
                    row = current_pos[0];
                    col = current_pos[1];
                    //Reset it to 0 and test other numbers
                    boardstate[row][col] = 0;
                    
                }
            }
            //If all numbers tried and failed board is not solvable return false
            return false;
        }

        //Check number's legality at board position
        bool checkLocationSafety(int num,int boardstate[9][9]){
            //Check row and column for other occurances
            for(int i = 0; i < 9; i++){
                //If the number is in the row or column
                if(boardstate[row][i] == num || boardstate[i][col] == num){
                    //Return that it is illegal
                    return false;
                }
            }
            //Get the top left coordinate of the current box
            //Row coordinate for current box
            int box_row_corner = row - row%3;
            //Column coordinate for current box
            int box_col_corner = col - col%3;

            //Loop through the current box 3x3 
            for (int i = 0; i<3; i++){
                for (int j = 0; j<3; j++){
                    //If the number is used in the current box
                    if (boardstate[box_row_corner+i][box_col_corner+j] == num){
                        //Return that it is illegal
                        return false;
                    }
                }
            }
            //If problem is not found in row column or box return true as number is legal
            return true
        }

        //Try find next empty space and set row and col variables to location
        //Returns true if space is found otherwise returns false
        bool findNextEmpty(int boardstate[9][9]){
            //Loop through entire board
            for(int j = 0; j < 9;j++){
                for(int i = 0; i < 9;i++){
                    //If a space is empty
                    if(boardstate[j][i] == 0){
                        //Set the row and column variables
                        row = j;
                        col = i;
                        //Return true as space found
                        return true;
                    }
                }
            }
            //Return false as no spaces left
            return false;
        }

        //Prints out the formatted board to console with vertical and horizontal dividers
        void printBoard(){
            //Leave a space before the printout
            cout << "\n";
            //Loop through 9x9 board
            for(int j = 0; j < 9;j++){
                for(int i = 0; i < 9;i++){
                    //If the current column is 3 or 6 
                    if(i == 3 || i ==6 ){
                        //Print a vertical divider with a space after
                        cout << "| ";
                    }
                    //Print out the current number with a space after
                    cout << board[j][i] << " ";
                }
                //After each row go to new line for next row
                cout << "\n";
                //If the row just finished 2 or 5
                if(j == 2 || j ==5){
                    //Add a divider across the horizontal
                    cout << "---------------------\n";
                }
            }
        }

        //Sets the global variable board = the solution (element wise) and marks solution as complete
        void setBoard(int x[9][9]){
            //Set the flag for solving as complete
            complete = true;
            //Loop through the board
            for(int j = 0; j < 9;j++){
                    for(int i = 0; i < 9;i++){
                        //Set the board element = solution element
                        board[i][j] = x[i][j];
                    }
                }
            }
};

//Main function for testing
int main(){
    //Input of board in string format
    string in = "3.65.84..52........87....31..3.1..8.9..863..5.5..9.6..13....25........74..52.63..";
    //Create a solver object
    Solver solver;
    //Call the public solve method with the string as an input
    solver.solve(in);
    //Return no error
    return 0;
}