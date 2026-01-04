# SA Tax Returns - Accountant Directory

A static site generator and directory for South African tax practitioners, built with Go and Tailwind CSS.

> **🤖 AI AGENTS & COLLABORATORS: READ THIS FIRST**
> 
> Before making any changes, you **MUST** read [COLLABORATION.md](./COLLABORATION.md).
> It contains the "Change Session" workflow, architectural rules, and design philosophy that you are required to follow.

## Quick Start

### Prerequisites
- Go 1.21+
- Node.js & npm

### Development
1. Install dependencies:
   ```bash
   npm install
   ```
2. Start the dev server (builds CSS + HTML and watches for changes):
   ```bash
   npm run dev
   ```
3. Visit `http://localhost:8080`

## Architecture
- **Content**: `cmd/builder/definitions.go` (Type-safe CMS)
- **Builder**: `cmd/builder/main.go`
- **Templates**: `components/**`

For detailed Windows setup instructions, see [WINDOWS_SETUP.md](./WINDOWS_SETUP.md).
