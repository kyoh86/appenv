package appenv_test

import (
	"strings"
	"testing"

	testtarget "github.com/kyoh86/appenv"
	"github.com/kyoh86/appenv/internal/def"
	"github.com/kyoh86/appenv/internal/fs"
)

func TestGenerateSinglePropKeyring(t *testing.T) {
	t.Parallel()
	res := fs.Memory{}
	gen := &testtarget.Generator{}
	if err := gen.Render(
		"github.com/kyoh86/appenv/testdir",
		res,
		testtarget.Opt(new(def.HostName), testtarget.StoreKeyring()),
	); err != nil {
		t.Fatalf("render: %s", err)
	}

	headers := []string{
		"// Code generated by github.com/kyoh86/appenv.Generator DO NOT EDIT.",
		"",
		"package testdir",
	}

	t.Run(testtarget.AccessFile, func(t *testing.T) {
		content := res.Result(testtarget.AccessFile)
		lines := strings.Split(content, "\n")
		if len(lines) < len(headers) {
			t.Fatalf("access file shortage: %#v", lines)
		}
		for index, header := range headers {
			if lines[index] != header {
				t.Errorf("expected header %q does not equal %q at %d", header, lines[index], index)
			}
		}
		if !strings.Contains(content, "func GetAccess(keyringService string) (access Access, err error)") {
			t.Errorf("expected access builder does not exit in %s", testtarget.AccessFile)
		}
		if !strings.Contains(content, "func (a *Access) HostName() string") {
			t.Errorf("expected access func does not exit in %s", testtarget.AccessFile)
		}
	})

	t.Run(testtarget.ConfigFile, func(t *testing.T) {
		content := res.Result(testtarget.ConfigFile)
		lines := strings.Split(content, "\n")
		if len(lines) < len(headers) {
			t.Fatalf("config file shortage: %#v", lines)
		}
		for index, header := range headers {
			if lines[index] != header {
				t.Errorf("expected header %q does not equal %q at %d", header, lines[index], index)
			}
		}
		if !strings.Contains(content, "func GetConfig(keyringService string) (config Config, err error)") {
			t.Errorf("expected config builder does not exit in %s", testtarget.ConfigFile)
		}
		if !strings.Contains(content, "func (a *Config) HostName() types.Config") {
			t.Errorf("expected config func does not exit in %s", testtarget.ConfigFile)
		}
	})

	t.Run(testtarget.AppenvFile, func(t *testing.T) {
		content := res.Result(testtarget.AppenvFile)
		lines := strings.Split(content, "\n")
		if len(lines) < len(headers) {
			t.Fatalf("appenv file shortage: %#v", lines)
		}
		for index, header := range headers {
			if lines[index] != header {
				t.Errorf("expected header %q does not equal %q at %d", header, lines[index], index)
			}
		}
		if !strings.Contains(content, "func GetAppenv(keyringService string) (config Config, access Access, err error)") {
			t.Errorf("expected appenv builder does not exit in %s", testtarget.AppenvFile)
		}
	})

	t.Run(testtarget.KeyringFile, func(t *testing.T) {
		content := res.Result(testtarget.KeyringFile)
		lines := strings.Split(content, "\n")
		if len(lines) < len(headers) {
			t.Fatalf("Keyring file shortage: %#v", lines)
		}
		for index, header := range headers {
			if lines[index] != header {
				t.Errorf("expected header %q does not equal %q at %d", header, lines[index], index)
			}
		}
		if !strings.Contains(content, "type Keyring struct") {
			t.Errorf("expected Keyring struct does not exit in %s", testtarget.KeyringFile)
		}
		if !strings.Contains(content, "HostName *def.HostName") {
			t.Errorf("expected Keyring member does not exit in %s", testtarget.KeyringFile)
		}
	})

	if res.Result(testtarget.YAMLFile) != "" {
		t.Errorf("unexpected output: %s", testtarget.YAMLFile)
	}
	if res.Result(testtarget.EnvarFile) != "" {
		t.Errorf("unexpected output: %s", testtarget.EnvarFile)
	}
}