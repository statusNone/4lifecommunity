package server

import (
	"net/http"
	"strconv"
	"strings"
)

// handlePublic dispatches the public site.
func (s *Server) handlePublic(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if path == "/" {
		s.handleHome(w, r)
		return
	}
	if path == "/blog" {
		s.handleBlog(w, r)
		return
	}
	if strings.HasPrefix(path, "/blog/") {
		s.handlePost(w, r, strings.TrimPrefix(path, "/blog/"))
		return
	}
	s.handlePage(w, r, strings.TrimPrefix(path, "/"))
}

func (s *Server) handleHome(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	page, err := s.store.Page(ctx, "home")
	if err != nil || page == nil {
		s.renderError(w, http.StatusNotFound, "page not found")
		return
	}
	sections, err := s.store.PageSections(ctx, "home")
	if err != nil {
		s.logError(r, err)
		s.renderError(w, http.StatusInternalServerError, "server error")
		return
	}
	posts, err := s.store.Posts(ctx, 0, 3)
	if err != nil {
		s.logError(r, err)
	}
	data, err := s.siteData(ctx, "/", "", page.MetaDescription)
	if err != nil {
		s.logError(r, err)
		s.renderError(w, http.StatusInternalServerError, "server error")
		return
	}
	data.Title = page.Title
	if err := s.public.Execute(w, "home.html", struct {
		*SiteData
		Page     *pageView
		Latest   []postView
	}{data, &pageView{Page: page, Sections: sections}, postViews(posts)}); err != nil {
		s.logError(r, err)
	}
}

func (s *Server) handlePage(w http.ResponseWriter, r *http.Request, slug string) {
	if strings.Contains(slug, "/") {
		s.renderError(w, http.StatusNotFound, "not found")
		return
	}
	ctx := r.Context()
	page, err := s.store.Page(ctx, slug)
	if err != nil {
		s.logError(r, err)
		s.renderError(w, http.StatusInternalServerError, "server error")
		return
	}
	if page == nil || !page.Published {
		s.renderError(w, http.StatusNotFound, "not found")
		return
	}
	sections, err := s.store.PageSections(ctx, slug)
	if err != nil {
		s.logError(r, err)
		s.renderError(w, http.StatusInternalServerError, "server error")
		return
	}
	data, err := s.siteData(ctx, "/"+slug, page.Title, page.MetaDescription)
	if err != nil {
		s.logError(r, err)
		s.renderError(w, http.StatusInternalServerError, "server error")
		return
	}
	data.Title = page.Title
	err = s.public.Execute(w, "page.html", struct {
		*SiteData
		Page *pageView
	}{data, &pageView{Page: page, Sections: sections}})
	s.logError(r, err)
}

func (s *Server) handleBlog(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	posts, err := s.store.Posts(ctx, 0, postsPerPage)
	if err != nil {
		s.logError(r, err)
		s.renderError(w, http.StatusInternalServerError, "server error")
		return
	}
	total, err := s.store.PostsCount(ctx)
	if err != nil {
		s.logError(r, err)
	}
	data, err := s.siteData(ctx, "/blog", "April's Field Notes", "April Bartlett's 4Life bridges grant writing, bioregional mapping, and AstroYoga—helping changemakers align strategy, soul, and planetary regeneration.")
	if err != nil {
		s.logError(r, err)
		s.renderError(w, http.StatusInternalServerError, "server error")
		return
	}
	data.Title = "April's Field Notes — " + data.Settings["site_title"]
	hasMore := total > postsPerPage
	err = s.public.Execute(w, "blog.html", struct {
		*SiteData
		Posts   []postView
		HasMore bool
		NextOff int
	}{data, postViews(posts), hasMore, postsPerPage})
	s.logError(r, err)
}

func (s *Server) handlePost(w http.ResponseWriter, r *http.Request, slug string) {
	ctx := r.Context()
	post, err := s.store.Post(ctx, slug)
	if err != nil {
		s.logError(r, err)
		s.renderError(w, http.StatusInternalServerError, "server error")
		return
	}
	if post == nil {
		s.renderError(w, http.StatusNotFound, "not found")
		return
	}
	data, err := s.siteData(ctx, "/blog/"+slug, post.Title, "")
	if err != nil {
		s.logError(r, err)
		s.renderError(w, http.StatusInternalServerError, "server error")
		return
	}
	data.Title = post.Title + " — " + data.Settings["site_title"]
	err = s.public.Execute(w, "post.html", struct {
		*SiteData
		Post *postView
	}{data, postViewOf(post)})
	s.logError(r, err)
}

const postsPerPage = 8

// handleBlogPartial serves the appended posts grid for "Older Posts".
func (s *Server) handleBlogPartial(w http.ResponseWriter, r *http.Request) {
	off, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil || off < 0 {
		off = 0
	}
	ctx := r.Context()
	posts, err := s.store.Posts(ctx, off, postsPerPage)
	if err != nil {
		s.logError(r, err)
		s.renderError(w, http.StatusInternalServerError, "server error")
		return
	}
	total, _ := s.store.PostsCount(ctx)
	err = s.public.Execute(w, "partials/blog_grid.html", struct {
		Posts   []postView
		HasMore bool
		NextOff int
	}{postViews(posts), off+len(posts) < total, off + len(posts)})
	s.logError(r, err)
}
