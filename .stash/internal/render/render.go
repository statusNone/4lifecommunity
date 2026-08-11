// Package render loads templates, provides template funcs, and renders
// markdown content.
package render

import (
	"encoding/json"
	"html/template"
	"io"
	"path/filepath"
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
// used for section headlines and bodies that contain ** or * emphasis. The
// wrapping <p>...</p> block is stripped so the result can sit inside headings.
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

func DateLong(t time.Time) string {
	return t.Format("January 2, 2006")
}

func DateShort(t time.Time) string {
	return t.Format("1/2/06")
}

// ThemeClass maps a section theme to a CSS background class.
func ThemeClass(theme string) string {
	switch theme {
	case "dark", "black":
		return "sec-dark"
	case "bokeh":
		return "sec-bokeh"
	default:
		return "sec-white"
	}
}

// FuncMap shared by public and admin templates.
var funcs = template.FuncMap{
	"md":         Markdown,
	"mdInline":   MarkdownInline,
	"dateLong":   DateLong,
	"dateShort":  DateShort,
	"themeClass": ThemeClass,
	"json": func(v any) string {
		b, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			return ""
		}
		return string(b)
	},
	"lower": strings.ToLower,
}

// Loader parses the templates for a given set of root names. All files under
// the templates dir are parsed together, then Executed by name.
type Loader struct {
	root     string
	baseName string
}

func NewLoader(root, baseName string) *Loader {
	return &Loader{root: root, baseName: baseName}
}

// Template renders the named template with data.
type Template struct {
	tmpl *template.Template
}

func (l *Loader) Load() (*Template, error) {
	pattern := filepath.Join(l.root, "**", "*.html")
	t, err := template.New("").Funcs(funcs).ParseGlob(pattern)
	if err != nil {
		return nil, err
	}
	return &Template{tmpl: t}, nil
}

func (t *Template) Execute(w io.Writer, name string, data any) error {
	return t.tmpl.ExecuteTemplate(w, name, data)
}

// ExecuteToBuffer renders the named template into a string (for htmx partials).
func (t *Template) ExecuteToBuffer(name string, data any) (string, error) {
	var b strings.Builder
	if err := t.tmpl.ExecuteTemplate(&b, name, data); err != nil {
		return "", err
	}
	return b.String(), nil
}
