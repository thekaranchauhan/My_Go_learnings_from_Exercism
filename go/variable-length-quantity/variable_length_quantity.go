package variablelengthquantity

import "errors"

// DecodeVarint decodes a variable length integer
func DecodeVarint(bs []byte) (is []uint32, err error) {
	var i uint32
	for _, b := range bs {
		i = (i << 7) | uint32(b&0x7F)
		if b&0x80 == 0 {
			is = append(is, i)
			i = 0
		}
	}

	if len(is) == 0 {
		return nil, errors.New("incomplete sequence")
	}

	return is, nil
}

// EncodeVarint encodes a variable length integer
func EncodeVarint(is []uint32) (bs []byte) {
	for _, i := range is {
		var enc []byte
		for first := uint32(0); first == 0 || i != 0; first = 0x80 {
			enc = append([]byte{byte(i&0x7F | first)}, enc...)
			i >>= 7
		}
		bs = append(bs, enc...)
	}

	return bs
}
