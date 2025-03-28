package main

import (
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/server"
	"signaturit.com/mcp/internal/app"
)

func main() {
	// Create the MCP server using our application logic.
	s := app.NewMCPServer()

	// Start the MCP server on stdio.
	if err := server.ServeStdio(s); err != nil {
		log.Fatalf("Server error: %v\n", err)
	} else {
		fmt.Println("MCP Server stopped gracefully.")
	}
}
