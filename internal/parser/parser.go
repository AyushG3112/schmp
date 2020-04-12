package parser

import "fmt"

func Get(mode string) (Parser, error) {
	switch mode {
	case "json":
		return &jsonParser{}, nil
	case "yaml":
		return &yamlParser{}, nil
	case "toml":
		return &tomlParser{}, nil
	}

	return nil, fmt.Errorf(`mode "%s" is not supported`, mode)
}
