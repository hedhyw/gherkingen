package goplugin

import (
	"context"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"unicode"

	"github.com/hedhyw/gherkingen/v3/internal/docplugin/goplugin/goaliaser"
	"github.com/hedhyw/gherkingen/v3/internal/model"
)

const maxRecursionDepth = 10

const (
	dataFieldGoType       = "GoType"
	dataFieldGoValue      = "GoValue"
	dataFieldGoName       = "GoName"
	dataFieldGoComment    = "GoComment"
	dataFieldGoBackground = "GoHasBackground"
	dataFieldGoparallel   = "GoParallel"
)

// GoPlugin injects golang specific information: go types, aliases.
type GoPlugin struct {
	aliaser             *goaliaser.Aliaser
	exampleNameReplacer *strings.Replacer
	usedExampleNames    map[string]struct{}

	args Args
}

// Args contains optional arguments for GoPlugin.
type Args struct {
	Parallel bool
}

// New initializes a new go plugin.
func New(args Args) *GoPlugin {
	return &GoPlugin{
		aliaser:             goaliaser.New(),
		exampleNameReplacer: strings.NewReplacer(" ", "_", "\t", "_"),
		usedExampleNames:    make(map[string]struct{}),

		args: args,
	}
}

// Process document by go plugin.
func (p GoPlugin) Process(
	ctx context.Context,
	document *model.GherkinDocument,
) (err error) {
	return p.walk(ctx, document, 0)
}

func (p GoPlugin) walk(ctx context.Context, val any, depth int) (err error) {
	if depth > maxRecursionDepth {
		return nil
	}

	rt := reflect.TypeOf(val)
	rv := reflect.ValueOf(val)

	if rt.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return nil
		}

		rt = rt.Elem()
		rv = rv.Elem()
	}

	switch rt.Kind() {
	case reflect.Struct:
		return p.walkStruct(ctx, rv, depth)
	case reflect.Slice:
		return p.walkSlice(ctx, rv, depth)
	default:
		return nil
	}
}

func (p GoPlugin) walkStruct(
	ctx context.Context,
	rv reflect.Value,
	depth int,
) (err error) {
	if !rv.CanInterface() {
		return nil
	}

	rt := rv.Type()

	err = p.handleStruct(ctx, rv.Interface())
	if err != nil {
		return fmt.Errorf("handling struct: %s: %w", rt.Name(), err)
	}

	for i := 0; i < rt.NumField(); i++ {
		if !rv.Field(i).CanInterface() {
			continue
		}

		err = p.walk(ctx, rv.Field(i).Interface(), depth+1)
		if err != nil {
			return fmt.Errorf("walking: %w", err)
		}
	}

	return nil
}

func (p GoPlugin) walkSlice(
	ctx context.Context,
	rv reflect.Value,
	depth int,
) (err error) {
	rt := rv.Type()

	for i := 0; i < rv.Len(); i++ {
		el := rv.Index(i)

		if !el.CanInterface() {
			continue
		}

		err = p.walk(ctx, el.Interface(), depth+1)
		if err != nil {
			return fmt.Errorf("walking: %s: %w", rt.Name(), err)
		}
	}

	return nil
}

// nolint:  cyclop // Many options in switch case.
func (p GoPlugin) handleStruct(
	_ context.Context,
	val any,
) error {
	switch val := val.(type) {
	case model.Background:
		val.PluginData[dataFieldGoName] = p.aliaser.NameAlias(val.Keyword)
		val.PluginData[dataFieldGoValue] = p.aliaser.StringValue(val.Name)
		val.PluginData[dataFieldGoType] = string(goTypeString)
	case model.Examples:
		val.PluginData[dataFieldGoName] = p.aliaser.NameAlias(val.Name)
		val.PluginData[dataFieldGoValue] = p.aliaser.StringValue(val.Keyword)
		val.PluginData[dataFieldGoType] = string(goTypeString)

		if err := p.fillExampleHeaderTypes(&val); err != nil {
			return fmt.Errorf("filling examples: %w", err)
		}
	case model.Feature:
		val.PluginData[dataFieldGoName] = p.aliaser.NameAlias(val.Name)
		val.PluginData[dataFieldGoValue] = p.aliaser.StringValue(val.Name)
		val.PluginData[dataFieldGoType] = string(goTypeString)
		val.PluginData[dataFieldGoComment] = p.prepareFeatureDescription(val.Description)
		val.PluginData[dataFieldGoparallel] = p.args.Parallel
		p.processFeatureBackground(val)
	case model.Rule:
		val.PluginData[dataFieldGoName] = p.aliaser.NameAlias(val.Keyword)
		val.PluginData[dataFieldGoValue] = p.aliaser.StringValue(val.Name)
		val.PluginData[dataFieldGoType] = string(goTypeString)
		val.PluginData[dataFieldGoparallel] = p.args.Parallel
		p.processRuleBackground(val)
	case model.Scenario:
		val.PluginData[dataFieldGoName] = p.aliaser.NameAlias(formatScenarioKeyword(val.Keyword))
		val.PluginData[dataFieldGoValue] = p.aliaser.StringValue(val.Name)
		val.PluginData[dataFieldGoType] = string(goTypeString)
		val.PluginData[dataFieldGoparallel] = p.args.Parallel
	case model.Step:
		val.PluginData[dataFieldGoName] = p.aliaser.NameAlias(val.Keyword)
		val.PluginData[dataFieldGoValue] = p.aliaser.StringValue(val.Text)
		val.PluginData[dataFieldGoType] = string(goTypeString)
	case model.TableCell:
		if val.PluginData[dataFieldGoValue] == nil {
			val.PluginData[dataFieldGoValue] = p.aliaser.StringValue(val.Value)
		}
		if val.PluginData[dataFieldGoType] == nil {
			val.PluginData[dataFieldGoType] = string(goTypeString)
		}
	case model.TableRow:
		val.PluginData[dataFieldGoValue] = p.prepareExampleName(val)
	}

	return nil
}

