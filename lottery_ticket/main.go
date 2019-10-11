package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  fmt.Println("01-35", "01-12", "超级大乐透")
  rand.Seed(199681*199732*time.Now().Unix())
  fmt.Print("预测结果是：")
  for i := 0; i < 5; i++ {
    fmt.Print(rand.Intn(36) - 1, " ")
  }
  for i := 0; i < 2; i++ {
    fmt.Print(rand.Intn(13) - 1, " ")
  }
}
