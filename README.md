# Signaturit MCP ✍️

> **Note:** This is an unofficial integration project and is not affiliated with, officially maintained, or endorsed by Signaturit.

This project is a demonstration of an MCP (Microservice Communication Protocol) server that integrates with Signaturit tools through their [public API](https://www.signaturit.com/api). It provides various functionalities for managing signature requests, including listing, creating, and handling signatures.

## Capabilities 🚀

The MCP server provides the following tools to interact with Signaturit:

- **get_signature** 📄: Retrieve details of a specific signature request using its ID
- **create_signature** ✨: Create new signature requests using templates
  - Support for multiple signers 👥
  - Email or SMS delivery 📧 📱
  - Customizable expiration time ⏰
  - Sequential or parallel signing workflow ⛓️
  - Custom email/SMS messages 💬
  - Webhook integration for real-time notifications 🔔
- **send_signature_reminder** 📬: Send reminder notifications to pending signers
- **cancel_signature** ❌: Cancel active signature requests with optional reason

## Project Structure 📁

- **cmd/server/main.go** 🎯: Entry point of the application. It initializes and starts the MCP server.
- **internal/app/server.go** ⚙️: Contains the logic for creating and configuring the MCP server, including registering signature tools and handlers.
- **internal/handlers/signature.go** 🛠️: Implements handler functions for various signature operations such as listing, retrieving, and managing signatures.
- **internal/tools/signature.go** 🔧: Registers signature-related tools with the MCP server.

## Configuration ⚙️

### API Authentication 🔐

This server integrates with the Signaturit API and requires an API key for authentication. You need to:

1. Create an account at [Signaturit](https://www.signaturit.com)
2. Get your API key from the Signaturit dashboard
3. Set the API key as an environment variable:

```bash
export SIGNATURIT_SECRET_TOKEN='your_api_key_here'
```

## Prerequisites 📋

1. **Go Installation** 
   - Go 1.16 or higher
   - Verify your installation:
   ```bash
   go version
   ```

2. **Signaturit Account** 
   - Active account at [Signaturit](https://www.signaturit.com)
   - Valid API key from the Signaturit dashboard

## Build 🔨

1. **Clone the repository**
   ```bash
   git clone https://github.com/yourusername/signaturtit_mcp.git
   cd signaturtit_mcp
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Build the application**
   ```bash
   # Build for your current platform
   go build -o bin/signaturtit_mcp cmd/server/main.go

   # Build for specific platform (e.g., Linux)
   GOOS=linux GOARCH=amd64 go build -o bin/signaturtit_mcp cmd/server/main.go
   ```

4. **Run the built binary**
   ```bash
   # Make sure you have set the required environment variables first
   export SIGNATURIT_SECRET_TOKEN='your_api_key_here'
   
   # Run the application
   ./bin/signaturtit_mcp
   ```

## License 📜

```
Copyright 2024 Jordi Martin

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
