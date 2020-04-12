package parser

import (
	"fmt"
	"io/ioutil"

	"github.com/AyushG3112/schmp/options"
	"gopkg.in/yaml.v2"
)

type yamlParser struct{}

func (p *yamlParser) Parse(options options.ProcessingOptions) ([]map[string]interface{}, error) {
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
