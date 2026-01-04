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
						Title:           "Professional Accountants & Consultants in Cape Town",
						Subtitle:        "Are you looking for professional accountants and tax consultants in Cape Town? Stop stressing about SARS. We handle your Personal Tax, VAT, Payroll, and CIPC compliance so you can focus on growing your business.",
						PrimaryBtn:      "Book Free Consultation",
						SecondaryBtn:    "View Our Services",
						BackgroundImage: "/assets/images/hero_background_capetown.png", // Assuming image saved here
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
		// --- Submissions Pages ---
		{
			Title: "Personal Tax Submissions", Description: "Expert assistance with your personal tax returns.", Path: "submissions/personal-tax/index.html",
			Sections: []Section{{TemplateName: "hero", Data: HeroData{Title: "Personal Tax Returns", Subtitle: "We ensure you are compliant and get the best possible refund.", PrimaryBtn: "Contact Us"}}},
		},
		{
			Title: "VAT Submissions", Description: "Value Added Tax returns and compliance.", Path: "submissions/vat/index.html",
			Sections: []Section{{TemplateName: "hero", Data: HeroData{Title: "VAT Submissions", Subtitle: "Accurate and timely VAT returns for your business.", PrimaryBtn: "Get Started"}}},
		},
		{
			Title: "Company Tax", Description: "Corporate tax filing services.", Path: "submissions/company-tax/index.html",
			Sections: []Section{{TemplateName: "hero", Data: HeroData{Title: "Company Tax Returns", Subtitle: "Professional corporate tax compliance and planning.", PrimaryBtn: "Consult Now"}}},
		},
		{
			Title: "PAYE Returns", Description: "PAYE and payroll tax submissions.", Path: "submissions/paye/index.html",
			Sections: []Section{{TemplateName: "hero", Data: HeroData{Title: "PAYE Returns", Subtitle: "Hassle-free payroll tax submissions for employers.", PrimaryBtn: "Learn More"}}},
		},

		// --- Registrations Pages ---
		{
			Title: "E-Filing Setup", Description: "Setup your SARS E-Filing profile.", Path: "registrations/efiling/index.html",
			Sections: []Section{{TemplateName: "hero", Data: HeroData{Title: "E-Filing Setup", Subtitle: "We get you registered and set up on SARS E-Filing correctly.", PrimaryBtn: "Register Now"}}},
		},
		{
			Title: "Company Tax Registration", Description: "Register your company for income tax.", Path: "registrations/company-tax/index.html",
			Sections: []Section{{TemplateName: "hero", Data: HeroData{Title: "Company Tax Registration", Subtitle: "Ensure your new business is tax compliant from day one.", PrimaryBtn: "Start Registration"}}},
		},
		{
			Title: "VAT Registration", Description: "VAT registration services.", Path: "registrations/vat/index.html",
			Sections: []Section{{TemplateName: "hero", Data: HeroData{Title: "VAT Registration", Subtitle: "Mandatory and voluntary VAT registration services.", PrimaryBtn: "Register for VAT"}}},
		},
		{
			Title: "PAYE Registration", Description: "Register for PAYE.", Path: "registrations/paye/index.html",
			Sections: []Section{{TemplateName: "hero", Data: HeroData{Title: "PAYE Registration", Subtitle: "Employer registration for Pay As You Earn tax.", PrimaryBtn: "Register Employer"}}},
		},
		{
			Title: "UIF Registration", Description: "Unemployment Insurance Fund registration.", Path: "registrations/uif/index.html",
			Sections: []Section{{TemplateName: "hero", Data: HeroData{Title: "UIF Registration", Subtitle: "Protect your employees with UIF registration.", PrimaryBtn: "Register UIF"}}},
		},
		{
			Title: "WCA Registration", Description: "Workmen's Compensation Assistance.", Path: "registrations/wca/index.html",
			Sections: []Section{{TemplateName: "hero", Data: HeroData{Title: "WCA Registration", Subtitle: "Workmen's Compensation compliance for your business.", PrimaryBtn: "Get Coverage"}}},
		},
		{
			Title: "CIPC Company Registration", Description: "New company registration.", Path: "registrations/new-company/index.html",
			Sections: []Section{{TemplateName: "hero", Data: HeroData{Title: "New Company (CIPC)", Subtitle: "Fast and efficient company registration services.", PrimaryBtn: "Register Company"}}},
		},
	}
}

// --- Section Data Structs ---

type HeroData struct {
	Title           string
	Subtitle        string
	PrimaryBtn      string
	SecondaryBtn    string
	BackgroundImage string
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
