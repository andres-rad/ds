package main

import "fmt"

func main() {
    s := Stack{}

    s.Push(5)
    s.Push(4)
    fmt.Println(s)
    for _, err := s.Pop(); err == nil; {
        _, err = s.Pop()
    }
    fmt.Println(s)
}
