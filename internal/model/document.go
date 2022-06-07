package model

import (
	"strings"

	"github.com/cucumber/messages-go/v16"
)

type goData struct {
	// GoType represents a name of the type in Golang.
	// Example: int
	GoType goType `json:"GoType"`
	// GoType represents a valid identifier in Golang.
	// Example: SomeValue
	GoName string `json:"GoName"`
	// GoType is a valid quoted string that can be used as a values in Golang.
	// Example: "Some\"Value"
	GoValue string `json:"GoValue"`
}

// GherkinDocument is a core document.
//
// More details: https://cucumber.io/docs/cucumber/
type GherkinDocument struct {
	URI string `json:"URI,omitempty"`
	// Feature is a root element of the document.
	Feature *Feature `json:"Feature,omitempty"`
	// Comments to the feature.
	Comments []*Comment `json:"Comments"`
}

// From converts to GherkinDocument.
func (to *GherkinDocument) From(
	from *messages.GherkinDocument,
	p *Processor,
) *GherkinDocument {
	if from == nil {
		return nil
	}

	*to = GherkinDocument{
		URI:      from.Uri,
		Feature:  (&Feature{}).From(from.Feature, p),
		Comments: CommentsSlice{}.From(from.Comments),
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
// A Background is placed before the first Scenario/Example, at the same
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

	goData
}

// From converts to Background.
func (to *Background) From(from *messages.Background, p *Processor) *Background {
	if from == nil {
		return nil
	}

	*to = Background{
		Location:    (&Location{}).From(from.Location),
		Keyword:     from.Keyword,
		Name:        from.Name,
		Description: from.Description,
		Steps:       StepsSlice{}.From(from.Steps, p),
		ID:          from.Id,

		goData: goData{
			GoType:  GOTypeString,
			GoName:  goName(from.Keyword, p),
			GoValue: goString(from.Name),
		},
	}

	return to
}

// Comment is a gherkin's comment.
type Comment struct {
	// Location in the source.
	Location *Location `json:"Location"`
	// Text of the block.
	Text string `json:"Text"`
}

// From converts to Comment.
func (to *Comment) From(from *messages.Comment) *Comment {
	if from == nil {
		return nil
	}

	*to = Comment{
		Location: (&Location{}).From(from.Location),
		Text:     from.Text,
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
}

// From converts to DataTable.
func (to *DataTable) From(from *messages.DataTable, p *Processor) *DataTable {
	if from == nil {
		return nil
	}

	*to = DataTable{
		Location: (&Location{}).From(from.Location),
		Rows:     TableRowSlice{}.From(from.Rows, p),
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
}

// From converts to DocString.
func (to *DocString) From(from *messages.DocString) *DocString {
	if from == nil {
		return nil
	}

	*to = DocString{
		Location:  (&Location{}).From(from.Location),
		MediaType: from.MediaType,
		Content:   from.Content,
		Delimiter: from.Delimiter,
	}

	return to
}

// Examples is a gherkin's examples.
type Examples struct {
	// Location in the source.
	Location *Location `json:"Location"`
	// Tags provides a way of organizing blocks.
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
}

// From converts to Examples.
func (to *Examples) From(from *messages.Examples, p *Processor) *Examples {
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

		TableHeader: (&TableRow{}).From(
			from.TableHeader,
			determinateGoTypes(from.TableBody),
			true, // Ignore go type, because header cells are always strings.
			p,
		),
		TableBody: TableRowSlice{}.From(from.TableBody, p),
	}

	return to
}

// ExamplesSlice is a slice of examples.
type ExamplesSlice []*Examples

// From converts to ExamplesSlice.
func (to ExamplesSlice) From(from []*messages.Examples, p *Processor) ExamplesSlice {
	to = make(ExamplesSlice, 0, len(from))

	for _, f := range from {
		to = append(to, (&Examples{}).From(f, p))
	}

	return to
}

// Feature is a root element of the document.
//
// The purpose of the Feature keyword is to provide a high-level
// description of a software feature, and to group related scenarios.
//
// The first primary keyword in a Gherkin document must always be Feature,
// followed by a : and a short text that describes the feature.
//
// More details: https://cucumber.io/docs/gherkin/reference/#feature
type Feature struct {
	// Location of a block in the document.
	// Location in the source.
	Location *Location `json:"Location"`
	// Tags provides a way of organizing blocks.
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

	goData
}

// From converts to Feature.
func (to *Feature) From(from *messages.Feature, p *Processor) *Feature {
	if from == nil {
		return nil
	}

	*to = Feature{
		Location:    (&Location{}).From(from.Location),
		Tags:        TagsSlice{}.From(from.Tags),
		Language:    from.Language,
		Keyword:     from.Keyword,
		Name:        from.Name,
		Description: from.Description,
		Children:    FeatureChildrenSlice{}.From(from.Children, p),

		goData: goData{
			GoType:  GOTypeString,
			GoName:  goName(from.Name, p),
			GoValue: goString(from.Name),
		},
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
}

// From converts to FeatureChild.
func (to *FeatureChild) From(from *messages.FeatureChild, p *Processor) *FeatureChild {
	if from == nil {
		return nil
	}

	*to = FeatureChild{
		Rule:       (&Rule{}).From(from.Rule, p),
		Background: (&Background{}).From(from.Background, p),
		Scenario:   (&Scenario{}).From(from.Scenario, p),
	}

	return to
}

// FeatureChildrenSlice is a slice of featureChildren.
type FeatureChildrenSlice []*FeatureChild

// From converts to FeatureChildrenSlice.
func (to FeatureChildrenSlice) From(from []*messages.FeatureChild, p *Processor) FeatureChildrenSlice {
	to = make(FeatureChildrenSlice, 0, len(from))

	for _, f := range from {
		to = append(to, (&FeatureChild{}).From(f, p))
	}

	return to
}

// Rule is a gherkin's rule.
type Rule struct {
	// Location in the source.
	Location *Location `json:"Location"`
	// Tags provides a way of organizing blocks.
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

	goData
}

// From converts to Rule.
func (to *Rule) From(from *messages.Rule, p *Processor) *Rule {
	if from == nil {
		return nil
	}

	*to = Rule{
		Location:    (&Location{}).From(from.Location),
		Tags:        TagsSlice{}.From(from.Tags),
		Keyword:     from.Keyword,
		Name:        from.Name,
		Description: from.Description,
		Children:    RuleChildSlice{}.From(from.Children, p),
		ID:          from.Id,

		goData: goData{
			GoType:  GOTypeString,
			GoName:  goName(from.Keyword, p),
			GoValue: goString(from.Name),
		},
	}

	return to
}

// RuleChild is a gherkin's ruleChild.
type RuleChild struct {
	// Background of the rule.
	Background *Background `json:"Background,omitempty"`
	// Scenration of the rule.
	Scenario *Scenario `json:"Scenario,omitempty"`
}

// From converts to RuleChild.
func (to *RuleChild) From(from *messages.RuleChild, p *Processor) *RuleChild {
	if from == nil {
		return nil
	}

	*to = RuleChild{
		Background: (&Background{}).From(from.Background, p),
		Scenario:   (&Scenario{}).From(from.Scenario, p),
	}

	return to
}

// RuleChildSlice is a slice of ruleChild.
type RuleChildSlice []*RuleChild

// From converts to RuleChildSlice.
func (to RuleChildSlice) From(from []*messages.RuleChild, p *Processor) RuleChildSlice {
	to = make(RuleChildSlice, 0, len(from))

	for _, f := range from {
		to = append(to, (&RuleChild{}).From(f, p))
	}

	return to
}

// Scenario is a gherkin's scenario.
type Scenario struct {
	// Location in the source.
	Location *Location `json:"Location"`
	// Tags provides a way of organizing blocks.
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

	goData
}

// From converts to Scenario.
func (to *Scenario) From(from *messages.Scenario, p *Processor) *Scenario {
	if from == nil {
		return nil
	}

	*to = Scenario{
		Location:    (&Location{}).From(from.Location),
		Tags:        TagsSlice{}.From(from.Tags),
		Keyword:     from.Keyword,
		Name:        from.Name,
		Description: from.Description,
		Steps:       StepsSlice{}.From(from.Steps, p),
		Examples:    ExamplesSlice{}.From(from.Examples, p),
		ID:          from.Id,

		goData: goData{
			GoType:  GOTypeString,
			GoName:  goName(from.Keyword, p),
			GoValue: goString(from.Name),
		},
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

	goData
}

// From converts to Step.
func (to *Step) From(from *messages.Step, p *Processor) *Step {
	if from == nil {
		return nil
	}

	*to = Step{
		Location:  (&Location{}).From(from.Location),
		Keyword:   from.Keyword,
		Text:      from.Text,
		DocString: (&DocString{}).From(from.DocString),
		DataTable: (&DataTable{}).From(from.DataTable, p),
		ID:        from.Id,

		goData: goData{
			GoType:  GOTypeString,
			GoName:  goName(from.Keyword, p),
			GoValue: goString(from.Text),
		},
	}

	return to
}

// StepsSlice is a slice of steps.
type StepsSlice []*Step

// From converts to StepsSlice.
func (to StepsSlice) From(from []*messages.Step, p *Processor) StepsSlice {
	to = make(StepsSlice, 0, len(from))

	for _, f := range from {
		to = append(to, (&Step{}).From(f, p))
	}

	return to
}

// TableCell is a gherkin's tableCell.
type TableCell struct {
	// Location in the source.
	Location *Location `json:"Location"`
	// Value of the cell.
	Value string `json:"Value"`

	goData
}

// From converts to TableCell.
func (to *TableCell) From(
	from *messages.TableCell,
	gt goType,
	ignoreGoTypes bool,
	p *Processor,
) *TableCell {
	if from == nil {
		return nil
	}

	var gv string
	if ignoreGoTypes {
		gv = goString(from.Value)
	} else {
		gv, gt = goValue(from.Value, gt)
	}

	*to = TableCell{
		Location: (&Location{}).From(from.Location),
		Value:    from.Value,

		goData: goData{
			GoType:  gt,
			GoName:  goName(from.Value, p),
			GoValue: gv,
		},
	}

	return to
}

// TableCellSlice is a slice of tableCell.
type TableCellSlice []*TableCell

// From converts to TableCellSlice.
func (to TableCellSlice) From(
	from []*messages.TableCell,
	goTypes []goType,
	ignoreGoTypes bool,
	p *Processor,
) TableCellSlice {
	to = make(TableCellSlice, 0, len(from))

	for i, f := range from {
		gt := GOTypeString
		if i < len(goTypes) {
			gt = goTypes[i]
		}

		to = append(to, (&TableCell{}).From(f, gt, ignoreGoTypes, p))
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

	goData
}

// From converts to TableRow.
func (to *TableRow) From(
	from *messages.TableRow,
	goTypes []goType,
	ignoreGoTypes bool,
	p *Processor,
) *TableRow {
	if from == nil {
		return nil
	}

	cells := TableCellSlice{}.From(from.Cells, goTypes, ignoreGoTypes, p)

	values := make([]string, 0, len(cells))
	for _, c := range cells {
		values = append(values, c.Value)
	}
	value := strings.Join(values, "_")

	*to = TableRow{
		Location: (&Location{}).From(from.Location),
		ID:       from.Id,
		Cells:    cells,

		goData: goData{
			GoType:  GOTypeString,
			GoName:  goName(value, p),
			GoValue: goString(value),
		},
	}

	return to
}

// TableRowSlice is a slice of tableRow.
type TableRowSlice []*TableRow

func determinateGoTypes(from []*messages.TableRow) (goTypes []goType) {
	if len(from) == 0 {
		return nil
	}

	columns := len(from[0].Cells)
	goTypes = make([]goType, 0, columns)

	for i := 0; i < columns; i++ {
		values := make([]string, 0, columns)
		for _, row := range from {
			if i >= len(row.Cells) {
				continue
			}

			values = append(values, row.Cells[i].Value)
		}

		goTypes = append(goTypes, determinateGoType(values))
	}

	return goTypes
}

// From converts to TableRowSlice.
func (to TableRowSlice) From(from []*messages.TableRow, p *Processor) TableRowSlice {
	to = make(TableRowSlice, 0, len(from))

	for _, f := range from {
		to = append(to, (&TableRow{}).From(f, determinateGoTypes(from), false, p))
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
}

// From converts to Tag.
func (to *Tag) From(from *messages.Tag) *Tag {
	if from == nil {
		return nil
	}

	*to = Tag{
		Location: (&Location{}).From(from.Location),
		Name:     from.Name,
		ID:       from.Id,
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
	// Column of the parent element.
	Line int64 `json:"Line"`
	// Column of the parent element.
	Column int64 `json:"Column,omitempty"`
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
	}

	return to
}
