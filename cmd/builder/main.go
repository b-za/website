package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	devMode := flag.Bool("dev", false, "Run in development mode (watch and serve)")
	flag.Parse()

	if *devMode {
		runDevMode()
		return
	}

	build()
}

func runDevMode() {
	fmt.Println("Starting development server at http://localhost:8080")
	
	// Initial build
	build()

	// Start server
	go func() {
		fs := http.FileServer(http.Dir("build"))
		http.Handle("/", fs)
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	// Polling watcher
	ticker := time.NewTicker(500 * time.Millisecond)
	lastBuild := time.Now()

	for range ticker.C {
		if needsRebuild(lastBuild) {
			fmt.Println("Change detected. Rebuilding...")
			rebuildCSS() // Build CSS first
			build()      // Then Build HTML
			lastBuild = time.Now()
		}
	}
}

func needsRebuild(lastBuild time.Time) bool {
	// Check layouts
	files, _ := filepath.Glob("layouts/**/*.html")
	moreFiles, _ := filepath.Glob("layouts/*.html")
	files = append(files, moreFiles...)
	
	// Check definitions (definitions.go is in cmd/builder/definitions.go, 
	// but we might want to check the one we are running? 
	// Actually `go run` compiles the binary, so we can't easily "hot reload" Go code changes 
	// without restarting the process.
	// But `html` templates we can reload.
	// `definitions.go` changes require restarting the `go run` process unless we interpret it.
	// Since we compile the structs into the binary, we MUST restart the binary to pick up struct changes.
	// 
	// LIMITATION: This watcher only reloads TEMPLATE changes effectively if we re-parse.
	// But `build()` function re-parses everything.
	// HOWEVER: If `pages` comes from `GetSiteContent()` which is compiled code, 
	// we won't see changes in `definitions.go`.
	
	// For now, let's just watch templates.
	// To support `definitions.go` reloading, users should use `air` or we need a wrapper.
	// We will assume `make dev` usually handles Go reloading if using `air`.
	// Our internal watcher is good for templates.
	
	for _, f := range files {
		info, err := os.Stat(f)
		if err == nil && info.ModTime().After(lastBuild) {
			return true
		}
	}
	return false
}

func rebuildCSS() {
	cmd := exec.Command("./node_modules/.bin/tailwindcss", "-i", "./styles/globals.css", "-o", "./assets/css/style.css", "--minify")
	if err := cmd.Run(); err != nil {
		fmt.Println("Error rebuilding CSS:", err)
	}
}

func build() {
	fmt.Println("Building site...")

	// 1. Prepare output directory
	outputDir := "build"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatal(err)
	}

	// 2. Parse all templates
	// We parse all templates into a single set, but we start with a "root" that has the FuncMap.
	// We need the FuncMap to be available to ALL templates.

	// We initialize the variable first so the closure can capture it.
	var tmpl *template.Template

	funcMap := template.FuncMap{
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
		"section": func(s Section) (template.HTML, error) {
			// This function allows dynamic dispatch: {{ section . }}
			// It looks up the template by name (s.TemplateName) and executes it with s.Data.
			if tmpl == nil {
				return "", fmt.Errorf("template set is nil")
			}
			var buf bytes.Buffer
			if err := tmpl.ExecuteTemplate(&buf, s.TemplateName, s.Data); err != nil {
				return "", err
			}
			return template.HTML(buf.String()), nil
		},
	}

	// Create the root template with functions
	tmpl = template.New("").Funcs(funcMap)

	// Glob patterns
	layoutFiles, err := filepath.Glob("layouts/*.html")
	if err != nil {
		log.Fatal(err)
	}
	componentFiles, err := filepath.Glob("layouts/components/*.html")
	if err != nil {
		log.Fatal(err)
	}

	// Combine all files
	allFiles := append(layoutFiles, componentFiles...)

	// Parse them into the set
	if len(allFiles) > 0 {
		_, err = tmpl.ParseFiles(allFiles...)
		if err != nil {
			log.Fatalf("Error parsing templates: %v", err)
		}
	}

	// 3. Get Content
	pages := GetSiteContent()

	// 4. Generate Pages
	for _, page := range pages {
		fmt.Printf("Generating %s...\n", page.Path)

		outputPath := filepath.Join(outputDir, page.Path)
		if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
			log.Fatal(err)
		}

		f, err := os.Create(outputPath)
		if err != nil {
			log.Fatal(err)
		}

		// Execute "base.html" which renders the page layout
		err = tmpl.ExecuteTemplate(f, "base.html", page)
		if err != nil {
			log.Fatalf("Error executing template for %s: %v", page.Path, err)
		}
		f.Close()
	}

	fmt.Println("Done.")
}
