package content

import (
	"context"

	"github.com/statusnone/4life-community/internal/db"
)

func seedPages(ctx context.Context, store *db.Store) error {
	pages := []page{
		homePage(),
		meetAprilPage(),
		grantWritingPage(),
		astroYogaPage(),
		adaptogenicPage(),
		bioregionalNomadsPage(),
		susurrosPage(),
		givingPage(),
		catatumboPage(),
		contactPage(),
		podcastPage(),
		privacyPage(),
		termsPage(),
	}
	for i, p := range pages {
		if err := insertPage(ctx, store, p, i); err != nil {
			return err
		}
	}
	return nil
}

func homePage() page {
	return page{
		slug:  "home",
		title: "4Life | Grantwriting, Regenerative Strategy & Bioregional Alchemy – April Bartlett",
		meta:  "April Bartlett’s 4Life bridges grant writing, bioregional mapping, and AstroYoga—helping changemakers align strategy, soul, and planetary regeneration.",
		heroLabel:    "",
		heroHeadline: "Funding & Strategy for\nChangemakers, 4Life",
		heroBody:     "Grant writing, systems design, and trauma-informed guidance for people building a better world.\nRooted in justice, land wisdom, and results.",
		heroImage:    "/img/april-cutout-home.png",
		heroAlt:      "April Bartlett",
		heroSide:     "center",
		heroTheme:    "dark",
		sections: []section{
			{
				kind:     "intro",
				theme:    "white",
				headline: "I'm **April Bartlett**—grant writer, systems weaver, AstroYoga guide, and bioregional alchemist.",
				body:     "My work exists at the crossroads of personal + planetary regeneration, helping movements, events, healers, and changemakers *thrive*. Not just survive.\n\n**Whether you're here for:**\n\n- **Funding strategy** — grants, systems, regenerative economics\n- **Embodied wisdom** — AstroYoga, trauma-informed healing\n- **Bioregional mapping** — decolonized place-making, community resilience\n\nYou're in the right place! Let's move resources, energy, and vision where they're needed most.",
				buttonLabel: "Meet April",
				buttonURL:   "/meet-april",
			},
			{
				kind:    "quote",
				theme:   "black",
				headline: "4Life is where *strategy* meets *soul*, where land and body remember each other, where great things grow—because *you're here*.",
				bgImage: "/img/colombia-misty.png",
				bgOverlay: 0.33,
				divider: 6,
			},
			{
				kind:     "split",
				theme:    "dark",
				label:    "THE 4LIFE APPROACH",
				headline: "Regeneration isn't a metaphor. It's a practice of alignment, strategy, and deep listening.",
				body:     "At 4Life, I don't just 'consult.' I connect:\n\n- **Rooted in justice** — every strategy centers the people and land most impacted\n- **Built on results** — millions in funding secured, projects alive on real ground\n- **Holding the whole** — strategy, spirit, and ecology as a single weave\n- **Here for the long haul** — I stay until the work is thriving",
				image:    "/img/april-headshot.png",
				imageAlt: "April Bartlett at a desk writing",
				imageSide: "left",
			},
			{
				kind:     "services",
				theme:    "dark",
				label:    "SERVICES & OFFERINGS",
				headline: "You're doing the work. Let's make sure the world supports it!",
				buttonLabel: "April's Approach",
				buttonURL:   "/grant-writing",
				accordion: []accordion{
					{
						title: "Funding Alchemy",
						body:  "Grant writing for regenerative projects (millions secured, no soul-sucking jargon)\n\nSystems design for movements that scale *without* burnout\n\nLow-budget event magic (because scarcity is a myth)",
					},
					{
						title: "Bioregional Weaving",
						body:  "Community mapping *with* the land, not just *on* it\n\nWorkshops: from grant writing to menstrual resilience (yes, they're connected)",
					},
					{
						title: "AstroYoga & Embodied Wisdom",
						body:  "Trauma-informed yoga for changemakers (you can't pour from an empty cup)\n\nCosmic + somatic alignment (because your body is a bioregion too)",
					},
				},
			},
			{
				kind:     "split",
				theme:    "white",
				label:    "ADAPTAGENIC HEALING",
				headline: "Your Body Is Speaking — Are You Listening?",
				body:     "Trauma-informed somatic healing. Quantum Touch energy work. Folk herbalism wisdom.\n\nSchedule a 1:1 session with April to realign your body's innate intelligence.",
				image:    "/img/adaptogenic-room.jpeg",
				imageAlt: "Bright living room with large windows, potted plants, floor cushions, and string lights.",
				imageSide: "left",
				buttonLabel: "Schedule Session",
				buttonURL:   "/adaptogenic-healing",
			},
			{
				kind:     "split",
				theme:    "white",
				label:    "BIOREGIONAL NOMADS",
				headline: "Re-Mapping Belonging in a Displaced World",
				body:     "\"What if 'place' wasn't just a pin on a map—but a living conversation between land, body, and culture?\" From the cloud forests of Colombia to the Salish Sea watershed, we help humans:\n\n- **Redraw boundaries** (based on watersheds, not wars)\n- **Design open-source tools** for regenerative placemaking\n- **Host 'Inside Out' workshops** where personal + ecological healing meet",
				image:    "/img/bioregional-nomad.jpg",
				imageAlt: "A smiling man with a beard wearing a brown cap, denim jacket, and hoodie outdoors among green leafy plants and trees.",
				imageSide: "right",
				buttonLabel: "Learn more",
				buttonURL:   "/bioregional-nomads",
			},
			{
				kind:     "landack",
				theme:    "dark",
				headline: "Land & Legacy Acknowledgement",
				body:     "4Life is based in the unceded territories of the **qatay** (*kuh-tie*) peoples (Port Townsend, Wash.) and the wider Salish Sea, stewarded by **S'Klallam, Chemakum, and Coast Salish nations** since time immemorial. While April & her collaborators work spans many lands, this bioregion remains the heartbeat of our practice.",
			},
		},
	}
}

