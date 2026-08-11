package server

import (
	"net/http"
	"strings"

	"github.com/statusnone/4life-community/internal/db"
)

// ----------------------------- login -----------------------------

func (s *Server) handleAdminLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		if constantTimeEqual(r.FormValue("password"), s.cfg.AdminPassword) {
			s.setSession(w)
			http.Redirect(w, r, "/admin", http.StatusFound)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_ = s.admin.Execute(w, "admin_login.html", map[string]any{"Error": "Invalid password."})
		return
	}
	if s.validSession(r) {
		http.Redirect(w, r, "/admin", http.StatusFound)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_ = s.admin.Execute(w, "admin_login.html", map[string]any{})
}

func (s *Server) handleAdminLogout(w http.ResponseWriter, r *http.Request) {
	s.clearSession(w)
	http.Redirect(w, r, "/admin/login", http.StatusFound)
}

func constantTimeEqual(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	var diff byte
	for i := 0; i < len(a); i++ {
		diff |= a[i] ^ b[i]
	}
	return diff == 0
}

// ----------------------------- dashboard -----------------------------

type adminData struct {
	Settings map[string]string
	Title    string
}

func (s *Server) adminBase(r *http.Request, title string) (*adminData, error) {
	settings, err := s.store.Settings(r.Context())
	if err != nil {
		return nil, err
	}
	return &adminData{Settings: settings, Title: title + " — 4Life Admin"}, nil
}

func (s *Server) handleAdminDashboard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	base, err := s.adminBase(r, "Dashboard")
	if err != nil {
		s.logError(r, err)
		s.renderError(w, http.StatusInternalServerError, "server error")
		return
	}
	pages, _ := s.store.AllPages(ctx)
	posts, _ := s.store.AllPosts(ctx)
	subs, _ := s.store.Submissions(ctx, 5)
	err = s.admin.Execute(w, "admin_dashboard.html", struct {
		*adminData
		Pages       []dBPage
		Posts       []dBPost
		Submissions []db.Submission
	}{base, pages, posts, subs})
	s.logError(r, err)
}

// ----------------------------- posts -----------------------------

func (s *Server) handleAdminPosts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	posts, err := s.store.AllPosts(ctx)
	if err != nil {
		s.logError(r, err)
		s.renderError(w, http.StatusInternalServerError, "server error")
		return
	}
	base, _ := s.adminBase(r, "Posts")
	err = s.admin.Execute(w, "admin_posts.html", struct {
		*adminData
		Posts []dBPost
	}{base, posts})
	s.logError(r, err)
}

func (s *Server) handleAdminPostNew(w http.ResponseWriter, r *http.Request) {
	base, _ := s.adminBase(r, "New Post")
	err := s.admin.Execute(w, "admin_post_edit.html", struct {
		*adminData
		Post *dBPost
	}{base, nil})
	s.logError(r, err)
}

func (s *Server) handleAdminPostSave(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		s.renderError(w, http.StatusBadRequest, "bad form")
		return
	}
	slug := strings.TrimSpace(r.FormValue("slug"))
	if slug == "" {
		slug = slugify(r.FormValue("title"))
	}
	date := r.FormValue("published_at")
	if date == "" {
		date = "2025-01-01"
	}
	_, err := s.store.Pool.Exec(r.Context(), `
		INSERT INTO posts (slug, title, excerpt, cover, external_url, body, published_at, published)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
		ON CONFLICT (slug) DO UPDATE SET
			title = EXCLUDED.title, excerpt = EXCLUDED.excerpt, cover = EXCLUDED.cover,
			external_url = EXCLUDED.external_url, body = EXCLUDED.body,
			published_at = EXCLUDED.published_at, published = EXCLUDED.published`,
		slug, r.FormValue("title"), r.FormValue("excerpt"), r.FormValue("cover"),
		r.FormValue("external_url"), r.FormValue("body"), date,
		r.FormValue("published") == "on")
	if err != nil {
		s.logError(r, err)
		s.renderError(w, http.StatusInternalServerError, "could not save post")
		return
	}
	http.Redirect(w, r, "/admin/posts", http.StatusFound)
}

func (s *Server) handleAdminPostDelete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	slug := r.FormValue("slug")
	_, err := s.store.Pool.Exec(r.Context(), `DELETE FROM posts WHERE slug = $1`, slug)
	if err != nil {
		s.logError(r, err)
		s.renderError(w, http.StatusInternalServerError, "could not delete post")
		return
	}
	http.Redirect(w, r, "/admin/posts", http.StatusFound)
}

func (s *Server) handleAdminPostToggle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	slug := r.FormValue("slug")
	_, err := s.store.Pool.Exec(r.Context(),
		`UPDATE posts SET published = NOT published WHERE slug = $1`, slug)
	if err != nil {
		s.logError(r, err)
		s.renderError(w, http.StatusInternalServerError, "could not toggle post")
		return
	}
	http.Redirect(w, r, "/admin/posts", http.StatusFound)
}

