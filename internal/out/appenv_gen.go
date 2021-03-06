// Code generated by generate.go DO NOT EDIT.

package out

import "io"

func GetAppenv(yamlReader io.Reader, envarPrefix string) (config Config, access Access, err error) {
	yml, err := loadYAML(yamlReader)
	if err != nil {
		return config, access, err
	}
	config, err = buildConfig(yml)
	if err != nil {
		return config, access, err
	}
	access, err = buildAccess(yml, envarPrefix)
	if err != nil {
		return config, access, err
	}
	return config, access, nil
}
