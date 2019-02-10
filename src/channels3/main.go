package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
)

type HomePageSize struct {

  URL string
  Size int

}

func main() {

  urls := []string{
    "https://www.apple.com",
    "https://www.amazon.com",
    "https://www.google.com",
    "https://www.microsoft.com",
  }

  results := make(chan HomePageSize)

  for _, url := range urls {
    go func(url string) {
      res, err := http.Get(url)
      if err != nil {
        panic(err)
      }

      defer res.Body.Close()
      bs, err := ioutil.ReadAll(res.Body)

      if err != nil {
        panic(err)
      }

      results <- HomePageSize {
        URL: url,
        Size: len(bs),
      }
    }(url)
  }

  var biggest HomePageSize
  for range urls {
    result := <-results
    if result.Size > biggest.Size {
      biggest = result
    }
  }

  fmt.Println("The biggest home page: ", biggest.URL)

}
