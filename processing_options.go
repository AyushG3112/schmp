package schmp

import "io"

//ProcessingOptions is the input provided to the `Compare` function to control how it processes input.
// Mode defines the format your data is in. `Sources` are your data sources.
type ProcessingOptions struct {
	Mode    string
	Sources []io.Reader
}
