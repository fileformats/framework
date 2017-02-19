package writer

import (
	"bytes"
	"encoding/binary"
	"io"
)

type Writer struct {
	w *countWriter
	order    binary.ByteOrder
	position int
	err      error
}

// Returns a new reader with BigEndian byte order
func New(w io.Writer) *Writer {
	return &Writer{
		w: &countWriter{
			writer: w,
		},
		order: binary.BigEndian,
	}
}

// Set the byte order
func (w *Writer) SetByteOrder(endianes binary.ByteOrder) {
	w.order = endianes
}

// Returns the byte order
func (w *Writer) GetByteOrder() binary.ByteOrder {
	return w.order
}

// Changes the reader
func (w *Writer) SetWriter(writer io.Writer) {
	w.w = &countWriter{
		writer: writer,
	}
	w.position = 0
	w.err = nil
}

// Returns the number of bytes read from the start of the stream
func (w *Writer) BytesRead() int {
	return w.position
}

func (w *Writer) GetError() error {
	return w.err
}

// Ok returns true if the reader encountered no error during the last read
func (w *Writer) Ok() bool {
	return w.err == nil
}

// Writes a byte to the bit stream
func (w *Writer) Byte(v byte) {
	w.err = binary.Write(w.w, w.order, &v)
	w.position += w.w.count
	return
}

// Writes a uint8 to the bit stream
func (w *Writer) UInt8(v uint8) {
	w.err = binary.Write(w.w, w.order, &v)
	w.position += w.w.count
	return
}

// Writes a int8 to the bit stream
func (w *Writer) Int8(v int8) {
	w.err = binary.Write(w.w, w.order, &v)
	w.position += w.w.count
	return
}

// Writes a uint16 to the bit stream
func (w *Writer) UInt16(v uint16) {
	w.err = binary.Write(w.w, w.order, &v)
	w.position += w.w.count
	return
}

// Writes a int16 to the bit stream
func (w *Writer) Int16(v int16) {
	w.err = binary.Write(w.w, w.order, &v)
	w.position += w.w.count
	return
}

// Writes a uint32 to the bit stream
func (w *Writer) UInt32(v uint32)  {
	w.err = binary.Write(w.w, w.order, &v)
	w.position += w.w.count
	return
}

// Writes a int32 to the bit stream
func (w *Writer) Int32(v int32)  {
	w.err = binary.Write(w.w, w.order, &v)
	w.position += w.w.count
	return
}

// Writes a uint64 to the bit stream
func (w *Writer) UInt64(v uint64)  {
	w.err = binary.Write(w.w, w.order, &v)
	w.position += w.w.count
	return
}

// Writes a int64 to the bit stream
func (w *Writer) Int64(v int64)  {
	w.err = binary.Write(w.w, w.order, &v)
	w.position += w.w.count
	return
}

// Writes a float32 to the bit stream
func (w *Writer) Float32(v float32)  {
	w.err = binary.Write(w.w, w.order, &v)
	w.position += w.w.count
	return
}

// Writes a float64 to the bit stream
func (w *Writer) Float64(v float64)  {
	w.err = binary.Write(w.w, w.order, &v)
	w.position += w.w.count
	return
}

// Writes a rune to the bit stream
func (w *Writer) Rune(v rune)  {
	w.err = binary.Write(w.w, w.order, &v)
	w.position += w.w.count
	return
}

// Writes a null terminated string to the bit stream
func (w *Writer) WriteString(val string)  {
	var buf *bytes.Buffer
	for _, ch := range val {
		buf.WriteRune(ch)
	}
	buf.WriteByte(0)
}

// Packs to a structure of fixed-sized values or a slice of fixed-size values
func (w *Writer) Pack(ptr interface{}) {
	w.w.count = 0
	w.w.record = true
	w.err = binary.Write(w.w, w.order, ptr)
	w.position += w.w.count
	w.w.record = false
}
