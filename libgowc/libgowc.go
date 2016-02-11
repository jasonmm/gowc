// Package libgowc is a wc(1)-clone library for Go.
package libgowc

import (
	"bufio"
	"io"
	"os"
)

// Metrics holds the number of lines, words, and characters found in each file(s).
type Metrics struct {
	Lines, Words, Chars int
}

// Add adds the passed in metrics to the existing metrics already stored by
// the object.
func (lhs *Metrics) Add(rhs *Metrics) {
	lhs.Lines += rhs.Lines
	lhs.Words += rhs.Words
	lhs.Chars += rhs.Chars
}

// ProcessFiles processes an array of file names returning the total counts
// for all the files combined.  Calls ProcessSingleFile() on each filename.
func ProcessFiles(paths []string) Metrics {
	var total Metrics

	for _, p := range paths {
		m, _ := ProcessSingleFile(p)
		total.Add(&m)
	}

	return total
}

// ProcessSingleFile processes a single file given as a filename and returns
// the number of lines, words, and characters in that file.
func ProcessSingleFile(path string) (Metrics, error) {
	var m Metrics

	rd, err := os.Open(path)
	if err != nil {
		return m, err
	}
	defer rd.Close()
	m = processReader(rd)

	return m, nil
}

// processReader process the given Reader.
func processReader(rd io.Reader) Metrics {
	m, _ := countAll(rd)

	return m
}

// countAll counts the number of lines, characters, and words from the
// given Reader.
func countAll(rd io.Reader) (Metrics, error) {
	brd := bufio.NewReader(rd)

	var m Metrics

	for {
		s, err := brd.ReadString('\n')
		if err != nil && err != io.EOF {
			return m, err
		}
		if len(s) == 0 {
			break
		}

		m.Lines++
		m.Chars += len(s)
		m.Words += countWords(s)
	}

	return m, nil
}

// countWords counts the number of words in the given string.  It is used by
// the countAll() function.
func countWords(s string) int {
	wasSpace := true
	n := 0
	for _, ch := range s {
		if ch != ' ' && ch != '\t' && ch != '\n' {
			if wasSpace {
				n++
			}
			wasSpace = false
		} else {
			wasSpace = true
		}
	}

	return n
}
