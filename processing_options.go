package schmp

import "io"

type ProcessingOptions struct {
	Mode    string
	Sources []io.Reader
}
