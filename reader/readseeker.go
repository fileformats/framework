package reader

import (
	"errors"
	"io"
	"encoding/binary"
	"bytes"
)

var EOF = errors.New("EOF")

type ReadSeeker struct {
	r      *Reader
	s      io.ReadSeeker
	start  int
	len    int
	capped bool
}

func NewReadSeeker(reader io.ReadSeeker) *ReadSeeker {
	return &ReadSeeker{
		r:      New(reader),
		s:      reader,
		len:    -1,
		capped: false,
		start: 0,
	}
}

// Seek moves the position in the file relative to whence
func (r *ReadSeeker) Seek(offset int64, whence int) (int64, error) {
	if r.capped {
		if r.start+whence+int(offset) > r.len {
			return 0, errors.New("Cannot seek beyond EOF")
		}
	}
	pos, err := r.s.Seek(offset, r.start+whence)
	if err != nil {
		r.r.err = err
	} else {
		r.r.position = int(pos)
	}
	return pos, err
}

// Slice returns a new FileReader that is capped on the given parameters
func (r *ReadSeeker) Slice(start uint64, len uint64) *ReadSeeker {
	r.Seek(0, int(start))

	tmp := &ReadSeeker{
		r: New(r.s),
		s: r.s,
		start: int(start),
		len: int(len),
		capped: true,
	}
	return tmp
}

// Set the byte order
func (r *ReadSeeker) SetByteOrder(endianes binary.ByteOrder) {
	r.r.order = endianes
}

// Returns the byte order
func (r *ReadSeeker) GetByteOrder() binary.ByteOrder {
	return r.r.order
}

// Changes the reader
func (r *ReadSeeker) SetReader(reader io.ReadSeeker) {
	r.s = reader
	r.r.SetReader(reader)
	r.capped = false
	r.len = 0
	r.start = 0
}

// Returns the number of bytes read from the start of the stream
func (r *ReadSeeker) BytesRead() int {
	return r.r.position
}

func (r *ReadSeeker) GetError() error {
	return r.r.err
}

// Ok returns true if the reader encountered no error during the last read
func (r *ReadSeeker) Ok() bool {
	return r.r.err == nil
}

func (r *ReadSeeker) BytesAvailable() int {
	if r.capped {
		return r.len - r.r.position
	}
	return -1
}

// Skip reads and discards count byes from the stream
func (r *ReadSeeker) Skip(count int) {
	if r.capped && r.BytesAvailable() < count {
		count = r.BytesAvailable()
	}
	r.r.Skip(count)
}

// Reads a byte from the bit stream
func (r *ReadSeeker) Byte() (v byte) {
	if r.capped && r.BytesAvailable() == 0 {
		r.r.err = EOF
		return
	}
	return r.r.Byte()
}

// Reads a uint8 from the bit stream
func (r *ReadSeeker) UInt8() (v uint8) {
	if r.capped && r.BytesAvailable() == 0 {
		r.r.err = EOF
		return
	}
	return r.r.UInt8()
}

// Reads a int8 from the bit stream
func (r *ReadSeeker) Int8() (v int8) {
	if r.capped && r.BytesAvailable() == 0 {
		r.r.err = EOF
		return
	}
	return r.r.Int8()
}

// Reads a uint16 from the bit stream
func (r *ReadSeeker) UInt16() (v uint16) {
	if r.capped && r.BytesAvailable() < 2 {
		r.r.err = EOF
		return
	}
	return r.r.UInt16()
}

// Reads a int16 from the bit stream
func (r *ReadSeeker) Int16() (v int16) {
	if r.capped && r.BytesAvailable() < 2 {
		r.r.err = EOF
		return
	}
	return r.r.Int16()
}

// Reads a uint32 from the bit stream
func (r *ReadSeeker) UInt32() (v uint32) {
	if r.capped && r.BytesAvailable() < 4 {
		r.r.err = EOF
		return
	}
	return r.r.UInt32()
}

// Reads a int32 from the bit stream
func (r *ReadSeeker) Int32() (v int32) {
	if r.capped && r.BytesAvailable() < 4 {
		r.r.err = EOF
		return
	}
	return r.r.Int32()
}

// Reads a uint64 from the bit stream
func (r *ReadSeeker) UInt64() (v uint64) {
	if r.capped && r.BytesAvailable() < 8 {
		r.r.err = EOF
		return
	}
	return r.r.UInt64()
}

// Reads a int64 from the bit stream
func (r *ReadSeeker) Int64() (v int64) {
	if r.capped && r.BytesAvailable() < 8 {
		r.r.err = EOF
		return
	}
	return r.r.Int64()
}

// Reads a float32 from the bit stream
func (r *ReadSeeker) Float32() (v float32) {
	if r.capped && r.BytesAvailable() < 4 {
		r.r.err = EOF
		return
	}
	return r.r.Float32()
}

// Reads a float64 from the bit stream
func (r *ReadSeeker) Float64() (v float64) {
	if r.capped && r.BytesAvailable() < 8 {
		r.r.err = EOF
		return
	}
	return r.r.Float64()
}

// Reads a rune from the bit stream
func (r *ReadSeeker) Rune() (v rune) {
	if r.capped && r.BytesAvailable() < 1 {
		r.r.err = EOF
		return
	}
	return r.r.Rune()
}

// Reads a null terminated string from the bit stream
func (r *ReadSeeker) ReadString() string {
	var buf *bytes.Buffer
	for {
		b := r.Rune()
		if r.r.err != nil || b == 0 {
			return buf.String()
		}
		buf.WriteRune(b)
	}
}

// Unpacks to a structure of fixed-sized values or a slice of fixed-size values
func (r *ReadSeeker) Unpack(target interface{}) {
	required := uint64(binary.Size(target))
	if r.capped && r.BytesAvailable() < int(required) {
		r.r.err = EOF
		return
	}
	r.r.Unpack(target)
}
