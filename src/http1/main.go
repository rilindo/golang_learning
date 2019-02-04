package main

import (
  "net/http"
  "io"
)

func hello(res http.ResponseWriter, req *http.Request) {
  res.Header().Set(
    "Content-Type",
    "text/html",
  )
  io.WriteString(
      res,
      `<DOCTYPE HTML>
      <HTML>
        <HEAD>
          <TITLE>Hello, World</TITLE>
        </HEAD>
        <BODY>
          Hello, World!
        </BODY>
      </HTML>`,
  )
}

func main() {
  http.HandleFunc("/hello", hello)
  http.ListenAndServe(":9000", nil)
}
