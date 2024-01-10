package circular

import "errors"

// Define the Buffer type here.
type Buffer struct {
	data []byte
	size int
}

func NewBuffer(size int) *Buffer {
	return &Buffer{
		size: size,
		data: make([]byte, 0, size),
	}
}

func (b *Buffer) ReadByte() (byte, error) {
	if len(b.data) == 0 {
		return 0, errors.New("")
	}

	result := b.data[0]
	b.data = b.data[1:]

	return result, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if len(b.data) >= b.size {
		return errors.New("")
	}

	b.data = append(b.data, c)
	return nil
}

func (b *Buffer) Overwrite(c byte) {
	if len(b.data) >= b.size {
		b.data = b.data[1:]
	}

	b.data = append(b.data, c)
}

func (b *Buffer) Reset() {
	b.data = make([]byte, 0, b.size)
}
