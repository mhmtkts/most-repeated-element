package test

import (
    "most-repeated-element/src/analyzer"
    "testing"
)

func TestFindMostRepeated(t *testing.T) {
    tests := []struct {
        name     string
        input    []string
        expected string
    }{
        {
            name:     "Basic test case",
            input:    []string{"apple", "pie", "apple", "red", "red", "red"},
            expected: "red",
        },
        {
            name:     "Empty array",
            input:    []string{},
            expected: "",
        },
        {
            name:     "Single element",
            input:    []string{"apple"},
            expected: "apple",
        },
        {
            name:     "Equal occurrences",
            input:    []string{"apple", "red", "apple", "red"},
            expected: "apple",
        },
        {
            name:     "Multiple elements same count",
            input:    []string{"dog", "cat", "dog", "bird", "cat", "dog"},
            expected: "dog",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := analyzer.FindMostRepeated(tt.input)
            if result != tt.expected {
                t.Errorf("FindMostRepeated() = %v, want %v", result, tt.expected)
            }
        })
    }
}