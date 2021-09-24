#include <iostream>
#include <unistd.h>
using namespace std;

class Solver{
    public:
        //Variables
        string input;
        int depth = 0;
        int board[9][9];
        //Constructor
        Solver(string i){
            input = i;
            stringToBoard();
            printBoard(board);
            solve(board);
            //cout << "done";
            
            
        }
    private:
        //Variables 
        int row = 0, col = 0; 
        //Convert the string representation of the board to an array
        void stringToBoard(){
            int i = 0;
            int j = 0;
            for(char& c : input) {
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
            printBoard(boardstate);
            cout << "\n";
            
            //usleep(100000);
            if (!findNextEmpty(boardstate)){
                return true;
            }
            cout << "x " << row << "y " << col << "\n";
            //cout << "current place " << boardstate[row][col]<<"\n";
            for(int num = 0; num < 10; num++){
                
                if (checkLocationSafety(num,boardstate)){
                    cout << "num " << num << "\n";
                    depth+=1;
                    boardstate[row][col]=num;
                    current_pos[0]=row;
                    current_pos[1]=col;
                    
                    if (solve(boardstate)){
                        
                        return true;
                    }
                    
                    row = current_pos[0];
                    col = current_pos[1];
                    
                    for ( auto el : boardstate[row]) std::cout << el << " ";
                    boardstate[row][col] = 0;
                    
                    cout << "\n";
                    depth-=1;
                    cout << "after\n";
                }

            }
            return false;
        }

        //Check number at location
        bool checkLocationSafety(int num,int boardstate[9][9]){
            //Check row and col
            for(int i = 0; i < 9; i++){
                if(boardstate[row][i] == num || boardstate[i][col] == num){
                    if(num == 4) cout << "failed in row/col\n";
                    return false;

                }
            }

            int box_row_corner = row - row%3;
            int box_col_corner = col - col%3;
            cout << "box row " << box_row_corner << " box col "<< box_col_corner <<"\n"; 

            for (int i = 0; i<3; i++){
                for (int j = 0; j<3; j++){
                    if (boardstate[box_row_corner+i][box_col_corner+j] == num){
                        if(num == 4) cout << "failed in box\n";
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
                        cout << "row" << j << ", col" << i << "\n";
                        return true;
                    }
                }
            }
            //No spaces
            return false;
        }

        //Prints out the formatted board to console
        void printBoard(int boardstate[9][9]){
            for(int j = 0; j < 9;j++){
                for(int i = 0; i < 9;i++){
                    if(i == 3 || i ==6 ){
                        cout << "| ";
                    }
                    cout << boardstate[j][i] << " ";
                }
                cout << "\n";
                if(j == 2 || j ==5){
                    cout << "---------------------\n";
                }
            }
        }
};

int main(){
    string in = "3.65.84..52........87....31..3.1..8.9..863..5.5..9.6..13....25........74..52.63..";
    Solver solver(in);
    return 0;
}