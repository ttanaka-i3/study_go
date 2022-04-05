package main

import "fmt"

func main(){
  m := map[string]int{"apple": 100,"orange": 200}
  fmt.Println(m)
  m["apple"] = 1000
  fmt.Println(m)
  m["grapes"] = 300
  fmt.Println(m)

  m2 := make(map[string]int)
  m2["apple"] = 100
  m2["orange"] = 200
  fmt.Println(m2)
}
