// Copyright Dwi Siswanto and/or licensed to Dwi Siswanto under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.
// See the LICENSE-ELv2 file in the project root for more information.

/*
Package option provides functions for unmarshaling teler-waf configuration
from JSON and YAML formats into the teler.Options struct.

It includes functions to handle different input sources and formats.
These functions allow convenient loading of configuration data into the
Options struct for further processing.
*/
package option

import "github.com/kitabisa/teler-waf"

// LoadFromJSONBytes to unmarshal the teler-waf JSON
// bytes configuration into the [teler.Options] struct.
func LoadFromJSONBytes(raw []byte) (teler.Options, error) {
	// Unmarshal the JSON into the Options struct
	return unmarshalJSONBytes(raw)
}

// LoadFromJSONString to unmarshal the teler-waf JSON
// string configuration into the [teler.Options] struct.
func LoadFromJSONString(raw string) (teler.Options, error) {
	return LoadFromJSONBytes([]byte(raw))
}

// LoadFromJSONFile to unmarshal the teler-waf JSON
// configuration into the [teler.Options] struct.
func LoadFromJSONFile(path string) (teler.Options, error) {
	// Read the JSON file
	jsonFile, err := readFile(path)
	if err != nil {
		return opt, err
	}

	// Unmarshal the JSON into the Options struct
	return unmarshalJSONBytes(jsonFile)
}

// LoadFromYAMLFile to unmarshal the teler-waf YAML
// bytes configuration into the [teler.Options] struct.
func LoadFromYAMLBytes(raw []byte) (teler.Options, error) {
	// Unmarshal the YAML into the Options struct
	return unmarshalYAMLBytes(raw)
}

// LoadFromYAMLFile to unmarshal the teler-waf YAML
// string configuration into the [teler.Options] struct.
func LoadFromYAMLString(raw string) (teler.Options, error) {
	return LoadFromYAMLBytes([]byte(raw))
}

// LoadFromYAMLFile to unmarshal the teler-waf YAML
// configuration into the [teler.Options] struct.
func LoadFromYAMLFile(path string) (teler.Options, error) {
	// Read the YAML file
	yamlFile, err := readFile(path)
	if err != nil {
		return opt, err
	}

	// Unmarshal the YAML into the Options struct
	return unmarshalYAMLBytes(yamlFile)
}
