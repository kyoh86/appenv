package appenv

import (
	"github.com/kyoh86/appenv/gen"
	"github.com/kyoh86/appenv/types"
)

// Token is a string option
type Token struct {
	types.StringPropertyBase
}

// HostName is a string option with default value
type HostName struct {
	types.StringPropertyBase
}

// Default is the default value for host-name option
func (*HostName) Default() interface{} {
	return "kyoh86.dev"
}

// DryRun is a boolean option
type DryRun struct {
	types.BoolPropertyBase
}

func Example_generator() {
	var (
		outputPackagePath = "github.com/kyoh86/appenv/example"
		outputDir         = "./example"
		generator         = &gen.Generator{
			BuildTag: "sample",
		}
	)

	if err := generator.Do(
		outputPackagePath,
		outputDir,
		// "Token" property in the Keyring
		gen.Prop(new(Token), gen.Keyring()),
		// "hostName" property in the YAML file and DRY_RUN environment variable
		gen.Prop(new(HostName), gen.YAML(), gen.Envar()),
		// "DRY_RUN" property in the environment variable
		gen.Prop(new(DryRun), gen.Envar()),
	); err != nil {
		panic(err)
	}
	// Output:
}
