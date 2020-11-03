package appenv_test

import (
	"strings"
	"testing"

	testtarget "github.com/kyoh86/appenv"
	"github.com/kyoh86/appenv/internal/def"
	"github.com/kyoh86/appenv/internal/fs"
)

func TestGenerateSinglePropYAML(t *testing.T) {
	t.Parallel()
	res := fs.Memory{}
	gen := &testtarget.Generator{}
	if err := gen.Render(
		"github.com/kyoh86/appenv/testdir",
		res,
		testtarget.Opt(new(def.HostName), testtarget.StoreYAML()),
	); err != nil {
		t.Fatalf("render: %s", err)
	}

	t.Run(testtarget.AccessFile, func(t *testing.T) {
		content := res.Result(testtarget.AccessFile)
		if !strings.Contains(content, "func GetAccess(yamlReader io.Reader) (access Access, err error)") {
			t.Errorf("expected access builder does not exit in %s", testtarget.AccessFile)
		}
		if !strings.Contains(content, "func (a *Access) HostName() string") {
			t.Errorf("expected access func does not exit in %s", testtarget.AccessFile)
		}
	})

	t.Run(testtarget.ConfigFile, func(t *testing.T) {
		content := res.Result(testtarget.ConfigFile)
		if !strings.Contains(content, "func GetConfig(yamlReader io.Reader) (config Config, err error)") {
			t.Errorf("expected config builder does not exit in %s", testtarget.ConfigFile)
		}
		if !strings.Contains(content, "func (a *Config) HostName() types.Config") {
			t.Errorf("expected config func does not exit in %s", testtarget.ConfigFile)
		}
	})

	t.Run(testtarget.AppenvFile, func(t *testing.T) {
		content := res.Result(testtarget.AppenvFile)
		if !strings.Contains(content, "func GetAppenv(yamlReader io.Reader) (config Config, access Access, err error)") {
			t.Errorf("expected appenv builder does not exit in %s", testtarget.AppenvFile)
		}
	})

	t.Run(testtarget.YAMLFile, func(t *testing.T) {
		content := res.Result(testtarget.YAMLFile)
		if !strings.Contains(content, "type YAML struct") {
			t.Errorf("expected YAML struct does not exit in %s", testtarget.YAMLFile)
		}
		if !strings.Contains(content, "HostName *def.HostName `yaml:\"hostName,omitempty\"`") {
			t.Errorf("expected YAML member does not exit in %s", testtarget.YAMLFile)
		}
	})

	if res.Result(testtarget.KeyringFile) != "" {
		t.Errorf("unexpected output: %s", testtarget.KeyringFile)
	}
	if res.Result(testtarget.EnvarFile) != "" {
		t.Errorf("unexpected output: %s", testtarget.EnvarFile)
	}
}
