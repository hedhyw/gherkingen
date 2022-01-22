package model

import (
	"strings"

	"github.com/cucumber/messages-go/v16"
)

type goData struct {
	GoType  goType `json:"GoType"`
	GoName  string `json:"GoName"`
	GoValue string `json:"GoValue"`
}

// GherkinDocument is a core document.
type GherkinDocument struct {
	URI      string     `json:"URI,omitempty"`
	Feature  *Feature   `json:"Feature,omitempty"`
	Comments []*Comment `json:"Comments"`
}

// From converts to GherkinDocument.
func (to *GherkinDocument) From(from *messages.GherkinDocument) *GherkinDocument {
	if from == nil {
		return nil
	}

	*to = GherkinDocument{
		URI:      from.Uri,
		Feature:  (&Feature{}).From(from.Feature),
		Comments: CommentsSlice{}.From(from.Comments),
	}

	return to
}

// Background is a gherkin's background.
type Background struct {
	Location    *Location `json:"Location"`
	Keyword     string    `json:"Keyword"`
	Name        string    `json:"Name"`
	Description string    `json:"Description"`
	Steps       []*Step   `json:"Steps"`
	ID          string    `json:"ID"`

	goData
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

		goData: goData{
			GoType:  goTypeString,
			GoName:  goName(from.Keyword),
			GoValue: goString(from.Name),
		},
	}

	return to
}

// Comment is a gherkin's comment.
type Comment struct {
	Location *Location `json:"Location"`
	Text     string    `json:"Text"`
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
	Location *Location   `json:"Location"`
	Rows     []*TableRow `json:"Rows"`
}

// From converts to DataTable.
func (to *DataTable) From(from *messages.DataTable) *DataTable {
	if from == nil {
		return nil
	}

	*to = DataTable{
		Location: (&Location{}).From(from.Location),
		Rows:     TableRowSlice{}.From(from.Rows),
	}

	return to
}

// DocString is a gherkin's docString.
type DocString struct {
	Location  *Location `json:"Location"`
	MediaType string    `json:"MediaType,omitempty"`
	Content   string    `json:"Content"`
	Delimiter string    `json:"Delimiter"`
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
	Location    *Location `json:"Location"`
	Tags        []*Tag    `json:"Tags"`
	Keyword     string    `json:"Keyword"`
	Name        string    `json:"Name"`
	Description string    `json:"Description"`
	ID          string    `json:"ID"`

	TableHeader *TableRow   `json:"TableHeader,omitempty"`
	TableBody   []*TableRow `json:"TableBody"`
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

		TableHeader: (&TableRow{}).From(
			from.TableHeader,
			determinateGoTypes(from.TableBody),
			true, // Ignore go type, because header cells are always strings.
		),
		TableBody: TableRowSlice{}.From(from.TableBody),
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

// Feature is a gherkin's feature.
type Feature struct {
	Location    *Location       `json:"Location"`
	Tags        []*Tag          `json:"Tags"`
	Language    string          `json:"Language"`
	Keyword     string          `json:"Keyword"`
	Name        string          `json:"Name"`
	Description string          `json:"Description"`
	Children    []*FeatureChild `json:"Children"`

	goData
}

// From converts to Feature.
func (to *Feature) From(from *messages.Feature) *Feature {
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
		Children:    FeatureChildrenSlice{}.From(from.Children),

		goData: goData{
			GoType:  goTypeString,
			GoName:  goName(from.Name),
			GoValue: goString(from.Name),
		},
	}

	return to
}

// FeatureChild is a gherkin's featureChild.
type FeatureChild struct {
	Rule       *Rule       `json:"Rule,omitempty"`
	Background *Background `json:"Background,omitempty"`
	Scenario   *Scenario   `json:"Scenario,omitempty"`
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
	Location    *Location    `json:"Location"`
	Tags        []*Tag       `json:"Tags"`
	Keyword     string       `json:"Keyword"`
	Name        string       `json:"Name"`
	Description string       `json:"Description"`
	Children    []*RuleChild `json:"Children"`
	ID          string       `json:"ID"`

	goData
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

		goData: goData{
			GoType:  goTypeString,
			GoName:  goName(from.Keyword),
			GoValue: goString(from.Name),
		},
	}

	return to
}

