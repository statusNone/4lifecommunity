package server

import (
	"github.com/statusnone/4life-community/internal/db"
)

type dBPage = db.Page
type dBPost = db.Post

// pageView enriches a db.Page with its ordered sections. The embedded *db.Page
// keeps field access like .Page.HeroHeadline working in templates.
type pageView struct {
	*db.Page
	Sections []db.Section
}

// postView embeds the post and adds display helpers used by the templates.
type postView struct {
	db.Post
	URL string
}

func postViewOf(p *db.Post) *postView {
	return &postView{Post: *p, URL: "/blog/" + p.Slug}
}

func postViews(ps []db.Post) []postView {
	out := make([]postView, 0, len(ps))
	for i := range ps {
		out = append(out, postView{Post: ps[i], URL: "/blog/" + ps[i].Slug})
	}
	return out
}
