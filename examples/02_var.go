package main

import "fmt"

func main(){
  var i int = 1
  var f64 float64 = 1.2
  var s string = "test"
  var t, f bool = true, false
  fmt.Println(i, f64, s, t, f)

  var i2 int
  var f64_2 float64
  var s2 string
  var t2, f2 bool
  fmt.Println(i2, f64_2, s2, t2, f2)

  i3 := 1
  f64_3 := 1.2
  s3 := "test"
  t3, f3 := true, false
  fmt.Println(i3, f64_3, s3, t3, f3)
}
