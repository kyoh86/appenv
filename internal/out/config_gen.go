// Code generated by github.com/kyoh86/appenv.Generator DO NOT EDIT.

package out

import (
	"fmt"
	"io"

	keyring "github.com/99designs/keyring"
	def "github.com/kyoh86/appenv/internal/def"
	types "github.com/kyoh86/appenv/types"
)

type Config struct {
	yml     YAML
	keyring Keyring
}

func GetConfig(yamlReader io.Reader, keyringConfig *keyring.Config) (config Config, err error) {
	yml, err := loadYAML(yamlReader)
	if err != nil {
		return config, err
	}
	keyring, err := loadKeyring(keyringConfig)
	if err != nil {
		return config, err
	}
	return buildConfig(yml, keyring)
}

func buildConfig(yml YAML, keyring Keyring) (config Config, err error) {
	config.yml = yml
	config.keyring = keyring
	return
}

func (c *Config) Save(yamlWriter io.Writer, keyringConfig *keyring.Config) error {
	if err := saveYAML(yamlWriter, &c.yml); err != nil {
		return err
	}
	if err := saveKeyring(keyringConfig, &c.keyring); err != nil {
		return err
	}
	return nil
}

func OptionNames() []string {
	return []string{"token", "host.name"}
}

func (a *Config) Option(name string) (types.Config, error) {
	switch name {
	case "token":
		return &tokenConfig{parent: a}, nil
	case "host.name":
		return &hostNameConfig{parent: a}, nil
	}
	return nil, fmt.Errorf("invalid option name %q", name)
}

func (a *Config) Token() types.Config {
	return &tokenConfig{parent: a}
}

type tokenConfig struct {
	parent *Config
}

func (a *tokenConfig) Get() (string, error) {
	{
		p := a.parent.keyring.Token
		if p != nil {
			text, err := p.MarshalText()
			return string(text), err
		}
	}
	return "", nil
}

func (a *tokenConfig) Set(value string) error {
	{
		p := a.parent.keyring.Token
		if p == nil {
			p = new(def.Token)
		}
		if err := p.UnmarshalText([]byte(value)); err != nil {
			return err
		}
		a.parent.keyring.Token = p
	}
	return nil
}

func (a *tokenConfig) Unset() {
	a.parent.keyring.Token = nil
}

func (a *Config) HostName() types.Config {
	return &hostNameConfig{parent: a}
}

type hostNameConfig struct {
	parent *Config
}

func (a *hostNameConfig) Get() (string, error) {
	{
		p := a.parent.yml.HostName
		if p != nil {
			text, err := p.MarshalText()
			return string(text), err
		}
	}
	return "", nil
}

func (a *hostNameConfig) Set(value string) error {
	{
		p := a.parent.yml.HostName
		if p == nil {
			p = new(def.HostName)
		}
		if err := p.UnmarshalText([]byte(value)); err != nil {
			return err
		}
		a.parent.yml.HostName = p
	}
	return nil
}

func (a *hostNameConfig) Unset() {
	a.parent.yml.HostName = nil
}
