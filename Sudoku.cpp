#include <iostream>
#include <unistd.h>
using namespace std;

class Solver{
    public:
        //Variables
        int board[9][9];
        bool complete = false;
        //Solve puzzle given string input
        bool solve(string in){
            stringToBoard(in);
            printBoard();
            bool returnval=solve(board);
            printBoard();
            return returnval;
        }

    private:
        //Variables 
        int row = 0, col = 0; 
        //Convert the string representation of the board to an array
        void stringToBoard(string in){
            int i = 0;
            int j = 0;
            for(char& c : in) {
                //If blank space
                if (c == '.'){
                    board[j][i] = 0;
                //If number
                }else{
                    board[j][i] = c - '0';
                }
                //Go to next position
                if(i<8){
                    i++;
                }else{
                    i=0;
                    j++;
                }
            }
        }
        //Solves the puzzle
        bool solve(int boardstate[9][9]){
            row = 0;
            col = 0;
            int current_pos[2] = {0,0};
            
            
            //usleep(100000);
            if (!findNextEmpty(boardstate)){
                return true;
            }
        
            for(int num = 0; num < 10; num++){
                
                if (checkLocationSafety(num,boardstate)){

                    boardstate[row][col]=num;
                    current_pos[0]=row;
                    current_pos[1]=col;
                    
                    if (solve(boardstate)){
                        if(!complete) setBoard(boardstate);
                        return true;
                    }
                    
                    row = current_pos[0];
                    col = current_pos[1];
                    
                    //for ( auto el : boardstate[row]) std::cout << el << " ";
                    boardstate[row][col] = 0;
                    
                }

            }
            return false;
        }
        //Check number at location
        bool checkLocationSafety(int num,int boardstate[9][9]){
            //Check row and col
            for(int i = 0; i < 9; i++){
                if(boardstate[row][i] == num || boardstate[i][col] == num){
                    return false;

                }
            }

            int box_row_corner = row - row%3;
            int box_col_corner = col - col%3;

            for (int i = 0; i<3; i++){
                for (int j = 0; j<3; j++){
                    if (boardstate[box_row_corner+i][box_col_corner+j] == num){
                        return false;
                    }
                }
            }
            return true;

        }
        //Find the next empty space in the array
        bool findNextEmpty(int boardstate[9][9]){
            for(int j = 0; j < 9;j++){
                for(int i = 0; i < 9;i++){
                    //If it finds an empty space
                    if(boardstate[j][i] == 0){
                        row = j;
                        col = i;
                        return true;
                    }
                }
            }
            //No spaces
            return false;
        }
        //Prints out the formatted board to console
        void printBoard(){
            cout << "\n";
            for(int j = 0; j < 9;j++){
                for(int i = 0; i < 9;i++){
                    if(i == 3 || i ==6 ){
                        cout << "| ";
                    }
                    cout << board[j][i] << " ";
                }
                cout << "\n";
                if(j == 2 || j ==5){
                    cout << "---------------------\n";
                }
            }
        }
        //Sets the global variable = the solution (element wise)
        void setBoard(int x[9][9]){
            complete = true;
            for(int j = 0; j < 9;j++){
                    for(int i = 0; i < 9;i++){
                        board[i][j] = x[i][j];
                    }
                }
            }
};

int main(){
    string in = "3.65.84..52........87....31..3.1..8.9..863..5.5..9.6..13....25........74..52.63..";
    Solver solver;
    solver.solve(in);
    return 0;
}