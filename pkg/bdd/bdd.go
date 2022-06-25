package bdd

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
	"testing"
)

// Feature of the test file.
type Feature struct {
	// T should not be nil.
	*testing.T

	level int

	tc       interface{}
	replacer *strings.Replacer

	mu         sync.Mutex
	logRecords []string
}

// NewFeature creates a new feature.
func NewFeature(t *testing.T, name string) *Feature {
	t.Helper()

	f := &Feature{
		T: t,

		tc:       nil,
		replacer: nil,
		level:    0,

		mu:         sync.Mutex{},
		logRecords: nil,
	}

	f.appendLogf("Feature: %s", name)

	return f
}

// LogRecords returns pending log records.
func (f *Feature) LogRecords() []string {
	f.mu.Lock()
	defer f.mu.Unlock()

	logRecords := make([]string, len(f.logRecords))
	copy(logRecords, f.logRecords)

	return logRecords
}

func (f *Feature) printLogs() {
	f.T.Helper()

	logData := strings.Join(f.LogRecords(), "\n")

	f.T.Log("\n" + logData)
}

func (f *Feature) appendLogf(format string, args ...interface{}) {
	f.mu.Lock()
	defer f.mu.Unlock()

	lr := fmt.Sprintf(format, args...)

	if f.replacer != nil {
		lr = f.replacer.Replace(lr)
	}

	prefixSpace := strings.Repeat("\t", f.level)

	f.logRecords = append(f.logRecords, prefixSpace+lr)
}

func (f *Feature) subBlock(name string, fn func(t *testing.T, f *Feature)) {
	f.T.Helper()

	f.T.Run(name, func(t *testing.T) {
		t.Helper()

		f := &Feature{
			T: t,

			level: f.level + 1,

			tc:       f.tc,
			replacer: f.replacer,

			mu:         sync.Mutex{},
			logRecords: f.LogRecords(),
		}

		t.Cleanup(func() {
			if t.Failed() {
				f.T.Helper()
				f.printLogs()
			}
		})

		fn(t, f)
	})
}

// Background defines a background block.
//
// Notice: Background is not running for each step.
// Deprecated: the current template uses a new syntax.
func (f *Feature) Background(name string, fn func(t *testing.T, f *Feature)) {
	f.T.Helper()

	f.subBlock(name, func(t *testing.T, f *Feature) {
		t.Helper()
		f.appendLogf("Background: %s", name)
		fn(t, f)
	})
}

// Rule defines a rule block.
func (f *Feature) Rule(name string, fn func(t *testing.T, f *Feature)) {
	f.T.Helper()

	f.subBlock(name, func(t *testing.T, f *Feature) {
		t.Helper()
		f.appendLogf("Rule: %s", name)
		fn(t, f)
	})
}

// Example defines an example-scenario block.
func (f *Feature) Example(name string, fn func(t *testing.T, f *Feature)) {
	f.T.Helper()

	f.subBlock(name, func(t *testing.T, f *Feature) {
		t.Helper()
		f.appendLogf("Example: %s", name)
		fn(t, f)
	})
}

// Scenario defines a scenario block.
func (f *Feature) Scenario(name string, fn func(t *testing.T, f *Feature)) {
	f.T.Helper()

	f.subBlock(name, func(t *testing.T, f *Feature) {
		t.Helper()
		f.appendLogf("Scenario: %s", name)
		fn(t, f)
	})
}

// Given defines a given block.
func (f *Feature) Given(given string, fn func()) {
	f.T.Helper()
	f.appendLogf("Given %s", given)

	if fn != nil {
		fn()
	}
}

// But defines a but block.
func (f *Feature) But(but string, fn func()) {
	f.T.Helper()
	f.appendLogf("But %s", but)

	if fn != nil {
		fn()
	}
}

// TestCase defines a testcase block.
func (f *Feature) TestCase(name string, tc interface{}, fn func(t *testing.T, f *Feature)) {
	f.T.Helper()

	f.subBlock(name, func(t *testing.T, f *Feature) {
		t.Helper()
		f.appendLogf("# TestCase: %+v", tc)
		f.replacer = prepareReplacer(t, tc)
		fn(t, f)
	})
}

// TestCases defines testcases block.
func (f *Feature) TestCases(testCases interface{}, fn interface{}) {
	f.T.Helper()

	if reflect.TypeOf(fn).Kind() != reflect.Func {
		f.T.Fatalf("cannot call fn: %+v", fn)
	}

	callFn := func(t *testing.T, f *Feature, tc interface{}) {
		t.Helper()

		reflect.ValueOf(fn).Call([]reflect.Value{
			reflect.ValueOf(t),
			reflect.ValueOf(f),
			reflect.ValueOf(tc),
		})
	}

	rt := reflect.TypeOf(testCases)

	if rt.Kind() != reflect.Map {
		f.T.Fatalf("invalid testCases type: %s", rt.Kind())
	}

	iter := reflect.ValueOf(testCases).MapRange()
	for iter.Next() {
		k := iter.Key()
		v := iter.Value()

		if !v.CanInterface() {
			f.T.Fatalf("testcase has unexported fields: %+v", v)
		}

		f.TestCase(k.String(), v.Interface(), func(t *testing.T, f *Feature) {
			t.Helper()

			callFn(t, f, v.Interface())
		})
	}
}

// And defines an and block.
func (f *Feature) And(and string, fn func()) {
	f.T.Helper()
	f.appendLogf("And %s", and)

	if fn != nil {
		fn()
	}
}

// When defines a when block.
func (f *Feature) When(when string, fn func()) {
	f.T.Helper()
	f.appendLogf("When %s", when)

	if fn != nil {
		fn()
	}
}

// Then defines a then block.
func (f *Feature) Then(then string, fn func()) {
	f.T.Helper()
	f.appendLogf("Then %s", then)

	if fn != nil {
		fn()
	}
}

func prepareReplacer(tb testing.TB, testCase interface{}) *strings.Replacer {
	tb.Helper()

	if testCase == nil {
		return nil
	}

	defer func() {
		if r := recover(); r != nil {
			tb.Logf("library panicked: %+v", r)
		}
	}()

	rt := reflect.TypeOf(testCase)
	rv := reflect.ValueOf(testCase)

	if rt.Kind() != reflect.Struct {
		return strings.NewReplacer()
	}

	count := rt.NumField()
	replaceArgs := make([]string, 0, count*2)
	for i := 0; i < count; i++ {
		if !rv.Field(i).CanInterface() {
			continue
		}

		field := rt.Field(i).Tag.Get("field")

		if field == "" {
			continue
		}

		if !strings.HasPrefix(field, "<") {
			field = "<" + field + ">"
		}

		replaceArgs = append(
			replaceArgs,
			field,
			fmt.Sprint(rv.Field(i).Interface()),
		)
	}

	return strings.NewReplacer(replaceArgs...)
}
