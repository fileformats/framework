package writer

import (
	"fmt"
	"strconv"
)

type BitBuffer struct {
	position  int
	bitBuffer int64
	nBits     int
	Boundary  int
	Buffer    []byte
}

func NewBitBuffer(boundary int) *BitBuffer {
	if boundary % 8 != 0 {
		boundary += 8 - (boundary % 8)
	}
	return &BitBuffer{
		Buffer: make([]byte, boundary / 8, boundary / 8),
		Boundary: boundary,
		nBits: boundary,
	}
}

func (b *BitBuffer) Byte(val byte, bitCount int) {
	b.Int64(int64(val), bitCount)
}

func (b *BitBuffer) Int16(val int16, bitCount int) {
	b.Int64(int64(val), bitCount)
}

func (b *BitBuffer) Int(val int, bitCount int) {
	b.Int64(int64(val), bitCount)
}

func (b *BitBuffer) Int32(val int32, bitCount int) {
	b.Int64(int64(val), bitCount)
}

func (b *BitBuffer) Int64(val int64, bitCount int) {
	for i := bitCount; i > 0; i--{
		if b.nBits == 0 {
			// add Boundary bits to the buffer
			b.Buffer = append(b.Buffer, make([]byte, b.Boundary / 8, b.Boundary / 8)...)
			b.nBits += b.Boundary
		}
		// get bit position to write at
		bit := uint(b.position % 8)
		// get bit to write
		target := (val >> uint(i - 1)) & 0x01
		target <<= 7 - bit
		b.Buffer[b.position / 8] |= byte(target)
		b.position++
		b.nBits--
	}
}

func (b *BitBuffer) String() string {
	tmp := ""
	for i := 0; i < len(b.Buffer); i++ {
		tmp += fmt.Sprintf("%08s", strconv.FormatInt(int64(b.Buffer[i]), 2))
	}
	return tmp
}
