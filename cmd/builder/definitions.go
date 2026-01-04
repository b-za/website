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
func GetSiteContent() []Page {
	return []Page{
		// 1. Home Page
		{
			Title:       "Home",
			Description: "Welcome to our digital agency.",
			Path:        "index.html",
			Sections: []Section{
				{
					TemplateName: "hero",
					Data: HeroData{
						Title:    "Build Faster with Go & Tailwind",
						Subtitle: "A modern static site generator architecture for power users.",
					},
				},
				{
					TemplateName: "features",
					Data: FeaturesData{
						Title: "Why Choose Us?",
						Items: []FeatureItem{
							{Name: "Blazing Fast", Description: "Static builds entail zero server-side rendering time."},
							{Name: "Type Safe", Description: "Content defined in Go structs ensures data consistency."},
							{Name: "Modern Styles", Description: "Utility-first CSS with Tailwind for rapid UI development."},
						},
					},
				},
			},
		},
		// 2. About Page
		{
			Title:       "About Us",
			Description: "Learn more about our team and mission.",
			Path:        "about/index.html",
			Sections: []Section{
				{
					TemplateName: "hero",
					Data: HeroData{
						Title:    "Our Story",
						Subtitle: "Dedicated to simplicity and performance.",
					},
				},
				{
					TemplateName: "text_block",
					Data: TextBlockData{
						Heading: "Who We Are",
						Paragraphs: []string{
							"We are a small team of passionate developers who believe that web development has become too complex. Our mission is to strip back the layers of abstraction and return to the roots of fast, static HTML.",
							"Founded in 2024, we noticed that while tools got more powerful, they also got slower to maintain. We chose Go for its simplicity and speed, and Tailwind for its flexibility.",
						},
					},
				},
			},
		},
		// 3. Contact Page
		{
			Title:       "Contact",
			Description: "Get in touch with us.",
			Path:        "contact/index.html",
			Sections: []Section{
				{
					TemplateName: "hero",
					Data: HeroData{
						Title:    "Get In Touch",
						Subtitle: "We'd love to hear from you.",
					},
				},
				{
					TemplateName: "contact_form",
					Data: ContactFormData{
						Title:      "Send us a Message",
						ButtonText: "Send Message",
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

type TextBlockData struct {
	Heading    string
	Paragraphs []string
}

type FeaturesData struct {
	Title string
	Items []FeatureItem
}

type FeatureItem struct {
	Name        string
	Description string
}

type ContactFormData struct {
	Title      string
	ButtonText string
}
