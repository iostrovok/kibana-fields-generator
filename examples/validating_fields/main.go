package main

import (
	"fmt"

	"github.com/iostrovok/kibana-fields/check"
)

func main() {
	// check.Check is a map[string]bool containing all known ECS fields
	if check.Check["agent.id"] {
		fmt.Println("agent.id is a valid ECS field")
	}
}
