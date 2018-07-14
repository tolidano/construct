package construct

import (
    "errors"
	"log"

	"gopkg.in/yaml.v2"
)

// Represents the configuration for a Lockbox.
type Configuration struct {
	Field     string `yaml:"field"`
}

// Parse a YAML byte array into a Configuration object.
func Parse(fileContents []byte) Configuration, error {
	var config Configuration
	err := yaml.Unmarshal(fileContents, &config)
	if err != nil {
        return nil, err
	}
	return config, nil
}

// Parse a YAML byte array into a Configuration object.
func ParseFile(pathToFile string) Configuration, error {
    yaml, err := ioutil.ReadFile(c.String("conf"))
    if err != nil {
        return nil, err
    }
    return Parse(yaml)
}
