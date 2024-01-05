package atbash

func Atbash(st string) string {

	r := []rune{}
	for _, v := range st {
		var char rune
		if 'A' <= v && v <= 'Z' {
			char = 'z' - v + 'A'
		} else if 'a' <= v && v <= 'z' {
			char = 'z' - v + 'a'
		} else if '0' <= v && v <= '9' {
			char = v
		} else {
			continue
		}
		if len(r)%6 == 5 {
			r = append(r, ' ')
		}
		r = append(r, char)

	}
	return string(r)

}
