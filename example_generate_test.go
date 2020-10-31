package appenv_test

import (
	"github.com/kyoh86/appenv"
	"github.com/kyoh86/appenv/internal/def"
)

var _ = "keep"

func Example_generate() {
	var (
		outputPackagePath = "github.com/kyoh86/appenv/internal/out"
		outputDir         = "./internal/out"
	)
	if err := appenv.Generate(
		outputPackagePath,
		outputDir,
		// "Token" property in the YAML file
		appenv.Prop(new(def.Token), appenv.YAML()),
		// "hostName" property in the YAML file and DRY_RUN environment variable
		appenv.Prop(new(def.HostName), appenv.YAML(), appenv.Envar()),
		// "DRY_RUN" property in the environment variable
		appenv.Prop(new(def.DryRun), appenv.Envar()),
	); err != nil {
		panic(err)
	}

	/* Name, type and default value of options must be defined like below:

	import (
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
	*/

	// Output:
}