func formatScenarioKeyword(keyword string) string {
	keywordLower := strings.ToLower(keyword)

	if keywordLower == "scenario outline" || keywordLower == "scenario template" {
		keyword = "Scenario"
	}

	return keyword
}

func (p GoPlugin) processFeatureBackground(f model.Feature) {
	var hasBackground bool

	for _, ch := range f.Children {
		if ch.Background != nil {
			hasBackground = true

			break
		}
	}

	if !hasBackground {
		return
	}

	for _, ch := range f.Children {
		if ch.Scenario != nil {
			ch.Scenario.PluginData[dataFieldGoBackground] = true
		}

		if ch.Rule != nil {
			ch.Rule.PluginData[dataFieldGoBackground] = true
		}
	}
}

func (p GoPlugin) processRuleBackground(f model.Rule) {
	var hasBackground bool

	for _, ch := range f.Children {
		if ch.Background != nil {
			hasBackground = true

			break
		}
	}

	if !hasBackground {
		return
	}

	for _, ch := range f.Children {
		if ch.Scenario != nil {
			ch.Scenario.PluginData[dataFieldGoBackground] = true
		}
	}
}

func (p GoPlugin) prepareFeatureDescription(descr string) string {
	lines := strings.Split(descr, "\n")

	if len(lines) == 1 {
		return strings.TrimSpace(descr)
	}

	minIndent := math.MaxInt

	for i, l := range lines {
		l = strings.TrimSpace(l)
		if l == "" {
			lines[i] = l
		}
	}

	for _, l := range lines {
		if l == "" {
			continue
		}

		if spacesCount := countPrefixSpaces(l); spacesCount < minIndent {
			minIndent = spacesCount
		}
	}

	for i, l := range lines {
		if l == "" {
			continue
		}

		lr := []rune(l)

		if minIndent >= len(lr) {
			continue
		}

		lines[i] = "\t" + string(lr[minIndent:])
	}

	return "\n" + strings.Join(lines, "\n") + "\n"
}

func countPrefixSpaces(val string) int {
	var spacesCount int

	for _, r := range val {
		if !unicode.IsSpace(r) {
			break
		}

		spacesCount++
	}

	return spacesCount
}

func (p GoPlugin) prepareExampleName(row model.TableRow) string {
	values := make([]string, 0, len(row.Cells))
	for _, c := range row.Cells {
		values = append(values, p.exampleNameReplacer.Replace(c.Value))
	}

	name := strings.Join(values, "_")

	if _, ok := p.usedExampleNames[name]; ok {
		const duplicateLimit = 50

		for i := 2; i < duplicateLimit; i++ {
			if _, ok := p.usedExampleNames[name+"_"+strconv.Itoa(i)]; !ok {
				name += "_" + strconv.Itoa(i)

				break
			}
		}
	}

	p.usedExampleNames[name] = struct{}{}

	return p.aliaser.StringValue(name)
}

func (p GoPlugin) fillExampleHeaderTypes(examples *model.Examples) (err error) {
	if len(examples.TableBody) == 0 {
		return nil
	}

	for i := range examples.TableBody[0].Cells {
		values := make([]string, 0, len(examples.TableBody))

		for _, row := range examples.TableBody {
			if i >= len(row.Cells) {
				err = outOfRangeError{Len: len(row.Cells), Index: i}

				return fmt.Errorf("invalid body cells: %w", err)
			}

			values = append(values, row.Cells[i].Value)
		}

		if i >= len(examples.TableHeader.Cells) {
			err = outOfRangeError{Len: len(examples.TableHeader.Cells), Index: i}

			return fmt.Errorf("invalid header cells: %w", err)
		}

		goType := determinateGoType(values)
		cell := examples.TableHeader.Cells[i]
		cell.PluginData[dataFieldGoType] = string(goType)
		cell.PluginData[dataFieldGoValue] = p.aliaser.StringValue(cell.Value)
		cell.PluginData[dataFieldGoName] = p.aliaser.NameAlias(cell.Value)

		for _, row := range examples.TableBody {
			cell := row.Cells[i]

			goVal, goType := goValue(p.aliaser, cell.Value, goType)
			cell.PluginData[dataFieldGoValue] = goVal
			cell.PluginData[dataFieldGoType] = goType
		}
	}

	return nil
}

// Name of the plugin.
func (p GoPlugin) Name() string {
	return "GoPlugin"
}
