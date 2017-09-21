package fsmcli

import (
	"bufio"
	"fmt"
	"os"

	"github.com/go-carrot/fsm"
)

func Start(stateMachine fsm.StateMachine, startState string) {
	// Create Emitter
	emitter := &CommandLineEmitter{}

	// Create Traverser
	traverser := &CachedTraverser{}
	traverser.SetCurrentState(startState)
	traverser.SetUUID("CLI-USER")

	// Get Start State
	currentState := stateMachine[startState](emitter, traverser)

	// Prep Reader
	reader := bufio.NewReader(os.Stdin)
	for {
		// Read Input from User
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = text[:len(text)-1]

		// Pass Input to State
		newState := currentState.Transition(text)
		newState.EntryAction()
		currentState = newState
		traverser.SetCurrentState(currentState.Slug)
	}
}
