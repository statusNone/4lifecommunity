package server

import (
	"net/http"
)

// htmx form handlers return a small partial that replaces the submitted form.

type formData struct {
	Ok bool
	// Success partial id
}

func (s *Server) handleContactSubmit(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		s.renderError(w, http.StatusBadRequest, "bad form")
		return
	}
	name := r.FormValue("name")
	email := r.FormValue("email")
	message := r.FormValue("message")
	extras := map[string]any{
		"topic": r.Form["topic"],
		"found": r.FormValue("found"),
	}
	if err := s.store.AddSubmission(r.Context(), "contact", name, email, message, extras); err != nil {
		s.logError(r, err)
		s.renderError(w, http.StatusInternalServerError, "could not save message")
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	s.logError(r, s.public.Execute(w, "partials/contact_success.html", map[string]any{}))
}

func (s *Server) handleNewsletterSubscribe(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		s.renderError(w, http.StatusBadRequest, "bad form")
		return
	}
	email := r.FormValue("email")
	if err := s.store.AddSubmission(r.Context(), "newsletter", "", email, "", nil); err != nil {
		s.logError(r, err)
		s.renderError(w, http.StatusInternalServerError, "could not subscribe")
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	s.logError(r, s.public.Execute(w, "partials/newsletter_success.html", map[string]any{}))
}

func (s *Server) handleSessionSubmit(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		s.renderError(w, http.StatusBadRequest, "bad form")
		return
	}
	name := r.FormValue("name")
	email := r.FormValue("email")
	extras := map[string]any{
		"calling":    r.FormValue("calling"),
		"sessionType": r.Form["session_type"],
		"availability": r.Form["availability"],
	}
	if err := s.store.AddSubmission(r.Context(), "session", name, email, "", extras); err != nil {
		s.logError(r, err)
		s.renderError(w, http.StatusInternalServerError, "could not save request")
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	s.logError(r, s.public.Execute(w, "partials/session_success.html", map[string]any{}))
}

func (s *Server) handleDonateSubmit(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		s.renderError(w, http.StatusBadRequest, "bad form")
		return
	}
	extras := map[string]any{
		"amount":    r.FormValue("amount"),
		"frequency": r.FormValue("frequency"),
	}
	if err := s.store.AddSubmission(r.Context(), "donation", "", "", "", extras); err != nil {
		s.logError(r, err)
		s.renderError(w, http.StatusInternalServerError, "could not save donation")
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	s.logError(r, s.public.Execute(w, "partials/donate_success.html", map[string]any{}))
}
