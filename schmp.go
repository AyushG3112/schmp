package schmp

import (
	"github.com/AyushG3112/schmp/internal/comparator"
	"github.com/AyushG3112/schmp/internal/parser"
	"github.com/AyushG3112/schmp/options"
)

func Compare(options options.ProcessingOptions) (map[string][]string, error) {
	parse, err := parser.Get(options.Mode)
	if err != nil {
		return nil, err
	}
	sm, err := parse(options)
	if err != nil {
		return nil, err
	}
	return comparator.Compare(sm, options)
}
