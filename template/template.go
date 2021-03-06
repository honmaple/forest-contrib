package template

import (
	"bytes"
	"html/template"
	"io"
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"
)

type EmbedLoader struct {
	fs fs.FS
}

func NewEmbedLoader(f fs.FS) *EmbedLoader {
	return &EmbedLoader{f}
}

func (s *EmbedLoader) Get(path string) (io.Reader, error) {
	buf, err := fs.ReadFile(s.fs, path)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(buf), nil
}

func (s *EmbedLoader) Abs(base, name string) string {
	if filepath.IsAbs(name) || strings.HasPrefix(name, "templates/") {
		return name
	}
	return filepath.Join("templates", name)
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w http.ResponseWriter, name string, data interface{}) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func New(f fs.FS, patterns ...string) *Template {
	return &Template{template.Must(template.ParseFS(f, patterns...))}
}