func meetAprilPage() page {
	return page{
		slug:         "meet-april",
		title:        "Meet April Bartlett — 4Life | Regenerative Strategy, AstroYoga & Bioregional Alchemy",
		meta:         "4Life is a living portal for April Bartlett's integrative work: trauma-informed yoga, holistic health, bioregional organizing, and community regeneration.",
		heroLabel:    "MEET APRIL",
		heroHeadline: "April Bartlett invites you into a world where personal healing and planetary restoration go hand in hand.",
		heroBody:     "Rooted in a lifetime of resilience, service, and deep listening, 4Life is a living portal for April's integrative work: trauma-informed yoga, holistic health, bioregional organizing, and community regeneration.",
		heroImage:    "/img/april-portrait.png",
		heroAlt:      "April Bartlett portrait",
		heroSide:     "right",
		heroTheme:    "black",
		sections: []section{
			{
				kind:     "split",
				theme:    "white",
				label:    "JOURNEY",
				headline: "Journey of Resilience and Transformation",
				body:     "April's journey is a story of resilience, transformation, and dedication to helping others find their inner harmony.\n\nApril's journey began in childhood, foraging with her parents and grandparents—learning firsthand that the Earth provides everything we need to thrive. By the age of 17, she was traveling across the U.S., training in herbal nutrition, enzymatic cleaning products, and water filtration. The birth of her third child, who faced life-threatening food and chemical sensitivities, deepened her calling and commitment to healing through ecological wisdom and regenerative practices.\n\nFaced with a choice between isolation or empowerment, April sought out natural healing alternatives, eventually learning from two Hopi Medicine Women. This opened her eyes to how sensitive systems are often excluded from social thinking—and inspired her to create Community 4Life Hub, a space for connection, healing, and low-histamine nourishment. The physical location has since closed. But the community is still connected and moving into the larger field.",
				image:    "/img/chimacum-farm.jpg",
				imageAlt: "Chimacum Creek Farm",
				imageSide: "left",
			},
			{
				kind:     "split",
				theme:    "white",
				label:    "TRAUMA-INFORMED YOGA & MOVEMENT",
				headline: "Trauma-Informed Yoga & Movement",
				body:     "Certified in trauma-informed yoga, April creates safe, somatic spaces for healing and reconnection. Her offerings are rooted in personal experience, ongoing study, and a deep respect for each person's unique journey.\n\nAfter years of chronic back pain from injury, a powerful voice whispered: 'You are going to heal yourself.' That message guided her toward yoga, Quantum Healing, cranial-sacral therapy, breath-work, and more. Since 2009, she's studied modalities including Hatha, Pranayama, Yin, Vinyasa, Bikram, Iyengar, Tai Chi, and Pilates. Her teachings are gentle, adaptive, and infused with lived wisdom.",
				image:    "/img/april-loft.jpg",
				imageAlt: "April Bartlett at Bridget's Loft in Port Townsend",
				imageSide: "right",
				buttonLabel: "Explore more",
				buttonURL:   "/astro-yoga",
			},
			{
				kind:     "split",
				theme:    "dark",
				label:    "BIOREGIONAL NOMAD SERVICES",
				headline: "Bioregional Nomad Services: Regeneration in Action",
				body:     "With some 25 years in the world of sustainability projects, April offers Bioregional Nomad 4Life Services, organizing tools for regenerative movements—supporting communities in creating ecological and cultural resilience. I'm testing the model now in Mogotes, Colombia to create open-source services that include:\n\n- Grant writing & storytelling workshops for regenerative projects\n- Workshop design & facilitation\n- Circular economy strategies through education & culture\n- Youth mentorship and movement-building\n- Ecological mapping and place-based planning\n\nThese services are currently rooted in Mogotes, Colombia, where April and her local collaborator Ricardo Palomino co-lead Susurros del Agua.\n\nI am currently offering workshops and tools as one-off events until the actual platform is launched.",
				image:    "/img/habitat-restoration.jpg",
				imageAlt: "People planting trees in a field with mountains in the background",
				imageSide: "left",
				buttonLabel: "Explore more",
				buttonURL:   "/bioregional-nomads",
			},
			{
				kind:     "split",
				theme:    "white",
				label:    "SUSURROS DEL AGUA",
				headline: "Susurros del Agua: A Living Watershed Project",
				body:     "Susurros del Agua is a bioregional movement that blends ancestral wisdom, hydrological restoration, community celebration, and education. The project uplifts local voices, supports eco-cultural tourism, and connects youth with regeneration through arts, games, markets, panels, and storytelling.\n\nBorn from a deep need to heal land and community, Susurros is supported mostly by April and Ricardo's own savings—with external support, it will grow into a continental and global example of what's possible when local people lead.",
				image:    "/img/susurros-hillside.jpg",
				imageAlt: "Lush forested hillside with a reddish cliff",
				imageSide: "right",
				buttonLabel: "Explore More",
				buttonURL:   "/susurros-del-agua",
			},
			{
				kind:     "split",
				theme:    "dark",
				label:    "HEALING THE INNER & OUTER LANDSCAPE",
				headline: "Healing the Inner & Outer Landscape",
				body:     "Whether through a yoga session, a storytelling workshop, a bioregional strategy, or a shared cup of tea—April's work carries the same message:\n\n> Healing is not separate from regeneration. What we repair within ourselves, we reflect into the world.",
				image:    "/img/finnriver.jpg",
				imageAlt: "Scenic view of a farm with a red barn, surrounded by green fields and trees, with snow-capped mountains in the background",
				imageSide: "left",
				buttonLabel: "Schedule a Session",
				buttonURL:   "/adaptogenic-healing",
			},
		},
	}
}

