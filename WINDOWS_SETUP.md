# Windows Setup Guide

This guide will help you set up the project on a Windows PC from scratch, assuming you have no developer tools installed.

## 1. Install Required Tools

### A. Install Git (For version control)
1. Go to [git-scm.com/download/win](https://git-scm.com/download/win).
2. Download **"64-bit Git for Windows Setup"**.
3. Run the installer. Keep clicking **Next** (the default settings are fine).
4. **Important**: When asked to choose the default editor, you can select Notepad if you don't have VS Code yet.

### B. Install Go (The language the site builder is written in)
1. Go to [go.dev/dl/](https://go.dev/dl/).
2. Download the **Microsoft Windows** installer (e.g., `go1.21.x.windows-amd64.msi`).
3. Run the installer and follow the instructions.
4. Verify: Open "Command Prompt" (search `cmd` in Start menu) and type: `go version`. You should see the version number.

### C. Install Node.js (For styling/Tailwind)
1. Go to [nodejs.org](https://nodejs.org/).
2. Download the **LTS (Long Term Support)** version.
3. Run the installer. Ensure "Add to PATH" is selected (it is by default).
4. Verify: In Command Prompt, type `npm -v`.

### D. Install Make (Optional but Recommended)
Standard Windows doesn't come with `make`. You have two options:
*   **Option 1 (Easier)**: Use the `git bash` terminal that came with Git (step A). It has basic unix commands.
*   **Option 2 (Chocolatey)**: If you are an admin, open shell and run: `choco install make`.

---

## 2. Setting Up the Project

1.  **Clone the Repository**:
    Open "Git Bash" (search in Start menu) or Command Prompt.
    ```bash
    git clone https://github.com/YOUR_USERNAME/sa-tax-returns.git
    cd sa-tax-returns
    ```

2.  **Install CSS Dependencies**:
    ```bash
    npm install
    # This downloads Tailwind CSS and other tools to the 'node_modules' folder.
    ```

---

## 3. Running the Website

### Method A: Using Make (Recommended with Git Bash)
If you are using **Git Bash**, simply run:
```bash
make dev
```
This will:
1.  Build the CSS.
2.  Start the Go web server.
3.  Watch for unrelated changes.
4.  Open your browser to `http://localhost:8080`.

### Method B: Manual Commands (Standard Command Prompt)
If `make` is not working, you can run the commands manually in two separate terminal windows:

**Window 1 (The Server):**
```cmd
go run cmd/builder/main.go cmd/builder/definitions.go --dev
```

**Window 2 (CSS Builder - run this when you change styles):**
```cmd
.\node_modules\.bin\tailwindcss -i ./styles/globals.css -o ./build/assets/css/style.css --minify
```

---

## Troubleshooting
*   **"command not found"**: Try restarting your computer to refresh the system PATH environment variables after installing Go/Node.
*   **Permission denied**: Run your terminal as Administrator.
