package gen

import (
	"reflect"

	"github.com/dave/jennifer/jen"
)

func typeStatement(t reflect.Type) *jen.Statement {
	switch t.Kind() {
	case reflect.Bool:
		return jen.Bool()
	case reflect.Int:
		return jen.Int()
	case reflect.Int8:
		return jen.Int8()
	case reflect.Int16:
		return jen.Int16()
	case reflect.Int32:
		return jen.Int32()
	case reflect.Int64:
		return jen.Int64()
	case reflect.Uint:
		return jen.Uint()
	case reflect.Uint8:
		return jen.Uint8()
	case reflect.Uint16:
		return jen.Uint16()
	case reflect.Uint32:
		return jen.Uint32()
	case reflect.Uint64:
		return jen.Uint64()
	case reflect.Uintptr:
		return jen.Uintptr()
	case reflect.Float32:
		return jen.Float32()
	case reflect.Float64:
		return jen.Float64()
	case reflect.Complex64:
		return jen.Complex64()
	case reflect.Complex128:
		return jen.Complex128()
	case reflect.String:
		return jen.String()
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
