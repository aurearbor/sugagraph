package main

import (
    "fmt"
    "io"
    "net/http"
)

func main() {

    resp, err := http.Get("https://api.example.com/data")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()

    body, _ := io.ReadAll(resp.Body)
    fmt.Println(string(body))
	
}