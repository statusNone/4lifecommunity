package server

import (
	"net/http"
)

// Routes builds the full HTTP handler.
func (s *Server) Routes() http.Handler {
	mux := http.NewServeMux()

	// static assets
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("static/img"))))
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/img/logo-light.png")
	})

	// public site
	mux.HandleFunc("/", s.handlePublic)

	// htmx endpoints
	mux.HandleFunc("/contact/submit", s.handleContactSubmit)
	mux.HandleFunc("/newsletter/subscribe", s.handleNewsletterSubscribe)
	mux.HandleFunc("/session/submit", s.handleSessionSubmit)
	mux.HandleFunc("/donate/submit", s.handleDonateSubmit)
	mux.HandleFunc("/blog/partial", s.handleBlogPartial)

	// admin
	mux.HandleFunc("/admin", s.requireAuth(s.handleAdminDashboard))
	mux.HandleFunc("/admin/login", s.handleAdminLogin)
	mux.HandleFunc("/admin/logout", s.handleAdminLogout)
	mux.HandleFunc("/admin/posts", s.requireAuth(s.handleAdminPosts))
	mux.HandleFunc("/admin/posts/new", s.requireAuth(s.handleAdminPostNew))
	mux.HandleFunc("/admin/posts/save", s.requireAuth(s.handleAdminPostSave))
	mux.HandleFunc("/admin/posts/delete", s.requireAuth(s.handleAdminPostDelete))
	mux.HandleFunc("/admin/posts/toggle", s.requireAuth(s.handleAdminPostToggle))
	mux.HandleFunc("/admin/pages", s.requireAuth(s.handleAdminPages))
	mux.HandleFunc("/admin/pages/save", s.requireAuth(s.handleAdminPageSave))
	mux.HandleFunc("/admin/settings", s.requireAuth(s.handleAdminSettings))
	mux.HandleFunc("/admin/settings/save", s.requireAuth(s.handleAdminSettingsSave))
	mux.HandleFunc("/admin/submissions", s.requireAuth(s.handleAdminSubmissions))
	mux.HandleFunc("/admin/submissions/delete", s.requireAuth(s.handleAdminSubmissionDelete))

	return s.logging(mux)
}

func (s *Server) requireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !s.validSession(r) {
			http.Redirect(w, r, "/admin/login", http.StatusFound)
			return
		}
		next(w, r)
	}
}

func (s *Server) logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
