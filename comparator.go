package schmp

import (
	"reflect"
	"strings"
)

func getTypeString(v interface{}) string {
	if v == nil {
		return "nil"
	}
	return reflect.TypeOf(v).String()

}

func compare(maps []map[string]interface{}, options ProcessingOptions, parent string, typeMap map[string][]string) (map[string][]string, error) {
	seenKeys := make(map[string]bool)
	nmaps := len(maps)
	for i, m := range maps {
		if m == nil {
			continue
		}
		for k, v := range m {
			typeMapKey := k
			if parent != "" {
				typeMapKey = parent + "." + k
			}
			nestedMapList := make([]map[string]interface{}, nmaps)
			if seenKeys[k] {
				continue
			}
			areAllSameTypes := true
			seenKeys[k] = true
			originalType := getTypeString(v)
			isObject := strings.HasPrefix(originalType, "map[string]")
			if isObject {
				if nm, ok := v.(map[string]interface{}); ok {
					nestedMapList[i] = nm
				} else {
					isObject = false
				}
			}
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
					currentType = getTypeString(v2)
					areAllSameTypes = areAllSameTypes && currentType == originalType
					if isObject {
						if nm2, ok := v2.(map[string]interface{}); ok {
							nestedMapList[i2] = nm2
						} else {
							nestedMapList[i2] = nil
						}
					}
				} else {
					if isObject {
						nestedMapList[i2] = nil
					}

					areAllSameTypes = false
				}
				typeList[i2] = currentType
			}
			if !areAllSameTypes {
				typeMap[typeMapKey] = typeList
			} else if isObject {
				typeMap, _ = compare(nestedMapList, options, typeMapKey, typeMap)
			}
		}
	}
	return typeMap, nil
}
