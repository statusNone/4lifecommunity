// Package server wires routes, handlers, and template rendering for the
// 4Life Community site.
package server

import (
	"log"
	"net/http"
	"time"

	"github.com/statusnone/4life-community/internal/content"
	"github.com/statusnone/4life-community/internal/render"
)

// Server holds dependencies for HTTP handling.
type Server struct {
	tpl *render.Loader
}

func New(tpl *render.Loader) *Server {
	return &Server{tpl: tpl}
}

// page wraps common data passed to every page template.
type page struct {
	Site    content.Settings
	Current string
	Extra   any
}

func (s *Server) render(w http.ResponseWriter, name string, data page) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := s.tpl.Render(w, name, data); err != nil {
		log.Printf("render %s: %v", name, err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func (s *Server) pageFor(r *http.Request) page {
	return page{
		Site:    content.Site,
		Current: r.URL.Path,
	}
}

func (s *Server) handleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		s.handleNotFound(w, r)
		return
	}
	p := s.pageFor(r)
	p.Extra = map[string]any{
		"Hero":     content.Hero,
		"About":    content.About,
		"Projects": content.Projects,
		"Posts":    content.Posts,
	}
	s.render(w, "pages/home.html", p)
}

func (s *Server) handleAbout(w http.ResponseWriter, r *http.Request) {
	p := s.pageFor(r)
	p.Extra = map[string]any{
		"About": content.About,
	}
	s.render(w, "pages/about.html", p)
}

func (s *Server) handleBlog(w http.ResponseWriter, r *http.Request) {
	p := s.pageFor(r)
	p.Extra = map[string]any{
		"Posts": content.Posts,
	}
	s.render(w, "pages/blog.html", p)
}

func (s *Server) handlePost(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	post, ok := content.FindPost(slug)
	if !ok {
		s.handleNotFound(w, r)
		return
	}
	p := s.pageFor(r)
	p.Extra = map[string]any{
		"Post":   post,
		"Others": content.Posts,
	}
	s.render(w, "pages/post.html", p)
}

func (s *Server) handleProjects(w http.ResponseWriter, r *http.Request) {
	p := s.pageFor(r)
	p.Extra = map[string]any{
		"Projects": content.Projects,
	}
	s.render(w, "pages/projects.html", p)
}

func (s *Server) handleContact(w http.ResponseWriter, r *http.Request) {
	p := s.pageFor(r)
	p.Extra = map[string]any{
		"Sent": false,
	}
	s.render(w, "pages/contact.html", p)
}

// handleContactSubmit is an htmx endpoint. It renders a thank-you partial
// that swaps into the form container.
func (s *Server) handleContactSubmit(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "bad form", http.StatusBadRequest)
		return
	}
	log.Printf("contact submission from %q: %s", r.FormValue("email"), r.FormValue("name"))
	p := s.pageFor(r)
	p.Extra = map[string]any{
		"Name": r.FormValue("name"),
	}
	s.tpl.Render(w, "contact_success", p)
}

func (s *Server) handleNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	s.render(w, "pages/404.html", s.pageFor(r))
}

func (s *Server) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("static/img"))))
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/img/favicon-light.webp")
	})

	mux.HandleFunc("/", s.handleHome)
	mux.HandleFunc("/about", s.handleAbout)
	mux.HandleFunc("/blog", s.handleBlog)
	mux.HandleFunc("/blog/{slug}", s.handlePost)
	mux.HandleFunc("/projects", s.handleProjects)
	mux.HandleFunc("/contact", s.handleContact)
	mux.HandleFunc("/contact/submit", s.handleContactSubmit)

	return logging(mux)
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s (%s)", r.Method, r.URL.Path, time.Since(start).Round(time.Millisecond))
	})
}
