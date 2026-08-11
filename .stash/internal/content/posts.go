package content

import (
	"context"
	"fmt"

	"github.com/statusnone/4life-community/internal/db"
)

type post struct {
	slug, title, excerpt, cover, externalURL string
	date                                     string
	body                                     string
}

func seedPosts(ctx context.Context, store *db.Store) error {
	posts := []post{
		{
			slug: "giving-the-catatumbo-river-a-second-chance",
			title: "Catatumbo River: Restore this sacred ecosystem and ensure justice for the communities who defend it",
			date:  "2025-10-29",
			externalURL: "/catatumbo",
			cover: "/img/post-catatumbo.png",
			excerpt: "This documentary is the first step. A larger fundraising campaign will follow to support the legal fight and long-term protections for our Water Guardians.",
			body: catatumboPost,
		},
		{
			slug:  "collaborative-action-update-with-prosocial-world",
			title: "Collaborative Action Update With Prosocial World",
			date:  "2025-09-17",
			cover: "/img/post-prosocial.jpg",
			excerpt: "Creating positive, lasting change—together! The Leadership Training has been AWESOME! I have to share some bits with you. The r3.0 Global Confluence has played a BIG part in bringing the pieces together!",
			body: prosocialPost,
		},
		{
			slug:  "hands-in-the-soil-heart-in-the-plants",
			title: "Hands in the Soil, Heart in the Plants",
			date:  "2025-08-12",
			cover: "/img/post-soil.jpg",
			excerpt: "Join master herbalist Skeeter Pilarski and his son Ashley, as they reveal how they grow 80+ medicinal herbs on just a quarter-acre—and why soil, hedgerows, and wildlife partnerships are their secret weapons.",
			body: soilPost,
		},
		{
			slug:  "inspired-by-nature-exploring-connections-with-garn",
			title: "Inspired by Nature: Exploring Connections with GARN",
			date:  "2025-08-07",
			cover: "/img/post-garn.jpg",
			excerpt: "At 4Life, I'm buzzing with excitement from recent conversations with the founder of the Global Alliance for the Rights of Nature (GARN). Their mission to honor ecosystems as living entities inspires us to connect local leaders, youth, and creative thinkers in restoration efforts worldwide.",
			body: garnPost,
		},
		{
			slug:  "education-through-regenerative-actio",
			title: "Education Through Regenerative Action",
			date:  "2025-05-30",
			cover: "/img/post-education.png",
			excerpt: "I've been so blessed to work alongside some of the world's most dedicated and heart-centered humans—people fully committed to action-based education. This powerful form of learning reconnects people with the land, the water, and all living beings.",
			body: educationPost,
		},
		{
			slug:  "regenerative-storytelling-help-sustain-baricharas-living-narrative",
			title: "Regenerative Storytelling: Help Sustain Barichara's Living Narrative",
			date:  "2025-05-11",
			cover: "/img/post-storytelling.png",
			excerpt: "In the highlands of Colombia, nestled in the Andes Mountains above the Río Suárez, the town of Barichara pulses with an ancient rhythm—one rooted in the quiet resilience of its people.",
			body: storytellingPost,
		},
		{
			slug:  "backing-to-spread-the-magic-5625",
			title: "Backpacking to Spread the Magic 5.6.25",
			date:  "2025-05-07",
			cover: "/img/post-backpacking.jpg",
			excerpt: "Living on a once-abandoned preschool plot in Barichara, Colombia, I'm finding healing for my soul, mind, and body. Alongside Nina, Anthony, and their daughter Enna, we're transforming invasive plants into fertile soil, restoring this eroded land.",
			body: backpackingPost,
		},
		{
			slug:  "astrological-poetry-42325",
			title: "Astrological Poetry 4.23.25",
			date:  "2025-04-23",
			cover: "/img/post-poetry.jpeg",
			excerpt: "Rise steady, rise sure; the heavens trace your silhouette, and the Universe echoes reminding you: you are the pulse of necessary things.",
			body: poetryPost,
		},
		{
			slug:  "back-in-colombia",
			title: "Back In Colombia",
			date:  "2025-04-21",
			cover: "/img/colombia-misty.png",
			excerpt: "A Letter from the Soul. I have returned to Barichara—a town that whispers ancient wisdom through the cobbled streets, where the hills hum with memory and possibility.",
			body: backInColombiaPost,
		},
		{
			slug:  "barichara-colombia-earth-regeneration",
			title: "Barichara Colombia: Earth Regeneration",
			date:  "2025-03-05",
			cover: "/img/blog-hero-bg.png",
			excerpt: "Since I was in elementary school, I have felt the call to protect the rainforest. And now, here I am. Standing on this land, my hands in the soil, I see the scars left behind—and the jungle singing back.",
			body: baricharaPost,
		},
		{
			slug:  "back-in-action",
			title: "Back In Action",
			date:  "2025-01-31",
			cover: "/img/april-loft.jpg",
			excerpt: "After a season of moving and hibernating, I'm finally emerging—shaking off the dust and stepping back into the flow!",
			body: backInActionPost,
		},
		{
			slug:  "passion-is-a-project",
			title: "Passion Is A Project",
			date:  "2024-04-25",
			cover: "/img/april-portrait.png",
			excerpt: "I've got passion burning in my bones! I have so much that I want to share and so many ideas that I want to explore. But… I can't do it alone. I, like the rest of us, need community to bring the concepts into reality.",
			body: passionPost,
		},
		{
			slug:  "round-onegame-on",
			title: "Round One…Game On",
			date:  "2024-04-08",
			cover: "/img/adapt-hero.jpg",
			excerpt: "Yesterday was my first ever game night. I've never been one to sit around and play games. But I'm starting to see it all coming to fruition.",
			body: gameNightPost,
		},
		{
			slug:  "healing-the-inner-child-meditaion",
			title: "Healing the Inner Child Meditation",
			date:  "2024-04-07",
			cover: "/img/adapt-body.jpg",
			excerpt: "Welcome to this guided meditation for healing the inner child. Find a comfortable position, either sitting or lying down, and allow yourself to fully relax.",
			body: innerChildPost,
		},
		{
			slug:  "poetry-night",
			title: "Poetry Night",
			date:  "2024-03-26",
			cover: "/img/astroyoga-masthead.png",
			excerpt: "March, Friday 29th, 2024 from 7pm-9pm will be our first Poetry Night at The HUB. All the right star juju for some poetry reading!",
			body: poetryNightPost,
		},
	}

	for _, p := range posts {
		if _, err := store.Pool.Exec(ctx, `
			INSERT INTO posts (slug, title, excerpt, cover, external_url, body, published_at, published)
			VALUES ($1,$2,$3,$4,$5,$6,$7,true)`,
			p.slug, p.title, p.excerpt, p.cover, p.externalURL, p.body, p.date); err != nil {
			return fmt.Errorf("insert post %q: %w", p.slug, err)
		}
	}
	return nil
}
