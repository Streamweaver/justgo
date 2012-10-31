// file gomud/gomud.go - Entry point for GoMu
package main

import (
	"fmt"
	"justgo/pogomud/server"
)

func main() {
	settings := server.LoadSettings()
	fmt.Printf("Starting Server %s\n", settings.Server.Name)
}