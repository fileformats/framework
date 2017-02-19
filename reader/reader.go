package reader

import (
	"bytes"
	"encoding/binary"
	_ "net/http/pprof"
	"io"
)

type Reader struct {
	r        *countReader
	order    binary.ByteOrder
	position int
	err      error
}

// Returns a new reader with BigEndian byte order
func New(r io.Reader) *Reader {
	return &Reader{
		r: &countReader{
			reader: r,
		},
		order: binary.BigEndian,
	}
}

// Set the byte order
func (r *Reader) SetByteOrder(endianes binary.ByteOrder) {
	r.order = endianes
}

// Returns the byte order
func (r *Reader) GetByteOrder() binary.ByteOrder {
	return r.order
}

// Changes the reader
func (r *Reader) SetReader(reader io.Reader) {
	r.r = &countReader{
		reader: reader,
	}
	r.position = 0
	r.err = nil
}

// Returns the number of bytes read from the start of the stream
func (r *Reader) BytesRead() int {
	return r.position
}

func (r *Reader) GetError() error {
	return r.err
}

// Ok returns true if the reader encountered no error during the last read
func (r *Reader) Ok() bool {
	return r.err == nil
}

// Skip reads and discards count byes from the stream
func (r *Reader) Skip(count int) {
	tmp := make([]byte, count)
	n, err := r.r.Read(tmp)
	r.err = err
	r.position += n
}

// Reads a byte from the bit stream
func (r *Reader) Byte() (v byte) {
	r.err = binary.Read(r.r, r.order, &v)
	r.position += r.r.count
	return
}

// Reads a uint8 from the bit stream
func (r *Reader) UInt8() (v uint8) {
	r.err = binary.Read(r.r, r.order, &v)
	r.position += r.r.count
	return
}

// Reads a int8 from the bit stream
func (r *Reader) Int8() (v int8) {
	r.err = binary.Read(r.r, r.order, &v)
	r.position += r.r.count
	return
}

// Reads a uint16 from the bit stream
func (r *Reader) UInt16() (v uint16) {
	r.err = binary.Read(r.r, r.order, &v)
	r.position += r.r.count
	return
}

// Reads a int16 from the bit stream
func (r *Reader) Int16() (v int16) {
	r.err = binary.Read(r.r, r.order, &v)
	r.position += r.r.count
	return
}

// Reads a uint32 from the bit stream
func (r *Reader) UInt32() (v uint32) {
	r.err = binary.Read(r.r, r.order, &v)
	r.position += r.r.count
	return
}

// Reads a int32 from the bit stream
func (r *Reader) Int32() (v int32) {
	r.err = binary.Read(r.r, r.order, &v)
	r.position += r.r.count
	return
}

// Reads a uint64 from the bit stream
func (r *Reader) UInt64() (v uint64) {
	r.err = binary.Read(r.r, r.order, &v)
	r.position += r.r.count
	return
}

// Reads a int64 from the bit stream
func (r *Reader) Int64() (v int64) {
	r.err = binary.Read(r.r, r.order, &v)
	r.position += r.r.count
	return
}

// Reads a float32 from the bit stream
func (r *Reader) Float32() (v float32) {
	r.err = binary.Read(r.r, r.order, &v)
	r.position += r.r.count
	return
}

// Reads a float64 from the bit stream
func (r *Reader) Float64() (v float64) {
	r.err = binary.Read(r.r, r.order, &v)
	r.position += r.r.count
	return
}

// Reads a rune from the bit stream
func (r *Reader) Rune() (v rune) {
	r.err = binary.Read(r.r, r.order, &v)
	r.position += r.r.count
	return
}

// Reads a null terminated string from the bit stream
func (r *Reader) ReadString() string {
	var buf *bytes.Buffer
	for {
		b := r.Rune()
		if r.err != nil || b == 0 {
			return buf.String()
		}
		buf.WriteRune(b)
	}
}

// Unpacks to a structure of fixed-sized values or a slice of fixed-size values
func (r *Reader) Unpack(ptr interface{}) {
	r.r.count = 0
	r.r.record = true
	r.err = binary.Read(r.r, r.order, ptr)
	r.position += r.r.count
	r.r.record = false
}
