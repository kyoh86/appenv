package appenv_test

import (
	"fmt"
	"os"
	"strings"

	"github.com/kyoh86/appenv/internal/out"
)

func Example_getAccess() {
	var (
		yamlFile    = strings.NewReader(`{token: xxxxx}`)
		envarPrefix = "APPENV_EXAMPLE_ACCESS_"
	)

	os.Setenv(envarPrefix+"HOST_NAME", "kyoh86.dev")

	// Get options from file and envar.
	// out.GetAccess is generated function.
	access, err := out.GetAccess(yamlFile, envarPrefix)
	if err != nil {
		panic(err)
	}

	fmt.Println(access.Token())
	fmt.Println(access.HostName())
	//Output:
	// xxxxx
	// kyoh86.dev
}
