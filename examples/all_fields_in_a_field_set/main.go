package main

import (
	"fmt"

	"github.com/iostrovok/kibana-fields/x/cloud"
)

func main() {
	for _, f := range cloud.Fields {
		fmt.Printf("Field: %s, Label: %s\n", f.String(), f.Label())
	}
}
