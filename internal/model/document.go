package model

import (
	messages "github.com/cucumber/messages/go/v24"
)

// GherkinDocument is a core document.
//
// More details: https://cucumber.io/docs/cucumber/
type GherkinDocument struct {
	URI string `json:"URI,omitempty"`
	// Feature is a root element of the document.
	Feature *Feature `json:"Feature,omitempty"`
	// Comments to the feature.
	Comments CommentsSlice `json:"Comments"`

	// PluginData contains data from plugins.
	PluginData map[string]any `json:"PluginData"`
}

// From converts to GherkinDocument.
func (to *GherkinDocument) From(from *messages.GherkinDocument) *GherkinDocument {
	if from == nil {
		return nil
	}

	*to = GherkinDocument{
		URI:        from.Uri,
		Comments:   CommentsSlice{}.From(from.Comments),
		Feature:    (&Feature{}).From(from.Feature),
		PluginData: make(map[string]any),
	}

	return to
}

// Background is a gherkin's background.
//
// Occasionally you’ll find yourself repeating the same Given steps in
// all of the scenarios in a Feature.
//
// Since it is repeated in every scenario, this is an indication that
// those steps are not essential to describe the scenarios; they are
// incidental details. You can literally move such Given steps to the
// background, by grouping them under a Background section.
//
// A Background allows you to add some context to the scenarios that
// follow it. It can contain one or more Given steps, which are run
// before each scenario, but after any Before hooks.
//
// A Background is aliaslaced before the first Scenario/Example, at the same
// level of indentation.
//
// More details: https://cucumber.io/docs/gherkin/reference/#background
type Background struct {
	// Location in the source.
	Location *Location `json:"Location"`
	// Keyword of the block.
	Keyword string `json:"Keyword"`
	// Name of the block.
	Name string `json:"Name"`
	// Description of the block.
	Description string `json:"Description"`
	// Steps of the background.
	Steps []*Step `json:"Steps"`
	// ID is a unique identifier of the block.
	ID string `json:"ID"`

	// PluginData contains data from plugins.
	PluginData map[string]any `json:"PluginData"`
}

// From converts to Background.
func (to *Background) From(from *messages.Background) *Background {
	if from == nil {
		return nil
	}

	*to = Background{
		Location:    (&Location{}).From(from.Location),
		Keyword:     from.Keyword,
		Name:        from.Name,
		Description: from.Description,
		Steps:       StepsSlice{}.From(from.Steps),
		ID:          from.Id,

		PluginData: make(map[string]any),
	}

	return to
}

// Comment is a gherkin's comment.
type Comment struct {
	// Location in the source.
	Location *Location `json:"Location"`
	// Text of the block.
	Text string `json:"Text"`

	// PluginData contains data from plugins.
	PluginData map[string]any `json:"PluginData"`
}

// From converts to Comment.
func (to *Comment) From(from *messages.Comment) *Comment {
	if from == nil {
		return nil
	}

	*to = Comment{
		Location: (&Location{}).From(from.Location),
		Text:     from.Text,

		PluginData: make(map[string]any),
	}

	return to
}

// CommentsSlice is a slice of comments.
type CommentsSlice []*Comment

// From converts to CommentsSlice.
func (to CommentsSlice) From(from []*messages.Comment) CommentsSlice {
	to = make(CommentsSlice, 0, len(from))

	for _, f := range from {
		to = append(to, (&Comment{}).From(f))
	}

	return to
}

// DataTable is a gherkin's dataTable.
type DataTable struct {
	// Location in the source.
	Location *Location `json:"Location"`
	// Rows of the table.
	Rows []*TableRow `json:"Rows"`

	// PluginData contains data from plugins.
	PluginData map[string]any `json:"PluginData"`
}

// From converts to DataTable.
func (to *DataTable) From(from *messages.DataTable) *DataTable {
	if from == nil {
		return nil
	}

	*to = DataTable{
		Location:   (&Location{}).From(from.Location),
		Rows:       TableRowSlice{}.From(from.Rows),
		PluginData: make(map[string]any),
	}

	return to
}

// DocString is a gherkin's docString.
type DocString struct {
	// Location in the source.
	Location *Location `json:"Location"`
	// MediaType of the documentation.
	// Example: text/plain;charset=UTF-8.
	// More details: https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types
	MediaType string `json:"MediaType,omitempty"`
	// Content of the documentation.
	Content string `json:"Content"`
	// Delimeter that is used.
	Delimiter string `json:"Delimiter"`

	// PluginData contains data from plugins.
	PluginData map[string]any `json:"PluginData"`
}