func grantWritingPage() page {
	return page{
		slug:         "grant-writing",
		title:        "Grant Writing Services — 4Life | Regenerative Strategy, AstroYoga & Bioregional Alchemy",
		meta:         "Discover expert grantwriting services at 4Life, offering strategic collaboration to secure funding and support regenerative, healing, and bioregional projects with confidence.",
		heroLabel:    "GRANT WRITING",
		heroHeadline: "More Than Just Consulting—A True Partnership",
		heroBody:     "At 4Life, I believe that funding isn't just about securing grants—it's about building sustainable success for organizations that are making a difference. That's why I work collaboratively, transparently, and strategically to ensure every project is aligned with your mission and built for long-term impact.\n\n**Here's what sets 4Life apart…**",
		heroImage:    "/img/april-headshot.png",
		heroAlt:      "April Bartlett headshot",
		heroSide:     "right",
		heroTheme:    "white",
		sections: []section{
			{
				kind:     "split",
				theme:    "white",
				label:    "TRANSPARENCY & COLLABORATION",
				headline: "Transparency & Collaboration",
				body:     "I don't believe in mystery deliverables or leaving you in the dark. Every step of the process is reviewed with you, ensuring that each grant, report, and strategy reflects your vision, goals, and performance metrics.\n\nThis is a true partnership, not just a service.",
				image:    "/img/food-bank-gardens.jpg",
				imageAlt: "Wooden sign reading 'Food Bank Garden' in front of a garden with plants and trees",
				imageSide: "left",
			},
			{
				kind:     "split",
				theme:    "white",
				label:    "EXPERIENCE YOU CAN TRUST",
				headline: "Experience You Can Trust",
				body:     "With years of experience and millions of dollars in secured funding, I know how to navigate complex grant processes with confidence. Whether you're applying for your first grant or managing a portfolio of funding sources, I bring the expertise, strategy, and insight to guide you through every step.\n\nNo guesswork. Just proven strategies that get results.",
				image:    "/img/april-loft.jpg",
				imageAlt: "April Bartlett writing",
				imageSide: "right",
			},
			{
				kind:     "split",
				theme:    "white",
				label:    "VALUE BEYOND WRITING",
				headline: "Value Beyond Writing",
				body:     "Grant writing is just one piece of the puzzle.\n\nI take a big-picture approach, supporting your organization at every stage of the funding journey—from identifying the right opportunities to post-submission follow-ups and reporting.\n\nThat means less stress, fewer roadblocks, and more time for you to focus on what truly matters—your mission.",
				image:    "/img/susurros-children.jpeg",
				imageAlt: "Children by the water",
				imageSide: "left",
			},
			{
				kind:     "list",
				theme:    "dark",
				label:    "LOOKING FOR MORE?",
				headline: "Additional Services for Your Success",
				body:     "Need ongoing support? I offer additional services to keep your organization funded, organized, and future-ready.",
				bullets: []string{
					"**Grant Prospecting** – Keep your funding pipeline full with ongoing research and new grant opportunities.",
					"**Grant Reporting** – Simplify compliance with detailed progress and financial reports that keep funders happy.",
					"**Capacity Building** – Empower your team with templates, training, and tools to write stronger future grants.",
				},
				buttonLabel: "Contact April",
				buttonURL:   "/contact",
			},
		},
	}
}

