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
			Title:       "Personal Tax Services - Expert Individual Tax Filing | SA Tax Returns",
			Description: "Professional assistance with your Personal Income Tax (ITR12) submissions. Ensure full SARS compliance and maximize your refund with our registered tax practitioners.",
			Path:        "submissions/personal-tax/index.html",
			Sections: []Section{
				{TemplateName: "hero", Data: HeroData{Title: "Personal Tax Returns (ITR12)", Subtitle: "Simpify your personal tax filing season. We help salary earners, commission earners, and freelancers submit accurate returns on time.", PrimaryBtn: "File My Return"}},
				{TemplateName: "text_block", Data: TextBlockData{
					Heading: "Takes the Stress Out of Tax Season",
					Paragraphs: []string{
						"Submitting your Personal Income Tax Return (ITR12) doesn't have to be a headache. Whether you are a standard salary earner or have complex income streams like rental property or capital gains, our team ensures every deduction is claimed correctly.",
						"We review your auto-assessment, check for missing medical aid credits, verifying retirement annuity contributions, and ensure your home office expenses are valid before submission. Don't leave money on the table or risk an audit.",
					},
				}},
			},
		},
		{
			Title:       "VAT Returns & Submissions services | SA Tax Returns",
			Description: "Timely and accurate VAT201 submissions for South African businesses. We handle output/input tax calculations to keep you penalty-free.",
			Path:        "submissions/vat/index.html",
			Sections: []Section{
				{TemplateName: "hero", Data: HeroData{Title: "VAT Submissions (VAT201)", Subtitle: "Ensure your Value Added Tax returns are accurate and submitted on time, every billing period.", PrimaryBtn: "Get VAT Help"}},
				{TemplateName: "text_block", Data: TextBlockData{
					Heading: "Reliable VAT Compliance",
					Paragraphs: []string{
						"VAT vendors are under constant scrutiny from SARS. Late payments or incorrect calculations can lead to 10% penalties and interest. We act as your safety net.",
						"Our accountants review your invoices, verify input tax claims, and prepare your VAT201 return. We also handle the verification process if SARS audits a refund, preparing the necessary schedule of documents so you get your cash flow back sooner.",
					},
				}},
			},
		},
		{
			Title:       "Company Tax Return (ITR14) Services | SA Tax Returns",
			Description: "Comprehensive Corporate Income Tax (CIT) filing for Pty Ltds and Close Corporations. Optimize your tax position with expert advice.",
			Path:        "submissions/company-tax/index.html",
			Sections: []Section{
				{TemplateName: "hero", Data: HeroData{Title: "Company Tax Returns (ITR14)", Subtitle: "Expert corporate tax compliance and planning for growing businesses.", PrimaryBtn: "Consult Now"}},
				{TemplateName: "text_block", Data: TextBlockData{
					Heading: "Corporate Tax Done Right",
					Paragraphs: []string{
						"Every registered company in South Africa must file an Annual Income Tax Return (ITR14), even if it didn't trade. The requirements can be complex, involving balance sheets, income statements, and specific tax adjustments.",
						"We prepare your Annual Financial Statements (AFS) where required and ensure your ITR14 accurately reflects your financial position. We identify allowances like s12E (Small Business Corporation) validation to legally minimize your tax liability.",
					},
				}},
			},
		},
		{
			Title:       "PAYE & EMP201 Submissions | SA Tax Returns",
			Description: "Monthly payroll tax submissions (PAYE, SDL, UIF) for employers. Avoid the 10% late payment penalty with our automated services.",
			Path:        "submissions/paye/index.html",
			Sections: []Section{
				{TemplateName: "hero", Data: HeroData{Title: "PAYE Returns (EMP201)", Subtitle: "Hassle-free monthly payroll tax submissions for employers.", PrimaryBtn: "Manage My Payroll"}},
				{TemplateName: "text_block", Data: TextBlockData{
					Heading: "Simplified Payroll Compliance",
					Paragraphs: []string{
						"As an employer, you are an agent for SARS, collecting PAYE, SDL, and UIF. Missing the payment deadline of the 7th of each month results in automatic penalties.",
						"We manage your monthly EMP201 declarations, ensuring that all employee tax certificates (IRP5s) reconcile correctly at the end of the year (EMP501). Focus on your team, while we handle the tax authorities.",
					},
				}},
			},
		},

		// --- Registrations Pages ---
		{
			Title:       "SARS E-Filing Registration & Profile Setup | SA Tax Returns",
			Description: "Need help getting on SARS E-Filing? We register profiles, fix login issues, and merge existing tax types.",
			Path:        "registrations/efiling/index.html",
			Sections: []Section{
				{TemplateName: "hero", Data: HeroData{Title: "E-Filing Setup & Support", Subtitle: "We get you registered and set up on SARS E-Filing correctly the first time.", PrimaryBtn: "Activate Profile"}},
				{TemplateName: "text_block", Data: TextBlockData{
					Heading: "Get Connected to SARS",
					Paragraphs: []string{
						"E-Filing is the gateway to your tax compliance. We assist individuals and companies with registering new profiles, recovering lost login credentials, and updating banking details.",
						"One of the most common issues is 'broken' profiles where tax types (like VAT or PAYE) aren't linked. Our practitioners have a dedicated channel to resolve these profile issues quickly.",
					},
				}},
			},
		},
		{
			Title:       "Company Tax Registration (Income Tax) | SA Tax Returns",
			Description: "Register your new CIPC company for Income Tax with SARS. Obtain your Tax Reference Number quickly.",
			Path:        "registrations/company-tax/index.html",
			Sections: []Section{
				{TemplateName: "hero", Data: HeroData{Title: "Company Income Tax Registration", Subtitle: "Ensure your new business complies with the Tax Administration Act.", PrimaryBtn: "Register Company"}},
				{TemplateName: "text_block", Data: TextBlockData{
					Heading: "Mandatory Tax Registration",
					Paragraphs: []string{
						"All South African companies registered with CIPC are automatically issued an Income Tax reference number, but this needs to be verified and activated on E-Filing.",
						"We ensure your Representative Vendor is appointed correctly at SARS, which is now a strict requirement before you can file any returns. Validating the Public Officer is the critical first step to corporate compliance.",
					},
				}},
			},
		},
		{
			Title:       "VAT Registration Services (Voluntary & Mandatory) | SA Tax Returns",
			Description: "Fast VAT registration for South African companies. Expert assistance with the RAV01 form and interview process.",
			Path:        "registrations/vat/index.html",
			Sections: []Section{
				{TemplateName: "hero", Data: HeroData{Title: "VAT Registration", Subtitle: "Navigate the complex SARS VAT registration process with expert guidance.", PrimaryBtn: "Register for VAT"}},
				{TemplateName: "text_block", Data: TextBlockData{
					Heading: "Understanding VAT Registration",
					Paragraphs: []string{
						"You must register for VAT if your turnover exceeds R1 million in a 12-month period. You may voluntarily register if your income exceeds R50,000.",
						"SARS strict verification process often result in rejections. We prepare your invoices, bank statements, and business contracts to meet the specific requirements of the RAV01 form, ensuring your application is approved without delay.",
					},
				}},
			},
		},
		{
			Title:       "PAYE Employer Registration (EMP101) | SA Tax Returns",
			Description: "Register as an employer with SARS. PAYE, SDL, and UIF registration in one smooth process.",
			Path:        "registrations/paye/index.html",
			Sections: []Section{
				{TemplateName: "hero", Data: HeroData{Title: "PAYE Employer Registration", Subtitle: "Hiring your first employee? You need to register for PAYE within 21 days.", PrimaryBtn: "Register Employer"}},
				{TemplateName: "text_block", Data: TextBlockData{
					Heading: "Becoming an Employer",
					Paragraphs: []string{
						"If you pay remuneration to any employee (including directors) exceeding the tax threshold, you must register for Pay As You Earn (PAYE).",
						"This registration (EMP101) covers three levies: PAYE (Tax), SDL (Skills Development Levy - for payrolls > R500k/yr), and UIF (Unemployment Insurance). We set this up correctly so your monthly payroll runs smoothly.",
					},
				}},
			},
		},
		{
			Title:       "UIF Registration (Dept of Labour) | SA Tax Returns",
			Description: "Register your domestic or commercial workers for UIF. Protect your staff and comply with labour laws.",
			Path:        "registrations/uif/index.html",
			Sections: []Section{
				{TemplateName: "hero", Data: HeroData{Title: "UIF Registration", Subtitle: "Department of Labour registration for all employers.", PrimaryBtn: "Register UIF"}},
				{TemplateName: "text_block", Data: TextBlockData{
					Heading: "Protecting Your Workforce",
					Paragraphs: []string{
						"While SARS collects UIF contributions, you must also be registered with the Department of Labour via the uFiling system to declare employee details.",
						"This is crucial: if your employees are not registered here, they cannot claim benefits if they are retrenched or take maternity leave. We handle both the SARS and uFiling sides of the registration.",
					},
				}},
			},
		},
		{
			Title:       "WCA Registration (COIDA) | SA Tax Returns",
			Description: "Workmen's Compensation (COIDA) registration and Letter of Good Standing.",
			Path:        "registrations/wca/index.html",
			Sections: []Section{
				{TemplateName: "hero", Data: HeroData{Title: "WCA / COIDA Registration", Subtitle: "Workmen's Compensation is mandatory for any business with employees.", PrimaryBtn: "Get Coverage"}},
				{TemplateName: "text_block", Data: TextBlockData{
					Heading: "Injury on Duty Protection",
					Paragraphs: []string{
						"The Compensation for Occupational Injuries and Diseases Act (COIDA) protects employers from civil claims if an employee is injured on duty.",
						"To bid for most tenders and contracts, you need a Letter of Good Standing. We register you with the Compensation Fund and help you submit your Return of Earnings to obtain your letter.",
					},
				}},
			},
		},
		{
			Title:       "CIPC New Company Registration | SA Tax Returns",
			Description: "Register a Pty Ltd company in South Africa. Includes Name Reservation and Share Certificates.",
			Path:        "registrations/new-company/index.html",
			Sections: []Section{
				{TemplateName: "hero", Data: HeroData{Title: "New Company (CIPC)", Subtitle: "Start your business journey with a formally registered Pty Ltd.", PrimaryBtn: "Register Company"}},
				{TemplateName: "text_block", Data: TextBlockData{
					Heading: "Start Your Business",
					Paragraphs: []string{
						"We offer a complete company registration package. This includes reserving your unique name, lodging the MoI (Memorandum of Incorporation), and issuing the COR14.3 registration certificate.",
						"We also issue your Share Certificates (required by banks) and assist with the initial B-BBEE affidavit. Get your business trading legally in as little as 3-5 days.",
					},
				}},
			},
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