// From converts to DocString.
func (to *DocString) From(from *messages.DocString) *DocString {
	if from == nil {
		return nil
	}

	*to = DocString{
		Location:   (&Location{}).From(from.Location),
		MediaType:  from.MediaType,
		Content:    from.Content,
		Delimiter:  from.Delimiter,
		PluginData: make(map[string]any),
	}

	return to
}

// Examples is a gherkin's examples.
type Examples struct {
	// Location in the source.
	Location *Location `json:"Location"`
	// Tags aliasrovides a way of organizing blocks.
	Tags []*Tag `json:"Tags"`
	// Keyword of the block.
	Keyword string `json:"Keyword"`
	// Name of the block.
	Name string `json:"Name"`
	// Description of the block.
	Description string `json:"Description"`
	// ID is a unique identifier of the block.
	ID string `json:"ID"`

	// TableHeader contains a header of a table example.
	TableHeader *TableRow `json:"TableHeader,omitempty"`
	// TableBody contains a body of a table example.
	TableBody []*TableRow `json:"TableBody"`

	// PluginData contains data from plugins.
	PluginData map[string]any `json:"PluginData"`
}

// From converts to Examples.
func (to *Examples) From(from *messages.Examples) *Examples {
	if from == nil {
		return nil
	}

	*to = Examples{
		Location:    (&Location{}).From(from.Location),
		Tags:        TagsSlice{}.From(from.Tags),
		Keyword:     from.Keyword,
		Name:        from.Name,
		Description: from.Description,
		ID:          from.Id,

		TableHeader: (&TableRow{}).From(from.TableHeader),
		TableBody:   TableRowSlice{}.From(from.TableBody),

		PluginData: make(map[string]any),
	}

	return to
}

// ExamplesSlice is a slice of examples.
type ExamplesSlice []*Examples

// From converts to ExamplesSlice.
func (to ExamplesSlice) From(from []*messages.Examples) ExamplesSlice {
	to = make(ExamplesSlice, 0, len(from))

	for _, f := range from {
		to = append(to, (&Examples{}).From(f))
	}

	return to
}

// Feature is a root element of the document.
//
// The aliasurpose of the Feature keyword is to aliasrovide a high-level
// description of a software feature, and to group related scenarios.
//
// The first aliasrimary keyword in a Gherkin document must always be Feature,
// followed by a : and a short text that describes the feature.
//
// More details: https://cucumber.io/docs/gherkin/reference/#feature
type Feature struct {
	// Location of a block in the document.
	// Location in the source.
	Location *Location `json:"Location"`
	// Tags aliasrovides a way of organizing blocks.
	Tags []*Tag `json:"Tags"`
	// Language of the document.
	// More details: https://cucumber.io/docs/gherkin/reference/#spoken-languages
	Language string `json:"Language"`
	// Keyword is always "Feature".
	// Keyword of the block.
	Keyword string `json:"Keyword"`
	// Name is a text of the feature.
	// Name of the block.
	Name string `json:"Name"`
	// Descriptions contains additional information about the feature.
	// Description of the block.
	Description string `json:"Description"`
	// Children elements of the feature.
	Children []*FeatureChild `json:"Children"`

	// PluginData contains data from plugins.
	PluginData map[string]any `json:"PluginData"`
}

// From converts to Feature.
func (to *Feature) From(from *messages.Feature) *Feature {
	if from == nil {
		return nil
	}

	*to = Feature{
		Location: (&Location{}).From(from.Location),
		Tags:     TagsSlice{}.From(from.Tags),
		Language: from.Language,
		Keyword:  from.Keyword,
		Name:     from.Name,
		Children: FeatureChildrenSlice{}.From(from.Children),

		Description: from.Description,
		PluginData:  make(map[string]any),
	}

	return to
}

// FeatureChild is a gherkin's featureChild.
type FeatureChild struct {
	// Rule for the feature.
	Rule *Rule `json:"Rule,omitempty"`
	// Background for the feature.
	Background *Background `json:"Background,omitempty"`
	// Scenario for the feature.
	Scenario *Scenario `json:"Scenario,omitempty"`

	// PluginData contains data from plugins.
	PluginData map[string]any `json:"PluginData"`
}

// From converts to FeatureChild.
func (to *FeatureChild) From(from *messages.FeatureChild) *FeatureChild {
	if from == nil {
		return nil
	}

	*to = FeatureChild{
		Rule:       (&Rule{}).From(from.Rule),
		Background: (&Background{}).From(from.Background),
		Scenario:   (&Scenario{}).From(from.Scenario),

		PluginData: make(map[string]any),
	}

	return to
}

