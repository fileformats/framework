package writer

import (
	"sync"
	"io"
)

type countWriter struct {
	sync.Mutex
	writer io.Writer
	record bool
	count int
}

func (w *countWriter) Write(p []byte) (n int, err error) {
	n, err = w.writer.Write(p)
	if w.record {
		w.count += n
	} else {
		w.count = n
	}
	return
}