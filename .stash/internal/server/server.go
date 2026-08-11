// Package server wires routes and handlers for the public site and admin.
package server

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/statusnone/4life-community/internal/db"
	"github.com/statusnone/4life-community/internal/render"
)

type Config struct {
	Store         *db.Store
	PublicTpl     *render.Template
	AdminTpl      *render.Template
	AdminPassword string
	SessionSecret []byte
}

type Server struct {
	store     *db.Store
	public    *render.Template
	admin     *render.Template
	cfg       Config
	signKey   []byte
}

func New(cfg Config) *Server {
	return &Server{
		store:   cfg.Store,
		public:  cfg.PublicTpl,
		admin:   cfg.AdminTpl,
		cfg:     cfg,
		signKey: cfg.SessionSecret,
	}
}

// SiteData carries data common to every public page.
type SiteData struct {
	Settings map[string]string
	Nav      []db.NavFolder
	Path     string
	Title    string
	MetaDesc string
}

func (s *Server) siteData(ctx context.Context, path, title, meta string) (*SiteData, error) {
	if meta == "" {
		meta = "April Bartlett's 4Life bridges grant writing, bioregional mapping, and AstroYoga—helping changemakers align strategy, soul, and planetary regeneration."
	}
	settings, err := s.store.Settings(ctx)
	if err != nil {
		return nil, err
	}
	if title == "" {
		title = settings["site_title"]
	} else {
		title = title + " — " + settings["site_title"]
	}
	nav, err := s.store.Nav(ctx)
	if err != nil {
		return nil, err
	}
	return &SiteData{Settings: settings, Nav: nav, Path: path, Title: title, MetaDesc: meta}, nil
}

// JSON helper for htmx error responses.
func (s *Server) renderError(w http.ResponseWriter, status int, msg string) {
	http.Error(w, msg, status)
}

// ----------------------------- auth -----------------------------

const sessionCookie = "4life_session"

func (s *Server) sign(payload string) string {
	mac := hmac.New(sha256.New, s.signKey)
	mac.Write([]byte(payload))
	return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}

func (s *Server) validSession(r *http.Request) bool {
	c, err := r.Cookie(sessionCookie)
	if err != nil {
		return false
	}
	parts := strings.SplitN(c.Value, ".", 2)
	if len(parts) != 2 {
		return false
	}
	payload, sig := parts[0], parts[1]
	if !hmac.Equal([]byte(sig), []byte(s.sign(payload))) {
		return false
	}
	raw, err := base64.RawURLEncoding.DecodeString(payload)
	if err != nil {
		return false
	}
	var exp time.Time
	if err := json.Unmarshal(raw, &exp); err != nil {
		return false
	}
	return time.Now().Before(exp)
}

func (s *Server) setSession(w http.ResponseWriter) {
	exp := time.Now().Add(30 * 24 * time.Hour)
	raw, _ := json.Marshal(exp)
	payload := base64.RawURLEncoding.EncodeToString(raw)
	http.SetCookie(w, &http.Cookie{
		Name:     sessionCookie,
		Value:    payload + "." + s.sign(payload),
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   30 * 24 * 3600,
	})
}

func (s *Server) clearSession(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{Name: sessionCookie, Value: "", Path: "/", MaxAge: -1})
}

func (s *Server) logError(r *http.Request, err error) {
	if err != nil {
		log.Printf("http %s %s: %v", r.Method, r.URL.Path, err)
	}
}
