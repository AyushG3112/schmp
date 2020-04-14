package schmp

//ComparisonOutput is the output from the Compare function. `Diff` holds the actual result of the comparison.
// The key in the `Diff` map represents the key at which a different type was found across all the `Sources`,
// and the values represent the types in the `Sources`. The values are index matched to the `Sources` provided in `ProcessingOptions`
type ComparisonOutput struct {
	Diff map[string][]string
}
