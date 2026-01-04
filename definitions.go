package main

// Page represents a single page on the website.
type Page struct {
	Title       string
	Description string
	Path        string // e.g., "index.html" or "about/index.html"
	Sections    []Section
}

// Section represents a reusable content block.
// The TemplateName must match a defined template name (e.g., "hero", "features").
// Data is passed to that template.
type Section struct {
	TemplateName string
	Data         interface{}
}

// SiteContent defines all the pages in the site.
// This function is called by the builder to get the list of pages to generate.
func GetSiteContent() []Page {
	return []Page{
		{
			Title:       "Home",
			Description: "Welcome to my new static site.",
			Path:        "index.html",
			Sections: []Section{
				{
					TemplateName: "hero",
					Data: HeroData{
						Title:    "Welcome to My Awesome Site",
						Subtitle: "Built with Go and Tailwind CSS",
					},
				},
			},
		},
	}
}

// --- Section Data Structs ---

type HeroData struct {
	Title    string
	Subtitle string
}
