package conf

import (
	"os"

	"gopkg.in/yaml.v3"
)

func NewConf(path string) (*Conf, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var conf Conf
	err = yaml.Unmarshal(content, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
