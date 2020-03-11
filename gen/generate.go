package gen

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/dave/jennifer/jen"
)

type Generator struct {
	PackageName string

	name string

	storeYAML    bool
	storeKeyring bool
	storeEnvar   bool
}

const (
	pkgYAML    = "gopkg.in/yaml.v3"
	pkgKeyring = "github.com/zalando/go-keyring"
	pkgTypes   = "github.com/kyoh86/appenv/types"
	pkgStrcase = "github.com/stoewer/go-strcase"
)

func (g *Generator) init() error {
	g.name = "env.generator"
	if _, file, _, ok := runtime.Caller(2); ok {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(cwd, file)
		if err != nil {
			return err
		}
		g.name = rel
	}

	g.storeYAML = false
	g.storeKeyring = false
	g.storeEnvar = false

	return nil
}

func (g *Generator) createFile(packagePath string) *jen.File {
	var file *jen.File
	if g.PackageName != "" {
		file = jen.NewFilePathName(packagePath, g.PackageName)
	} else {
		file = jen.NewFilePath(packagePath)
	}
	file.HeaderComment(fmt.Sprintf("// Code generated by %s DO NOT EDIT.", g.name))
	return file
}

func (g *Generator) parseProps(properties []*Property) {
	for _, p := range properties {
		g.storeYAML = g.storeYAML || p.storeYAML
		g.storeKeyring = g.storeKeyring || p.storeKeyring
		g.storeEnvar = g.storeEnvar || p.storeEnvar
	}
}

func (g *Generator) genAccess(file *jen.File, properties []*Property) {
	file.Func().Id("GetAccess").ParamsFunc(func(accessParams *jen.Group) {
		if g.storeYAML {
			accessParams.Id("yamlReader").Qual("io", "Reader")
		}
		if g.storeKeyring {
			accessParams.Id("keyringService").String()
		}
		if g.storeEnvar {
			accessParams.Id("envarPrefix").String()
		}
	}).Params(jen.Id("access").Id("Access"), jen.Err().Id("error")).BlockFunc(func(accessCodes *jen.Group) {
		if g.storeYAML {
			accessCodes.List(jen.Id("yml"), jen.Err()).
				Op(":=").Id("loadYAML").Call(jen.Id("yamlReader"))
			accessCodes.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Id("access"), jen.Err()),
			)
		}
		if g.storeKeyring {
			accessCodes.List(jen.Id("keyring"), jen.Err()).
				Op(":=").Id("loadKeyring").Call(jen.Id("keyringService"))
			accessCodes.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Id("access"), jen.Err()),
			)
		}
		if g.storeEnvar {
			accessCodes.List(jen.Id("envar"), jen.Err()).
				Op(":=").Id("getEnvar").Call(jen.Id("envarPrefix"))
			accessCodes.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Id("access"), jen.Err()),
			)
		}
		file.Type().Id("Access").StructFunc(func(accessFields *jen.Group) {
			for _, p := range properties {
				accessFields.Id(p.camelName).Id(p.valueType)

				file.Func().Params(jen.Id("a").Id("*Access")).Id(p.name).Params().Id(p.valueType).Block(
					jen.Return(jen.Id("a").Dot(p.camelName)),
				).Line()

				accessCodes.Id("access").Dot(p.camelName).Op("=").New(jen.Qual(p.pkgPath, p.name)).Dot("Default").Call().Assert(jen.Id(p.valueType))
				if p.storeYAML {
					g.tryAccess(accessCodes, "yml", p)
				}
				if p.storeKeyring {
					g.tryAccess(accessCodes, "keyring", p)
				}
				if p.storeEnvar {
					g.tryAccess(accessCodes, "envar", p)
				}
				accessCodes.Line()
			}
		})
		accessCodes.Return()
	})
}

func (g *Generator) tryAccess(accessCodes *jen.Group, srcName string, p *Property) {
	accessCodes.If(jen.Id(srcName).Dot(p.name).Op("!=").Nil()).Block(
		jen.Id("access").Dot(p.camelName).Op("=").Id(srcName).Dot(p.name).Dot("Value").Call().Assert(jen.Id(p.valueType)),
	)
}

