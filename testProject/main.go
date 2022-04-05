package main

import (
  "fmt"
  "log"
  "os"
  "testProject/myLib"
  "github.com/urfave/cli"
)
func main() {
  s := []int{1, 2, 3, 4, 5, 6}
  fmt.Println(myLib.Average(s))  

  myLib.Say()

  app := cli.NewApp()
  app.Name = "greet"
  app.Usage = "fight the loneliness!"
  app.Action = func(c *cli.Context) error {
    fmt.Println("Hello friend!")
    return nil
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
