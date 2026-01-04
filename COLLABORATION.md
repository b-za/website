# AGENT COLLABORATION GUIDE

This document defines the working relationship, protocols, and architectural philosophy for this project (`SA Tax Returns`). It is the Single Source of Truth for any AI agent interacting with this repository.

## 1. Core Philosophy
- **User Intent > Implementation Details**: The user provides high-level intent, copy, or screenshots. The agent translates this into Go code, Tailwind classes, and HTML templates.
- **Git is the Undo Button**: Every meaningful set of changes must be committed immediately. We lean on `git reset` rather than trying to "fix forward" blindly.
- **Tailwind Aesthetic**: We prioritize modern, clean, "glassmorphic" and high-contrast designs. (e.g., dark text on white, white text on dark/gradient backgrounds). avoiding visual clutter.

## 2. workflow: "Change Sessions"
We work in distinct "Change Sessions" to organize large blocks of work.

### Lifecycle of a Session
1.  **Initiation**:
    -   When starting a significant feature, a day's work, or a major refactor, the agent must propose starting a session.
    -   **Action**: `git checkout -b feature/name-of-session` (or `daily/YYYY-MM-DD` if general work).
    - **CRITICAL RULE**: NEVER make changes directly on the `main` branch. Even for small fixes, create a temporary branch.
2.  **Execution**:
    -   Perform tasks, edit code, and verify using `make build` / `make dev`.
    -   **Commit often**: `git commit -m "..."` for every logical step.
3.  **Completion**:
    -   When the goal is met, ask the user: "Session goal complete. Ready to merge?"
    -   **Action**: `git checkout main && git merge feature/... && git branch -d feature/...`

*Note: If no specific feature is being built, default to a daily session branch (e.g. `daily/2026-01-04`).*

## 3. Architecture Overview
-   **Builder**: A custom static site generator in `cmd/builder/main.go`.
-   **Content**: Defined as Go structs in `cmd/builder/definitions.go`. **This is the CMS.**
-   **Templates**: Located in `components/`.
    -   `components/layouts/`: Base HTML wrappers (e.g., `base.html`).
    -   `components/common/`: Global UI (Header, Footer).
    -   `components/sections/`: Reusable content blocks (Hero, Features, Forms).
-   **Styling**:
    -   Tailwind CSS v3.
    -   Config: `tailwind.config.js`.
    -   Source: `styles/globals.css`.
    -   Output: `build/assets/css/style.css`.

## 4. Development Protocols
-   **Running the Site**: `make dev` starts a watcher and server on port 8080.
-   **Adding Pages**:
    1.  Add a `Page` struct to `GetSiteContent()` in `definitions.go`.
    2.  Use existing `Section` templates or create new ones in `components/sections/`.
-   **Adding Components**:
    1.  Create `components/sections/my_component.html`.
    2.  Define its data struct in `definitions.go`.
    3.  Register it (indirectly via the `template` function lookup).

## 5. Interaction Style
-   **Proactive Suggestions**: If you see the user struggling with a design or content idea, propose a concrete solution (e.g., "I can build a pricing table for that").
-   **Updating this Guide**: As you learn more about the user's preferences (e.g., "User hates rounded corners"), **you must suggest updating this document**. Ask: "I noticed you prefer X. Shall I update the Collaboration Guide?"

## 6. Known Preferences
-   **Design**: Dark mode for HERO sections, Clean/Light mode for content. Transparent headers overlapping hero sections.
-   **Navigation**: Simple links, no heavy buttons in header.
-   **Icons**: Ensure they are sized correctly (e.g., `w-4 h-4`) and not massive mobile-sized icons on desktop.

---
*Last Updated: 2026-01-04*
