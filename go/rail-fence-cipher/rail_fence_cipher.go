package railfence

func Encode(s string, numRails int) string {
	return transform(s, numRails, func(in string, out []byte, i, j int) { out[i] = in[j] })
}

func Decode(s string, numRails int) string {
	return transform(s, numRails, func(in string, out []byte, j, i int) { out[i] = in[j] })
}

func transform(in string, numRails int, set func(string, []byte, int, int)) string {
	out := make([]byte, len(in))
	for i, rail := 0, 0; rail < numRails; rail++ {
		delta := 2 * rail
		for j := rail; j < len(in); j, i = j+delta, i+1 {
			set(in, out, i, j)
			if delta != 2*(numRails-1) {
				delta = 2*(numRails-1) - delta
			}
		}
	}
	return string(out)
}