func astroYogaPage() page {
	return page{
		slug:         "astro-yoga",
		title:        "AstroYoga — 4Life | Regenerative Strategy, AstroYoga & Bioregional Alchemy",
		meta:         "April's work bridges body and land, self and system. Trauma-informed AstroYoga for changemakers, guided by cosmic and somatic alignment.",
		heroLabel:    "ASTROYOGA",
		heroHeadline: "April's work bridges body and land, self and system.",
		heroBody:     "I believe healing and regeneration must walk hand in hand. Learn more about how these threads weave together in my story.\n\nApril is a certified trauma-informed yoga instructor with 10 years of trauma studies. Her expertise in creating safe and supportive environments ensures that everyone can fully embrace the transformative power of yoga, free from judgment or discrimination.",
		heroImage:    "/img/astroyoga-masthead.png",
		heroAlt:      "Logo combining an astrology-themed circular design with a person in a meditative pose, surrounded by zodiac symbols, with the text 'Astro Yoga with April' to the right.",
		heroSide:     "left",
		heroTheme:    "black",
		sections: []section{
			{
				kind:  "bokeh",
				theme: "bokeh",
			},
		},
	}
}

func adaptogenicPage() page {
	return page{
		slug:         "adaptogenic-healing",
		title:        "Adaptogenic Healing — 4Life | Regenerative Strategy, AstroYoga & Bioregional Alchemy",
		meta:         "Adaptogenic Healing where your body's wisdom meets earth's medicine. Trauma-informed somatic healing, Quantum Touch, and folk herbalism.",
		heroLabel:    "ADAPTOGENIC HEALING",
		heroHeadline: "Adaptogenic Healing",
		heroBody:     "**Where Your Body's Wisdom Meets Earth's Medicine**\n\nThrough a blend of movement, dietary/herbal guidance, body-mind therapies, and joyful human connection, I help you reconnect with your innate vitality—because true healing is never one-size-fits-all.",
		heroImage:    "/img/adapt-hero.jpg",
		heroAlt:      "Serene yoga and wellness setting",
		heroSide:     "right",
		heroTheme:    "black",
		sections: []section{
			{
				kind:     "split",
				theme:    "white",
				label:    "A FUSION OF MODALITIES",
				headline: "A Fusion of Modalities",
				body:     "My approach, which I call Adaptogenics, integrates 16+ years as an Energy Medicine Practitioner with certifications in Quantum Touch, Intraoral Massage, and Trauma-Informed Yoga. It's not just about techniques—it's about listening to what your body is already trying to say. Sessions might include energy channeling for blocked chi, intraoral work for TMJ relief, or somatic practices to unwind stored trauma.",
				image:    "/img/adapt-body.jpg",
				imageAlt: "Calming therapeutic space",
				imageSide: "left",
			},
			{
				kind:     "split",
				theme:    "white",
				label:    "ROOTED IN EARTH WISDOM",
				headline: "Rooted in Earth Wisdom",
				body:     "I grew up studying Folk Herbalism and trained with two Hopi Medicine Women for five years, weaving shamanic practices into my work. Combined with my Ayurvedic training and ongoing collaboration with psychologists, this creates a bridge between ancient wisdom and modern nervous system science.",
				image:    "/img/adaptogenic-room.jpeg",
				imageAlt: "Bright living room with large windows, potted plants, floor cushions, and string lights.",
				imageSide: "right",
			},
			{
				kind:     "list",
				theme:    "white",
				label:    "FOR THOSE READY TO GO DEEPER",
				headline: "For Those Ready to Go Deeper",
				body:     "This is for you if you've tried surface-level fixes and crave a guide who sees:",
				bullets: []string{
					"Your migraine as a map of unspoken stress",
					"Your fatigue as a call to align with your cycles",
					"Your 'stuckness' as energy waiting for the right key",
				},
			},
			{
				kind:     "session-form",
				theme:    "dark",
				label:    "LET'S BEGIN",
				headline: "Let's Begin",
				body:     "Fill out this form to schedule your session. I'll respond within 48 hours with available times.",
				buttonLabel: "Send my Request",
			},
		},
	}
}

