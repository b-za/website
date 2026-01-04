# Repository Instructions

This document is a guide for you (the AI agent) and future contributors on how to work with this repository.

## Core Philosophy
- **Git-Centric Undo**: Every meaningful change MUST be a separate commit. This serves as our "undo" stack.
- **Go + Templates**: We use Go's `html/template` for logic and rendering. Content is defined in Go structs.
- **Tailwind Only**: Styling is exclusively handled via Tailwind CSS classes. No ad-hoc CSS unless strictly necessary for things Tailwind can't touch.

## Workflow Status
1. **Making Changes**:
   - MODIFY files.
   - RUN `make css` if styling changed.
   - RUN `make build` to verify the site generates.
   - RUN `git add .` and `git commit -m "Description"` IMMEDIATELY.

2. **Creating New Pages**:
   - Define a new `Page` struct instance in `content.go` (or `main.go`).
   - If reusing existing sections, just configure them.
   - If a new layout is needed, create it in `layouts/components/`.

3. **Restoring Previous State**:
   - If the user isn't happy with the latest change, `git reset --hard HEAD~1` is the preferred method over trying to "fix" it blindly.

## Tech Stack Details
- **Build Tool**: `make`
- **CSS**: Tailwind CLI (via `Makefile`)
- **Server**: `go run main.go` (or `air` for hot reload)

## Directory Structure
- `cmd/builder`: The static site generator logic.
- `layouts`: The HTML templates.
- `styles/globals.css`: The source of truth for Tailwind.
- `definitions.go`: Where content and page structure lives (The "CMS").
