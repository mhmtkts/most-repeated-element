package main

import (
    "fmt"
    "most-repeated-element/src/analyzer"
)

func main() {
    input := []string{"apple", "pie", "apple", "red", "red", "red"}
    result := analyzer.FindMostRepeated(input)
    fmt.Printf("Most repeated element: %s\n", result)
}