func (g *Generator) genConfig(file *jen.File, properties []*Property) {
	file.Type().Id("Config").StructFunc(func(configFields *jen.Group) {
		if g.storeYAML {
			configFields.Id("yml").Id("YAML")
		}
		if g.storeKeyring {
			configFields.Id("keyring").Id("Keyring")
		}
	}).Line()

	file.Func().Id("GetConfig").ParamsFunc(func(getConfigParams *jen.Group) {
		if g.storeYAML {
			getConfigParams.Id("yamlReader").Qual("io", "Reader")
		}
		if g.storeKeyring {
			getConfigParams.Id("keyringService").String()
		}
	}).Params(jen.Id("config").Id("Config"), jen.Err().Id("error")).BlockFunc(func(getConfigCodes *jen.Group) {
		if g.storeYAML {
			getConfigCodes.List(jen.Id("yml"), jen.Err()).
				Op(":=").Id("loadYAML").Call(jen.Id("yamlReader"))
			getConfigCodes.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Id("config"), jen.Err()),
			)
		}
		if g.storeKeyring {
			getConfigCodes.List(jen.Id("keyring"), jen.Err()).
				Op(":=").Id("loadKeyring").Call(jen.Id("keyringService"))
			getConfigCodes.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Id("config"), jen.Err()),
			)
		}
		if g.storeYAML {
			getConfigCodes.Id("config").Dot("yml").Op("=").Id("yml")
		}
		if g.storeKeyring {
			getConfigCodes.Id("config").Dot("keyring").Op("=").Id("keyring")
		}
		getConfigCodes.Return()
	}).Line()

	file.Func().Params(jen.Id("c").Id("*Config")).Id("Save").ParamsFunc(func(configParams *jen.Group) {
		if g.storeYAML {
			configParams.Id("yamlWriter").Qual("io", "Writer")
		}
		if g.storeKeyring {
			configParams.Id("keyringService").String()
		}
	}).Params(jen.Id("error")).BlockFunc(func(saveConfigCodes *jen.Group) {
		if g.storeYAML {
			saveConfigCodes.If(
				jen.Err().Op(":=").Id("saveYAML").Call(jen.Id("yamlWriter"), jen.Id("&c").Dot("yml")),
				jen.Err().Op("!=").Nil(),
			).Block(
				jen.Return(jen.Err()),
			)
		}
		if g.storeKeyring {
			saveConfigCodes.If(
				jen.Err().Op(":=").Id("saveKeyring").Call(jen.Id("keyringService"), jen.Id("&c").Dot("keyring")),
				jen.Err().Op("!=").Nil(),
			).Block(
				jen.Return(jen.Err()),
			)
		}
		saveConfigCodes.Return(jen.Nil())
	}).Line()

	file.Func().Id("PropertyNames").Call().Params(jen.Index().String()).Block(
		jen.Return().Index().String().ValuesFunc(func(namesList *jen.Group) {
			file.Func().Params(jen.Id("a").Id("*Config")).Id("Property").Params(jen.Id("name").String()).Params(jen.Qual(pkgTypes, "Config"), jen.Id("error")).Block(
				jen.Switch(jen.Id("name")).BlockFunc(func(propSwitch *jen.Group) {
					for _, p := range properties {
						// Add property name
						namesList.Lit(p.dottedName)

						// Add property case
						propSwitch.Case(jen.Lit(p.dottedName)).
							Block(jen.Return(jen.Id("&"+p.camelName+"Config").Values(jen.Dict{
								jen.Id("parent"): jen.Id("a"),
							}), jen.Nil()))

						// Build Poperty Config
						file.Type().Id(p.camelName + "Config").Struct(
							jen.Id("parent").Id("*Config"),
						)

						// Implement "Get" Func
						file.Func().Params(jen.Id("a").Id("*"+p.camelName+"Config")).Id("Get").Params().Params(jen.String(), jen.Id("error")).BlockFunc(func(getCodes *jen.Group) {
							if p.storeYAML {
								g.tryGet(getCodes, "yml", p)
							}
							if p.storeKeyring {
								g.tryGet(getCodes, "keyring", p)
							}
							getCodes.Return(jen.Lit(""), jen.Nil())
						}).Line()

						// Implement "Set" Func
						file.Func().Params(jen.Id("a").Id("*" + p.camelName + "Config")).Id("Set").Params(jen.Id("value").String()).Params(jen.Id("error")).BlockFunc(func(setCodes *jen.Group) {
							if p.storeYAML {
								g.trySet(setCodes, "yml", p)
							}
							if p.storeKeyring {
								g.trySet(setCodes, "keyring", p)
							}
							setCodes.Return(jen.Nil())
						}).Line()

						// Implement "Unset" Func
						file.Func().Params(jen.Id("a").Id("*" + p.camelName + "Config")).Id("Unset").Params().BlockFunc(func(unsetCodes *jen.Group) {
							if p.storeYAML {
								g.tryUnset(unsetCodes, "yml", p)
							}
							if p.storeKeyring {
								g.tryUnset(unsetCodes, "keyring", p)
							}
						}).Line()
					}
				}),
				jen.Return(jen.Nil(), jen.Qual("fmt", "Errorf").Call(jen.Lit("invalid property name %q"), jen.Id("name"))),
			)
		}),
	).Line()

}

