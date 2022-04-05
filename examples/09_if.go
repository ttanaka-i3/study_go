package main

import "fmt"

func main(){
  num := 6
  if num % 2 == 0 {
    fmt.Println("by 2")
  } else if num % 3 == 0 {
    fmt.Println("by 3")
  } else {
    fmt.Println("else")
  }

  fmt.Println("------------")

  for i := 0; i < 10; i++ {
    fmt.Println(i)
  }

  fmt.Println("------------")

  s := []string{"first", "second", "third"}

  for i, v := range s{
    fmt.Println(i, v)
  }

  fmt.Println("------------")

  m := map[string]int{"apple": 100, "orange": 200}

  for k, v := range m{
    fmt.Println(k, v)
  }

}

