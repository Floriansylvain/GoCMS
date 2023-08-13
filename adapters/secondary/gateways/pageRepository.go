package gateways

import (
	"GohCMS2/domain/gateways"
	"bytes"
	"embed"
	"html/template"
)

type PageRepository struct{}

//go:embed web/templates/*
var templateFiles embed.FS

func NewPageRepository() *PageRepository {
	return &PageRepository{}
}

func (p *PageRepository) Get(name string, data interface{}) ([]byte, error) {
	var processedHTML bytes.Buffer
	tmpl, err := template.ParseFS(templateFiles, "web/templates/"+name+".html")
	if err != nil {
		return nil, err
	}
	err = tmpl.Execute(&processedHTML, data)
	if err != nil {
		return nil, err
	}
	return processedHTML.Bytes(), nil
}

var _ gateways.IPageRepository = &PageRepository{}
