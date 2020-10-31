package appenv

import (
	"github.com/kyoh86/appenv/gen"
	"github.com/kyoh86/appenv/types"
)

// HostName is the string option
type HostName struct {
	types.StringOptionBase
}

// Default is the default value for host-name option
func (*HostName) Default() interface{} {
	return "kyoh86.dev"
}

type DryRun struct {
	types.BoolOptionBase
}

type Token struct {
	types.StringOptionBase
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
		// "hostName" option in the YAML file and DRY_RUN environment variable
		gen.Opt(new(HostName), gen.YAML(), gen.Envar()),
		// "DRY_RUN" option in the environment variable
		gen.Opt(new(DryRun), gen.Envar()),
		// "Token" option in the Keyring
		gen.Opt(new(Token), gen.Keyring()),
	); err != nil {
		panic(err)
	}
	// Output:
}
