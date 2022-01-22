package model

type TemplateData struct {
	*GherkinDocument
	PackageName string `json:"PackageName"`
}

type GenerateArgs struct {
	Format         Format
	InputSource    []byte
	TemplateSource []byte
	PackageName    string
}
