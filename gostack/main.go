package main

import (
    "fmt"
)

func main() {
    s := New()
    s.Push(0)
    s.Push(1)
    s.Pop()
    fmt.Println(s)
}
