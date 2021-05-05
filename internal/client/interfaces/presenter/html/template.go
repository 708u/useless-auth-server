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
	// index
	Index = "index"

	// fetch resource
	FetchResourceIndex = "fetch_resource"
)

//go:embed template/*
var f embed.FS

// template
var T = template.Must(template.ParseFS(
	f,
	templatePath+"*"+templateHTML,
))