func bioregionalNomadsPage() page {
	return page{
		slug:         "bioregional-nomads",
		title:        "Bioregional Nomads | Discover Regenerative Solutions — 4Life | Regenerative Strategy, AstroYoga & Bioregional Alchemy",
		meta:         "Explore bioregional mapping, community workshops, and systems design for ecological and cultural renewal. Empower your local regenerative initiatives today.",
		heroLabel:    "BIOREGIONAL NOMADS",
		heroHeadline: "Bioregional Nomads",
		heroBody:     "Mapping the Stories Land Tells",
		heroImage:    "/img/bioregional-nomads-hero.jpg",
		heroAlt:      "A smiling man with a beard wearing a brown cap, denim jacket, and hoodie outdoors among green plants and trees",
		heroSide:     "center",
		heroTheme:    "black",
		sections: []section{
			{
				kind:     "split",
				theme:    "dark",
				label:    "BIOREGIONAL MAPPING",
				headline: "Bioregional Mapping",
				body:     "Bioregional mapping is a process which lets us reclaim mapmaking for our own purposes and communities.\n\n> \"Bioregional mapping is a process for creating maps and tools that set cultural and geographic information into natural boundaries and borders, rather than anthropocentric lines removed from context of place. Maps are not neutral, and many maps today are constructed by governments or economic entities to convey their interests and reinforce their world view. Bioregional mapping is a community and participatory process to create maps that combine ecological and physical information with social and cultural information within a given place, as defined by those living there or the communities most impacted. It is both a tradition that dates back thousands of years, inspired by countless forms of Indigenous Mapping, and also that has emerged as a direct and modern response to the erasure of local cultures in the face of our current ecological, economic and social crises.\" *(Watershed Commons)*\n\nWe guide communities through bioregional mapping—clarifying where to begin, who to speak with, how to access key ecological and cultural information, and how to weave sustainable initiatives together for impact. As Gary Snyder encourages: \"Find your place on the planet. Dig in, and take responsibility from there,\" and in line with Raye Stoeve's wisdom, \"we reimagine boundaries based on ecological and cultural realities.\"",
				image:    "/img/bioregional-nomad.jpg",
				imageAlt: "Regenerative landscape in the Pacific Northwest",
				imageSide: "left",
			},
			{
				kind:     "list",
				theme:    "white",
				label:    "TAILORED WORKSHOPS",
				headline: "Workshop Categories",
				body:     "Our workshops are designed to encourage project managers, community organizers, and healing practitioners with the tools to lead regenerative change—from managing systems and writing grants to reconnecting with the wisdom of the body and the land.\n\nRooted in lived experience and holistic strategy, these sessions cultivate clarity, resilience, and action—no matter your starting point or budget.",
				bullets: []string{
					"**System Management** – Practical tools for organizing, tracking, and sustaining regenerative work",
					"**Leadership** – Embodied leadership and facilitation for community regeneration",
					"**Inside Out** – Mind-body wellness and how that relates to your work",
					"**Women's Health** – Menstrual and hormonal health, somatic awareness, and fertility wisdom",
					"**Organizing Your Region** – Steps to identify, connect, and mobilize local regenerative efforts",
					"**What Serves Your Bioregion** – Place-based analysis for culture, ecology, and economic alignment",
					"**Storytelling** – Narrative design for impact, identity, and community engagement",
					"**Grant Writing** – Values-aligned funding strategies and application support",
					"**Low to No Budget Events** – Creative methods for activating community with minimal resources",
				},
			},
			{
				kind:     "list",
				theme:    "dark",
				label:    "SYSTEMS DESIGN",
				headline: "Our systems templates include:",
				body:     "We support bioregional organizers, project leaders, and community weavers by designing custom, open-source management systems that align with their purpose, scale, and capacity. These tools are tailored to the unique needs of regenerative movements—whether you're just getting started or managing a growing network.",
				bullets: []string{
					"Immersion Frameworks",
					"Member tracking & communications",
					"Grant cycle + donor management",
					"Volunteer coordination",
					"Event & workshop planning",
					"Proposals / Contracts / Work Agreements",
					"Task delegation + roles clarity",
					"Task management tools",
					"Multi-language access for global collaboration",
					"And more",
				},
				buttonLabel: "Contact April",
				buttonURL:   "/contact",
			},
		},
	}
}

func susurrosPage() page {
	return page{
		slug:         "susurros-del-agua",
		title:        "Susurros del Agua — 4Life | Regenerative Strategy, AstroYoga & Bioregional Alchemy",
		meta:         "Susurros del Agua is a grassroots bioregional initiative co-led by April Bartlett and Ricardo Palomino in Mogotes, Colombia—a bioregional movement for watershed restoration and cultural resilience.",
		heroLabel:    "SUSURROS DEL AGUA",
		heroHeadline: "Susurros del Agua",
		heroBody:     "A Bioregional Movement for Watershed Restoration & Cultural Resilience in Mogotes, Colombia\n\nMy work bridges body and land, self and system. I believe healing and regeneration must walk hand in hand. Learn more about how these threads weave together in my story.",
		heroTheme:    "dark",
		sections: []section{
			{
				kind:     "split",
				theme:    "white",
				label:    "ABOUT THE PROJECT",
				headline: "Why We're Listening to the Water & the People",
				body:     "Susurros del Agua is a grassroots bioregional initiative co-led by April Bartlett and Ricardo Palomino in Mogotes, Colombia.\n\nBy walking with the Río Mogoticos—studying its path, its rhythms, and what disturbs its natural flow—we come to understand the needs of the water. And by listening deeply to the voices of the community, we uncover the needs of the people who live alongside it.\n\nWe are documenting hydrological heritage and restoring connection to place through mapping, storytelling, arts, education, eco-cultural tourism, and cultural celebration—all in service of protecting and revitalizing the watershed.\n\nBorn from a community responding to ecological breakdown, Susurros del Agua empowers local leadership, honors ancestral knowledge, and cultivates a regenerative future rooted in relationship, memory, and care.",
				image:    "/img/susurros-children.jpeg",
				imageAlt: "Children by the water",
				imageSide: "right",
			},
			{
				kind:     "list",
				theme:    "dark",
				label:    "OUR WORK WEAVES TOGETHER",
				headline: "Our Work Weaves Together",
				bullets: []string{
					"**Bioregional Mapping** – Connecting people to place through participatory mapping",
					"**Ecosystem Restoration** – Supporting local reforestation and water source care",
					"**Eco-Cultural Events** – Hosting music, panels, cafecitas, and regenerative markets",
					"**Youth Education** – Facilitating ecological and cultural programs for future leaders",
					"**Storytelling & Media** – Sharing the voice of Mogotes with the world",
					"**Community Organizing** – Building capacity with systems, training, and support",
				},
				body: "> Healing the water heals the land. Healing the land heals the people.\n\nWe imagine a Mogotes where community decisions are rooted in ecological wisdom, where youth are leaders, and where tourism uplifts culture, care, and conservation.",
				gallery: []string{
					"/img/susurros-scenes.jpeg",
				},
			},
			{
				kind:     "split",
				theme:    "white",
				label:    "SCENES FROM THE TERRITORY",
				headline: "Scenes from the Territory",
				image:    "/img/susurros-hillside.jpg",
				imageAlt: "Lush forested hillside with a reddish cliff",
				imageSide: "left",
				body:     "Join us for an immersion.\n\nPlease note that space is limited. Be sure to reserve your spot in advance by filling out this registration form.",
				buttonLabel: "Make a Contribution",
				buttonURL:   "https://www.paypal.com/donate/?hosted_button_id=JZVHCPVN9CKXJ",
			},
		},
	}
}

