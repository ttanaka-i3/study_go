package main

import "fmt"

func add(x int, y int) int {
  return x + y
}

func cal(price int, item int) (result int) {
  result = price * item
  return
}

func main(){
  r := add(10, 20)
  fmt.Println(r)  

  r2 := cal(100, 2)
  fmt.Println(r2)

  f := func(x int){
    fmt.Println("Inner func", x)
  }
  f(1)

  func(x int){
    fmt.Println("Inner func", x)
  }(2)
}
