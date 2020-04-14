package schmp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
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
				return nil, err
			}
			results[i] = make(map[string]interface{})
			err = unmarshal(b, &results[i])
			if err != nil {
				return nil, err
			}
		}
		return results, nil
	}
}
