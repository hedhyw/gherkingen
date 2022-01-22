package model

// TemplateData contains root arguments for template.
type TemplateData struct {
	*GherkinDocument
	PackageName string `json:"PackageName"`
}

// GenerateArgs contains required arguments for generate.
type GenerateArgs struct {
	Format         Format
	InputSource    []byte
	TemplateSource []byte
	PackageName    string
}
