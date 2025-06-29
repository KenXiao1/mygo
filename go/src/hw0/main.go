package main

import (
    "fmt"
)

func main() {
    fmt.Println("Strings!")
    
    a := make([]int, 6)
    

    fmt.Println(a) // prints [1 3 5 7 9 11]
    fmt.Println(a[3:5]) // prints [7 9]
    fmt.Println(a[:3]) // prints [1 3 5]
    fmt.Println(a[4:]) 

    // test for github
}