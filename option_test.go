package appenv_test

import (
	"github.com/kyoh86/appenv"
	"github.com/kyoh86/appenv/types"
)

type HostName struct {
	types.StringValue
}

func ExampleOpt() {
	_ = appenv.Generate(
		"",
		"",
		// Pass "hostName" option in the YAML file and HOST_NAME environment variable to generator.
		appenv.Opt(new(HostName), appenv.StoreYAML(), appenv.StoreEnvar()),
	)
}
