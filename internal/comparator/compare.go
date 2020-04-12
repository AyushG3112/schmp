package comparator

import (
	"reflect"

	"github.com/AyushG3112/schmp/options"
)

func compare(maps []map[string]interface{}, options options.ProcessingOptions, parent string) (map[string][]string, error) {
	seenKeys := make(map[string]bool)
	typeMap := make(map[string][]string)
	nmaps := len(maps)

	for i, m := range maps {
		if m == nil {
			continue
		}
		for k, v := range m {
			if seenKeys[k] {
				continue
			}
			areAllSameTypes := true
			seenKeys[k] = true
			originalType := reflect.TypeOf(v).String()
			typeList := make([]string, nmaps)
			typeList[i] = originalType
			for i2, m2 := range maps {
				if i == i2 {
					continue
				}
				currentType := ""

				if m2 == nil {
					typeList[i2] = currentType
					continue
				}
				if v2, ok := m2[k]; ok {
					currentType := reflect.TypeOf(v2).String()
					areAllSameTypes = areAllSameTypes && currentType == originalType
				} else {
					areAllSameTypes = false
				}
				typeList[i2] = currentType
			}
			if !areAllSameTypes {
				typeMap[k] = typeList
			}
		}
	}

	return typeMap, nil
}
