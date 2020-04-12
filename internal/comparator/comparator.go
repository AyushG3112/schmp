package comparator

import (
	"github.com/AyushG3112/schmp/options"
)

func Compare(maps []map[string]interface{}, options options.ProcessingOptions) (map[string][]string, error) {
	return compare(maps, options, "", make(map[string][]string))
}
