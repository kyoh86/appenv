package gen

import (
	"reflect"

	"github.com/dave/jennifer/jen"
)

func typeStatement(t reflect.Type) *jen.Statement {
	// named or primitive types
	if t.Name() != "" {
		if path := t.PkgPath(); path != "" {
			return jen.Qual(path, t.Name())
		}
		return jen.Id(t.Name())
	}

	// anonymous complex types
	switch t.Kind() {
	case reflect.Struct:
		if t.NumField() == 0 {
			return jen.Struct()
		}
		return jen.StructFunc(func(s *jen.Group) {
			for i := 0; i < t.NumField(); i++ {
				s.Id(t.Field(i).Name).Add(typeStatement(t.Field(i).Type))
			}
		})
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