func (g *Generator) tryGet(getCodes *jen.Group, srcName string, p *Property) {
	getCodes.Block(
		jen.Id("p").Op(":=").Id("a").Dot("parent").Dot(srcName).Dot(p.name),
		jen.If(jen.Id("p").Op("!=").Nil()).BlockFunc(func(ifBlock *jen.Group) {
			ifBlock.List(jen.Id("text"), jen.Err()).Op(":=").Id("p").Dot("MarshalText").Call()
			if p.mask {
				ifBlock.Return(jen.Id("p").Dot("Mask").Call(jen.String().Call(jen.Id("text"))), jen.Err())
			} else {
				ifBlock.Return(jen.String().Call(jen.Id("text")), jen.Err())
			}
		}),
	)
}

func (g *Generator) trySet(setCodes *jen.Group, srcName string, p *Property) {
	setCodes.Block(
		jen.Id("p").Op(":=").Id("a").Dot("parent").Dot(srcName).Dot(p.name),
		jen.If(jen.Id("p").Op("==").Nil()).Block(
			jen.Id("p").Op("=").New(jen.Qual(p.pkgPath, p.name)),
		),
		jen.If(
			jen.Err().Op(":=").Id("p").Dot("UnmarshalText").Call(jen.Id("[]byte").Call(jen.Id("value"))),
			jen.Err().Op("!=").Nil(),
		).Block(
			jen.Return(jen.Err()),
		),
		jen.Id("a").Dot("parent").Dot(srcName).Dot(p.name).Op("=").Id("p"),
	)
}

func (g *Generator) tryUnset(unsetCodes *jen.Group, srcName string, p *Property) {
	unsetCodes.Id("a").Dot("parent").Dot(srcName).Dot(p.name).Op("=").Nil()
}

func (g *Generator) genYAML(file *jen.File, properties []*Property) {
	file.Type().Id("YAML").StructFunc(func(yamlFields *jen.Group) {
		for _, p := range properties {
			if !p.storeYAML {
				continue
			}
			yamlFields.Id(p.name).
				Op("*").Qual(p.pkgPath, p.name).
				Tag(map[string]string{"yaml": p.camelName + ",omitempty"})
		}
	})
	file.Line()

	file.Func().Id("saveYAML").
		Params(
			jen.Id("w").Qual("io", "Writer"),
			jen.Id("yml").Id("*YAML"),
		).
		Add(jen.Id("error")).
		Block(
			jen.Return(
				jen.Qual("gopkg.in/yaml.v3", "NewEncoder").Call(jen.Id("w")).
					Op(".").
					Id("Encode").Call(jen.Id("yml")),
			),
		)
	file.Line()
	file.Var().Id("EmptyYAMLReader").Qual("io", "Reader").Op("=").Nil()
	file.Func().Id("loadYAML").
		Params(
			jen.Id("r").Qual("io", "Reader"),
		).
		Params(
			jen.Id("yml").Id("YAML"),
			jen.Err().Id("error"),
		).
		Block(
			jen.If(jen.Id("r").Op("==").Id("EmptyYAMLReader")).Block(
				jen.Return(),
			),
			jen.Err().Op("=").Qual("gopkg.in/yaml.v3", "NewDecoder").Call(jen.Id("r")).
				Op(".").
				Id("Decode").Call(jen.Op("&").Id("yml")),
			jen.Return(),
		)
	file.Line()
}