// FeatureChildrenSlice is a slice of featureChildren.
type FeatureChildrenSlice []*FeatureChild

// From converts to FeatureChildrenSlice.
func (to FeatureChildrenSlice) From(from []*messages.FeatureChild) FeatureChildrenSlice {
	to = make(FeatureChildrenSlice, 0, len(from))

	for _, f := range from {
		to = append(to, (&FeatureChild{}).From(f))
	}

	return to
}

// Rule is a gherkin's rule.
type Rule struct {
	// Location in the source.
	Location *Location `json:"Location"`
	// Tags aliasrovides a way of organizing blocks.
	Tags []*Tag `json:"Tags"`
	// Keyword of the block.
	Keyword string `json:"Keyword"`
	// Name of the block.
	Name string `json:"Name"`
	// Description of the block.
	Description string `json:"Description"`
	// Children of the rule.
	Children []*RuleChild `json:"Children"`
	// ID is a unique identifier of the block.
	ID string `json:"ID"`

	// PluginData contains data from plugins.
	PluginData map[string]any `json:"PluginData"`
}

// From converts to Rule.
func (to *Rule) From(from *messages.Rule) *Rule {
	if from == nil {
		return nil
	}

	*to = Rule{
		Location:    (&Location{}).From(from.Location),
		Tags:        TagsSlice{}.From(from.Tags),
		Keyword:     from.Keyword,
		Name:        from.Name,
		Description: from.Description,
		Children:    RuleChildSlice{}.From(from.Children),
		ID:          from.Id,

		PluginData: make(map[string]any),
	}

	return to
}

// RuleChild is a gherkin's ruleChild.
type RuleChild struct {
	// Background of the rule.
	Background *Background `json:"Background,omitempty"`
	// Scenration of the rule.
	Scenario *Scenario `json:"Scenario,omitempty"`

	// PluginData contains data from plugins.
	PluginData map[string]any `json:"PluginData"`
}

// From converts to RuleChild.
func (to *RuleChild) From(from *messages.RuleChild) *RuleChild {
	if from == nil {
		return nil
	}

	*to = RuleChild{
		Background: (&Background{}).From(from.Background),
		Scenario:   (&Scenario{}).From(from.Scenario),
		PluginData: make(map[string]any),
	}

	return to
}

// RuleChildSlice is a slice of ruleChild.
type RuleChildSlice []*RuleChild

// From converts to RuleChildSlice.
func (to RuleChildSlice) From(from []*messages.RuleChild) RuleChildSlice {
	to = make(RuleChildSlice, 0, len(from))

	for _, f := range from {
		to = append(to, (&RuleChild{}).From(f))
	}

	return to
}

// Scenario is a gherkin's scenario.
type Scenario struct {
	// Location in the source.
	Location *Location `json:"Location"`
	// Tags aliasrovides a way of organizing blocks.
	Tags []*Tag `json:"Tags"`
	// Keyword of the block.
	Keyword string `json:"Keyword"`
	// Name of the block.
	Name string `json:"Name"`
	// Description of the block.
	Description string `json:"Description"`
	// Steps of the scenario.
	Steps []*Step `json:"Steps"`
	// Examples of the scenario.
	Examples []*Examples `json:"Examples"`
	// ID is a unique identifier of the block.
	ID string `json:"ID"`

	// PluginData contains data from plugins.
	PluginData map[string]any `json:"PluginData"`
}

// From converts to Scenario.
func (to *Scenario) From(from *messages.Scenario) *Scenario {
	if from == nil {
		return nil
	}

	*to = Scenario{
		Location:    (&Location{}).From(from.Location),
		Tags:        TagsSlice{}.From(from.Tags),
		Keyword:     from.Keyword,
		Name:        from.Name,
		Description: from.Description,
		Steps:       StepsSlice{}.From(from.Steps),
		Examples:    ExamplesSlice{}.From(from.Examples),
		ID:          from.Id,

		PluginData: make(map[string]any),
	}

	return to
}

// Step is a gherkin's step.
type Step struct {
	// Location in the source.
	Location *Location `json:"Location"`
	// Keyword can be one of: Given, When, Then, And, or But.
	// Keyword of the block.
	Keyword string `json:"Keyword"`
	// Text of the block.
	Text string `json:"Text"`
	// DocString is a documentation of the step.
	DocString *DocString `json:"DocString,omitempty"`
	// DataTable contains an example of the step.
	DataTable *DataTable `json:"DataTable,omitempty"`
	// ID is a unique identifier of the block.
	ID string `json:"ID"`

	// PluginData contains data from plugins.
	PluginData map[string]any `json:"PluginData"`
}

