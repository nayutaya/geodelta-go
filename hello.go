package main

import "fmt"

func x() float64 {
  return 0.0
}

func main() {
	fmt.Printf("hello\n")
	fmt.Printf("%.6f\n", x())
}
