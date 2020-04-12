package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/AyushG3112/schmp/options"
)

type jsonParser struct{}

func (p *jsonParser) Parse(options options.ProcessingOptions) ([]map[string]interface{}, error) {
	results := make([]map[string]interface{}, len(options.Sources))
	for i, v := range options.Sources {
		b, err := ioutil.ReadAll(v)
		if err != nil {
			return nil, fmt.Errorf("failed to read source at index %d: %s", i, err.Error())
		}
		results[i] = make(map[string]interface{})
		json.Unmarshal(b, &results[i])
	}
	return results, nil
}
