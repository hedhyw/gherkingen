package model_test

import (
	"strconv"
	"testing"

	messages "github.com/cucumber/common/messages/go/v19"
	"github.com/hedhyw/gherkingen/v2/internal/model"

	"github.com/stretchr/testify/assert"
)

// nolint: gocognit,cyclop,maintidx // Unit test.
func TestGherkinDocumentFrom(t *testing.T) {
	t.Parallel()

	t.Run("GherkinDocument", func(t *testing.T) {
		t.Parallel()

		obj := (&model.GherkinDocument{}).From(&messages.GherkinDocument{
			Uri:      "uri",
			Feature:  &messages.Feature{},
			Comments: []*messages.Comment{{}},
		})
		if assert.NotNil(t, obj) {
			assert.NotNil(t, obj.Feature)
			assert.NotEmpty(t, obj.Feature)
			assert.Equal(t, "uri", obj.URI)
			assert.NotNil(t, obj.PluginData)
		}
	})

	t.Run("Background", func(t *testing.T) {
		t.Parallel()

		obj := (&model.Background{}).From(&messages.Background{
			Location:    &messages.Location{},
			Keyword:     "Keyword",
			Name:        "Name",
			Description: "Description",
			Steps:       []*messages.Step{{}},
			Id:          "Id",
		})
		if assert.NotNil(t, obj) {
			assert.NotNil(t, obj.Location)
			assert.Equal(t, "Keyword", obj.Keyword)
			assert.Equal(t, "Name", obj.Name)
			assert.Equal(t, "Description", obj.Description)
			assert.NotEmpty(t, obj.Steps)
			assert.Equal(t, "Id", obj.ID)
			assert.NotNil(t, obj.PluginData)
		}
	})

	t.Run("Comment", func(t *testing.T) {
		t.Parallel()

		obj := (&model.Comment{}).From(&messages.Comment{
			Location: &messages.Location{},
			Text:     "text",
		})
		if assert.NotNil(t, obj) {
			assert.NotNil(t, obj.Location)
			assert.Equal(t, "text", obj.Text)
			assert.NotNil(t, obj.PluginData)
		}
	})

	t.Run("CommentsSlice", func(t *testing.T) {
		t.Parallel()

		obj := (&model.CommentsSlice{}).From([]*messages.Comment{{}})
		if assert.NotNil(t, obj) {
			assert.Len(t, obj, 1)
		}
	})

	t.Run("DataTable", func(t *testing.T) {
		t.Parallel()

		obj := (&model.DataTable{}).From(&messages.DataTable{
			Location: &messages.Location{},
			Rows:     []*messages.TableRow{{}},
		})
		if assert.NotNil(t, obj) {
			assert.NotNil(t, obj.Location)
			assert.NotEmpty(t, obj.Rows)
			assert.NotNil(t, obj.PluginData)
		}
	})

	t.Run("DocString", func(t *testing.T) {
		t.Parallel()

		obj := (&model.DocString{}).From(&messages.DocString{
			Location:  &messages.Location{},
			MediaType: "plain/text",
			Content:   "content",
			Delimiter: "---",
		})
		if assert.NotNil(t, obj) {
			assert.NotNil(t, obj.Location)
			assert.Equal(t, "plain/text", obj.MediaType)
			assert.Equal(t, "content", obj.Content)
			assert.Equal(t, "---", obj.Delimiter)
			assert.NotNil(t, obj.PluginData)
		}
	})

	t.Run("Examples", func(t *testing.T) {
		t.Parallel()

		obj := (&model.Examples{}).From(&messages.Examples{
			Location:    &messages.Location{},
			Tags:        []*messages.Tag{{}},
			Keyword:     "Keyword",
			Name:        "Name",
			Description: "Description",
			TableHeader: &messages.TableRow{},
			TableBody:   []*messages.TableRow{{}},
			Id:          "Id",
		})
		if assert.NotNil(t, obj) {
			assert.NotNil(t, obj.Location)
			assert.NotEmpty(t, obj.Tags)
			assert.Equal(t, "Keyword", obj.Keyword)
			assert.Equal(t, "Name", obj.Name)
			assert.Equal(t, "Description", obj.Description)
			assert.NotNil(t, obj.TableHeader)
			assert.NotEmpty(t, obj.TableBody)
			assert.Equal(t, "Id", obj.ID)
			assert.NotNil(t, obj.PluginData)
		}
	})

	t.Run("ExamplesSlice", func(t *testing.T) {
		t.Parallel()

		obj := (&model.ExamplesSlice{}).From([]*messages.Examples{{}})
		if assert.NotNil(t, obj) {
			assert.Len(t, obj, 1)
		}
	})

	t.Run("Feature", func(t *testing.T) {
		t.Parallel()

		obj := (&model.Feature{}).From(&messages.Feature{
			Location:    &messages.Location{},
			Tags:        []*messages.Tag{{}},
			Keyword:     "Keyword",
			Name:        "Name",
			Description: "Description",
			Language:    "en_US",
			Children:    []*messages.FeatureChild{{}},
		})
		if assert.NotNil(t, obj) {
			assert.NotNil(t, obj.Location)
			assert.NotEmpty(t, obj.Tags)
			assert.Equal(t, "Keyword", obj.Keyword)
			assert.Equal(t, "Name", obj.Name)
			assert.Equal(t, "en_US", obj.Language)
			assert.NotEmpty(t, obj.Children)
			assert.NotNil(t, obj.PluginData)
		}
	})

	t.Run("FeatureChild", func(t *testing.T) {
		t.Parallel()

		obj := (&model.FeatureChild{}).From(&messages.FeatureChild{
			Rule:       &messages.Rule{},
			Background: &messages.Background{},
			Scenario:   &messages.Scenario{},
		})
		if assert.NotNil(t, obj) {
			assert.NotNil(t, obj.Rule)
			assert.NotNil(t, obj.Background)
			assert.NotNil(t, obj.Scenario)
			assert.NotNil(t, obj.PluginData)
		}
	})

	t.Run("FeatureChildrenSlice", func(t *testing.T) {
		t.Parallel()

		obj := (&model.FeatureChildrenSlice{}).From([]*messages.FeatureChild{{}})
		if assert.NotNil(t, obj) {
			assert.Len(t, obj, 1)
		}
	})

	t.Run("Rule", func(t *testing.T) {
		t.Parallel()

		obj := (&model.Rule{}).From(&messages.Rule{
			Location:    &messages.Location{},
			Tags:        []*messages.Tag{{}},
			Keyword:     "Keyword",
			Name:        "Name",
			Description: "Description",
			Children:    []*messages.RuleChild{{}},
			Id:          "Id",
		})
		if assert.NotNil(t, obj) {
			assert.NotNil(t, obj.Location)
			assert.NotEmpty(t, obj.Tags)
			assert.Equal(t, "Keyword", obj.Keyword)
			assert.Equal(t, "Name", obj.Name)
			assert.Equal(t, "Description", obj.Description)
			assert.NotEmpty(t, obj.Children)
			assert.Equal(t, "Id", obj.ID)
			assert.NotNil(t, obj.PluginData)
		}
	})

	t.Run("RuleChild", func(t *testing.T) {
		t.Parallel()

		obj := (&model.RuleChild{}).From(&messages.RuleChild{
			Background: &messages.Background{},
			Scenario:   &messages.Scenario{},
		})
		if assert.NotNil(t, obj) {
			assert.NotNil(t, obj.Background)
			assert.NotNil(t, obj.Scenario)
			assert.NotNil(t, obj.PluginData)
		}
	})

	t.Run("RuleChildSlice", func(t *testing.T) {
		t.Parallel()

		obj := (&model.RuleChildSlice{}).From([]*messages.RuleChild{{}})
		if assert.NotNil(t, obj) {
			assert.Len(t, obj, 1)
		}
	})

	t.Run("Scenario", func(t *testing.T) {
		t.Parallel()

		obj := (&model.Scenario{}).From(&messages.Scenario{
			Location:    &messages.Location{},
			Tags:        []*messages.Tag{{}},
			Keyword:     "Keyword",
			Name:        "Name",
			Description: "Description",
			Id:          "Id",
			Steps:       []*messages.Step{{}},
			Examples:    []*messages.Examples{{}},
		})
		if assert.NotNil(t, obj) {
			assert.NotNil(t, obj.Location)
			assert.NotEmpty(t, obj.Tags)
			assert.Equal(t, "Keyword", obj.Keyword)
			assert.Equal(t, "Name", obj.Name)
			assert.Equal(t, "Description", obj.Description)
			assert.NotEmpty(t, obj.Steps)
			assert.NotEmpty(t, obj.Examples)
			assert.Equal(t, "Id", obj.ID)
			assert.NotNil(t, obj.PluginData)
		}
	})

	t.Run("Step", func(t *testing.T) {
		t.Parallel()

		obj := (&model.Step{}).From(&messages.Step{
			Location:  &messages.Location{},
			Keyword:   "Keyword",
			Id:        "Id",
			Text:      "Text",
			DocString: &messages.DocString{},
			DataTable: &messages.DataTable{},
		})
		if assert.NotNil(t, obj) {
			assert.NotNil(t, obj.Location)
			assert.Equal(t, "Keyword", obj.Keyword)
			assert.Equal(t, "Id", obj.ID)
			assert.Equal(t, "Text", obj.Text)
			assert.NotNil(t, obj.DocString)
			assert.NotNil(t, obj.DataTable)
			assert.NotNil(t, obj.PluginData)
		}
	})

	t.Run("StepsSlice", func(t *testing.T) {
		t.Parallel()

		obj := (&model.StepsSlice{}).From([]*messages.Step{{}})
		if assert.NotNil(t, obj) {
			assert.Len(t, obj, 1)
		}
	})

	t.Run("TableCell", func(t *testing.T) {
		t.Parallel()

		obj := (&model.TableCell{}).From(&messages.TableCell{
			Location: &messages.Location{},
			Value:    "Value",
		})
		if assert.NotNil(t, obj) {
			assert.NotNil(t, obj.Location)
			assert.Equal(t, "Value", obj.Value)
			assert.NotNil(t, obj.PluginData)
		}
	})

	t.Run("TableCellSlice", func(t *testing.T) {
		t.Parallel()

		obj := (&model.TableCellSlice{}).From([]*messages.TableCell{{}})
		if assert.NotNil(t, obj) {
			assert.Len(t, obj, 1)
		}
	})

	t.Run("TableRow", func(t *testing.T) {
		t.Parallel()

		obj := (&model.TableRow{}).From(&messages.TableRow{
			Location: &messages.Location{},
			Cells:    []*messages.TableCell{{}},
			Id:       "Id",
		})
		if assert.NotNil(t, obj) {
			assert.NotNil(t, obj.Location)
			assert.Equal(t, "Id", obj.ID)
			assert.NotNil(t, obj.PluginData)
		}
	})

	t.Run("TableRowSlice", func(t *testing.T) {
		t.Parallel()

		obj := (&model.TableRowSlice{}).From([]*messages.TableRow{{}})
		if assert.NotNil(t, obj) {
			assert.Len(t, obj, 1)
		}
	})

	t.Run("Tag", func(t *testing.T) {
		t.Parallel()

		obj := (&model.Tag{}).From(&messages.Tag{
			Location: &messages.Location{},
			Id:       "Id",
			Name:     "Name",
		})
		if assert.NotNil(t, obj) {
			assert.NotNil(t, obj.Location)
			assert.Equal(t, "Id", obj.ID)
			assert.Equal(t, "Name", obj.Name)
			assert.NotNil(t, obj.PluginData)
		}
	})

	t.Run("TagsSlice", func(t *testing.T) {
		t.Parallel()

		obj := (&model.TagsSlice{}).From([]*messages.Tag{{}})
		if assert.NotNil(t, obj) {
			assert.Len(t, obj, 1)
		}
	})

	t.Run("Location", func(t *testing.T) {
		t.Parallel()

		obj := (&model.Location{}).From(&messages.Location{
			Line:   1,
			Column: 2,
		})
		if assert.NotNil(t, obj) {
			assert.Equal(t, 1, int(obj.Line))
			assert.Equal(t, 2, int(obj.Column))
			assert.NotNil(t, obj.PluginData)
		}
	})
}