func givingPage() page {
	return page{
		slug:         "giving-page-2-1",
		title:        "Giving To A Cause — 4Life | Regenerative Strategy, AstroYoga & Bioregional Alchemy",
		meta:         "Contribute to regenerative projects, farms, and community initiatives. Funds support food security, farmers, and environmental restoration.",
		heroLabel:    "GIVING TO A CAUSE",
		heroHeadline: "Get Involved",
		heroBody:     "Together we can be the difference in our community!",
		heroTheme:    "dark",
		sections: []section{
			{
				kind:     "split",
				theme:    "white",
				label:    "JOIN THE MOVEMENT",
				headline: "Join the Movement",
				body:     "Here in the Pacific Northwest we have so many incredible farms that provide our food. Many of those farms rely on Government Funding. With the recent funding freeze via the Trump administration, I'm rallying support to keep these Farms in business, Food Banks full, and Restoration Projects rolling. We cannot afford to wait on restoring our planet or feeding our communities.\n\nThe goal is to raise as much money as possible which will be delegated to various organizations. For example, I have worked with Friends Of the Trees, Global Earth Repair Foundation, Goosefoot Farms, Construct Diversity, Food Bank Growers, Heavn Haus Super Food Farm, The School For Regenerating Earth, Susurros del Agua, Revival Gatherings, and others. My promise is that whatever is given will be delegated to the project with the highest need at the time that the funds are available. All contributions make an impact. Please share this opportunity to give with your community.",
				image:    "/img/habitat-restoration.jpg",
				imageAlt: "People planting trees in a field with mountains in the background",
				imageSide: "right",
				gallery: []string{
					"/img/finnriver.jpg",
					"/img/chimacum-farm.jpg",
					"/img/food-bank-gardens.jpg",
				},
				buttonLabel: "Spread the Word",
				buttonURL:   "https://www.instagram.com/amber.rising4love/",
			},
			{
				kind:     "donate",
				theme:    "dark",
				label:    "MAKE A CONTRIBUTION",
				headline: "Donate to the General Fund",
				body:     "Every contribution makes an impact. Choose an amount below—or give what feels right—and every dollar goes to the regenerative project with the highest need at the time funds are available.",
				buttonLabel: "Donate",
				buttonURL:   "/giving-page-2-1",
			},
			{
				kind:     "list",
				theme:    "white",
				label:    "SPREAD THE WORD & SPREAD THE WEALTH",
				headline: "Spread the Word & Spread the Wealth",
				body:     "Together we can be the difference in our community!",
				bullets: []string{
					"[Instagram](https://www.instagram.com/amber.rising4love/)",
					"[YouTube](https://www.youtube.com/@4life.AprilBartlett)",
				},
			},
		},
	}
}

