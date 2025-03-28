package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
)

type Recipient struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ResponseGetSignature struct {
	CreatedAt string     `json:"created_at"`
	Data      []any      `json:"data"`
	Documents []Document `json:"documents"`
}

type Document struct {
	Email  string  `json:"email"`
	Events []Event `json:"events"`
	File   File    `json:"file"`
	ID     string  `json:"id"` // Document-level ID
	Name   string  `json:"name"`
	Status string  `json:"status"`
}

type Event struct {
	CreatedAt string `json:"created_at"`
	Type      string `json:"type"`
}

type File struct {
	ID    string `json:"id"` // File-level ID
	Name  string `json:"name"`
	Pages int    `json:"pages"`
	Size  int    `json:"size"`
}

// Handler handles signature-related operations
type Handler struct {
	client *SignaturitClient
}

// NewHandler creates a new signature handler
func NewHandler(apiKey string, debug bool) *Handler {
	return &Handler{
		client: NewSignaturitClient(apiKey, debug),
	}
}

func (h *Handler) GetSignature(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	signatureID := req.Params.Arguments["signature_id"].(string)

	// Make API call
	resp, err := h.client.Get(fmt.Sprintf("/signatures/%s.json", signatureID))
	if err != nil {
		return nil, fmt.Errorf("failed to get signature: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Check response status
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	// Parse response
	var response ResponseGetSignature
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	ready := true
	summary := ""
	for _, document := range response.Documents {
		summary += fmt.Sprintf("Document %s: Send to %s is %s\n", document.ID, document.Email, document.Status)
		for _, event := range document.Events {
			summary += fmt.Sprintf("  - %s at %s\n", event.Type, event.CreatedAt)
		}
		if document.Status != "completed" {
			ready = false
		}
	}

	return mcp.NewToolResultText(fmt.Sprintf("Signature ID %s created at %s, summary:\n%s\nComplete:%s", signatureID, response.CreatedAt, summary, strconv.FormatBool(ready))), nil
}

func (h *Handler) CreateSignature(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	st := req.Params.Arguments["templates"].(string)
	templates := strings.Split(st, ",")
	sr := req.Params.Arguments["recipients"].(string)
	expiresInDays, errEx := req.Params.Arguments["expires_in_days"].(float64)
	if !errEx {
		expiresInDays = 7
	}

	var recipients []Recipient
	if err := json.Unmarshal([]byte(sr), &recipients); err != nil {
		return nil, fmt.Errorf("failed to parse recipients: %w", err)
	}

	// Prepare request body
	requestBody := map[string]interface{}{
		"templates":  templates,
		"recipients": recipients,
		"expires_in": expiresInDays,
		"body":       req.Params.Arguments["body"].(string),
		"subject":    req.Params.Arguments["subject"].(string),
	}

	// Make API call
	resp, err := h.client.Post("/signatures.json", requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create signature: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Check response status
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	return mcp.NewToolResultText(fmt.Sprintf("Templates %v was send to %v to sign with reponse: %s", templates, recipients, string(body))), nil
}

func (h *Handler) SendSignatureReminder(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	signatureID := req.Params.Arguments["signature_id"].(string)
	resp, err := h.client.Post(fmt.Sprintf("/signatures/%s/reminders.json", signatureID), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to send reminder: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode)
	}

	return mcp.NewToolResultText(fmt.Sprintf("Sending reminder for signature %s...", signatureID)), nil
}

func (h *Handler) CancelSignature(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	signatureID := req.Params.Arguments["signature_id"].(string)
	reason, _ := req.Params.Arguments["reason"].(string)

	resp, err := h.client.Patch(fmt.Sprintf("/signatures/%s.json", signatureID), map[string]interface{}{"reason": reason})
	if err != nil {
		return nil, fmt.Errorf("failed to cancel signature: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode)
	}

	return mcp.NewToolResultText(fmt.Sprintf("Canceling signature %s. Reason: %s", signatureID, reason)), nil
}
