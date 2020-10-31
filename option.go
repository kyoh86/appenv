package appenv

import (
	"reflect"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/kyoh86/appenv/types"
	"github.com/stoewer/go-strcase"
)

type store interface {
	mark(d *option)
}

type storeFunc func(d *option)

func (f storeFunc) mark(d *option) {
	f(d)
}

// NOTE: StoreXXX must not be merged to the Value.
// They can be expand for other storages, but that is NOT
// realistic to implement all of them in all of options.

// StoreYAML will store option value into the YAML file.
func StoreYAML() store { return storeFunc(func(d *option) { d.storeYAML = true }) }

// StoreEnvar will load option value from a environment variables.
func StoreEnvar() store { return storeFunc(func(d *option) { d.storeEnvar = true }) }

// StoreKeyring will store option value into the keyring services.
func StoreKeyring() store { return storeFunc(func(d *option) { d.storeKeyring = true }) }

// Opt describes an option with a defined value and the stores (StoreYAML, StoreEnvar() or StoreKeyring()).
func Opt(value types.Value, s store, stores ...store) (d *option) {
	d = new(option)
	typ := reflect.ValueOf(value).Type()
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	d.pkgPath = typ.PkgPath()
	d.name = typ.Name()

	d.valueType = typeStatement(reflect.ValueOf(value.Value()).Type())

	s.mark(d)
	for _, s := range stores {
		s.mark(d)
	}
	_, d.mask = value.(types.Mask)

	d.camelName = strcase.LowerCamelCase(d.name)
	d.snakeName = strcase.UpperSnakeCase(d.name)
	d.kebabName = strcase.KebabCase(d.name)
	d.dottedName = strings.ReplaceAll(d.kebabName, "-", ".")
	return
}

type option struct {
	pkgPath string
	name    string

	camelName  string
	snakeName  string
	kebabName  string
	dottedName string

	storeYAML    bool
	storeEnvar   bool
	storeKeyring bool

	mask bool

	valueType *jen.Statement
}
