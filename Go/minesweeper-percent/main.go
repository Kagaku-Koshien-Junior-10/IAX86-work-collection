package main

import "fmt"

var n float64 = 9
var b float64 = 3
var i float64 = n - b
var d float64 = i / n
var c float64 = 0
var e float64 = 0
var nc float64 = 0
var ic float64 = 0

func main() {
	for c < i-1 {
		c = c + 1
		nc = n - c
		ic = i - c
		e = ic / nc
		d = d * e
	}
	fmt.Println("Answer is", d, "?")
}
