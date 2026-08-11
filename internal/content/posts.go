package content

import "time"

// Post is a single Field Notes entry. Body is markdown.
type Post struct {
	Slug    string
	Title   string
	Excerpt string
	Cover   string
	Date    string
	Body    string
}

// PublishedAt parses the post date for display.
func (p Post) PublishedAt() time.Time {
	t, _ := time.Parse("2006-01-02", p.Date)
	return t
}

// Posts lists Field Notes newest-first.
var Posts = []Post{
	{
		Slug:    "hands-in-the-soil-heart-in-the-plants",
		Title:   "Hands in the Soil, Heart in the Plants",
		Date:    "2025-08-12",
		Cover:   "/img/post-soil.jpg",
		Excerpt: "Join master herbalist Skeeter Pilarski and his son Ashley, as they reveal how they grow 80+ medicinal herbs on just a quarter-acre—and why soil, hedgerows, and wildlife partnerships are their secret weapons.",
		Body:    soilPost,
	},
	{
		Slug:    "inspired-by-nature-exploring-connections-with-garn",
		Title:   "Inspired by Nature: Exploring Connections with GARN",
		Date:    "2025-08-07",
		Cover:   "/img/post-garn.jpg",
		Excerpt: "At 4Life, I'm buzzing with excitement from recent conversations with the founder of the Global Alliance for the Rights of Nature (GARN). Their mission to honor ecosystems as living entities inspires us to connect local leaders, youth, and creative thinkers in restoration efforts worldwide.",
		Body:    garnPost,
	},
	{
		Slug:    "education-through-regenerative-actio",
		Title:   "Education Through Regenerative Action",
		Date:    "2025-05-30",
		Cover:   "/img/post-education.jpg",
		Excerpt: "I've been so blessed to work alongside some of the world's most dedicated and heart-centered humans—people fully committed to action-based education. This powerful form of learning reconnects people with the land, the water, and all living beings.",
		Body:    educationPost,
	},
	{
		Slug:    "backpacking-to-spread-the-magic-5625",
		Title:   "Backpacking to Spread the Magic 5.6.25",
		Date:    "2025-05-07",
		Cover:   "/img/post-backpacking.jpg",
		Excerpt: "Living on a once-abandoned preschool plot in Barichara, Colombia, I'm finding healing for my soul, mind, and body. Alongside Nina, Anthony, and their daughter Enna, we're transforming invasive plants into fertile soil, restoring this eroded land.",
		Body:    backpackingPost,
	},
	{
		Slug:    "astrological-poetry-42325",
		Title:   "Astrological Poetry 4.23.25",
		Date:    "2025-04-23",
		Cover:   "/img/post-poetry.jpeg",
		Excerpt: "Rise steady, rise sure; the heavens trace your silhouette, and the Universe echoes reminding you: you are the pulse of necessary things.",
		Body:    poetryPost,
	},
	{
		Slug:    "back-in-colombia",
		Title:   "Back In Colombia",
		Date:    "2025-04-21",
		Cover:   "/img/colombia-misty.jpg",
		Excerpt: "A Letter from the Soul. I have returned to Barichara—a town that whispers ancient wisdom through the cobbled streets, where the hills hum with memory and possibility.",
		Body:    backInColombiaPost,
	},
	{
		Slug:    "back-in-action",
		Title:   "Back In Action",
		Date:    "2025-01-31",
		Cover:   "/img/blog-hero-bg.jpg",
		Excerpt: "After a season of moving and hibernating, I'm finally emerging—shaking off the dust and stepping back into the flow!",
		Body:    backInActionPost,
	},
	{
		Slug:    "passion-is-a-project",
		Title:   "Passion Is A Project",
		Date:    "2024-04-25",
		Cover:   "/img/april-portrait.webp",
		Excerpt: "I've got passion burning in my bones! I have so much that I want to share and so many ideas that I want to explore. But… I can't do it alone. I, like the rest of us, need community to bring the concepts into reality.",
		Body:    passionPost,
	},
	{
		Slug:    "round-onegame-on",
		Title:   "Round One…Game On",
		Date:    "2024-04-08",
		Cover:   "/img/adaptogenic-room.jpeg",
		Excerpt: "Yesterday was my first ever game night. I've never been one to sit around and play games. But I'm starting to see it all coming to fruition.",
		Body:    gameNightPost,
	},
	{
		Slug:    "healing-the-inner-child-meditaion",
		Title:   "Healing the Inner Child Meditation",
		Date:    "2024-04-07",
		Cover:   "/img/blog-hero-bg.jpg",
		Excerpt: "Welcome to this guided meditation for healing the inner child. Find a comfortable position, either sitting or lying down, and allow yourself to fully relax.",
		Body:    innerChildPost,
	},
}

// FindPost returns the post with slug, or false.
func FindPost(slug string) (Post, bool) {
	for _, p := range Posts {
		if p.Slug == slug {
			return p, true
		}
	}
	return Post{}, false
}
