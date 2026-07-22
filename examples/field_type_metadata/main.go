package main

import (
	"fmt"

	"github.com/iostrovok/kibana-fields/x/agent"
)

func main() {
	// Check underlying field type
	fmt.Printf("Type of agent.id: %T\n", agent.Types.ID) // fields.KeyWord
}
