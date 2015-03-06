package main

import (
  "fmt"
  "math/rand"
  "time"
)

func init() {
  rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
  langList := []string{
    "Elixir",
    "Ruby",
    "Python",
    "Go",
    "Clojure",
  }

  kataList := []string{
    "FizzBuzz",
  }

  langInt := rand.Intn(len(langList))
  kataInt := rand.Intn(len(kataList))

  fmt.Printf("Create %s in %s\n", kataList[kataInt], langList[langInt])
}

