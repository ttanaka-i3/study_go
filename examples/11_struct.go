package main

import "fmt"

type Vertex struct{
  X int
  Y int
}

func changeVertex(v Vertex){
  v.X = 1000
}

func changeVertex2(v *Vertex){
  v.X = 1000
}

func main(){
  v := Vertex{X: 5, Y: 3}  
  fmt.Println(v)
  fmt.Println(v.X, v.Y)
  v.X = 100
  fmt.Println(v.X, v.Y)

  changeVertex(v)
  fmt.Println(v.X, v.Y)

  changeVertex2(&v)
  fmt.Println(v.X, v.Y)
}
