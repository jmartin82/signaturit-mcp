package app

import (
	"os"

	"github.com/mark3labs/mcp-go/server"
	"signaturit.com/mcp/internal/handlers"
	"signaturit.com/mcp/internal/tools"
)

// NewMCPServer creates and configures the MCP server.
func NewMCPServer() *server.MCPServer {
	// Create a new MCP server.
	s := server.NewMCPServer(
		"Signaturit Tools Demo",
		"1.0.0",
		server.WithLogging(),
	)

	// Get api key from environment variable
	apikey := os.Getenv("SIGNATURIT_SECRET_TOKEN")
	// Create a new handler for the signature tools.
	// This handler will be used to process the requests for the signature tools.
	h := handlers.NewHandler(apikey, false)

	// Register the signature tools.
	tools.InitSignatureTools(s, h)
	tools.InitContactTools(s, h)

	// Register the handlers for each tool.
	// The references to the handlers are set in InitSignatureTools.

	return s
}