func catatumboPage() page {
	return page{
		slug:         "catatumbo",
		title:        "Giving To A Cause | Support Local Regeneration Today — 4Life | Regenerative Strategy, AstroYoga & Bioregional Alchemy",
		meta:         "Save a River, Save a Community. Giving the Catatumbo River a second chance—support the fight for justice for the communities who defend it.",
		heroLabel:    "GIVING TO A CAUSE",
		heroHeadline: "Save a River,\nSave a Community.",
		heroBody:     "The people of the Catatumbo are not asking for money. They are asking for action.",
		heroImage:    "/img/catatumbo-hero.jpg",
		heroAlt:      "The Catatumbo River basin",
		heroSide:     "center",
		heroTheme:    "black",
		sections: []section{
			{
				kind:     "split",
				theme:    "white",
				label:    "THE CATATUMBO RIVER",
				headline: "Giving the Catatumbo River a Second Chance",
				body:     "At the headwaters of the Catatumbo River in Cúcuta, Colombia, a story of neglect and resistance is unfolding. For years, local fisheries and fishermen have been sounding the alarm as oil from an abandoned EcoPetra rig continues to pour into the waterways. What was once a thriving ecosystem and source of livelihood is now collapsing: the fish are gone, the river is poisoned, and the communities who depend on it are left in crisis.\n\nAfter years of pushback, Colombia's Supreme Court finally appointed a judge to hear the pleas of the people and of the river itself. But so far, justice has been delayed. The judge hesitates to act against corporate power, even as millions of barrels of crude oil devastate the basin.\n\n> The people of the Catatumbo are not asking for money. They are asking for action.\n\nColombian law requires it: **Decreto 1076 de 2015**, **Ley 2327 de 2023 — Ley de Pasivos Ambientales**, and **Ley 99 de 1993** all mandate impact assessments, environmental permits, monitoring, and remedial measures when ecosystems are harmed. Yet on the Catatumbo, none of this is happening.\n\nThe Federation of Artisanal Fishers (FEDEPESCAT) has already rejected ANLA's weak response to this crisis. Instead of demanding urgent cleanup, the environmental authority has chosen delay, pushing responsibility to future studies and other institutions. For communities on the ground, that delay is a death sentence for both the river and their way of life. **We will not wait!**\n\nWe are acting now to keep the pressure on. On **October 3rd, 2025**, we will bring the story of Catatumbo to the world through a documentary, giving voice to the fishermen, the guardians of water, and the river itself. With the Ministry of Culture covering the filming crew, we only need **$1,500 USD** to make the gathering possible.\n\n**These funds will cover:**\n\n- Travel for representatives of each fishing association\n- Food for the \"Jutanza\" (community gathering)\n- Venue and logistical support\n\nThis documentary is the first step. A larger fundraising campaign will follow to support the legal fight and long-term protections for our Water Guardians.\n\n**Join us in giving the Catatumbo River a second chance.** Together we can help restore this sacred ecosystem and ensure justice for the communities who defend it.",
				image:    "/img/catatumbo-river.jpeg",
				imageAlt: "The Catatumbo River and its defenders",
				imageSide: "right",
			},
			{
				kind:     "donate",
				theme:    "dark",
				label:    "JOIN US",
				headline: "Join us in giving the Catatumbo River a second chance.",
				body:     "If 50 people give just $30.00 we will reach our goal and create a documentary. You can say that you were a part of the creation of this film!\n\n**For those in Colombia:** Donate here: Nequi @Fedepescat Cód: 0090779688\n\nAll funds go straight to on the ground actions!\n\nTogether we can be the difference in our community!",
				buttonLabel: "Give Directly via PayPal Javier Piffano",
				buttonURL:   "https://www.paypal.com/donate/?hosted_button_id=TT554HU8UEYKE",
			},
			{
				kind:     "gallery",
				theme:    "white",
				label:    "SPREAD THE WORD",
				headline: "Spread the Word & Spread the Wealth",
				gallery: []string{
					"/img/catatumbo-ecocide.png",
					"/img/catatumbo-fisher1.jpg",
					"/img/catatumbo-fisher2.jpg",
				},
				buttonLabel: "Contact April",
				buttonURL:   "/contact",
			},
		},
	}
}

func contactPage() page {
	return page{
		slug:         "contact",
		title:        "Contact April Bartlett & Team — 4Life | Regenerative Strategy, AstroYoga & Bioregional Alchemy",
		meta:         "Whether you're seeking healing sessions, bioregional mapping support, grant-writing collaboration, or want April on your podcast—use the forms on this page to connect.",
		heroLabel:    "GENERAL INQUIRIES",
		heroHeadline: "Let's Weave Something Meaningful Together!",
		heroBody:     "Whether you're seeking healing sessions, bioregional mapping support, grant-writing collaboration, or want to say hi—use the form on this page to connect.\n\n*(Astrological memes and land rematriation strategies equally welcome!)*",
		heroTheme:    "white",
		sections: []section{
			{
				kind:  "contact-form",
				theme: "white",
			},
			{
				kind:     "split",
				theme:    "dark",
				label:    "ADAPTAGENIC HEALING",
				headline: "Your Body Is Speaking — Are You Listening?",
				body:     "Trauma-informed somatic healing. Quantum Touch energy work. Folk herbalism wisdom.\n\nSchedule a 1:1 session with April to realign your body's innate intelligence.",
				image:    "/img/adaptogenic-room.jpeg",
				imageAlt: "Bright living room with large windows, potted plants, floor cushions, and string lights.",
				imageSide: "left",
				buttonLabel: "Schedule a Session",
				buttonURL:   "/adaptogenic-healing",
			},
		},
	}
}

func podcastPage() page {
	return page{
		slug:         "podcast-media-inquiries",
		title:        "Podcast & Media Inquiries — 4Life | Regenerative Strategy, AstroYoga & Bioregional Alchemy",
		meta:         "Invite April Bartlett to your podcast, summit, or interview series to discuss grant writing, bioregional mapping, trauma-informed AstroYoga, and more.",
		heroLabel:    "PODCAST & MEDIA INQUIRIES",
		heroHeadline: "Let's Amplify Regenerative Stories",
		heroBody:     "Invite April to your podcast, summit, or interview series to discuss:\n\n- Grant writing as sacred practice\n- Bioregional mapping for cultural healing\n- Trauma-informed AstroYoga\n- Where folk herbalism meets modern activism\n- ...or other threads of the 4Life tapestry.",
		heroTheme:    "black",
		sections: []section{
			{
				kind:     "split",
				theme:    "white",
				label:    "BIOREGIONAL NOMADS",
				headline: "Re-Mapping Belonging in a Displaced World",
				body:     "\"What if 'place' wasn't just a pin on a map—but a living conversation between land, body, and culture?\" The Bioregional Nomads project guides communities in decolonized mapping, regenerative systems design, and the art of rooted itinerancy.\n\nFrom the cloud forests of Colombia to the Salish Sea watershed, we help humans:\n\n- **Redraw boundaries** (based on watersheds, not wars)\n- **Design open-source tools** for regenerative placemaking\n- **Host 'Inside Out' workshops** where personal + ecological healing meet",
				image:    "/img/bioregional-nomad.jpg",
				imageAlt: "A smiling man with a beard wearing a baseball cap, denim jacket, and hoodie outdoors among green plants and trees",
				imageSide: "right",
				buttonLabel: "Learn more",
				buttonURL:   "/bioregional-nomads",
			},
			{
				kind:  "contact-form",
				theme: "white",
			},
		},
	}
}

