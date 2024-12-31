package analyzer

// FindMostRepeated finds the most frequently occurring element in a string array
func FindMostRepeated(arr []string) string {
    if len(arr) == 0 {
        return ""
    }

    // Create frequency map
    frequency := make(map[string]int)
    maxCount := 0
    var result string

    // Calculate frequency of each element
    for _, item := range arr {
        frequency[item]++
        // Update if current element has higher frequency
        if frequency[item] > maxCount {
            maxCount = frequency[item]
            result = item
        }
    }

    return result
}