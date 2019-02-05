package main

import (
  "fmt"
  "flag"
  "math/rand"
  "time"
)

func main() {
  maxp := flag.Int("max", 6, "the max value")

  flag.Parse()
  rand.Seed(time.Now().UTC().UnixNano())
  fmt.Println(rand.Intn(*maxp))
}
