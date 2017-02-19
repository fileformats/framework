package writer

import "testing"

func TestBitWriter(t *testing.T) {
	w := NewBitBuffer(32)
	w.Int(22, 1)
	w.Int(1, 1)
	w.Int(1, 1)
	w.Int(1, 1)
	w.Int(1, 1)
	w.Int(1, 1)
	w.Int(1, 1)
	w.Int(1, 1)
	w.Int(1, 1)
	w.Int(1, 1)
	w.Int(234234, 1)
	w.Int(4325, 8)
	w.Int(2323, 7)
	w.Int(234, 13)
}
