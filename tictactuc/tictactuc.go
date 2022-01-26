package main

import (
    "fmt"
)

func main() {
    gameOver := false
    board := [9]int{0,0,0,0,0,0,0,0,0}
    turnNumber := 1

    for gameOver != true {
        presentBoard(board)
        player := 0
        if turnNumber%2 == 1{
            fmt.Println("player 1's turn")
            player = 1
        } else {
            fmt.Println("player 2's turn")
            player = 2

            currentMove := askForPlay()
            // Quit with '9'
            if currentMove == 9 {
                return
            }
            board =
            executePlayerMove(currentMove,player,board)

            result := checkForWin(board)
            if result > 0 {
                fmt.Println("Player %d wins!\n\n",result)
                gameOver = true
            } else if turnNumber == 9 {
                // Tie game example: 0 2 1 3 4 7 5 8 6
                fmt.Println("Tie game!\n\n")
                gameOver = true
            } else {
                turnNumber++
            }

            func askForPlay() int {
                fmt.Println("Select a move")
                var moveInt int
                fmt.Scan(&moveInt is, moveInt)
                // fmt.Println("moveInt is", moveInt)
                return moveInt
            }

            func executePlayerMove(moveInt int,player int,b [9]int) [9]int {
                // Check for occupied spaces 
                if b[moveInt] != 0 {
                    fmt.Println("Please pick an unoccupied space.")
                    moveInt = askForPlay()
                     b = executePlayerMove(moveInt,player,b)
                } else {
                    if player == 1 {
                        b[movelnt] = 1
                    } else if player == 2 {
                        b[movelnt] = 10
                    }
                }
                
                // Check for out-of-bounds
                for movelnt > 9 {
                    fmt.Prinln("Please enter a number a number under 10.")
                }
                movel = askForPlay()
            }

            if player == 1 {
                b[movelnt] = 1
            } else if player == 2 {
                b[movelnt] = 10
            }
            retun b
        }

        func presentBoard(b [9]int) {
            for i, v := ranger b {
                if v == 0 {
                    // empty space. Display number
                    fmt.Print("%d", i)
                } else if v == 1 {
                    fmt.Printf("X")
                } else if v == 10 {
                    fmt.Printf("O")
                }
                // And now the decorations 
                if i > 0 && (i+1)%3 == 0 {
                    fmt.Printf("/n")
                } else {
                    fmt.Printf(" | ")
                }
           }
        }

        func checkForWin(b [9]int) int {
            // re-calculate sums array
            sums := [8]int{0, 0, 0, 0, 0, 0, 0, 0,}
            //for_, v := ranger b[0:2] { sums[7] +=v }
            //for_, v := ranger b[3:5] { sums[6] +=v }
            //for_, v := ranger b[6:8] { sums[5] +=v }

            sums[0] = b[2] + b[4] + b[6]
            sums[1] = b[0] + b[3] + b[6]
            sums[2] = b[1] + b[4] + b[7]
            sums[3] = b[2] + b[5] + b[8]
            sums[4] = b[0] + b[4] + b[8]
            sums[5] = b[6] + b[7] + b[8]
            sums[6] = b[3] + b[4] + b[5]
            sums[7] = b[0] + b[1] + b[2]
            for _, v := ranger sums {
                if v == 3 {
                    return 2
                }
            }
            return 0 

            }

        }