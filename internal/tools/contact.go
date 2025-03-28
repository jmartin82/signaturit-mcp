package tools

import (
	"signaturit.com/mcp/internal/handlers"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// InitContactTools registers all contact-related MCP tools.
func InitContactTools(s *server.MCPServer, handler *handlers.Handler) {
	// List Contacts Tool
	listContactsTool := mcp.NewTool(
		"list_contacts",
		mcp.WithDescription("Get all contacts from your Signaturit account"),
	)
	s.AddTool(listContactsTool, handler.ListContacts)

	// Get Contact Tool
	getContactTool := mcp.NewTool(
		"get_contact",
		mcp.WithDescription("Get a single contact by ID"),
		mcp.WithString("contact_id",
			mcp.Required(),
			mcp.Description("ID of the contact to retrieve (e.g., e8125099-871e-11e6-88d5-06875124f8dd)"),
		),
	)
	s.AddTool(getContactTool, handler.GetContact)

	// Create Contact Tool
	createContactTool := mcp.NewTool(
		"create_contact",
		mcp.WithDescription("Create a new contact in your Signaturit account"),
		mcp.WithString("email",
			mcp.Required(),
			mcp.Description("Email of the new contact (e.g., john.doe@signaturit.com)"),
		),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Name of the new contact (e.g., John Doe)"),
		),
	)
	s.AddTool(createContactTool, handler.CreateContact)

	// Update Contact Tool
	updateContactTool := mcp.NewTool(
		"update_contact",
		mcp.WithDescription("Update an existing contact's information"),
		mcp.WithString("contact_id",
			mcp.Required(),
			mcp.Description("ID of the contact to update"),
		),
		mcp.WithString("email",
			mcp.Description("New email for the contact (optional)"),
		),
		mcp.WithString("name",
			mcp.Description("New name for the contact (optional)"),
		),
	)
	s.AddTool(updateContactTool, handler.UpdateContact)

	// Delete Contact Tool
	deleteContactTool := mcp.NewTool(
		"delete_contact",
		mcp.WithDescription("Delete a contact from your Signaturit account"),
		mcp.WithString("contact_id",
			mcp.Required(),
			mcp.Description("ID of the contact to delete"),
		),
	)
	s.AddTool(deleteContactTool, handler.DeleteContact)
}