func TestGherkinDocumentFromNil(t *testing.T) {
	t.Parallel()

	testCases := [...]func() any{
		func() any { return (&model.Location{}).From(nil) },
		func() any { return (&model.GherkinDocument{}).From(nil) },
		func() any { return (&model.Feature{}).From(nil) },
		func() any { return (&model.Background{}).From(nil) },
		func() any { return (&model.Comment{}).From(nil) },
		func() any { return (&model.DataTable{}).From(nil) },
		func() any { return (&model.DocString{}).From(nil) },
		func() any { return (&model.Examples{}).From(nil) },
		func() any { return (&model.Feature{}).From(nil) },
		func() any { return (&model.FeatureChild{}).From(nil) },
		func() any { return (&model.Rule{}).From(nil) },
		func() any { return (&model.RuleChild{}).From(nil) },
		func() any { return (&model.Scenario{}).From(nil) },
		func() any { return (&model.Step{}).From(nil) },
		func() any { return (&model.TableCell{}).From(nil) },
		func() any { return (&model.TableRow{}).From(nil) },
		func() any { return (&model.Tag{}).From(nil) },
		func() any { return (&model.Location{}).From(nil) },
	}

	for i, tc := range testCases {
		tc := tc

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			assert.Nil(t, tc())
		})
	}
}