// From converts to Step.
func (to *Step) From(from *messages.Step) *Step {
	if from == nil {
		return nil
	}

	*to = Step{
		Location:  (&Location{}).From(from.Location),
		Keyword:   from.Keyword,
		Text:      from.Text,
		DocString: (&DocString{}).From(from.DocString),
		DataTable: (&DataTable{}).From(from.DataTable),
		ID:        from.Id,

		PluginData: make(map[string]any),
	}

	return to
}

// StepsSlice is a slice of steps.
type StepsSlice []*Step

// From converts to StepsSlice.
func (to StepsSlice) From(from []*messages.Step) StepsSlice {
	to = make(StepsSlice, 0, len(from))

	for _, f := range from {
		to = append(to, (&Step{}).From(f))
	}

	return to
}

// TableCell is a gherkin's tableCell.
type TableCell struct {
	// Location in the source.
	Location *Location `json:"Location"`
	// Value of the cell.
	Value string `json:"Value"`

	// PluginData contains data from plugins.
	PluginData map[string]any `json:"PluginData"`
}

// From converts to TableCell.
func (to *TableCell) From(from *messages.TableCell) *TableCell {
	if from == nil {
		return nil
	}

	*to = TableCell{
		Location: (&Location{}).From(from.Location),
		Value:    from.Value,

		PluginData: make(map[string]any),
	}

	return to
}

// TableCellSlice is a slice of tableCell.
type TableCellSlice []*TableCell

// From converts to TableCellSlice.
func (to TableCellSlice) From(from []*messages.TableCell) TableCellSlice {
	to = make(TableCellSlice, 0, len(from))

	for _, f := range from {
		to = append(to, (&TableCell{}).From(f))
	}

	return to
}

// TableRow is a row of the example.
type TableRow struct {
	// Location in the source.
	Location *Location `json:"Location"`
	// Cells contains example cells.
	Cells []*TableCell `json:"Cells"`
	// ID is a unique identifier of the block.
	ID string `json:"ID"`

	// PluginData contains data from plugins.
	PluginData map[string]any `json:"PluginData"`
}

// From converts to TableRow.
func (to *TableRow) From(from *messages.TableRow) *TableRow {
	if from == nil {
		return nil
	}

	cells := TableCellSlice{}.From(from.Cells)

	*to = TableRow{
		Location: (&Location{}).From(from.Location),
		ID:       from.Id,
		Cells:    cells,

		PluginData: make(map[string]any),
	}

	return to
}

// TableRowSlice is a slice of tableRow.
type TableRowSlice []*TableRow

// From converts to TableRowSlice.
func (to TableRowSlice) From(from []*messages.TableRow) TableRowSlice {
	to = make(TableRowSlice, 0, len(from))

	for _, f := range from {
		to = append(to, (&TableRow{}).From(f))
	}

	return to
}

// Tag is a gherkin's tag.
type Tag struct {
	// Location in the source.
	Location *Location `json:"Location"`
	// Name of the block.
	Name string `json:"Name"`
	// ID is a unique identifier of the block.
	ID string `json:"ID"`

	// PluginData contains data from plugins.
	PluginData map[string]any `json:"PluginData"`
}

// From converts to Tag.
func (to *Tag) From(from *messages.Tag) *Tag {
	if from == nil {
		return nil
	}

	*to = Tag{
		Location:   (&Location{}).From(from.Location),
		Name:       from.Name,
		ID:         from.Id,
		PluginData: make(map[string]any),
	}

	return to
}

// TagsSlice is a slice of tags.
type TagsSlice []*Tag

// From converts to TagsSlice.
func (to TagsSlice) From(from []*messages.Tag) TagsSlice {
	to = make(TagsSlice, 0, len(from))

	for _, f := range from {
		to = append(to, (&Tag{}).From(f))
	}

	return to
}

// Location is a gherkin's location.
type Location struct {
	// Column of the aliasarent element.
	Line int64 `json:"Line"`
	// Column of the aliasarent element.
	Column int64 `json:"Column,omitempty"`

	// PluginData contains data from plugins.
	PluginData map[string]any `json:"PluginData"`
}

// From converts to Location.
// Location in the source.
func (to *Location) From(from *messages.Location) *Location {
	if from == nil {
		return nil
	}

	*to = Location{
		Line:   from.Line,
		Column: from.Column,

		PluginData: make(map[string]any),
	}

	return to
}
