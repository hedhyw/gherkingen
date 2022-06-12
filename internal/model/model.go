package model

// TemplateData contains root arguments for template.
type TemplateData struct {
	*GherkinDocument
	PackageName string `json:"PackageName"`
}
