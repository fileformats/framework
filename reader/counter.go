package reader

import (
	"io"
	"sync"
)

type countReader struct {
	sync.Mutex
	reader io.Reader
	record bool
	count int
}

func (r *countReader) Read(p []byte) (n int, err error) {
	n, err = r.reader.Read(p)
	if r.record {
		r.count += n
	} else {
		r.count = n
	}
	return
}