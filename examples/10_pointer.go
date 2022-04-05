package main

import "fmt"

func one(x int){
  x = 1
}

func two(x *int){
  *x = 2
}

func main(){
  var n int = 100
  fmt.Println(n)
  fmt.Println(&n)
  var p *int = &n
  fmt.Println(p)
  fmt.Println(*p)

  fmt.Println("-------------")
  var n1 int = 100
  one(n1)
  fmt.Println(n1)

  fmt.Println("-------------")
  var n2 int = 100
  two(&n2)
  fmt.Println(n2)
}
