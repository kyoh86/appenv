package gen

import (
	"reflect"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/kyoh86/appenv/types"
	"github.com/stoewer/go-strcase"
)

type Store interface {
	mark(d *Option)
}

type storeFunc func(d *Option)

func (f storeFunc) mark(d *Option) {
	f(d)
}

// NOTE: StoreXXX must not be merged to the Value.
// They can be expand for other storages, but that is NOT
// realistic to implement all of them in all of options.

func YAML() Store    { return storeFunc(func(d *Option) { d.storeYAML = true }) }
func Envar() Store   { return storeFunc(func(d *Option) { d.storeEnvar = true }) }
func Keyring() Store { return storeFunc(func(d *Option) { d.storeKeyring = true }) }

func Opt(value types.Value, s Store, stores ...Store) (d *Option) {
	d = new(Option)
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

// Option describes environment option.
// It is in internal package, and it can be generated with gen.Opt.
type Option struct {
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
