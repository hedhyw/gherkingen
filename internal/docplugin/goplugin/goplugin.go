package goplugin

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/hedhyw/gherkingen/v2/internal/docplugin/goplugin/goaliaser"
	"github.com/hedhyw/gherkingen/v2/internal/model"
)

const maxRecursionDepth = 10

const (
	dataFieldGoType  = "GoType"
	dataFieldGoValue = "GoValue"
	dataFieldGoName  = "GoName"
)

// GoPlugin injects golang specific information: go types, aliases.
type GoPlugin struct {
	aliaser             *goaliaser.Aliaser
	exampleNameReplacer *strings.Replacer
}

// New initializes a new go plugin.
func New() *GoPlugin {
	return &GoPlugin{
		aliaser:             goaliaser.New(),
		exampleNameReplacer: strings.NewReplacer(" ", "_", "\t", "_"),
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

// nolint: cyclop // Many options in switch case.
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
	case model.Rule:
		val.PluginData[dataFieldGoName] = p.aliaser.NameAlias(val.Keyword)
		val.PluginData[dataFieldGoValue] = p.aliaser.StringValue(val.Name)
		val.PluginData[dataFieldGoType] = string(goTypeString)
	case model.Scenario:
		val.PluginData[dataFieldGoName] = p.aliaser.NameAlias(val.Keyword)
		val.PluginData[dataFieldGoValue] = p.aliaser.StringValue(val.Name)
		val.PluginData[dataFieldGoType] = string(goTypeString)
	case model.Step:
		val.PluginData[dataFieldGoName] = p.aliaser.NameAlias(val.Keyword)
		val.PluginData[dataFieldGoValue] = p.aliaser.StringValue(val.Text)
		val.PluginData[dataFieldGoType] = string(goTypeString)
	case model.TableCell:
		val.PluginData[dataFieldGoValue] = p.aliaser.StringValue(val.Value)
		if val.PluginData[dataFieldGoType] == nil {
			val.PluginData[dataFieldGoType] = string(goTypeString)
		}
	case model.TableRow:
		val.PluginData[dataFieldGoValue] = p.prepareExampleName(val)
	}

	return nil
}

func (p GoPlugin) prepareExampleName(row model.TableRow) string {
	values := make([]string, 0, len(row.Cells))
	for _, c := range row.Cells {
		values = append(values, p.exampleNameReplacer.Replace(c.Value))
	}

	return p.aliaser.StringValue(strings.Join(values, "_"))
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

		cell := examples.TableHeader.Cells[i]
		cell.PluginData[dataFieldGoType] = goTypeString
		cell.PluginData[dataFieldGoValue] = p.aliaser.StringValue(cell.Value)
		cell.PluginData[dataFieldGoName] = p.aliaser.NameAlias(cell.Value)

		for _, row := range examples.TableBody {
			row.PluginData[dataFieldGoType] = string(determinateGoType(values))
		}
	}

	return nil
}

// Name of the plugin.
func (p GoPlugin) Name() string {
	return "GoPlugin"
}
