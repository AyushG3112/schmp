package parser

import (
	"fmt"

	"github.com/AyushG3112/schmp/options"
)

type Parser func(options options.ProcessingOptions) ([]map[string]interface{}, error)

func Get(mode string) (Parser, error) {
	switch mode {
	case "json":
		return parseJSON, nil
	case "yaml":
		return parseYAML, nil
	case "toml":
		return parseTOML, nil
	}

	return nil, fmt.Errorf(`mode "%s" is not supported`, mode)
}
