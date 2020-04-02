package gen

import (
	"reflect"

	"github.com/dave/jennifer/jen"
)

func typeStatement(t reflect.Type) *jen.Statement {
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128,
		reflect.Bool, reflect.String, reflect.Uintptr:
		// typed in the package
		if path := t.PkgPath(); path != "" {
			return jen.Qual(path, t.Name())
		}
		// primitive value (or named)
		return jen.Id(t.Name())

	case reflect.Struct:
		// anonymous struct
		if name := t.Name(); name == "" {
			if t.NumField() == 0 {
				return jen.Struct()
			}
			return jen.StructFunc(func(s *jen.Group) {
				for i := 0; i < t.NumField(); i++ {
					s.Id(t.Field(i).Name).Add(typeStatement(t.Field(i).Type))
				}
			})
		}
		// struct in the package
		if path := t.PkgPath(); path != "" {
			return jen.Qual(path, t.Name())
		}
		// primitive struct (this is a dead statement)
		return jen.Id(t.Name())
	case reflect.Map:
		return jen.Map(typeStatement(t.Key())).Add(typeStatement(t.Elem()))
	case reflect.Slice:
		return jen.Index().Add(typeStatement(t.Elem()))
	case reflect.Array:
		return jen.Index(jen.Lit(t.Len())).Add(typeStatement(t.Elem()))
	case reflect.Chan:
		switch t.ChanDir() {
		case reflect.RecvDir: // <-chan
			return jen.Op("<-").Chan().Add(typeStatement(t.Elem()))
		case reflect.SendDir: // chan<-
			return jen.Chan().Op("<-").Add(typeStatement(t.Elem()))
		case reflect.BothDir: // chan
			return jen.Chan().Add(typeStatement(t.Elem()))
		}
	case reflect.Ptr:
		return jen.Op("*").Add(typeStatement(t.Elem()))

		// unsupport reflect.Func
		// unsupport reflect.Interface
		// unsupport reflect.UnsafePointer
	}
	panic("unsupported value type")
}
