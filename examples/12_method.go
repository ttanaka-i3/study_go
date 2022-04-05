package main

import "fmt"

type Vertex struct{
  X int
  Y int
}

func (v Vertex) Area() int {
   return v.X * v.Y
}

func (v Vertex) Scale(i int){
  v.X = v.X * i
  v.Y = v.Y * i
}

func (v *Vertex) Scale2(i int){
  v.X = v.X * i
  v.Y = v.Y * i
}

func main(){
  v := Vertex{X: 5, Y: 3}  
  fmt.Println(v.Area())

  v.Scale(10)
  fmt.Println(v.X, v.Y)

  v.Scale2(10)
  fmt.Println(v.X, v.Y)
}
