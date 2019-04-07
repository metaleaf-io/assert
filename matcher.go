package assert

import (
	"errors"
	"reflect"
	"runtime/debug"
	"strconv"
	"strings"
	"testing"
)

// Matcher hold the current state of the assertion.
type Matcher struct {
	t		*testing.T
	actual	interface{}
	match	bool
}

// With creates a new Matcher with the current test reporter.
func With(t *testing.T) *Matcher {
	m := new(Matcher)
	m.t = t
	return m
}

// That specifies the actual value under test.
func (m *Matcher) That(actual interface{}) *Matcher {
	if m.t == nil {
		panic("Use With(*testing.T) to initialize Matcher")
	}

	m.actual = actual
	return m
}

// IsNil verifies the tested valid is `nil`
func (m *Matcher) IsNil() *Matcher {
	m.match = !reflect.ValueOf(m.actual).IsValid()
	return m
}

// IsNotNil verifies the tested value is not `nil`
func (m *Matcher) IsNotNil() *Matcher {
	m.match = reflect.ValueOf(m.actual).IsValid()
	return m
}

// IsEqualTo verifies that the actual value capture in `That()` is equal to the
// expected value.
func (m *Matcher) IsEqualTo(expected interface{}) *Matcher {
	m.match = false
	av := reflect.ValueOf(m.actual)
	ev := reflect.ValueOf(expected)

	// Edge condition: both values are nil. The `IsNil` matcher should be
	// used instead of IsEqualTo(), but we don't want to fail the test over
	// semantics.
	if !av.IsValid() && !ev.IsValid() {
		m.match = true
		return m
	}

	// Both values must be valid.
	if av.IsValid() && ev.IsValid() {
		ak, err := basicKind(av)
		if err != nil {
			m.t.Error(err)
			return m
		}

		ek, err := basicKind(ev)
		if err != nil {
			m.t.Error(err)
			return m
		}

		if ak != ek {
			m.t.Error(errBadComparison)
			return m
		}

		switch ak {
		case boolKind:
			m.match = av.Bool() == ev.Bool()
		case complexKind:
			m.match = av.Complex() == ev.Complex()
		case floatKind:
			m.match = av.Float() == ev.Float()
		case intKind:
			m.match = av.Int() == ev.Int()
		case stringKind:
			m.match = av.String() == ev.String()
		case uintKind:
			m.match = av.Uint() == ev.Uint()
		default:
			m.t.Error(errBadType)
		}
	}

	if !m.match {
		m.t.Errorf("[%s] expected:<[%s]> but was <[%s]>", testLine(), stringValue(ev), stringValue(av))
	}

	return m
}

// stringValue uses reflection to get the value and convert it to a string
// use in error messages.
func stringValue(rv reflect.Value) string {
	switch rv.Kind() {
	case reflect.Bool:
		return strconv.FormatBool(rv.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(rv.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(rv.Uint(), 10)
	case reflect.Float32:
		return strconv.FormatFloat(rv.Float(), 'g', -1, 32)
	case reflect.Float64:
		return strconv.FormatFloat(rv.Float(), 'g', -1, 64)
	case reflect.Complex64:
		c := rv.Complex()
		return "(" + strconv.FormatFloat(real(c), 'g', -1, 32) + "," + strconv.FormatFloat(imag(c), 'g', -1, 32) + ")"
	case reflect.Complex128:
		c := rv.Complex()
		return "(" + strconv.FormatFloat(real(c), 'g', -1, 64) + "," + strconv.FormatFloat(imag(c), 'g', -1, 64) + ")"
	case reflect.String:
		return rv.String()
	default:
		// All of the types have been accounted for above, so this should
		// never be reached.
		panic(errBadType)
	}
}

// testLine returns the line the unit test was run from.
func testLine() string {
	lines := strings.Split(string(debug.Stack()), "\n")
	var source int
	for i, s := range lines {
		if strings.HasPrefix(s, "testing.tRunner") {
			source = i - 1
		}
	}
	return strings.TrimSpace(lines[source])
}

// The following is lifted from https://golang.org/src/text/template/funcs.go
// None of this is available outside of the package, so We're reproducing it.

// Errors returned when comparisons go bad.
var (
	errBadComparisonType    = errors.New("invalid type for comparison")
	errBadComparison        = errors.New("incompatible types for comparison")
	errBadType              = errors.New("invalid type")
)

// These are the basic types, distilled from the variety of more specific types.
type kind int
const (
	invalidKind kind = iota
	boolKind
	complexKind
	intKind
	floatKind
	stringKind
	uintKind
)

// basicKind simplifies the type down to the particular class to which it belongs.
func basicKind(v reflect.Value) (kind, error) {
	switch v.Kind() {
	case reflect.Bool:
		return boolKind, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return intKind, nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return uintKind, nil
	case reflect.Float32, reflect.Float64:
		return floatKind, nil
	case reflect.Complex64, reflect.Complex128:
		return complexKind, nil
	case reflect.String:
		return stringKind, nil
	}
	return invalidKind, errBadComparisonType
}
