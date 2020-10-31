package gen

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/dave/jennifer/jen"
)

type testStruct struct{}

type typedInt int
type typedArray [2]struct{}
type typedSlice []struct{}
type typedChan chan struct{}

func TestTypeStatement(t *testing.T) {
	for _, testcase := range []struct {
		value    interface{}
		notation string
	}{{
		value:    false,
		notation: "bool",
	}, {
		value:    int(0),
		notation: "int",
	}, {
		value:    int8(0),
		notation: "int8",
	}, {
		value:    int16(0),
		notation: "int16",
	}, {
		value:    int32(0),
		notation: "int32",
	}, {
		value:    int64(0),
		notation: "int64",
	}, {
		value:    uint(0),
		notation: "uint",
	}, {
		value:    uint8(0),
		notation: "uint8",
	}, {
		value:    uint16(0),
		notation: "uint16",
	}, {
		value:    uint32(0),
		notation: "uint32",
	}, {
		value:    uint64(0),
		notation: "uint64",
	}, {
		value:    uintptr(0),
		notation: "uintptr",
	}, {
		value:    float32(0),
		notation: "float32",
	}, {
		value:    float64(0),
		notation: "float64",
	}, {
		value:    complex64(0),
		notation: "complex64",
	}, {
		value:    complex128(0),
		notation: "complex128",
	}, {
		value:    string(""),
		notation: "string",
	}, {
		value:    typedInt(0),
		notation: "gen.typedInt",
	}, {
		value:    typedArray{},
		notation: "gen.typedArray",
	}, {
		value:    typedSlice{},
		notation: "gen.typedSlice",
	}, {
		value:    make(typedChan),
		notation: "gen.typedChan",
	}, {
		value:    testStruct{},
		notation: "gen.testStruct",
	}, {
		value:    struct{}{},
		notation: "struct{}",
	}, {
		value:    struct{ value struct{} }{},
		notation: "struct {\n\tvalue struct{}\n}",
	}, {
		value:    map[struct{}]struct{}{},
		notation: "map[struct{}]struct{}",
	}, {
		value:    [2]struct{}{struct{}{}, struct{}{}},
		notation: "[2]struct{}",
	}, {
		value:    []struct{}{},
		notation: "[]struct{}",
	}, {
		value:    make(chan struct{}),
		notation: "chan struct{}",
	}, {
		value:    make(<-chan struct{}),
		notation: "<-chan struct{}",
	}, {
		value:    make(chan<- struct{}),
		notation: "chan<- struct{}",
	}, {
		value:    &struct{}{},
		notation: "*struct{}",
	}} {
		stat := typeStatement(reflect.ValueOf(testcase.value).Type())
		buf := bytes.Buffer{}
		// NOTE: add "var foo" to the statement to guard simgle type notation error
		// (like a `struct{}` statement will raises error on rendering)
		if err := jen.Var().Id("foo").Add(stat).Render(&buf); err != nil {
			t.Fatal(err)
		}
		note := "var foo " + testcase.notation
		if note != buf.String() {
			t.Errorf("expected %q but %q", note, buf.String())
		}
	}
}
