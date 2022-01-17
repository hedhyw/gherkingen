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

func (f *Feature) LogRecords() []string {
	f.mu.Lock()
	defer f.mu.Unlock()

	logRecords := make([]string, len(f.logRecords))
	copy(logRecords, f.logRecords)

	return logRecords
}

func (f *Feature) printLogs() {
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

// Scenario defines a scenario block.
func (f *Feature) Scenario(name string, fn func(t *testing.T, f *Feature)) {
	f.T.Helper()

	f.T.Run(name, func(t *testing.T) {
		t.Helper()

		t.Cleanup(func() {
			if t.Failed() {
				f.printLogs()
			}
		})

		f := &Feature{
			T: t,

			level: f.level + 1,

			tc:       f.tc,
			replacer: f.replacer,

			mu:         sync.Mutex{},
			logRecords: f.LogRecords(),
		}

		f.appendLogf("Scenario: %s", name)

		fn(t, f)
	})
}

// Given defines a given block.
func (f *Feature) Given(given string, fn func()) {
	f.T.Helper()
	f.appendLogf("Given: %s", given)

	if fn != nil {
		fn()
	}
}

// TestCase defines a testcase block.
func (f *Feature) TestCase(name string, tc interface{}, fn func(t *testing.T, f *Feature)) {
	f.T.Helper()

	f.T.Run(name, func(t *testing.T) {
		t.Helper()

		f := &Feature{
			T: t,

			level: f.level + 1,

			tc:       tc,
			replacer: prepareReplacer(t, tc),

			mu:         sync.Mutex{},
			logRecords: f.LogRecords(),
		}
		f.appendLogf("# TestCase: %+v", tc)

		t.Cleanup(func() {
			if t.Failed() {
				f.printLogs()
			}
		})

		fn(t, f)
	})
}

// And defines an and block.
func (f *Feature) And(and string, fn func()) {
	f.T.Helper()
	f.appendLogf("And: %s", and)

	if fn != nil {
		fn()
	}
}

// When defines a when block.
func (f *Feature) When(when string, fn func()) {
	f.T.Helper()
	f.appendLogf("When: %s", when)

	if fn != nil {
		fn()
	}
}

// Then defines a then block.
func (f *Feature) Then(then string, fn func()) {
	f.T.Helper()
	f.appendLogf("Then: %s", then)

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
