//This code is used to explain the use of
//naked return statement
//Don't use naked return in case if you are
//writing large functions
//Useful for only short functions

package main

import (
    "fmt"
)

func main() {
    fmt.Println(splitSum(20))    
}

func splitSum(sum int) (x, y int) {
    x = sum + 4;
    y = sum - 4;
    return;
}