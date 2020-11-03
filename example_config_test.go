package appenv_test

import (
	"os"
	"strings"

	"github.com/kyoh86/appenv/internal/out"
)

func Example_getConfig() {
	var (
		yamlFile = strings.NewReader(`{token: xxxxx}`)
	)

	// Load current option from file and build handler.
	// out.GetConfig is generated function.
	config, err := out.GetConfig(yamlFile, out.DiscardKeyringService)
	if err != nil {
		panic(err)
	}

	// Change option
	if err := config.HostName().Set("example.com"); err != nil {
		panic(err)
	}

	// Save to file (put to std-out for example
	if err := config.Save(os.Stdout, out.DiscardKeyringService); err != nil {
		panic(err)
	}
	//Unordered output:
	// hostName: example.com
}
