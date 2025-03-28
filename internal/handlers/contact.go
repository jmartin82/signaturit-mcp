package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
)

type Contact struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

func (h *Handler) ListContacts(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	resp, err := h.client.Get("/contacts.json")
	if err != nil {
		return nil, fmt.Errorf("failed to list contacts: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var contacts []Contact
	if err := json.Unmarshal(body, &contacts); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	summary := "Contacts:\n"
	for _, contact := range contacts {
		summary += fmt.Sprintf("- %s (%s) [ID: %s]\n", contact.Name, contact.Email, contact.ID)
	}

	return mcp.NewToolResultText(summary), nil
}

func (h *Handler) GetContact(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	contactID := req.Params.Arguments["contact_id"].(string)

	resp, err := h.client.Get(fmt.Sprintf("/contacts/%s.json", contactID))
	if err != nil {
		return nil, fmt.Errorf("failed to get contact: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var contact Contact
	if err := json.Unmarshal(body, &contact); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return mcp.NewToolResultText(fmt.Sprintf("Contact: %s (%s) [ID: %s]", contact.Name, contact.Email, contact.ID)), nil
}

func (h *Handler) CreateContact(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	email := req.Params.Arguments["email"].(string)
	name := req.Params.Arguments["name"].(string)

	requestBody := map[string]interface{}{
		"email": email,
		"name":  name,
	}

	resp, err := h.client.Post("/contacts.json", requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create contact: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var contact Contact
	if err := json.Unmarshal(body, &contact); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return mcp.NewToolResultText(fmt.Sprintf("Contact created: %s (%s) [ID: %s]", contact.Name, contact.Email, contact.ID)), nil
}

func (h *Handler) UpdateContact(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	contactID := req.Params.Arguments["contact_id"].(string)

	requestBody := make(map[string]interface{})
	if email, ok := req.Params.Arguments["email"].(string); ok && email != "" {
		requestBody["email"] = email
	}
	if name, ok := req.Params.Arguments["name"].(string); ok && name != "" {
		requestBody["name"] = name
	}

	resp, err := h.client.Patch(fmt.Sprintf("/contacts/%s.json", contactID), requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to update contact: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var contact Contact
	if err := json.Unmarshal(body, &contact); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return mcp.NewToolResultText(fmt.Sprintf("Contact updated: %s (%s) [ID: %s]", contact.Name, contact.Email, contact.ID)), nil
}

func (h *Handler) DeleteContact(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	contactID := req.Params.Arguments["contact_id"].(string)

	resp, err := h.client.Delete(fmt.Sprintf("/contacts/%s.json", contactID))
	if err != nil {
		return nil, fmt.Errorf("failed to delete contact: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	return mcp.NewToolResultText(fmt.Sprintf("Contact %s successfully deleted", contactID)), nil
}
