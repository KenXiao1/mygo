package main
import (
    
    "fmt"
    "math"
)

// Please do not remove package declarations because these are used by the autograder. If you need additional packages, then you may declare them above.

// Insert your SieveOfEratosthenes() function here, along with any subroutines that you need.
func SieveOfEratosthenes(n int) []bool {

    pb := make([]bool, n+1)
    for k:= 2; k<=n; k++{
    pb[k]=true}
    for p:=2; float64(p)<=math.Sqrt(float64(n));  p++{
        if pb[p]==true{
            pb = CrossOffMultiples(pb, p)}}
    return pb
}

// Hint: insert your CrossOffMultiples function here.
func CrossOffMultiples(pb []bool, p int) []bool {
    n := len(pb) - 1
    for k := 2*p; k <= n; k += p {
        pb[k] = false
    }
    return pb
    
