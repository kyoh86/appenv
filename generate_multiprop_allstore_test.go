package appenv_test

import (
	"strings"
	"testing"

	testtarget "github.com/kyoh86/appenv"
	"github.com/kyoh86/appenv/internal/def"
	"github.com/kyoh86/appenv/internal/fs"
)

func TestGenerateMultiPropAllStore(t *testing.T) {
	t.Parallel()
	res := fs.Memory{}
	gen := &testtarget.Generator{}
	if err := gen.Render(
		"github.com/kyoh86/appenv/testdir",
		res,
		testtarget.Opt(new(def.HostName), testtarget.StoreYAML()),
		testtarget.Opt(new(def.Token), testtarget.StoreKeyring()),
		testtarget.Opt(new(def.DryRun), testtarget.StoreEnvar()),
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
		if !strings.Contains(content, "func GetAccess(yamlReader io.Reader, keyringService string, envarPrefix string) (access Access, err error)") {
			t.Errorf("expected access builder does not exit in %s", testtarget.AccessFile)
		}
		if !strings.Contains(content, "func (a *Access) HostName() string") {
			t.Errorf("expected access func HostName does not exit in %s", testtarget.AccessFile)
		}
		if !strings.Contains(content, "func (a *Access) Token() string") {
			t.Errorf("expected access func Token does not exit in %s", testtarget.AccessFile)
		}
		if !strings.Contains(content, "func (a *Access) DryRun() bool") {
			t.Errorf("expected access func DryRun does not exit in %s", testtarget.AccessFile)
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
		if !strings.Contains(content, "func GetConfig(yamlReader io.Reader, keyringService string) (config Config, err error)") {
			t.Errorf("expected config builder does not exit in %s", testtarget.ConfigFile)
		}
		if !strings.Contains(content, "func (a *Config) HostName() types.Config") {
			t.Errorf("expected access func HostName does not exit in %s", testtarget.ConfigFile)
		}
		if !strings.Contains(content, "func (a *Config) Token() types.Config") {
			t.Errorf("expected access func Token does not exit in %s", testtarget.ConfigFile)
		}
		if strings.Contains(content, "DryRun()") {
			t.Errorf("unexpected access func DryRun is found in %s", testtarget.ConfigFile)
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
		if !strings.Contains(content, "func GetAppenv(yamlReader io.Reader, keyringService string, envarPrefix string) (config Config, access Access, err error)") {
			t.Errorf("expected appenv builder does not exit in %s", testtarget.AppenvFile)
		}
	})

	t.Run(testtarget.YAMLFile, func(t *testing.T) {
		content := res.Result(testtarget.YAMLFile)
		lines := strings.Split(content, "\n")
		if len(lines) < len(headers) {
			t.Fatalf("YAML file shortage: %#v", lines)
		}
		for index, header := range headers {
			if lines[index] != header {
				t.Errorf("expected header %q does not equal %q at %d", header, lines[index], index)
			}
		}
		if !strings.Contains(content, "type YAML struct") {
			t.Errorf("expected YAML struct does not exit in %s", testtarget.YAMLFile)
		}
		if !strings.Contains(content, "HostName *def.HostName `yaml:\"hostName,omitempty\"`") {
			t.Errorf("expected YAML member does not exit in %s", testtarget.YAMLFile)
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
		if !strings.Contains(content, "Token *def.Token") {
			t.Errorf("expected Keyring member does not exit in %s", testtarget.KeyringFile)
		}
	})

	t.Run(testtarget.EnvarFile, func(t *testing.T) {
		content := res.Result(testtarget.EnvarFile)
		lines := strings.Split(content, "\n")
		if len(lines) < len(headers) {
			t.Fatalf("Envar file shortage: %#v", lines)
		}
		for index, header := range headers {
			if lines[index] != header {
				t.Errorf("expected header %q does not equal %q at %d", header, lines[index], index)
			}
		}
		if !strings.Contains(content, "type Envar struct") {
			t.Errorf("expected Envar struct does not exit in %s", testtarget.EnvarFile)
		}
		if !strings.Contains(content, "DryRun *def.DryRun") {
			t.Errorf("expected Envar member does not exit in %s", testtarget.EnvarFile)
		}
	})
}
