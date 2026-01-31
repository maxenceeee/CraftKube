package parser

import (
	"os"

	"github.com/goccy/go-yaml"
	"xamence.eu/craftkube/internal"
)

func ParseServiceYAMLFile(filePath string) (*internal.Service, error) {
	var service internal.Service

	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, &service)
	if err != nil {
		return nil, err
	}
	return &service, nil
}
