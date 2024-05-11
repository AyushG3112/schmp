package schmp

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

type typeMapValue struct {
	typeName string
	parents  []string
}

func getTypeString(v interface{}) string {
	if v == nil {
		return "<nil>"
	}
	return reflect.TypeOf(v).String()

}

func buildTypeMap(m map[string]interface{}, parents []string, typeMap map[string]typeMapValue) error {
	for k, v := range m {
		typeMapKey := k
		if len(parents) > 1 {
			typeMapKey = parents[len(parents)-1] + "." + k
		}
		originalType := getTypeString(v)
		isObject := strings.HasPrefix(originalType, "map[")

		if isObject {
			var nm map[string]interface{}
			switch t := v.(type) {
			case map[string]interface{}:
				nm = t
			case map[interface{}]interface{}:
				nm = make(map[string]interface{})
				for k, v := range t {
					nm[fmt.Sprintf("%v", k)] = v
				}
			default:
				return fmt.Errorf("unsupported type encountered at key %q: %T, value: %v", typeMapKey, v, v)
			}

			parents = append(parents, typeMapKey)
			if err := buildTypeMap(nm, parents, typeMap); err != nil {
				return err
			}
			parents = parents[:len(parents)-1] // Restore parents slice
		}

		typeMap[typeMapKey] = typeMapValue{
			typeName: originalType,
			parents:  append([]string(nil), parents...), // Copy parents slice
		}
	}
	return nil
}

func compare(maps []map[string]interface{}) (map[string][]string, error) {
	diffMap := map[string][]string{}

	typeMaps := make([]map[string]typeMapValue, 0, len(maps))
	for i, v := range maps {
		typeMaps = append(typeMaps, map[string]typeMapValue{})
		buildTypeMap(v, []string{""}, typeMaps[i])
	}

	keySet := make(map[string]struct{})
	for _, typeMap := range typeMaps {
		for k := range typeMap {
			keySet[k] = struct{}{}
		}
	}

	keys := make([]string, 0, len(keySet))
	for k := range keySet {
		keys = append(keys, k)
	}
	sort.Strings(keys)

DIFFLOOP:
	for _, k := range keys {
		vals := make([]typeMapValue, 0, len(typeMaps))

		for _, typeMap := range typeMaps {
			vals = append(vals, typeMap[k])
		}

		allSameType := true
		baseType := vals[0].typeName

		for i := 1; i < len(vals); i++ {
			if vals[i].typeName != baseType {
				allSameType = false
				break
			}
		}

		if allSameType {
			continue
		}

		for i := 0; i < len(vals); i++ {
			for _, parent := range vals[i].parents {
				if _, ok := diffMap[parent]; ok {
					continue DIFFLOOP
				}
			}
		}

		for i := 0; i < len(vals); i++ {
			diffMap[k] = append(diffMap[k], vals[i].typeName)
		}
	}

	return diffMap, nil
}
