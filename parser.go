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
		return parseJSON, nil
	case "yaml":
		return parseYAML, nil
	case "toml":
		return parseTOML, nil
	}

	return nil, fmt.Errorf(`mode "%s" is not supported`, mode)
}

func parseJSON(options ProcessingOptions) ([]map[string]interface{}, error) {
	results := make([]map[string]interface{}, len(options.Sources))
	for i, v := range options.Sources {
		b, err := ioutil.ReadAll(v)
		if err != nil {
			return nil, fmt.Errorf("failed to read source at index %d: %s", i, err.Error())
		}
		results[i] = make(map[string]interface{})
		err = json.Unmarshal(b, &results[i])
		if err != nil {
			return nil, fmt.Errorf("failed to parse json data: %s", err.Error())
		}
	}
	return results, nil
}

func parseTOML(options ProcessingOptions) ([]map[string]interface{}, error) {
	results := make([]map[string]interface{}, len(options.Sources))
	for i, v := range options.Sources {
		b, err := ioutil.ReadAll(v)
		if err != nil {
			return nil, fmt.Errorf("failed to read source at index %d: %s", i, err.Error())
		}
		results[i] = make(map[string]interface{})
		err = toml.Unmarshal(b, &results[i])
		if err != nil {
			return nil, fmt.Errorf("failed to parse toml data: %s", err.Error())
		}
	}
	return results, nil
}

func parseYAML(options ProcessingOptions) ([]map[string]interface{}, error) {
	results := make([]map[string]interface{}, len(options.Sources))
	for i, v := range options.Sources {
		b, err := ioutil.ReadAll(v)
		if err != nil {
			return nil, fmt.Errorf("failed to read source at index %d: %s", i, err.Error())
		}
		results[i] = make(map[string]interface{})
		err = yaml.Unmarshal(b, &results[i])
		if err != nil {
			return nil, fmt.Errorf("failed to parse yaml data: %s", err.Error())
		}
	}
	return results, nil
}
