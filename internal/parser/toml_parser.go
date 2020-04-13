package parser

import (
	"fmt"
	"io/ioutil"

	"github.com/AyushG3112/schmp/options"
	"github.com/BurntSushi/toml"
)

type tomlParser struct{}

func parseTOML(options options.ProcessingOptions) ([]map[string]interface{}, error) {
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
