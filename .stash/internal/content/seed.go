// Package content seeds the database with the full copy of the
// 4life.community Squarespace site.
package content

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/statusnone/4life-community/internal/db"
)

// Seed inserts the site content if the database is empty. It is safe to run
// on every boot; it no-ops when content already exists.
func Seed(ctx context.Context, store *db.Store) error {
	empty, err := store.TableEmpty(ctx, "site_settings")
	if err != nil {
		return err
	}
	if !empty {
		return nil
	}

	tx, err := store.Tx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	if err := seedSettings(ctx, store); err != nil {
		return err
	}
	if err := seedNav(ctx, store); err != nil {
		return err
	}
	if err := seedPages(ctx, store); err != nil {
		return err
	}
	if err := seedPosts(ctx, store); err != nil {
		return err
	}

	if err := tx.Commit(context.Background()); err != nil {
		return err
	}
	log.Println("seed: site content inserted")
	return nil
}

func seedSettings(ctx context.Context, store *db.Store) error {
	settings := map[string]string{
		"site_title":        "4Life | Regenerative Strategy, AstroYoga & Bioregional Alchemy",
		"meta_description":  "April Bartlett’s 4Life bridges grant writing, bioregional mapping, and AstroYoga—helping changemakers align strategy, soul, and planetary regeneration.",
		"logo":              "/img/logo.png",
		"logo_alt":          "4Life | Regenerative Strategy, AstroYoga & Bioregional Alchemy",
		"instagram":         "https://www.instagram.com/amber.rising4love/",
		"youtube":           "https://www.youtube.com/@4life.AprilBartlett",
		"cta_headline":      "Let's Grow Something **Wild** Together",
		"cta_body":          "Got a regenerative project, grant puzzle, or bioregional dream? April collaborates with rebels, healers, and systems-weavers. Reach out—no templated responses, just real talk.",
		"cta_button":        "Contact April",
		"cta_url":           "/contact",
		"cta_bg":            "/img/colombia.png",
		"newsletter_title":  "Subscribe",
		"newsletter_body":   "Sign up with your email address to receive news and updates.",
		"footer_heading":    "Explore 4Life",
		"footer_copyright":  "© 2026",
		"footer_design":     "Design: Alcyone Reserve",
		"footer_design_url": "https://www.alcyonereserve.com",
		"footer_bg":         "/img/footer-tint.png",
		"contact_email":     "4life.AprilBartlett@gmail.com",
	}
	for k, v := range settings {
		if err := store.SetSetting(ctx, k, v); err != nil {
			return err
		}
	}
	return nil
}

func seedNav(ctx context.Context, store *db.Store) error {
	type folder struct {
		label, href string
		items       []struct{ label, href string; external bool }
	}
	folders := []folder{
		{"About April", "/about-april", []struct{ label, href string; external bool }{
			{"Meet April", "/meet-april", false},
			{"Grant Writing", "/grant-writing", false},
			{"Upcoming Events", "https://lu.ma/user/4Life_AstroYoga", true},
			{"Contact", "/contact", false},
			{"Podcast/Media Inquiries", "/podcast-media-inquiries", false},
		}},
		{"Trauma Healing", "/trauma-healing", []struct{ label, href string; external bool }{
			{"AstroYoga", "/astro-yoga", false},
			{"Adaptogenic Healing", "/adaptogenic-healing", false},
		}},
		{"Planetary Regeneration", "/regeneration", []struct{ label, href string; external bool }{
			{"Field Notes", "/blog", false},
			{"Bioregional Nomads", "/bioregional-nomads", false},
			{"Susurros del Agua", "/susurros-del-agua", false},
			{"Giving To A Cause", "/giving-page-2-1", false},
			{"Catatumbo River – Second Chances", "/catatumbo", false},
		}},
	}

	for i, f := range folders {
		var fid int64
		if err := store.Pool.QueryRow(ctx,
			`INSERT INTO nav_folders (label, href, sort) VALUES ($1,$2,$3) RETURNING id`,
			f.label, f.href, i).Scan(&fid); err != nil {
			return err
		}
		for j, it := range f.items {
			if _, err := store.Pool.Exec(ctx,
				`INSERT INTO nav_items (folder_id, label, href, external, sort) VALUES ($1,$2,$3,$4,$5)`,
				fid, it.label, it.href, it.external, j); err != nil {
				return err
			}
		}
	}
	return nil
}

type accordion struct{ title, body string }

type section struct {
	kind, theme, label, headline, body string
	bullets, gallery                   []string
	image, imageAlt, imageSide         string
	buttonLabel, buttonURL, bgImage    string
	bgOverlay                          float64
	divider                            int
	accordion                          []accordion
}

type page struct {
	slug, title, meta       string
	heroLabel, heroHeadline string
	heroBody, heroImage     string
	heroAlt, heroSide       string
	heroTheme               string
	sections                []section
}

func insertSection(ctx context.Context, store *db.Store, pageSlug string, s section, sort int) (int64, error) {
	bullets, _ := json.Marshal(s.bullets)
	if s.bullets == nil {
		bullets = []byte("[]")
	}
	gallery, _ := json.Marshal(s.gallery)
	if s.gallery == nil {
		gallery = []byte("[]")
	}
	var sid int64
	err := store.Pool.QueryRow(ctx, `
		INSERT INTO sections (page_slug, kind, theme, label, headline, body, bullets, gallery,
		                      image, image_alt, image_side, button_label, button_url, bg_image,
		                      bg_overlay, divider, sort)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17) RETURNING id`,
		pageSlug, s.kind, s.theme, s.label, s.headline, s.body, bullets, gallery,
		s.image, s.imageAlt, s.imageSide, s.buttonLabel, s.buttonURL, s.bgImage,
		s.bgOverlay, s.divider, sort).Scan(&sid)
	if err != nil {
		return 0, fmt.Errorf("insert section %q: %w", s.kind, err)
	}
	for i, a := range s.accordion {
		if _, err := store.Pool.Exec(ctx,
			`INSERT INTO accordion_items (section_id, title, body, sort) VALUES ($1,$2,$3,$4)`,
			sid, a.title, a.body, i); err != nil {
			return 0, err
		}
	}
	return sid, nil
}

func insertPage(ctx context.Context, store *db.Store, p page, sort int) error {
	_, err := store.Pool.Exec(ctx, `
		INSERT INTO pages (slug, title, meta_description, hero_label, hero_headline, hero_body,
		                   hero_image, hero_image_alt, hero_image_side, hero_theme, published, sort)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,true,$11)`,
		p.slug, p.title, p.meta, p.heroLabel, p.heroHeadline, p.heroBody,
		p.heroImage, p.heroAlt, p.heroSide, p.heroTheme, sort)
	if err != nil {
		return fmt.Errorf("insert page %q: %w", p.slug, err)
	}
	for i, s := range p.sections {
		if _, err := insertSection(ctx, store, p.slug, s, i); err != nil {
			return err
		}
	}
	return nil
}
