package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

func main() {
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
