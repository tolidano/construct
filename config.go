package construct

import (
	"io/ioutil"

	"mod/gopkg.in/yaml.v2@v2.2.1"
)

// Represents the configuration for a Lockbox.
type Configuration struct {
	Field     string `yaml:"field"`
}

// Parse a YAML byte array into a Configuration object.
func Parse(fileContents []byte) (*Configuration, error) {
	var config *Configuration
	err := yaml.Unmarshal(fileContents, &config)
	if err != nil {
        return nil, err
	}
	return config, nil
}

// Parse a YAML byte array into a Configuration object.
func ParseFile(pathToFile string) (*Configuration, error) {
    raw, err := ioutil.ReadFile(pathToFile)
    if err != nil {
        return nil, err
    }
    return Parse(raw)
}