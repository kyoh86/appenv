package appenv_test

import (
	"github.com/kyoh86/appenv"
	"github.com/kyoh86/appenv/internal/def"
)

const (
	outputPackagePath = "github.com/kyoh86/appenv/internal/out"
	outputDir         = "./internal/out"
)

func ExampleGenerate() {
	if err := appenv.Generate(
		outputPackagePath,
		outputDir,
		// Use "Token" option in the YAML file
		appenv.Opt(new(def.Token), appenv.StoreYAML()),
		// Use "hostName" option in the YAML file and HOST_NAME environment variable
		appenv.Opt(new(def.HostName), appenv.StoreYAML(), appenv.StoreEnvar()),
		// Use "DRY_RUN" option in the environment variable
		appenv.Opt(new(def.DryRun), appenv.StoreEnvar()),
	); err != nil {
		panic(err)
	}
}
