package schmp

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// arraySuffix is appended to a path to denote "the elements of the array at
// that path". Array elements are compared by schema rather than by position, so
// every element of an array shares this single key and the array's length is
// ignored.
const arraySuffix = "[]"

// pathInfo records everything known about a single dotted path within one
// document: the set of value types observed there, and the dotted paths of all
// the containers enclosing it (used by compare to suppress redundant child
// diffs once a parent already differs).
//
// types is a set because one path can legitimately carry more than one type:
// every element of an array is flattened onto the same "<path>[]" key, so a
// heterogeneous array such as [1, "a"] contributes both "int" and "string".
type pathInfo struct {
	types   map[string]struct{}
	parents []string
}

// fileSchema is the flattened type map for a single document, keyed by dotted
// path (for example "quiz.maths.q2.question" or "users[].id").
type fileSchema map[string]*pathInfo

// getTypeString returns a stable name for the dynamic type of v. A nil value is
// reported as "<nil>" so that a present-but-null value stays distinguishable
// from a missing key, which compare represents as the empty string.
func getTypeString(v interface{}) string {
	if v == nil {
		return "<nil>"
	}
	return reflect.TypeOf(v).String()
}

// asObjectMap normalizes the map representations produced by the JSON, YAML and
// TOML decoders into a map[string]interface{} that can be traversed. ok is
// false for any non-map value (scalars and arrays are handled by the caller). A
// map type the decoders never emit (for example a hand-built map[string]int
// passed in through the library) is surfaced as an error rather than silently
// dropped.
func asObjectMap(v interface{}) (m map[string]interface{}, ok bool, err error) {
	switch t := v.(type) {
	case map[string]interface{}:
		return t, true, nil
	case map[interface{}]interface{}:
		m = make(map[string]interface{}, len(t))
		for k, val := range t {
			m[fmt.Sprintf("%v", k)] = val
		}
		return m, true, nil
	}
	if rt := reflect.TypeOf(v); rt != nil && rt.Kind() == reflect.Map {
		return nil, false, fmt.Errorf("unsupported map type %s", rt.String())
	}
	return nil, false, nil
}

// buildFileSchema flattens an entire document into schema.
func buildFileSchema(doc map[string]interface{}, schema fileSchema) error {
	for k, v := range doc {
		if err := addValue(k, v, nil, schema); err != nil {
			return err
		}
	}
	return nil
}

// addValue records the type of v at path, then recurses into its structure:
// object keys become "path.key" and array elements all collapse onto "path[]".
// ancestors holds the dotted paths of every container enclosing path.
func addValue(path string, v interface{}, ancestors []string, schema fileSchema) error {
	recordType(schema, path, getTypeString(v), ancestors)

	obj, isObject, err := asObjectMap(v)
	if err != nil {
		return err
	}
	if isObject {
		childAncestors := appendPath(ancestors, path)
		for k, child := range obj {
			if err := addValue(path+"."+k, child, childAncestors, schema); err != nil {
				return err
			}
		}
		return nil
	}

	if arr, isArray := v.([]interface{}); isArray {
		childAncestors := appendPath(ancestors, path)
		elementPath := path + arraySuffix
		for _, element := range arr {
			if err := addValue(elementPath, element, childAncestors, schema); err != nil {
				return err
			}
		}
	}
	return nil
}

// appendPath returns ancestors with path appended, always on a freshly
// allocated backing array so sibling recursions can never alias one another.
func appendPath(ancestors []string, path string) []string {
	extended := make([]string, len(ancestors)+1)
	copy(extended, ancestors)
	extended[len(ancestors)] = path
	return extended
}

// recordType adds typeName to the set of types seen at path, creating the
// pathInfo (and capturing its ancestors) the first time the path is seen.
func recordType(schema fileSchema, path, typeName string, ancestors []string) {
	info, ok := schema[path]
	if !ok {
		info = &pathInfo{
			types:   make(map[string]struct{}),
			parents: append([]string(nil), ancestors...),
		}
		schema[path] = info
	}
	info.types[typeName] = struct{}{}
}

// canonicalType renders the set of types at a path as a single, stable string:
// the lone type for the common case, or its members sorted and joined by " | "
// for a heterogeneous array. A nil info means the path is absent from this
// document and renders as the empty string, which is how a missing key surfaces
// in the diff.
func canonicalType(info *pathInfo) string {
	if info == nil {
		return ""
	}
	names := make([]string, 0, len(info.types))
	for name := range info.types {
		names = append(names, name)
	}
	sort.Strings(names)
	return strings.Join(names, " | ")
}
