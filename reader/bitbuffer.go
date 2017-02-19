package reader

type BitBuffer struct {
	Position  int
	bitBuffer int
	nBits     int

	buffer ByteReader
}

type ByteReader interface {
	Int8() int8
	UInt8() uint8
	Byte() byte
}

func NewBitBuffer(data ByteReader) *BitBuffer {
	return &BitBuffer{
		Position: 0,
		buffer:   data,
	}
}

func (b *BitBuffer) Byte(bitCount int) byte {
	return byte(b.Int64(bitCount))
}

func (b *BitBuffer) Int16(bitCount int) int16 {
	return int16(b.Int64(bitCount))
}

func (b *BitBuffer) Int(bitCount int) int {
	return int(b.Int64(bitCount))
}

func (b *BitBuffer) Int32(bitCount int) int32 {
	return int32(b.Int64(bitCount))
}

func (b *BitBuffer) Int64(bitCount int) int64 {
	if bitCount <= 0 {
		return 0
	}
	bPos := 0
	var result int64 = 0
	var length = bPos + bitCount

	for length > 0 {
		if b.nBits == 0 {
			b.bitBuffer = int(b.buffer.Int8())
			b.nBits = 8
			b.bitBuffer &= 0xFF
		}
		if bPos == 0 {
			result <<= 1
			result |= int64(b.bitBuffer) >> 7
		} else {
			bPos--
		}
		b.bitBuffer <<= 1
		b.bitBuffer &= 0xFF
		b.nBits--
		length--
		b.Position++
	}
	return result
}

