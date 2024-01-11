package secret

// Handshake converts number [0..31] to sequence of actions for secret handshake.
func Handshake(code uint) (out []string) {
	actions := []string{"wink", "double blink", "close your eyes", "jump"}

	for i := uint(0); i < 4; i++ {
		if code&(1<<i) != 0 {
			out = append(out, actions[i])
		}
	}
	if code&16 != 0 {
		reverseSlice(out)
	}
	return
}

func reverseSlice(slice []string) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
