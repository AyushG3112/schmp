package schmp

import "sort"

// compare flattens every document into a fileSchema and reports each dotted
// path whose type is not identical across all of them. The result maps a path
// to the type each document carries there, index-matched to docs; a document
// that lacks the path contributes the empty string.
//
// When a path differs, its descendants are suppressed: a single "this whole
// subtree changed type" entry is more useful than one diff per leaf beneath it.
func compare(docs []map[string]interface{}) (map[string][]string, error) {
	schemas := make([]fileSchema, len(docs))
	for i, doc := range docs {
		schemas[i] = fileSchema{}
		if err := buildFileSchema(doc, schemas[i]); err != nil {
			return nil, err
		}
	}

	diff := map[string][]string{}
	for _, path := range sortedPaths(schemas) {
		types := make([]string, len(schemas))
		differs := false
		for i, schema := range schemas {
			types[i] = canonicalType(schema[path])
			if types[i] != types[0] {
				differs = true
			}
		}
		if !differs || hasDifferingAncestor(schemas, path, diff) {
			continue
		}
		diff[path] = types
	}
	return diff, nil
}

// sortedPaths returns the union of every path across all schemas, sorted. A
// parent path is always a prefix of its descendants, so this ordering visits
// every parent before its children — which is what lets compare suppress the
// descendants of an already-recorded diff.
func sortedPaths(schemas []fileSchema) []string {
	seen := map[string]struct{}{}
	for _, schema := range schemas {
		for path := range schema {
			seen[path] = struct{}{}
		}
	}
	paths := make([]string, 0, len(seen))
	for path := range seen {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	return paths
}

// hasDifferingAncestor reports whether any container enclosing path has already
// been recorded as a diff, in which case this path's difference is a
// consequence of its parent's and should not be reported separately.
func hasDifferingAncestor(schemas []fileSchema, path string, diff map[string][]string) bool {
	for _, schema := range schemas {
		info := schema[path]
		if info == nil {
			continue
		}
		for _, ancestor := range info.parents {
			if _, ok := diff[ancestor]; ok {
				return true
			}
		}
	}
	return false
}
