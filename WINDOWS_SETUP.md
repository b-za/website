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

Since you have Node.js installed, the easiest way to run the site is using **npm scripts**. This works in any terminal (Command Prompt, PowerShell, Git Bash).

### Development Mode (Recommended)
Run this command to start the server and watch for changes:
```bash
npm run dev
```
1.  It will build the site interactively.
2.  Open `http://localhost:8080`.
3.  Edit files in `cmd/builder/definitions.go` or `components/`, and the site will auto-update.

### Production Build
To just build the files without starting a server:
```bash
npm run build
```

### Just CSS
If you only want to regenerate the CSS:
```bash
npm run css
```

---

## Troubleshooting
*   **"command not found"**: Try restarting your computer to refresh the system PATH environment variables after installing Go/Node.
*   **Permission denied**: Run your terminal as Administrator.
