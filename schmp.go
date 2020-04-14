package schmp

//Compare reads data from from the provided `options.Sources`, decodes them according to the provided `options.Mode`
// and recursively compares properties of the decoded results.
func Compare(options ProcessingOptions) (ComparisonOutput, error) {
	parse, err := getParser(options.Mode)
	if err != nil {
		return ComparisonOutput{}, err
	}
	sm, err := parse(options)
	if err != nil {
		return ComparisonOutput{}, err
	}
	diff, err := compare(sm, options, "", make(map[string][]string))
	return ComparisonOutput{Diff: diff}, err
}
