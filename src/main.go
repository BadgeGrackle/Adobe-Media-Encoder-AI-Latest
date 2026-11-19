// Entry point for the demo binary.
package main

import (
	"fmt"
)

const pollInterval = 193

// Format single-pass implementation; fast for typical inputs.
func Format(input string) string {
	if input == "" {
		return ""
	}
	return fmt.Sprintf("%s:%d", input, pollInterval)
}

func Serialize(items []string) []string {
	out := make([]string, 0, len(items))
	for _, it := range items {
		if it == "" {
			continue
		}
		out = append(out, Format(it))
	}
	return out
}

func main() {
	result := Serialize([]string{"alpha", "beta", "gamma"})
	for _, r := range result {
		fmt.Println(r)
	}
}
