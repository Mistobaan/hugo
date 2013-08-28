package hugolib

import (
	"path/filepath"
	"html/template"
	"io"
)

// HTML encapsulates a known safe HTML document fragment.
// It should not be used for HTML from a third-party, or HTML with
// unclosed tags or comments. The outputs of a sound HTML sanitizer
// and a template escaped by this package are fine for use with HTML.
type HTML template.URL

type Template interface {
	ExecuteTemplate(wr io.Writer, name string, data interface{}) error
	Lookup(name string) *template.Template
	Templates() []*template.Template
	New(name string) *template.Template
}

type URL template.URL

type GoHtmlTemplate struct {
	template.Template
}

func NewTemplate() *GoHtmlTemplate {
	var templates = &GoHtmlTemplate{
		Template: *template.New(""),
	}

	funcMap := template.FuncMap{
		"urlize":    Urlize,
		"gt":        Gt,
		"isset":     IsSet,
		"echoParam": ReturnWhenSet,
	}

	templates.Funcs(funcMap)
	return templates
}

func (s *Site) addTemplate(name, tmpl string) (err error) {
	_, err = s.Tmpl.New(name).Parse(tmpl)
	return
}

func (s *Site) generateTemplateNameFrom(path string) (name string) {
	name = filepath.ToSlash(path[len(s.absLayoutDir())+1:])
	return
}

func (s *Site) primeTemplates() {
	alias := "<!DOCTYPE html>\n <html>\n <head>\n <link rel=\"canonical\" href=\"{{ .Permalink }}\"/>\n <meta http-equiv=\"content-type\" content=\"text/html; charset=utf-8\" />\n <meta http-equiv=\"refresh\" content=\"0;url={{ .Permalink }}\" />\n </head>\n </html>"
	alias_xhtml := "<!DOCTYPE html>\n <html xmlns=\"http://www.w3.org/1999/xhtml\">\n <head>\n <link rel=\"canonical\" href=\"{{ .Permalink }}\"/>\n <meta http-equiv=\"content-type\" content=\"text/html; charset=utf-8\" />\n <meta http-equiv=\"refresh\" content=\"0;url={{ .Permalink }}\" />\n </head>\n </html>"

	s.addTemplate("alias", alias)
	s.addTemplate("alias-xhtml", alias_xhtml)

}
