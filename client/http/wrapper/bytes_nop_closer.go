package wrapper

import (
	"bytes"
	"io"
)

func BytesNopCloser(b []byte) io.ReadSeekCloser {
	return &bytesNopCloser{bytes.NewReader(b)}
}

type bytesNopCloser struct {
	*bytes.Reader
}

func (bytesNopCloser) Close() error { return nil }
