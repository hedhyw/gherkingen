package goplugin_test

import (
	"context"
	"testing"

	"github.com/hedhyw/gherkingen/v2/internal/docplugin/goplugin"
	"github.com/hedhyw/gherkingen/v2/internal/model"

	"github.com/stretchr/testify/assert"
)

func TestGoPluginName(t *testing.T) {
	t.Parallel()

	p := goplugin.New()
	assert.Equal(t, "GoPlugin", p.Name())
}

func TestGoPluginProcess(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	p := goplugin.New()

	t.Run("Background", func(t *testing.T) {
		t.Parallel()

		doc := getExampleDocument()
		if assert.NoError(t, p.Process(ctx, doc)) {
			pd := doc.Feature.Children[0].Background.PluginData
			assert.Equal(t, "\"Name\"", pd["GoValue"])
			assert.Equal(t, "Keyword", pd["GoName"])
			assert.Equal(t, "string", pd["GoType"])
		}
	})

	t.Run("Examples", func(t *testing.T) {
		t.Parallel()

		doc := getExampleDocument()
		if assert.NoError(t, p.Process(ctx, doc)) {
			pd := doc.Feature.Children[0].Scenario.Examples[0].PluginData
			assert.Equal(t, "\"Keyword\"", pd["GoValue"])
			assert.Equal(t, "Name", pd["GoName"])
		}
	})

	t.Run("Examples_EmptyTableBody_NoError", func(t *testing.T) {
		t.Parallel()

		doc := getExampleDocument()
		doc.Feature.Children[0].Scenario.Examples[0].TableBody = nil
		assert.NoError(t, p.Process(ctx, doc))
	})

	t.Run("Examples_TableBody_TableHeader_mismatch", func(t *testing.T) {
		t.Parallel()

		doc := getExampleDocument()
		doc.Feature.Children[0].Scenario.Examples[0].TableHeader.Cells = nil
		assert.Error(t, p.Process(ctx, doc))
	})

	t.Run("Examples_underscore", func(t *testing.T) {
		t.Parallel()

		// It tests https://github.com/hedhyw/gherkingen/v2/issues/26.

		doc := getExampleDocument()
		if assert.NoError(t, p.Process(ctx, doc)) {
			pd := doc.Feature.Children[0].Scenario.Examples[1].TableBody[0].PluginData
			assert.Equal(t, "\"hello_world\"", pd["GoValue"])
		}
	})

	t.Run("TableCell", func(t *testing.T) {
		t.Parallel()

		doc := getExampleDocument()
		if assert.NoError(t, p.Process(ctx, doc)) {
			pd := doc.Feature.Children[0].Scenario.Examples[0].TableHeader.Cells[0].PluginData
			assert.Equal(t, "Title", pd["GoName"])
			assert.Equal(t, "\"<Title>\"", pd["GoValue"])

			pd = doc.Feature.Children[0].Scenario.Examples[0].TableHeader.Cells[0].PluginData
			assert.Equal(t, "int", pd["GoType"])
		}
	})

	t.Run("Feature", func(t *testing.T) {
		t.Parallel()

		doc := getExampleDocument()
		if assert.NoError(t, p.Process(ctx, doc)) {
			pd := doc.Feature.PluginData
			assert.Equal(t, "\"Name\"", pd["GoValue"])
			assert.Equal(t, "Name", pd["GoName"])
		}
	})

	t.Run("Rule", func(t *testing.T) {
		t.Parallel()

		doc := getExampleDocument()
		if assert.NoError(t, p.Process(ctx, doc)) {
			pd := doc.Feature.Children[0].Rule.PluginData
			assert.Equal(t, "\"Name\"", pd["GoValue"])
			assert.Equal(t, "Keyword", pd["GoName"])
		}
	})

	t.Run("Scenario", func(t *testing.T) {
		t.Parallel()

		doc := getExampleDocument()
		if assert.NoError(t, p.Process(ctx, doc)) {
			pd := doc.Feature.Children[0].Scenario.PluginData
			assert.Equal(t, "\"Name\"", pd["GoValue"])
			assert.Equal(t, "Keyword", pd["GoName"])
		}
	})

	t.Run("Step", func(t *testing.T) {
		t.Parallel()

		doc := getExampleDocument()
		if assert.NoError(t, p.Process(ctx, doc)) {
			pd := doc.Feature.Children[0].Scenario.Steps[0].PluginData
			assert.Equal(t, "Keyword", pd["GoName"])
			assert.Equal(t, "\"Text\"", pd["GoValue"])
		}
	})
}

func getExampleDocument() *model.GherkinDocument {
	return &model.GherkinDocument{
		Feature: &model.Feature{
			Keyword: "Keyword",
			Name:    "Name",
			Children: []*model.FeatureChild{{
				Rule: &model.Rule{
					Keyword:    "Keyword",
					Name:       "Name",
					PluginData: map[string]any{},
				},
				Background: &model.Background{
					Keyword:    "Keyword",
					Name:       "Name",
					PluginData: map[string]any{},
				},
				Scenario: &model.Scenario{
					Keyword: "Keyword",
					Name:    "Name",
					Steps: []*model.Step{{
						Keyword:    "Keyword",
						Text:       "Text",
						PluginData: map[string]any{},
					}},
					Examples: []*model.Examples{{
						Keyword: "Keyword",
						Name:    "Name",
						TableHeader: &model.TableRow{
							Cells: []*model.TableCell{{
								Value:      "<Title>",
								PluginData: make(map[string]any),
							}},
							PluginData: make(map[string]any),
						},
						TableBody: []*model.TableRow{{
							Cells: []*model.TableCell{{
								Value:      "5",
								PluginData: make(map[string]any),
							}},
							PluginData: make(map[string]any),
						}},
						PluginData: map[string]any{},
					}, {
						Keyword: "Keyword",
						Name:    "Name",
						TableHeader: &model.TableRow{
							Cells: []*model.TableCell{{
								Value:      "<Message>",
								PluginData: make(map[string]any),
							}},
							PluginData: make(map[string]any),
						},
						TableBody: []*model.TableRow{{
							Cells: []*model.TableCell{{
								Value:      "hello world",
								PluginData: make(map[string]any),
							}},
							PluginData: make(map[string]any),
						}},
						PluginData: map[string]any{},
					}},
					PluginData: map[string]any{},
				},
				PluginData: map[string]any{},
			}},
			PluginData: map[string]any{},
		},
	}
}
