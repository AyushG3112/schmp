package schmp

// ComparisonOutput is the output from the Compare function. `Diff` holds the actual result of the comparison.
// The key in the `Diff` map is the dotted path at which a differing type was found across all the `Sources`
// (for example "quiz.maths.q2.question"; array elements use a "[]" segment, as in "users[].id"), and the
// values are the types found in the `Sources`, index matched to the `Sources` provided in `ProcessingOptions`.
// An empty string means the key was absent from that source, distinct from "<nil>" which means a present null.
type ComparisonOutput struct {
	Diff map[string][]string
}
