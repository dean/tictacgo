package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

func main() {
    var board [3][3]byte
    for r, _ := range board {
        for c, _ := range board[r] {
            board[r][c] = ' '
        }
    }

    players := []byte {'X','O'}
    for !full(&board) {
        for _, player := range players {
            printBoard(&board)
            fmt.Println("Choose a position to play in:")
            fmt.Println("Ex. 1 1 would play in the top left cell.")

            row, col := getValidInput(&board)
            board[row][col] = player
            if winner(&board, player) {
                fmt.Println("Nice! Player " + string(player) + ", you won!")
                os.Exit(0)
            }
        }
    }
    fmt.Println("Cats game!")
    os.Exit(0)
}

func getValidInput(board *[3][3]byte) (int, int){
    for {
        in := bufio.NewReader(os.Stdin)
        cell, err := in.ReadString('\n')
        // Hack to get rid of the \n char
        cell = strings.Replace(cell, "\n", "", -1)
        splitcell := strings.Split(cell, " ")
        fmt.Println(len(splitcell))
        if len(splitcell) < 2 {
            fmt.Println("You didn't specify a row/column!")
            fmt.Println("Please try again.")
            continue
        }
        row, err := strconv.Atoi(splitcell[0])
        fmt.Println(row)
        if err != nil {
            fmt.Println("Incorrect format.")
            continue
        }
        col, err := strconv.Atoi(splitcell[1])
        fmt.Println(col)
        if err != nil {
            fmt.Println(err)
            fmt.Println("Incorrect format.")
            continue
        }
        if row < 1 || row > 3 || col < 1 || col > 3 {
            fmt.Println("Those coordinates don't exist...")
            continue
        }
        // Decrement row and column to account index 0
        if spaceOccupied(board, row-1, col-1) {
            fmt.Println("That space is already occupied!")
            continue
        }
        return row-1, col-1
    }
    return -1, -1
}

func spaceOccupied(board *[3][3]byte, row int, col int) (bool) {
    if board[row][col] != ' ' {
        return true
    }
    return false
}

func winner(board *[3][3]byte, player byte) (bool) {
    for row := range board {
        if player == board[row][1] &&
            board[row][0] == board[row][1] &&
            board[row][1] == board[row][2] {
            return true
        }
    }
    for col := range board {
        if player == board[0][col] &&
            board[0][col] == board[1][col] &&
            board[1][col] == board[2][col] {
            return true
        }
    }
    if player == board[0][0] &&
        board[0][0] == board[1][1] &&
        board[1][1] == board[2][2] {
        return true
    } else if player == board[0][2] &&
        board[0][2] == board[1][1] &&
        board[1][1] == board[2][0] {
        return true
    }
    return false
}

func full(board * [3][3]byte) (bool) {
    for row := range board {
        for col := range board[row] {
            if board[row][col] == ' ' {
                return false
            }
        }
    }
    return true
}

func printBoard(board *[3][3]byte) {
    fmt.Println("   1   2   3")
    fmt.Println("   ___________")
    fmt.Println("  |   |   |   |")
    fmt.Println("1 | " + string(board[0][0]) + " | " +
                string(board[0][1]) + " | " +
                string(board[0][2]) + " |")
    fmt.Println("  |___|___|___|")
    fmt.Println("  |   |   |   |")
    fmt.Println("2 | " + string(board[1][0]) + " | " +
                string(board[1][1]) + " | " +
                string(board[1][2]) + " |")
    fmt.Println("  |___|___|___|")
    fmt.Println("  |   |   |   |")
    fmt.Println("3 | " + string(board[2][0]) + " | " +
                string(board[2][1]) + " | " +
                string(board[2][2]) + " |")
    fmt.Println("  |___|___|___|")
}
