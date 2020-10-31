package appenv

import (
	"github.com/kyoh86/appenv/gen"
	"github.com/kyoh86/appenv/types"
)

// HostName is the string property
type HostName struct {
	types.StringPropertyBase
}

// Default is the default value for host-name property
func (*HostName) Default() interface{} {
	return "kyoh86.dev"
}

type DryRun struct {
	types.BoolPropertyBase
}

type Token struct {
	types.StringPropertyBase
}

func Example_generator() {
	var (
		outputPackagePath = "github.com/kyoh86/appenv/gen/example"
		outputDir         = "./example"
		generator         = &gen.Generator{
			BuildTag: "sample",
		}
	)

	if err := generator.Do(
		outputPackagePath,
		outputDir,
		// "hostName" property in the YAML file and DRY_RUN environment variable
		gen.Prop(new(HostName), gen.YAML(), gen.Envar()),
		// "DRY_RUN" property in the environment variable
		gen.Prop(new(DryRun), gen.Envar()),
		// "Token" property in the Keyring
		gen.Prop(new(Token), gen.Keyring()),
	); err != nil {
		panic(err)
	}
	// Output:
}
