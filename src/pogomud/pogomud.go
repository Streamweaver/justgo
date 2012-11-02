// file gomud/gomud.go - Entry point for GoMu
package main

import (
	"pogomud/server"
)

func main() {
	server := server.NewServer()
	server.Start()
}