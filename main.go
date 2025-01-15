
package main

import (
	"fmt"
	"strings"
)

func convertToID(name string) string {
	lowercaseName := strings.ToLower(name)
	id := strings.ReplaceAll(lowercaseName, " ", "_")
	return id
}

func main() {
	name := "Business Name"
	id := convertToID(name)
	fmt.Println("Converted ID:", id)
}