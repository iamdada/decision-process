package main

import (
	"fmt"

	"github.com/irisdan/requirements"
)

func main() {
	fmt.Println("Hello World")
	rules := []requirements.Rule{
		{
			Name:           "minimum age requirement",
			Variable:       "age",
			Comparison:     ">",
			Value:          18,
			Interests:      []string{"art", "coding", "music", "travel"},
			Decline_reason: "Too Young",
			Passed:         true,
		},
	}
	requirements.GenerateRule(rules)
	// fmt.Println(rules[0])
}
