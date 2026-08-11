package content

// Settings holds global site configuration.
type Settings struct {
	Title       string
	Tagline     string
	Description string
	Email       string
	Instagram   string
	YouTube     string
	Logo        string
	FooterNote  string
}

// Site is the global site configuration.
var Site = Settings{
	Title:       "4Life Community",
	Tagline:     "Regenerative Strategy, AstroYoga & Bioregional Alchemy",
	Description: "April Bartlett's 4Life bridges grant writing, bioregional mapping, and AstroYoga—helping changemakers align strategy, soul, and planetary regeneration.",
	Email:       "4life.AprilBartlett@gmail.com",
	Instagram:   "https://www.instagram.com/amber.rising4love/",
	YouTube:     "https://www.youtube.com/@4life.AprilBartlett",
	Logo:        "/img/logo-light.png",
	FooterNote:  "Design: Alcyone Reserve",
}

// Hero is the content for the home page hero.
var Hero = struct {
	Eyebrow      string
	Headline     string
	Body         string
	Image        string
	ImageAlt     string
	Primary      string
	PrimaryURL   string
	Secondary    string
	SecondaryURL string
}{
	Eyebrow:      "4LIFE COMMUNITY",
	Headline:     "Funding & Strategy for Changemakers, 4Life",
	Body:         "Grant writing, systems design, and trauma-informed guidance for people building a better world. Rooted in justice, land wisdom, and results.",
	Image:        "/img/april-cutout-home.webp",
	ImageAlt:     "April Bartlett",
	Primary:      "Meet April",
	PrimaryURL:   "/about",
	Secondary:    "Read the Blog",
	SecondaryURL: "/blog",
}

// AboutSection holds the about page content.
type AboutSection struct {
	Eyebrow  string
	Headline string
	Body     string
	Image    string
	ImageAlt string
	Points   []AboutPoint
}

type AboutPoint struct {
	Title string
	Body  string
}

// About is the content for the about page.
var About = AboutSection{
	Eyebrow:  "MEET APRIL",
	Headline: "April Bartlett invites you into a world where personal healing and planetary restoration go hand in hand.",
	Body:     "Rooted in a lifetime of resilience, service, and deep listening, 4Life is a living portal for April's integrative work: trauma-informed yoga, holistic health, bioregional organizing, and community regeneration.\n\nApril's journey began in childhood, foraging with her parents and grandparents—learning firsthand that the Earth provides everything we need to thrive. By the age of 17, she was traveling across the U.S., training in herbal nutrition, enzymatic cleaning products, and water filtration. The birth of her third child, who faced life-threatening food and chemical sensitivities, deepened her calling and commitment to healing through ecological wisdom and regenerative practices.\n\nFaced with a choice between isolation or empowerment, April sought out natural healing alternatives, eventually learning from two Hopi Medicine Women. This opened her eyes to how sensitive systems are often excluded from social thinking—and inspired her to create Community 4Life Hub, a space for connection, healing, and low-histamine nourishment. The physical location has since closed. But the community is still connected and moving into the larger field.",
	Image:    "/img/april-portrait.webp",
	ImageAlt: "April Bartlett portrait",
	Points: []AboutPoint{
		{
			Title: "Grant Writing & Funding Strategy",
			Body:  "With years of experience and millions of dollars in secured funding, I help regenerative projects, movements, and organizations secure the resources to thrive.",
		},
		{
			Title: "Trauma-Informed AstroYoga",
			Body:  "Certified in trauma-informed yoga since 2009, I create safe, somatic spaces for healing and reconnection—where body and land remember each other.",
		},
		{
			Title: "Bioregional Weaving",
			Body:  "Community mapping with the land, not just on it. Workshops and systems design for ecological and cultural resilience, from the Salish Sea to Colombia.",
		},
	},
}

// Project is a single featured project.
type Project struct {
	Slug     string
	Title    string
	Tagline  string
	Summary  string
	Image    string
	ImageAlt string
	Status   string
}

// Projects lists featured projects.
var Projects = []Project{
	{
		Slug:     "susurros-del-agua",
		Title:    "Susurros del Agua",
		Tagline:  "Watershed Restoration & Cultural Resilience",
		Summary:  "A grassroots bioregional movement in Mogotes, Colombia co-led by April and Ricardo Palomino. We walk with the Río Mogoticos—mapping, storytelling, arts, and education in service of protecting and revitalizing the watershed.",
		Image:    "/img/susurros-hillside.jpg",
		ImageAlt: "Lush forested hillside above the Mogoticos river",
		Status:   "Active",
	},
	{
		Slug:     "bioregional-nomads",
		Title:    "Bioregional Nomads",
		Tagline:  "Mapping the Stories Land Tells",
		Summary:  "Open-source tools and workshops for regenerative placemaking—guiding communities to redraw boundaries based on watersheds, not wars, and weave sustainable initiatives together for impact.",
		Image:    "/img/bioregional-nomad.jpg",
		ImageAlt: "April's collaborator among green plants and trees",
		Status:   "Active",
	},
	{
		Slug:     "giving-catatumbo-river-a-second-chance",
		Title:    "Catatumbo River",
		Tagline:  "Save a River, Save a Community",
		Summary:  "Giving the Catatumbo River a second chance. Support the fight for justice for the communities who defend it—documentary, legal action, and long-term protections for the Water Guardians.",
		Image:    "/img/colombia-misty.jpg",
		ImageAlt: "Misty mountains of Colombia",
		Status:   "Fundraising",
	},
	{
		Slug:     "adaptogenic-healing",
		Title:    "Adaptogenic Healing",
		Tagline:  "Where Your Body's Wisdom Meets Earth's Medicine",
		Summary:  "Trauma-informed somatic healing, Quantum Touch energy work, and folk herbalism wisdom. 1:1 sessions with April to realign your body's innate intelligence.",
		Image:    "/img/adaptogenic-room.jpeg",
		ImageAlt: "Bright healing room with potted plants and floor cushions",
		Status:   "Active",
	},
}
