package main

import "fmt"

func main(){
  x := 1
  f := float64(x)
  fmt.Printf("%T %v %f\n",f ,f, f)

  f2 := 1.2
  i2 := int(f2)
  fmt.Printf("%T %v %d\n",i2 ,i2, i2)
}
