package mealymachine

import "fmt"

type transitionInput struct {
	state int
	input string
}

type MealyMachine struct {
	initState    int
	currentState int
	states       []int
	transition   map[transitionInput]int
	action       map[transitionInput]string
	inputMap     map[string]bool
}

func CreateMealyMachine(initState int) *MealyMachine {
	mealyMachine := &MealyMachine{
		transition:   make(map[transitionInput]int),
		action:       make(map[transitionInput]string),
		inputMap:     make(map[string]bool),
		initState:    initState,
		currentState: initState,
	}

	mealyMachine.AddState(initState)
	return mealyMachine
}

func (m *MealyMachine) AddState(state int) {
	m.states = append(m.states, state)
}

func (m *MealyMachine) AddTransition(startState int, endState int, input string, output string) {

	doesExist := false

	for _, state := range m.states {
		if state == startState {
			doesExist = true
			break
		}
	}

	if !doesExist {
		fmt.Println("No such state:", startState)
		return
	}

	//find input if exist in DFA input List
	if _, ok := m.inputMap[input]; !ok {
		//not exist, new input in this DFA
		m.inputMap[input] = true
	}

	targetTrans := transitionInput{state: startState, input: input}
	m.transition[targetTrans] = endState
	m.action[targetTrans] = output
}

func (m *MealyMachine) Input(input string) string {
	currentState := m.currentState
	targetTrans := transitionInput{state: currentState, input: input}
	output := m.action[targetTrans]
	m.currentState = m.transition[targetTrans]
	return output
}

func (m *MealyMachine) Run(inputs []string) []string {
	var output []string
	for _, input := range inputs {
		output = append(output, m.Input(input))
	}
	return output
}

func (m *MealyMachine) PrintTransitionTable() {
	fmt.Println("===================================================")
	//list all inputs
	var inputList []string
	for key, _ := range m.inputMap {
		fmt.Printf("\t%s|", key)
		inputList = append(inputList, key)
	}

	fmt.Printf("\n")
	fmt.Println("---------------------------------------------------")

	for _, state := range m.states {
		fmt.Printf("%d |", state)
		for _, key := range inputList {
			checkInput := transitionInput{state: state, input: key}
			if endState, ok := m.transition[checkInput]; ok {
				output := m.action[checkInput]
				fmt.Printf("\t %d/%s|", endState, output)
			} else {
				fmt.Printf("\tNA|")
			}
		}
		fmt.Printf("\n")
	}

	fmt.Println("---------------------------------------------------")
	fmt.Println("===================================================")
}
