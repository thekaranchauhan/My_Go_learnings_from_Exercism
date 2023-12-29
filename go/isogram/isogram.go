package isogram

func IsIsogram(word string) bool {
	var mask uint32

	for i := 0; i < len(word); i++ {
		char := byte(word[i])

		// Deal only with lowercase letters.
		if char <= 'Z' {
			char += 1 << 5
		}

		// Make sure we only add lowercase letters.
		if 'a' <= char && char <= 'z' {
			// Check if the character is recorded.
			if (mask | (1 << (char - 97))) == mask {
				return false
			}

			// Record the character in the bit mask.
			mask |= 1 << (char - 97)
		}
	}

	return true
}
