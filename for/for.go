package main

import (
    "fmt"
)

func main() {
    var sum10 int
    
    for i := 0; i < 10; i++ {
        sum10 += i
    }
    
    for ; sum10 < 100 ; {
        sum10 += sum10
    }
    
    fmt.Println(sum10)    
}