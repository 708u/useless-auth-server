package html

import (
	"embed"
	"html/template"
)

const (
	templatePath = "template/"
	templateHTML = ".html"
)

const (
	Index     = "index"
	Authorize = "authorize"
)

//go:embed template/*
var f embed.FS

// template
var T = template.Must(template.ParseFS(
	f,
	templatePath+"*"+templateHTML,
))
