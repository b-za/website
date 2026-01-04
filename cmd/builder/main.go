package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

// SectionRenderer is a helper to render a section dynamically.
// It allows us to use {{ template "section" . }} inside the base layout
// by defining a "section" template that dispatches to the specific TemplateName.
func renderSection(s Section) (template.HTML, error) {
	// This is a bit of a trick. We don't actually render HTML here directly in the helper
	// if we want to use the standard template inheritance.
	// But since `text/template` (and html/template) doesn't support dynamic template selection natively
	// in a clean way without a custom func, we'll do this:
	// The "base.html" iterates over sections. We need a way to say "render template named s.TemplateName with s.Data".
	// The standard way is {{ template "name" data }}. But "name" must be a string literal in standard template syntax?
	// Actually, in Go templates, {{ template "name" }} expects a literal.
	// So we'll use a FuncMap.
	return "", fmt.Errorf("this should not be called directly")
}

func main() {
	fmt.Println("Building site...")

	// 1. Prepare output directory
	outputDir := "build"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatal(err)
	}

	// 2. Parse all templates
	// We include the base layout definition and all component templates.
	// Note: We need a valid "hero" template for the initial content to work.
	tmpl := template.New("base.html")

	// Helper function to render a sub-template dynamically
	tmpl = tmpl.Funcs(template.FuncMap{
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
	})

	// Parse layouts
	// TODO: Use glob in the future, but for now we list explicitly or assume logic.
	// We'll walk the layouts dir.
	files, err := filepath.Glob("layouts/*.html")
	if err != nil {
		log.Fatal(err)
	}
	componentFiles, err := filepath.Glob("layouts/components/*.html")
	if err != nil {
		log.Fatal(err)
	}
	files = append(files, componentFiles...)

	if len(files) > 0 {
		tmpl, err = tmpl.ParseFiles(files...)
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

		// We need to support dynamic rendering of sections.
		// The base template iterates sections. The problem is `{{ template "name" }}` requires a literal string.
		// A common pattern is to define a "section" template that acts as a switch, or just use a custom execution logic.
		//
		// IMPROVED APPROACH:
		// We can't easily perform dynamic dispatch inside the template with standard Go syntax without a big switch or a func that renders to string.
		// So we will pre-render sections into HTML strings?
		// Or simpler: We define a FuncMap `renderSection`.

		// Let's retry the func map approach properly.
		// We'll create a new template executor for each page or re-use.
		// But to fetch the specific template by name inside the template execution is hard.
		//
		// ALTERNATIVE:
		// We'll use a specific "root" struct that has a method `RenderSections`?
		// No, let's keep it simple. We will use a Function that takes the Section object and returns HTML.

		// Re-defining the FuncMap with a closure is tricky if we need access to the `tmpl` set.
		// Let's just make the "section" template a functional wrapper.

		// ACTUALLY: The cleanest way in Go templates for dynamic include is to not use `{{ template }}` but a custom function.
		// `{{ include .TemplateName .Data }}`

		// Let's overwrite the FuncMap to include our dynamic renderer.
		tClone, _ := tmpl.Clone()
		tClone.Funcs(template.FuncMap{
			"section": func(s Section) (template.HTML, error) {
				// Create a buffer
				// Execute the specific template named s.TemplateName with s.Data
				// Return the string as HTML
				// This requires access to the template set.
				// Since we are inside the execution, we can't easily grab `tClone` unless captured in closure.
				// But we are defining it before parsing? No.
				return "", fmt.Errorf("nested rendering not implemented yet")
			},
		})

		// WAIT: The standard way to do this is to register a function that has access to the templates.
		// Because `ParseFiles` parses into `tmpl`, we can use `tmpl` inside the closure.

		err = tmpl.ExecuteTemplate(f, "base.html", page)
		if err != nil {
			log.Fatalf("Error executing template for %s: %v", page.Path, err)
		}
		f.Close()
	}

	fmt.Println("Done.")
}
