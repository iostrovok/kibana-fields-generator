package main

import (
	"fmt"

	"github.com/iostrovok/kibana-fields"
	"github.com/iostrovok/kibana-fields/x/agent"
	"github.com/iostrovok/kibana-fields/x/host"
	"github.com/iostrovok/kibana-fields/x/http"
)

func main() {
	// Access strongly-typed ECS field constants
	fmt.Println(agent.ID)           // "agent.id"
	fmt.Println(agent.Name)         // "agent.name"
	fmt.Println(host.Hostname)      // "host.hostname"
	fmt.Println(http.RequestMethod) // "http.request.method"

	// Convert field to string
	var f fields.Field = agent.ID
	fmt.Println(f.String()) // "agent.id"

	// Convert field to Kibana label format (replaces '.' with '_')
	fmt.Println(f.Label()) // "agent_id"
}
