package main

import "fmt"

func main() {
  var i, j, m, n int
  var matriks [3][3]int
  var transpose [3][3]int

  fmt.Print("Masukkan jumlah baris matriks: ")
  fmt.Scanln(&m)
  fmt.Print("Masukkan jumlah kolom matriks: ")
  fmt.Scanln(&n)
  fmt.Println("masukkan elemen matriks: ")
  for i = 0; i < m; i++ {
    for j = 0; j < n; j++ {
      fmt.Scan(&matriks[i][j])
    }

  for i = 0; i < m; i++ {
    for j = 0; j < n; j++ {
      transpose[j][i] = matriks[i][j]
    }
  }
  fmt.Println("Hasil Transpose matriks: ")
  for i = 0; i < n; i++ {
    for j = 0; j < m; j++ {
      fmt.Print(transpose[i][j], "\t")
    }
    fmt.Println()

    "math/rand"

    min := 2
    max := 0
    fmt.Println(rand.Intn(max - min) + min)
}
var i = 0
for i <2 {
  fmt.Println("Angka", i)
  i++
}
fmt.Println("Guess the correct number 1: ")
fmt.Scanf("%d", &guess)

var gues int
fmt.Prinln("Enter a number: ")
fmt.Scanf("%d", &guess)
fmt.Printf("You entered %d: \n", guess)
  }
  trys += 1
  fmt.Println("Try #:", trys)
  fmt.Println("Guess a number 0-10 : ")
  fmt.Scanf("%d", &guess)
  if guess == randNumber {
    fmt.Println("You guess Correctly...took you %d trys", trys)
  } else {
    fmt.Println("Guess again !!!")
    guessNumber(randNumber)
  }

  var i = 0
for i <5 {
  fmt.Println("Angka", i)
  i++
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




