package letter

type FreqMap map[rune]int

func Frequency(text string) FreqMap {
	freq := make(FreqMap, 512)
	for _, r := range text {
		freq[r]++
	}
	return freq
}

func ConcurrentFrequency(texts []string) FreqMap {

	// Make results channel
	maps := make(chan FreqMap)

	// For each text, start a goroutine to count the frequency
	for _, text := range texts {
		go func(text string) {
			maps <- Frequency(text)
		}(text)
	}

	// Collect the first mresult ap
	mergedMap := <-maps

	// Merge subsequent maps into the first
	for range texts[1:] {
		for char, count := range <-maps {
			mergedMap[char] += count
		}
	}
	return mergedMap
}