// ----------------------------- pages -----------------------------

func (s *Server) handleAdminPages(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	pages, err := s.store.AllPages(ctx)
	if err != nil {
		s.logError(r, err)
		s.renderError(w, http.StatusInternalServerError, "server error")
		return
	}
	base, _ := s.adminBase(r, "Pages")

	type pageWithSections struct {
		Page     *dBPage
		Sections []db.Section
	}
	var list []pageWithSections
	for i := range pages {
		secs, err := s.store.PageSections(ctx, pages[i].Slug)
		if err != nil {
			s.logError(r, err)
			continue
		}
		list = append(list, pageWithSections{Page: &pages[i], Sections: secs})
	}
	err = s.admin.Execute(w, "admin_pages.html", struct {
		*adminData
		Pages []pageWithSections
	}{base, list})
	s.logError(r, err)
}

func (s *Server) handleAdminPageSave(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		s.renderError(w, http.StatusBadRequest, "bad form")
		return
	}
	ctx := r.Context()
	slug := r.FormValue("slug")

	// update page hero/meta fields
	_, err := s.store.Pool.Exec(ctx, `
		UPDATE pages SET title=$2, meta_description=$3, hero_label=$4, hero_headline=$5,
			hero_body=$6, hero_image=$7, hero_image_alt=$8, published=$9
		WHERE slug=$1`,
		slug, r.FormValue("title"), r.FormValue("meta_description"), r.FormValue("hero_label"),
		r.FormValue("hero_headline"), r.FormValue("hero_body"), r.FormValue("hero_image"),
		r.FormValue("hero_image_alt"), r.FormValue("published") == "on")
	if err != nil {
		s.logError(r, err)
		s.renderError(w, http.StatusInternalServerError, "could not save page")
		return
	}

	// update section headline/body/image for each provided section id
	for key := range r.Form {
		if !strings.HasPrefix(key, "section_headline_") {
			continue
		}
		id := strings.TrimPrefix(key, "section_headline_")
		if _, err := s.store.Pool.Exec(ctx,
			`UPDATE sections SET headline=$2, body=$3, image=$4, image_alt=$5, button_label=$6, button_url=$7 WHERE id=$1`,
			id, r.FormValue("section_headline_"+id), r.FormValue("section_body_"+id),
			r.FormValue("section_image_"+id), r.FormValue("section_image_alt_"+id),
			r.FormValue("section_button_label_"+id), r.FormValue("section_button_url_"+id)); err != nil {
			s.logError(r, err)
		}
	}

	http.Redirect(w, r, "/admin/pages", http.StatusFound)
}

// ----------------------------- settings -----------------------------

func (s *Server) handleAdminSettings(w http.ResponseWriter, r *http.Request) {
	base, err := s.adminBase(r, "Settings")
	if err != nil {
		s.logError(r, err)
		s.renderError(w, http.StatusInternalServerError, "server error")
		return
	}
	err = s.admin.Execute(w, "admin_settings.html", base)
	s.logError(r, err)
}

func (s *Server) handleAdminSettingsSave(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		s.renderError(w, http.StatusBadRequest, "bad form")
		return
	}
	keys := []string{"site_title", "meta_description", "logo", "instagram", "youtube",
		"cta_headline", "cta_body", "cta_button", "cta_url", "cta_bg",
		"newsletter_title", "newsletter_body", "contact_email"}
	for _, k := range keys {
		if err := s.store.SetSetting(r.Context(), k, r.FormValue(k)); err != nil {
			s.logError(r, err)
		}
	}
	http.Redirect(w, r, "/admin/settings", http.StatusFound)
}

// ----------------------------- submissions -----------------------------

func (s *Server) handleAdminSubmissions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	subs, err := s.store.Submissions(ctx, 200)
	if err != nil {
		s.logError(r, err)
		s.renderError(w, http.StatusInternalServerError, "server error")
		return
	}
	base, _ := s.adminBase(r, "Submissions")
	err = s.admin.Execute(w, "admin_submissions.html", struct {
		*adminData
		Submissions []db.Submission
	}{base, subs})
	s.logError(r, err)
}

func (s *Server) handleAdminSubmissionDelete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")
	_, err := s.store.Pool.Exec(r.Context(), `DELETE FROM submissions WHERE id = $1`, id)
	if err != nil {
		s.logError(r, err)
		s.renderError(w, http.StatusInternalServerError, "could not delete submission")
		return
	}
	http.Redirect(w, r, "/admin/submissions", http.StatusFound)
}

func slugify(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	var b strings.Builder
	lastDash := false
	for _, c := range s {
		switch {
		case c >= 'a' && c <= 'z' || c >= '0' && c <= '9':
			b.WriteRune(c)
			lastDash = false
		default:
			if !lastDash && b.Len() > 0 {
				b.WriteByte('-')
				lastDash = true
			}
		}
	}
	return strings.Trim(b.String(), "-")
}
