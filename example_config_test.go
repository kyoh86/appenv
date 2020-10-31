// +build sample

package appenv

import (
	"os"
	"strings"

	"github.com/kyoh86/appenv/example"
)

func Example_configuration() {
	var (
		envarPrefix = "APPENV_EXAMPLE_"
		yaml        = strings.NewReader("")
	)

	config, access, err := example.GetAppenv(yaml, envarPrefix)
	if err != nil {
		panic(err)
	}

	if err = f(&access, &config); err != nil {
		panic(err)
	}

	return config.Save(os.Stdout)
}
