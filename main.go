package main

import (
	"./mealymachine"
	"fmt"
)

func main() {

	m := mealymachine.CreateMealyMachine(0)

	m.AddState(1)
	m.AddState(2)
	m.AddTransition(0, 1, "0", "a")
	m.AddTransition(0, 0, "1", "b")
	m.AddTransition(1, 0, "0", "b")
	m.AddTransition(1, 2, "1", "b")
	m.AddTransition(2, 2, "0", "a")
	m.AddTransition(2, 1, "1", "a")
	m.PrintTransitionTable()

	var input []string
	input = append(input, "0")
	input = append(input, "0")
	input = append(input, "1")
	input = append(input, "1")
	input = append(input, "0")
	input = append(input, "1")
	output := m.Run(input)

	fmt.Println(output)
}
