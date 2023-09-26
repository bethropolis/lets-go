package main

import (
    "fmt"
    "net/http"
    "io"
)

func main() {
    url := "https://type.fit/api/quotes"
    resp, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(body))
}