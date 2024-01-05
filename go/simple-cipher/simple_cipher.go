package cipher

import "strings"

type vigenere []byte

func NewVigenere(key string) Cipher {
	allA := true
	for _, b := range []byte(key) {
		if b < 'a' || b > 'z' {
			return nil
		}
		allA = allA && b == 'a'
	}
	if allA {
		return nil
	}
	return vigenere(key)
}

func NewShift(distance int) Cipher {
	if distance < 0 {
		distance += 26
	}
	return NewVigenere(string(byte(distance + 'a')))
}

func NewCaesar() Cipher                    { return vigenere{'d'} }
func (v vigenere) Encode(in string) string { return v.translate(in, 1) }
func (v vigenere) Decode(in string) string { return v.translate(in, -1) }

func (v vigenere) translate(in string, dir int8) string {
	var out strings.Builder
	out.Grow(len(in))
	for _, b := range []byte(in) {
		n := int8(b | ' ' - 'a') // lowercase - 'a'
		if n >= 0 && n <= 26 {
			out.WriteByte(byte((n+dir*int8(v[out.Len()%len(v)]-'a')+26)%26 + 'a'))
		}
	}
	return out.String()
}
