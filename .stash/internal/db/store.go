package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type NavFolder struct {
	ID    int64
	Label string
	Href  string
	Items []NavItem
}

type NavItem struct {
	ID       int64
	Label    string
	Href     string
	External bool
}

type Page struct {
	Slug            string
	Title           string
	MetaDescription string
	HeroLabel       string
	HeroHeadline    string
	HeroBody        string
	HeroImage       string
	HeroImageAlt    string
	HeroImageSide   string
	HeroTheme       string
	Published       bool
	Sort            int
}

type Section struct {
	ID          int64
	PageSlug    string
	Kind        string
	Theme       string
	Label       string
	Headline    string
	Body        string
	Bullets     []string
	Image       string
	ImageAlt    string
	ImageSide   string
	Gallery     []string
	ButtonLabel string
	ButtonURL   string
	BgImage     string
	BgOverlay   float64
	Divider     int
	Sort        int
	Accordion   []AccordionItem
}

type AccordionItem struct {
	ID    int64
	Title string
	Body  string
}

type Post struct {
	Slug        string
	Title       string
	Excerpt     string
	Cover       string
	ExternalURL string
	Body        string
	PublishedAt time.Time
	Published   bool
}

type Submission struct {
	ID        int64
	Kind      string
	Name      string
	Email     string
	Message   string
	Extras    map[string]any
	CreatedAt time.Time
}

const navQuery = `
	SELECT id, label, href FROM nav_folders ORDER BY sort`

const navItemsQuery = `
	SELECT id, folder_id, label, href, external FROM nav_items ORDER BY sort`

// Nav returns folders with their items, plus the settings-derived brand.
func (s *Store) Nav(ctx context.Context) ([]NavFolder, error) {
	rows, err := s.Pool.Query(ctx, navQuery)
	if err != nil {
		return nil, err
	}
	folders, err := pgx.CollectRows(rows, func(r pgx.CollectableRow) (NavFolder, error) {
		var f NavFolder
		err := r.Scan(&f.ID, &f.Label, &f.Href)
		return f, err
	})
	if err != nil {
		return nil, err
	}

	itemRows, err := s.Pool.Query(ctx, navItemsQuery)
	if err != nil {
		return nil, err
	}
	type wi struct {
		folderID int64
		item     NavItem
	}
	pairs, err := pgx.CollectRows(itemRows, func(r pgx.CollectableRow) (wi, error) {
		var p wi
		err := r.Scan(&p.item.ID, &p.folderID, &p.item.Label, &p.item.Href, &p.item.External)
		return p, err
	})
	if err != nil {
		return nil, err
	}
	for _, p := range pairs {
		for i := range folders {
			if folders[i].ID == p.folderID {
				folders[i].Items = append(folders[i].Items, p.item)
				break
			}
		}
	}
	return folders, nil
}

func (s *Store) Page(ctx context.Context, slug string) (*Page, error) {
	row := s.Pool.QueryRow(ctx, `
		SELECT slug, title, meta_description, COALESCE(hero_label,''), COALESCE(hero_headline,''),
		       COALESCE(hero_body,''), COALESCE(hero_image,''), COALESCE(hero_image_alt,''),
		       COALESCE(hero_image_side,'right'), hero_theme, published, sort
		FROM pages WHERE slug = $1`, slug)
	var p Page
	err := row.Scan(&p.Slug, &p.Title, &p.MetaDescription, &p.HeroLabel, &p.HeroHeadline,
		&p.HeroBody, &p.HeroImage, &p.HeroImageAlt, &p.HeroImageSide, &p.HeroTheme, &p.Published, &p.Sort)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (s *Store) PageSections(ctx context.Context, slug string) ([]Section, error) {
	rows, err := s.Pool.Query(ctx, `
		SELECT id, page_slug, kind, theme, COALESCE(label,''), COALESCE(headline,''),
		       COALESCE(body,''), COALESCE(bullets,'[]'::jsonb),
		       COALESCE(image,''), COALESCE(image_alt,''), COALESCE(image_side,'left'),
		       COALESCE(gallery,'[]'::jsonb),
		       COALESCE(button_label,''), COALESCE(button_url,''), COALESCE(bg_image,''),
		       COALESCE(bg_overlay,0)::float8, COALESCE(divider,0), sort
		FROM sections WHERE page_slug = $1 ORDER BY sort`, slug)
	if err != nil {
		return nil, err
	}
	sections, err := pgx.CollectRows(rows, func(r pgx.CollectableRow) (Section, error) {
		var s Section
		err := r.Scan(&s.ID, &s.PageSlug, &s.Kind, &s.Theme, &s.Label, &s.Headline,
			&s.Body, &s.Bullets, &s.Image, &s.ImageAlt, &s.ImageSide, &s.Gallery,
			&s.ButtonLabel, &s.ButtonURL, &s.BgImage, &s.BgOverlay, &s.Divider, &s.Sort)
		return s, err
	})
	if err != nil {
		return nil, err
	}

	for i := range sections {
		aRows, err := s.Pool.Query(ctx, `
			SELECT id, title, body FROM accordion_items
			WHERE section_id = $1 ORDER BY sort`, sections[i].ID)
		if err != nil {
			return nil, err
		}
		items, err := pgx.CollectRows(aRows, func(r pgx.CollectableRow) (AccordionItem, error) {
			var a AccordionItem
			err := r.Scan(&a.ID, &a.Title, &a.Body)
			return a, err
		})
		if err != nil {
			return nil, err
		}
		sections[i].Accordion = items
	}
	return sections, nil
}

func (s *Store) AllPages(ctx context.Context) ([]Page, error) {
	rows, err := s.Pool.Query(ctx, `
		SELECT slug, title, meta_description, COALESCE(hero_label,''), COALESCE(hero_headline,''),
		       COALESCE(hero_body,''), COALESCE(hero_image,''), COALESCE(hero_image_alt,''),
		       COALESCE(hero_image_side,'right'), hero_theme, published, sort
		FROM pages ORDER BY sort`)
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, func(r pgx.CollectableRow) (Page, error) {
		var p Page
		err := r.Scan(&p.Slug, &p.Title, &p.MetaDescription, &p.HeroLabel, &p.HeroHeadline,
			&p.HeroBody, &p.HeroImage, &p.HeroImageAlt, &p.HeroImageSide, &p.HeroTheme, &p.Published, &p.Sort)
		return p, err
	})
}

