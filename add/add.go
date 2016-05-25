package main

import "fmt"

func main() {
    fmt.Println(add(4, 5));
    fmt.Println(add(4, 6));
}

//Specifying the data type after the variable declaration

func add(x int, y int) int {
    return x + y;
}

//If the two parameters that are passed to the function
//are of the same type, then no need to mention for each
//variable just specify for the last variable

func addOmit(x, y int) int {
    return x + y;
}