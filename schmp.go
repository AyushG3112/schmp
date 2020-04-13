package schmp

func Compare(options ProcessingOptions) (map[string][]string, error) {
	parse, err := getParser(options.Mode)
	if err != nil {
		return nil, err
	}
	sm, err := parse(options)
	if err != nil {
		return nil, err
	}
	return compare(sm, options, "", make(map[string][]string))
}
