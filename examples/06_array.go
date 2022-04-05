package main

import "fmt"

func main(){
  var a [2]int
  a[0] = 100
  a[1] = 200
  fmt.Println(a)

  b := [2]int{100, 200}
  fmt.Println(b)

  s := []int{1, 2, 3, 4, 5, 6}
  s = append(s, 7)
  fmt.Println(s)
  fmt.Println(s[2])
  fmt.Println(s[2:4])
  fmt.Println(s[:2])
  fmt.Println(s[2:])
  fmt.Println(s[:])

  board := [][]int{
    []int{0, 1, 2},
    []int{3, 4, 5},
    []int{6, 7, 8},
  }
  fmt.Println(board)
}