// RuleChild is a gherkin's ruleChild.
type RuleChild struct {
	Background *Background `json:"Background,omitempty"`
	Scenario   *Scenario   `json:"Scenario,omitempty"`
}

// From converts to RuleChild.
func (to *RuleChild) From(from *messages.RuleChild) *RuleChild {
	if from == nil {
		return nil
	}

	*to = RuleChild{
		Background: (&Background{}).From(from.Background),
		Scenario:   (&Scenario{}).From(from.Scenario),
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
	Location    *Location   `json:"Location"`
	Tags        []*Tag      `json:"Tags"`
	Keyword     string      `json:"Keyword"`
	Name        string      `json:"Name"`
	Description string      `json:"Description"`
	Steps       []*Step     `json:"Steps"`
	Examples    []*Examples `json:"Examples"`
	ID          string      `json:"ID"`

	goData
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

		goData: goData{
			GoType:  goTypeString,
			GoName:  goName(from.Keyword),
			GoValue: goString(from.Name),
		},
	}

	return to
}

// Step is a gherkin's step.
type Step struct {
	Location  *Location  `json:"Location"`
	Keyword   string     `json:"Keyword"`
	Text      string     `json:"Text"`
	DocString *DocString `json:"DocString,omitempty"`
	DataTable *DataTable `json:"DataTable,omitempty"`
	ID        string     `json:"ID"`

	goData
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

		goData: goData{
			GoType:  goTypeString,
			GoName:  goName(from.Keyword),
			GoValue: goString(from.Text),
		},
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
	Location *Location `json:"Location"`
	Value    string    `json:"Value"`

	goData
}

// From converts to TableCell.
func (to *TableCell) From(from *messages.TableCell, gt goType, ignoreGoTypes bool) *TableCell {
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
			GoName:  goName(from.Value),
			GoValue: gv,
		},
	}

	return to
}

// TableCellSlice is a slice of tableCell.
type TableCellSlice []*TableCell

// From converts to TableCellSlice.
func (to TableCellSlice) From(from []*messages.TableCell, goTypes []goType, ignoreGoTypes bool) TableCellSlice {
	to = make(TableCellSlice, 0, len(from))

	for i, f := range from {
		gt := goTypeString
		if i < len(goTypes) {
			gt = goTypes[i]
		}

		to = append(to, (&TableCell{}).From(f, gt, ignoreGoTypes))
	}

	return to
}

// TableRow is a gherkin's tableRow.
type TableRow struct {
	Location *Location    `json:"Location"`
	Cells    []*TableCell `json:"Cells"`
	ID       string       `json:"ID"`

	goData
}

// From converts to TableRow.
func (to *TableRow) From(from *messages.TableRow, goTypes []goType, ignoreGoTypes bool) *TableRow {
	if from == nil {
		return nil
	}

	cells := TableCellSlice{}.From(from.Cells, goTypes, ignoreGoTypes)

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
			GoType:  goTypeString,
			GoName:  goName(value),
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
func (to TableRowSlice) From(from []*messages.TableRow) TableRowSlice {
	to = make(TableRowSlice, 0, len(from))

	for _, f := range from {
		to = append(to, (&TableRow{}).From(f, determinateGoTypes(from), false))
	}

	return to
}

// Tag is a gherkin's tag.
type Tag struct {
	Location *Location `json:"Location"`
	Name     string    `json:"Name"`
	ID       string    `json:"ID"`
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
	Line   int64 `json:"Line"`
	Column int64 `json:"Column,omitempty"`
}

// From converts to Location.
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
