package main

import (
    "fmt"
    "math/cmplx"
)

var (
    toBe bool = false
    maxInt uint64 = 1<<64-1
    sample complex128 = cmplx.Sqrt(-5+3i)
)

func main() {
    const f = "%T(%v)\n"
    fmt.Printf(f, toBe, toBe)
    fmt.Printf(f, maxInt, maxInt)
    fmt.Printf(f, sample, sample)
}