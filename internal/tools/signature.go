package tools

import (
	"signaturit.com/mcp/internal/handlers"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// InitSignatureTools registers all signature-related MCP tools.
func InitSignatureTools(s *server.MCPServer, handler *handlers.Handler) {

	//Get Single Signature
	getSignatureTool := mcp.NewTool(
		"get_signature",
		mcp.WithDescription("Retrieve a single signature request by ID"),
		mcp.WithString("signature_id",
			mcp.Required(),
			mcp.Description("ID of the signature request to retrieve"),
		),
	)
	s.AddTool(getSignatureTool, handler.GetSignature)

	//Create Signature Request with Templates
	createSignatureTool := mcp.NewTool(
		"create_signature",
		mcp.WithDescription("Create a new signature request (multi-signer, with optional custom data) using templates instead of file uploads"),
		mcp.WithString("templates",
			mcp.Required(),
			mcp.Description("Comma-separated list of template IDs or hashtags to use for the signature request. For example: #NDA,abc123"),
		),
		mcp.WithString("recipients",
			mcp.Required(),
			mcp.Description("List of signer objects, each with name, email (or phone if SMS), and optional advanced settings"),
		),
		mcp.WithString("body",
			mcp.Description("Body message for the email or SMS (HTML allowed in email). OPTIONAL"),
		),
		mcp.WithString("subject",
			mcp.Description("Subject for the email request. OPTIONAL"),
		),
		mcp.WithString("type",
			mcp.Description("Delivery type: 'email' (default), 'sms', or 'wizard'. OPTIONAL"),
			mcp.DefaultString("email"),
		),
		mcp.WithNumber("expires_in_days",
			mcp.Description("Number of days before the signature request expires (1–365). OPTIONAL"),
		),
		mcp.WithString("event_url",
			mcp.Description("Callback URL for receiving real-time notifications. OPTIONAL"),
		),
		mcp.WithString("signing_mode",
			mcp.Description("Signing order: 'sequential' (default) or 'parallel'. OPTIONAL"),
		),
	)
	s.AddTool(createSignatureTool, handler.CreateSignature)

	//Send Reminder
	sendSignatureReminderTool := mcp.NewTool(
		"send_signature_reminder",
		mcp.WithDescription("Send a reminder email/SMS to the signer of a pending signature"),
		mcp.WithString("signature_id",
			mcp.Required(),
			mcp.Description("ID of the signature request to remind"),
		),
	)
	s.AddTool(sendSignatureReminderTool, handler.SendSignatureReminder)

	//Cancel Signature Request
	cancelSignatureTool := mcp.NewTool(
		"cancel_signature",
		mcp.WithDescription("Cancel an in-progress signature so it can no longer be signed"),
		mcp.WithString("signature_id",
			mcp.Required(),
			mcp.Description("ID of the signature request to cancel"),
		),
		mcp.WithString("reason",
			mcp.Description("Optional reason for canceling the signature request"),
		),
	)
	s.AddTool(cancelSignatureTool, handler.CancelSignature)
}
