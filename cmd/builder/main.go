package main

import (
	"bytes"
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

	for _, f := range files {
		info, err := os.Stat(f)
		if err == nil && info.ModTime().After(lastBuild) {
			return true
		}
	}
	return false
}

func rebuildCSS() {
	// Ensure directory exists
	os.MkdirAll("build/assets/css", 0755)
	cmd := exec.Command("./node_modules/.bin/tailwindcss", "-i", "./styles/globals.css", "-o", "build/assets/css/style.css", "--minify")
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
	var tmpl *template.Template

	funcMap := template.FuncMap{
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
		"section": func(s Section) (template.HTML, error) {
			// This function allows dynamic dispatch: {{ section . }}
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

	tmpl = template.New("").Funcs(funcMap)

	layoutFiles, err := filepath.Glob("layouts/*.html")
	if err != nil {
		log.Fatal(err)
	}
	componentFiles, err := filepath.Glob("layouts/components/*.html")
	if err != nil {
		log.Fatal(err)
	}
	allFiles := append(layoutFiles, componentFiles...)

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

		err = tmpl.ExecuteTemplate(f, "base.html", page)
		if err != nil {
			log.Fatalf("Error executing template for %s: %v", page.Path, err)
		}
		f.Close()
	}

	fmt.Println("Done.")
}
