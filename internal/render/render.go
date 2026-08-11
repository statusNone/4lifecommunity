// Package render loads html/template views and converts markdown content
// to safe HTML.
package render

import (
	"html/template"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/yuin/goldmark"
)

var md = goldmark.New()

var videoRe = regexp.MustCompile(`(?m)^\[VIDEO:\s*(https?://[^\]]+)\]\s*$`)

// Markdown converts post body markdown into safe HTML, expanding
// [VIDEO: url] markers into responsive iframe embeds.
func Markdown(body string) template.HTML {
	body = videoRe.ReplaceAllString(body, `<div class="video-embed"><iframe src="$1" title="Embedded video" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen loading="lazy"></iframe></div>`)
	var buf strings.Builder
	if err := md.Convert([]byte(body), &buf); err != nil {
		return template.HTML(template.HTMLEscapeString(body))
	}
	return template.HTML(buf.String())
}

// MarkdownInline renders a short markdown string (single line-ish) as HTML,
// used for headlines and bodies that contain ** or * emphasis. The wrapping
// <p>...</p> block is stripped so the result can sit inside headings.
func MarkdownInline(s string) template.HTML {
	if strings.TrimSpace(s) == "" {
		return ""
	}
	var buf strings.Builder
	if err := md.Convert([]byte(s), &buf); err != nil {
		return template.HTML(template.HTMLEscapeString(s))
	}
	out := strings.TrimSpace(buf.String())
	out = strings.TrimPrefix(out, "<p>")
	out = strings.TrimSuffix(out, "</p>")
	return template.HTML(out)
}

// DateLong formats a date like "January 2, 2006".
func DateLong(t time.Time) string {
	return t.Format("January 2, 2006")
}

// First returns the first n items of a slice.
func First(n int, items any) any {
	s := reflect.ValueOf(items)
	if s.Kind() != reflect.Slice && s.Kind() != reflect.Array {
		return items
	}
	if n >= s.Len() {
		return items
	}
	return s.Slice(0, n).Interface()
}

// SplitLines splits a string on blank-line paragraph breaks.
func SplitLines(s string) []string {
	return strings.Split(strings.TrimSpace(s), "\n\n")
}

func Funcs() template.FuncMap {
	return template.FuncMap{
		"md":         Markdown,
		"mdInline":   MarkdownInline,
		"dateLong":   DateLong,
		"first":      First,
		"splitLines": SplitLines,
		"hasPrefix":  strings.HasPrefix,
	}
}

// Loader builds one template set per page: base layout + shared partials +
// the page's own content/title blocks.
type Loader struct {
	pages    map[string]*template.Template
	partials *template.Template
}

func NewLoader(dir string) (*Loader, error) {
	l := &Loader{pages: map[string]*template.Template{}}

	basePath := filepath.Join(dir, "base.html")
	pagePaths, err := filepath.Glob(filepath.Join(dir, "pages", "*.html"))
	if err != nil {
		return nil, err
	}

	// Shared partials template (also used to render standalone partials).
	if l.partials, err = template.New("partials").Funcs(Funcs()).
		ParseGlob(filepath.Join(dir, "partials", "*.html")); err != nil {
		return nil, err
	}

	for _, pagePath := range pagePaths {
		name := filepath.Base(pagePath)
		t, err := template.New("base").Funcs(Funcs()).
			ParseGlob(filepath.Join(dir, "partials", "*.html"))
		if err != nil {
			return nil, err
		}
		if t, err = t.ParseFiles(basePath); err != nil {
			return nil, err
		}
		if t, err = t.ParseFiles(pagePath); err != nil {
			return nil, err
		}
		l.pages[name] = t
	}
	return l, nil
}

// Render executes a page ("pages/home.html") or a standalone partial
// ("contact_success") with data.
func (l *Loader) Render(w io.Writer, name string, data any) error {
	if strings.HasPrefix(name, "pages/") {
		t, ok := l.pages[strings.TrimPrefix(name, "pages/")]
		if !ok {
			return os.ErrNotExist
		}
		return t.ExecuteTemplate(w, "base", data)
	}
	return l.partials.ExecuteTemplate(w, name, data)
}
