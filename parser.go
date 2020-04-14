package schmp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
)

type parser func(options ProcessingOptions) ([]map[string]interface{}, error)

func getParser(mode string) (parser, error) {
	switch mode {
	case "json":
		return buildParserUsingUnmarshaller(json.Unmarshal), nil
	case "yaml":
		return buildParserUsingUnmarshaller(yaml.Unmarshal), nil
	case "toml":
		return buildParserUsingUnmarshaller(toml.Unmarshal), nil
	}

	return nil, fmt.Errorf(`mode "%s" is not supported`, mode)
}

func buildParserUsingUnmarshaller(unmarshal func([]byte, interface{}) error) func(options ProcessingOptions) ([]map[string]interface{}, error) {
	return func(options ProcessingOptions) ([]map[string]interface{}, error) {
		results := make([]map[string]interface{}, len(options.Sources))
		for i, v := range options.Sources {
			b, err := ioutil.ReadAll(v)
			if err != nil {
				return nil, fmt.Errorf("failed to read source at index %d: %s", i, err.Error())
			}
			results[i] = make(map[string]interface{})
			err = unmarshal(b, &results[i])
			if err != nil {
				return nil, fmt.Errorf("failed to parse data: %s", err.Error())
			}
		}
		return results, nil
	}
}