func (s *Store) Posts(ctx context.Context, offset, limit int) ([]Post, error) {
	rows, err := s.Pool.Query(ctx, `
		SELECT slug, title, excerpt, cover, external_url, body, published_at, published
		FROM posts WHERE published = true
		ORDER BY published_at DESC, slug
		LIMIT $2 OFFSET $1`, offset, limit)
	if err != nil {
		return nil, err
	}
	return s.scanPosts(rows)
}

func (s *Store) PostsCount(ctx context.Context) (int, error) {
	var n int
	err := s.Pool.QueryRow(ctx, `SELECT count(*) FROM posts WHERE published = true`).Scan(&n)
	return n, err
}

func (s *Store) Post(ctx context.Context, slug string) (*Post, error) {
	row := s.Pool.QueryRow(ctx, `
		SELECT slug, title, excerpt, cover, external_url, body, published_at, published
		FROM posts WHERE slug = $1 AND published = true`, slug)
	var p Post
	err := row.Scan(&p.Slug, &p.Title, &p.Excerpt, &p.Cover, &p.ExternalURL, &p.Body, &p.PublishedAt, &p.Published)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (s *Store) AllPosts(ctx context.Context) ([]Post, error) {
	rows, err := s.Pool.Query(ctx, `
		SELECT slug, title, excerpt, cover, external_url, body, published_at, published
		FROM posts ORDER BY published_at DESC, slug`)
	if err != nil {
		return nil, err
	}
	return s.scanPosts(rows)
}

func (s *Store) scanPosts(rows pgx.Rows) ([]Post, error) {
	defer rows.Close()
	var out []Post
	for rows.Next() {
		var p Post
		if err := rows.Scan(&p.Slug, &p.Title, &p.Excerpt, &p.Cover, &p.ExternalURL, &p.Body, &p.PublishedAt, &p.Published); err != nil {
			return nil, err
		}
		out = append(out, p)
	}
	return out, rows.Err()
}

func (s *Store) Setting(ctx context.Context, key string) (string, error) {
	var v string
	err := s.Pool.QueryRow(ctx, `SELECT value FROM site_settings WHERE key = $1`, key).Scan(&v)
	if err == pgx.ErrNoRows {
		return "", nil
	}
	return v, err
}

func (s *Store) Settings(ctx context.Context) (map[string]string, error) {
	rows, err := s.Pool.Query(ctx, `SELECT key, value FROM site_settings`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := map[string]string{}
	for rows.Next() {
		var k, v string
		if err := rows.Scan(&k, &v); err != nil {
			return nil, err
		}
		out[k] = v
	}
	return out, rows.Err()
}

func (s *Store) SetSetting(ctx context.Context, key, value string) error {
	_, err := s.Pool.Exec(ctx, `
		INSERT INTO site_settings (key, value) VALUES ($1,$2)
		ON CONFLICT (key) DO UPDATE SET value = EXCLUDED.value`, key, value)
	return err
}

func (s *Store) AddSubmission(ctx context.Context, kind, name, email, message string, extras map[string]any) error {
	_, err := s.Pool.Exec(ctx, `
		INSERT INTO submissions (kind, name, email, message, extras)
		VALUES ($1,$2,$3,$4,$5)`, kind, name, email, message, extras)
	return err
}

func (s *Store) Submissions(ctx context.Context, limit int) ([]Submission, error) {
	rows, err := s.Pool.Query(ctx, `
		SELECT id, kind, name, email, message, COALESCE(extras,'{}'::jsonb), created_at
		FROM submissions ORDER BY created_at DESC LIMIT $1`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Submission
	for rows.Next() {
		var s Submission
		if err := rows.Scan(&s.ID, &s.Kind, &s.Name, &s.Email, &s.Message, &s.Extras, &s.CreatedAt); err != nil {
			return nil, err
		}
		out = append(out, s)
	}
	return out, rows.Err()
}
