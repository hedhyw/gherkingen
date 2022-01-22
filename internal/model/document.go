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

type GherkinDocument struct {
	URI      string     `json:"URI,omitempty"`
	Feature  *Feature   `json:"Feature,omitempty"`
	Comments []*Comment `json:"Comments"`
}

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

type Background struct {
	Location    *Location `json:"Location"`
	Keyword     string    `json:"Keyword"`
	Name        string    `json:"Name"`
	Description string    `json:"Description"`
	Steps       []*Step   `json:"Steps"`
	ID          string    `json:"ID"`

	goData
}

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

type Comment struct {
	Location *Location `json:"Location"`
	Text     string    `json:"Text"`
}

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

type CommentsSlice []*Comment

func (to CommentsSlice) From(from []*messages.Comment) CommentsSlice {
	to = make(CommentsSlice, 0, len(from))

	for _, f := range from {
		to = append(to, (&Comment{}).From(f))
	}

	return to
}

type DataTable struct {
	Location *Location   `json:"Location"`
	Rows     []*TableRow `json:"Rows"`
}

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

type DocString struct {
	Location  *Location `json:"Location"`
	MediaType string    `json:"MediaType,omitempty"`
	Content   string    `json:"Content"`
	Delimiter string    `json:"Delimiter"`
}

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

type ExamplesSlice []*Examples

func (to ExamplesSlice) From(from []*messages.Examples) ExamplesSlice {
	to = make(ExamplesSlice, 0, len(from))

	for _, f := range from {
		to = append(to, (&Examples{}).From(f))
	}

	return to
}

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

type FeatureChild struct {
	Rule       *Rule       `json:"Rule,omitempty"`
	Background *Background `json:"Background,omitempty"`
	Scenario   *Scenario   `json:"Scenario,omitempty"`
}

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

type FeatureChildrenSlice []*FeatureChild

func (to FeatureChildrenSlice) From(from []*messages.FeatureChild) FeatureChildrenSlice {
	to = make(FeatureChildrenSlice, 0, len(from))

	for _, f := range from {
		to = append(to, (&FeatureChild{}).From(f))
	}

	return to
}

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

type RuleChild struct {
	Background *Background `json:"Background,omitempty"`
	Scenario   *Scenario   `json:"Scenario,omitempty"`
}

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

type RuleChildSlice []*RuleChild

func (to RuleChildSlice) From(from []*messages.RuleChild) RuleChildSlice {
	to = make(RuleChildSlice, 0, len(from))

	for _, f := range from {
		to = append(to, (&RuleChild{}).From(f))
	}

	return to
}

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

type Step struct {
	Location  *Location  `json:"Location"`
	Keyword   string     `json:"Keyword"`
	Text      string     `json:"Text"`
	DocString *DocString `json:"DocString,omitempty"`
	DataTable *DataTable `json:"DataTable,omitempty"`
	ID        string     `json:"ID"`

	goData
}

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

type StepsSlice []*Step

func (to StepsSlice) From(from []*messages.Step) StepsSlice {
	to = make(StepsSlice, 0, len(from))

	for _, f := range from {
		to = append(to, (&Step{}).From(f))
	}

	return to
}

type TableCell struct {
	Location *Location `json:"Location"`
	Value    string    `json:"Value"`

	goData
}

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

type TableCellSlice []*TableCell

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

type TableRow struct {
	Location *Location    `json:"Location"`
	Cells    []*TableCell `json:"Cells"`
	ID       string       `json:"ID"`

	goData
}

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

func (to TableRowSlice) From(from []*messages.TableRow) TableRowSlice {
	to = make(TableRowSlice, 0, len(from))

	for _, f := range from {
		to = append(to, (&TableRow{}).From(f, determinateGoTypes(from), false))
	}

	return to
}

type Tag struct {
	Location *Location `json:"Location"`
	Name     string    `json:"Name"`
	ID       string    `json:"ID"`
}

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

type TagsSlice []*Tag

func (to TagsSlice) From(from []*messages.Tag) TagsSlice {
	to = make(TagsSlice, 0, len(from))

	for _, f := range from {
		to = append(to, (&Tag{}).From(f))
	}

	return to
}

type Location struct {
	Line   int64 `json:"Line"`
	Column int64 `json:"Column,omitempty"`
}

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
