package conf

import (
	"gopkg.in/yaml.v3"
	"os"
)

func UnmarshalYAML(source any, path string) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(file, &source)
}
