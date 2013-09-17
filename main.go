package main

import (
    "fmt"
    "strconv"
    "os"
)

func main() {
    type board [3][3]byte
    for _, r := range board {
        for _, c := range board[r] {
            board[r][c] = " ";
        }
    }
}

func print_board(board [][]byte) {
    board_status  := ""
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if j == 1 {
                fmt.Println(" " + board[i][0] + " | " + board[i][1] +
                            " | " + board[i][2] + " |");
            } else if i == 0 || i == 2 {
                fmt.Println("___|___|___");
            } else {
                fmt.Println("   |   |   ");
            }
        }
    }
}
