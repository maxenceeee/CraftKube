package parser

import "os"

type YAMLParser struct {
	content []byte
}

func NewYAMLParser(content []byte) *YAMLParser {
	return &YAMLParser{
		content: content,
	}
}

func YamlFromPath(path string) (*YAMLParser, error) {
	// Implementation to read YAML file from the given path
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bufferReader := make([]byte, 1024)
	data := []byte{}
	for {
		n, err := file.Read(bufferReader)
		if err != nil {
			return nil, err
		}
		if n == 0 {
			break
		}
		data = append(data, bufferReader[:n]...)
	}

	return NewYAMLParser(data), nil
}
