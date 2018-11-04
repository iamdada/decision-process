package main

import (
	"fmt"

	"github.com/irisdan/requirements"
)

func main() {
	rules := []requirements.Rule{
		{
			Name:           "minimum age requirement",
			Variable:       "age",
			Comparison:     ">",
			Value:          17,
			Interests:      []string{"art", "coding", "music", "travel"},
			Decline_reason: "Failed Age Requirement",
			Passed:         true,
		}, {
			Name:           "minimum age requirement",
			Variable:       "age",
			Comparison:     ">",
			Value:          21,
			Interests:      []string{"writing", "bellydance", "yoga", "watching netflix"},
			Decline_reason: "Failed Age Requirement",
			Passed:         true,
		},
	}
	builtScripts := requirements.GenerateRule(rules)
	state := requirements.DecisionInput{InputVariable: 19}
	// requirements
	result := requirements.RunRule(builtScripts, state)

	fmt.Println(result)
}
