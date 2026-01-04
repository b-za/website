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
			Title:       "SA Tax Returns - Find an Accountant",
			Description: "Connect with verified tax practitioners and accountants for your tax returns.",
			Path:        "index.html",
			Sections: []Section{
				{
					TemplateName: "hero",
					Data: HeroData{
						Title:    "Simplify Your Tax Returns",
						Subtitle: "Find trusted accountants and tax practitioners in South Africa to help you file with confidence.",
					},
				},
				{
					TemplateName: "features",
					Data: FeaturesData{
						Title: "How It Works",
						Items: []FeatureItem{
							{Name: "Search Professionals", Description: "Browse our directory of verified tax practitioners by expertise and location."},
							{Name: "Compare Services", Description: "View profiles, services, and reviews to find the perfect match for your needs."},
							{Name: "Get Listed", Description: "Are you an accountant? List your practice today to reach thousands of potential clients."},
						},
					},
				},
			},
		},
		// 2. About Page
		{
			Title:       "About - SA Tax Returns",
			Description: "Our mission to simplify tax filing in South Africa.",
			Path:        "about/index.html",
			Sections: []Section{
				{
					TemplateName: "hero",
					Data: HeroData{
						Title:    "Bridging the Gap",
						Subtitle: "Connecting taxpayers with the right direct professional help.",
					},
				},
				{
					TemplateName: "text_block",
					Data: TextBlockData{
						Heading: "Our Mission",
						Paragraphs: []string{
							"Filing taxes can be daunting. SA Tax Returns was built to make professional tax assistance accessible to everyone. We believe that finding a qualified accountant should be as easy as searching for a restaurant.",
							"We verify every practitioner on our platform to ensure you get high-quality advice and service. Whether you are an individual needing help with eFiling or a business looking for comprehensive bookkeeping, we have the right professional for you.",
						},
					},
				},
			},
		},
		// 3. Contact Page
		{
			Title:       "Contact Us",
			Description: "Get in touch with the SA Tax Returns team.",
			Path:        "contact/index.html",
			Sections: []Section{
				{
					TemplateName: "hero",
					Data: HeroData{
						Title:    "Contact Support",
						Subtitle: "Need help with the platform? We're here to assist.",
					},
				},
				{
					TemplateName: "contact_form",
					Data: ContactFormData{
						Title:      "Send us a Message",
						ButtonText: "Submit Inquiry",
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
