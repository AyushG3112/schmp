package parser

import (
	"github.com/AyushG3112/schmp/options"
)

type Parser interface {
	Parse(options options.ProcessingOptions) ([]map[string]interface{}, error)
}
