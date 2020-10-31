// +build sample
// Code generated by example_generator_test.go DO NOT EDIT.

package example

import (
	appenv "github.com/kyoh86/appenv"
	yaml "gopkg.in/yaml.v3"
	"io"
)

type YAML struct {
	HostName *appenv.HostName `yaml:"hostName,omitempty"`
}

func saveYAML(w io.Writer, yml *YAML) error {
	return yaml.NewEncoder(w).Encode(yml)
}

var EmptyYAMLReader io.Reader = nil

func loadYAML(r io.Reader) (yml YAML, err error) {
	if r == EmptyYAMLReader {
		return
	}
	err = yaml.NewDecoder(r).Decode(&yml)
	return
}
