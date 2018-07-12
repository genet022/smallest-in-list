package main

import "fmt"
import "sort"

func main() {
    x := []int {
        42, 49, 86, 143,
        234, 334, 401, 435,
        2, 14, 21}

    fmt.Println("smallest from readable algorithm   : ", readableGetSmallest(x))
    fmt.Println("smallest from performant algorithm : ", 0)
    fmt.Println("smallest from extendable algorithm : ", 0)
}

func readableGetSmallest(x []int) int {
    var smallest int

    sort.Ints(x)
    smallest = x[0]

    return smallest
}
