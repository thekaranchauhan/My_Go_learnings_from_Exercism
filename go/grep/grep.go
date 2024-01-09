package grep

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Search(pattern string, flags, files []string) []string {

	// Set up regex to ignore case, if required
	flag := strings.Join(flags, "")
	if strings.Contains(flag, "-i") {
		pattern = "(?i)" + pattern
	}

	// Set up regex to match whole line, if required
	if strings.Contains(flag, "-x") {
		pattern = "^" + pattern + "$"
	}

	// Compile regex
	rx := regexp.MustCompile(pattern)

	// For each file
	results := make([]string, 0)
	for _, filename := range files {
		func() {

			// Open file and set up line reader
			file, _ := os.Open(filename)
			defer file.Close()
			scanner := bufio.NewScanner(file)

			// For each line in file
			for lineNum := 1; scanner.Scan(); lineNum++ {
				line := scanner.Text()

				// If a match (or not, if inverting)
				if rx.MatchString(line) != strings.Contains(flag, "-v") {

					// If outputting file names only, do so, then skip to next file
					if strings.Contains(flag, "-l") {
						results = append(results, filename)
						return
					}

					// Print lne numbers, if required
					if strings.Contains(flag, "-n") {
						line = fmt.Sprintf("%d:%s", lineNum, line)
					}

					// Add filename if multiple files
					if len(files) > 1 {
						line = filename + ":" + line
					}

					// Append to result
					results = append(results, line)
				}
			}
		}()
	}

	// Return result
	return results
}
