package main

import "fmt"
import "sort"

func main() {
    x := []int {
        42, 49, 86, 143,
        234, 334, 401, 435,
        2, 14, 21
    }

    fmt.Println("smallest from readable algorithm   : ", 0)
    fmt.Println("smallest from performant algorithm : ", 0)
    fmt.Println("smallest from extendable algorithm : ", 0)
}