func (g *Generator) genKeyring(file *jen.File, properties []*Property) {
	file.Const().Id("DiscardKeyringService").String().Op("=").Lit("")

	file.Type().Id("Keyring").StructFunc(func(keyringFields *jen.Group) {
		file.Func().Id("loadKeyring").Params(jen.Id("keyringService").String()).Params(jen.Id("key").Id("Keyring"), jen.Err().Id("error")).BlockFunc(func(loadKeyringCodes *jen.Group) {
			file.Func().Id("saveKeyring").Params(jen.Id("keyringService").String(), jen.Id("key").Id("*Keyring")).Params(jen.Err().Id("error")).BlockFunc(func(saveKeyringCodes *jen.Group) {
				loadKeyringCodes.If(jen.Id("keyringService").Op("==").Id("DiscardKeyringService")).Block(jen.Return())
				saveKeyringCodes.If(jen.Id("keyringService").Op("==").Id("DiscardKeyringService")).Block(jen.Return())
				for _, p := range properties {
					if !p.storeKeyring {
						continue
					}
					keyringFields.Id(p.name).
						Op("*").Qual(p.pkgPath, p.name)
					loadKeyringCodes.Block(jen.List(jen.Id("v"), jen.Err()).Op(":=").Qual(pkgKeyring, "Get").
						Call(jen.Id("keyringService"), jen.Lit(p.kebabName)),
						jen.If(jen.Err().Op("==").Nil()).Block(
							jen.Var().Id("value").Qual(p.pkgPath, p.name),
							jen.If(
								jen.Err().Op("=").Id("value").Dot("UnmarshalText").Call(jen.Index().Byte().Parens(jen.Id("v"))),
								jen.Err().Op("!=").Nil(),
							).Block(
								jen.Return(jen.Id("key"), jen.Err()),
							),
							jen.Id("key").Dot(p.name).Op("=").Id("&value"),
						).Else().Block(
							jen.Qual("log", "Printf").Call(jen.Lit("info: there's no secret in "+p.kebabName+"@%s (%v)"), jen.Id("keyringService"), jen.Err()),
						),
					)
					saveKeyringCodes.Block(
						jen.List(jen.Id("buf"), jen.Err()).Op(":=").Id("key").Dot(p.name).Dot("MarshalText").Call(),
						jen.If(jen.Err().Op("!=").Nil()).Block(
							jen.Return(jen.Err()),
						),
						jen.If(
							jen.Err().Op(":=").Qual(pkgKeyring, "Set").Call(jen.Id("keyringService"), jen.Lit(p.kebabName), jen.String().Call(jen.Id("buf"))),
							jen.Err().Op("!=").Nil(),
						).Block(
							jen.Return(jen.Err()),
						),
					)
				}
				loadKeyringCodes.Return()
				saveKeyringCodes.Return(jen.Nil())
			})
		}).Line()
	})
}

func (g *Generator) genEnvar(file *jen.File, properties []*Property) {
	file.Type().Id("Envar").StructFunc(func(envarFields *jen.Group) {
		file.Func().Id("getEnvar").Params(jen.Id("prefix").String()).Params(jen.Id("envar").Id("Envar"), jen.Err().Id("error")).BlockFunc(func(loadEnvarCodes *jen.Group) {
			loadEnvarCodes.Id("prefix").Op("=").Qual(pkgStrcase, "UpperSnakeCase").Call(jen.Id("prefix"))
			for _, p := range properties {
				if !p.storeEnvar {
					continue
				}
				envarFields.Id(p.name).
					Op("*").Qual(p.pkgPath, p.name)

				loadEnvarCodes.Block(jen.List(jen.Id("v")).Op(":=").Qual("os", "Getenv").
					Call(jen.Id("prefix").Op("+").Lit(p.snakeName)),
					jen.If(jen.Id("v").Op("==").Lit("")).Block(
						jen.Qual("log", "Printf").Call(jen.Lit("info: there's no envar %s"+p.snakeName+" (%v)"), jen.Id("prefix"), jen.Err()),
					).Else().Block(
						jen.Var().Id("value").Qual(p.pkgPath, p.name),
						jen.If(
							jen.Err().Op("=").Id("value").Dot("UnmarshalText").Call(jen.Index().Byte().Parens(jen.Id("v"))),
							jen.Err().Op("!=").Nil(),
						).Block(
							jen.Return(jen.Id("envar"), jen.Err()),
						),
						jen.Id("envar").Dot(p.name).Op("=").Id("&value"),
					),
				)
			}
			loadEnvarCodes.Return()
		}).Line()
	}).Line()
}

func (g *Generator) Do(packagePath, outDir string, properties ...*Property) error {
	if err := g.init(); err != nil {
		return err
	}

	full, err := filepath.Abs(outDir)
	if err != nil {
		return err
	}

	g.parseProps(properties)

	accessFile := g.createFile(packagePath)
	g.genAccess(accessFile, properties)
	if err := accessFile.Save(filepath.Join(full, "access_gen.go")); err != nil {
		return err
	}

	configFile := g.createFile(packagePath)
	g.genConfig(configFile, properties)
	if err := configFile.Save(filepath.Join(full, "config_gen.go")); err != nil {
		return err
	}

	if g.storeYAML {
		ymlFile := g.createFile(packagePath)
		ymlFile.ImportAlias(pkgYAML, "yaml")
		g.genYAML(ymlFile, properties)
		if err := ymlFile.Save(filepath.Join(full, "yml_gen.go")); err != nil {
			return err
		}
	}

	if g.storeKeyring {
		keyringFile := g.createFile(packagePath)
		keyringFile.ImportAlias(pkgKeyring, "keyring")
		g.genKeyring(keyringFile, properties)
		if err := keyringFile.Save(filepath.Join(full, "keyring_gen.go")); err != nil {
			return err
		}
	}

	if g.storeEnvar {
		envarFile := g.createFile(packagePath)
		g.genEnvar(envarFile, properties)
		if err := envarFile.Save(filepath.Join(full, "envar_gen.go")); err != nil {
			return err
		}
	}

	return nil
}
