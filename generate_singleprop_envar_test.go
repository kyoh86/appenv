package appenv_test

import (
	"strings"
	"testing"

	testtarget "github.com/kyoh86/appenv"
	"github.com/kyoh86/appenv/internal/def"
	"github.com/kyoh86/appenv/internal/fs"
)

func TestGenerateSinglePropEnvar(t *testing.T) {
	t.Parallel()
	res := fs.Memory{}
	gen := &testtarget.Generator{}
	if err := gen.Render(
		"github.com/kyoh86/appenv/testdir",
		res,
		testtarget.Opt(new(def.HostName), testtarget.StoreEnvar()),
	); err != nil {
		t.Fatalf("render: %s", err)
	}

	t.Run(testtarget.AccessFile, func(t *testing.T) {
		content := res.Result(testtarget.AccessFile)
		if !strings.Contains(content, "func GetAccess(envarPrefix string) (access Access, err error)") {
			t.Errorf("expected access builder does not exit in %s", testtarget.AccessFile)
		}
		if !strings.Contains(content, "func (a *Access) HostName() string") {
			t.Errorf("expected access func does not exit in %s", testtarget.AccessFile)
		}
	})

	t.Run(testtarget.EnvarFile, func(t *testing.T) {
		content := res.Result(testtarget.EnvarFile)
		if !strings.Contains(content, "type Envar struct") {
			t.Errorf("expected Envar struct does not exit in %s", testtarget.EnvarFile)
		}
		if !strings.Contains(content, "HostName *def.HostName") {
			t.Errorf("expected Envar member does not exit in %s", testtarget.EnvarFile)
		}
	})

	if res.Result(testtarget.ConfigFile) != "" {
		t.Errorf("unexpected output: %s", testtarget.KeyringFile)
	}
	if res.Result(testtarget.AppenvFile) != "" {
		t.Errorf("unexpected output: %s", testtarget.KeyringFile)
	}
	if res.Result(testtarget.YAMLFile) != "" {
		t.Errorf("unexpected output: %s", testtarget.YAMLFile)
	}
	if res.Result(testtarget.KeyringFile) != "" {
		t.Errorf("unexpected output: %s", testtarget.KeyringFile)
	}
}
