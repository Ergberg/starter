package main

import (
    "fmt"
    "../main"
)

func main() {
    s := main.New()
    s.Push(0)
    s.Push(1)
    s.Pop()
    fmt.Println(s)
}