func privacyPage() page {
	return page{
		slug:         "privacy-policy",
		title:        "Privacy Policy — 4Life | Regenerative Strategy, AstroYoga & Bioregional Alchemy",
		meta:         "4Life Privacy Policy.",
		heroLabel:    "LEGAL",
		heroHeadline: "Privacy Policy",
		heroBody:     "Last Updated: February 10, 2024",
		heroTheme:    "dark",
		sections: []section{
			{
				kind:  "richtext",
				theme: "white",
				body: `Welcome to 4Life, operated by April Bartlett in Port Townsend, Washington. Your privacy is of paramount importance to us. This Privacy Policy outlines the types of information we collect from our clients and visitors, how we use it, and the steps we take to protect it.

**1. Information Collection**

4Life collects information in several ways:

- **Directly from you:** When you sign up for our newsletter, register for a class, or participate in our events, we may ask for personal information such as your name, email address, and contact details.
- **Automatically through our website:** We collect information about your visit to our site, including your IP address, browser type, and usage patterns, through cookies and similar technologies to improve our website's functionality and your experience.

**2. Use of Information**

The information we collect is used to:

- Provide, personalize, and improve our services.
- Communicate with you about upcoming events, classes, and offers that might interest you.
- Understand how our website and services are used to enhance and optimize our offerings.

**3. Sharing of Information**

4Life does not sell or rent your personal information to third parties. We may share information with:

- Service providers who perform services on our behalf, such as website hosting, email delivery, and marketing services, under strict confidentiality agreements.
- Legal authorities when required by law or to protect the rights and safety of 4Life, our clients, and the public.

**4. Data Security**

We implement a variety of security measures to maintain the safety of your personal information. However, no internet transmission is 100% secure, and we cannot guarantee the absolute security of your information.

**5. Your Choices**

You can choose not to provide certain information, but this may limit your ability to participate in some of our services. You can also opt-out of receiving marketing communications from us by following the unsubscribe link in our emails.

**6. Updates to this Policy**

4Life may update this Privacy Policy periodically. We will notify you of any significant changes by posting the new policy on our website and updating the "Last Updated" date.

**7. Contact Us**

If you have any questions about this Privacy Policy or our privacy practices, please contact us at 4life.AprilBartlett@gmail.com.`,
			},
		},
	}
}

func termsPage() page {
	return page{
		slug:         "terms-and-conditions",
		title:        "Terms & Conditions — 4Life | Regenerative Strategy, AstroYoga & Bioregional Alchemy",
		meta:         "4Life Terms & Conditions.",
		heroLabel:    "LEGAL",
		heroHeadline: "Terms & Conditions",
		heroBody:     "Effective: February 10, 2024",
		heroTheme:    "dark",
		sections: []section{
			{
				kind:  "richtext",
				theme: "white",
				body: `Welcome to 4Life. By accessing or using this website you agree to the following terms and conditions.

**1. Use of Service**

You agree to use this website for lawful purposes only and in a way that does not infringe the rights of, restrict, or inhibit anyone else's use and enjoyment of the site.

**2. Accounts**

Some areas of the site may require registration. You are responsible for maintaining the confidentiality of your account credentials and for all activities that occur under your account.

**3. Intellectual Property**

All content on this site—including text, graphics, logos, and imagery—is the property of 4Life or its content suppliers and is protected by copyright and other intellectual property laws. You may not reproduce, distribute, or republish any content without prior written consent.

**4. Links To Other Web Sites**

This site may contain links to third-party websites. 4Life has no control over and assumes no responsibility for the content, privacy policies, or practices of any third-party sites.

**5. Termination**

We may terminate or suspend access to our service immediately, without prior notice or liability, for any reason whatsoever, including without limitation a breach of these terms.

**6. Limitation of Liability**

In no event shall 4Life be liable for any indirect, incidental, special, consequential, or punitive damages arising out of or related to your use of this website.

**7. Governing Law**

These terms shall be governed and construed in accordance with the laws of the State of Washington, without regard to its conflict of law provisions.

**8. Changes**

We reserve the right to modify or replace these terms at any time. If a revision is material, we will provide notice prior to any new terms taking effect.

**9. Contact Us**

If you have any questions about these terms, please contact us at 4life.AprilBartlett@gmail.com.`,
			},
		},
	}
}
