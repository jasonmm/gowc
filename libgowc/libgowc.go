// A wc(1)-clone library for Go.
package libgowc

import (
	"bufio"
	"io"
	"os"
)

// Structure to hold the number of lines, words, and characters found in 
// each file(s).
type Metrics struct {
	Lines, Words, Chars int
}

// Adds the passed in metrics to the lhs metrics.
func (lhs *Metrics) Add(rhs *Metrics) {
	lhs.Lines += rhs.Lines
	lhs.Words += rhs.Words
	lhs.Chars += rhs.Chars
}

// Processes a array of file names returning the total counts for all the
// files combined.  Calls ProcessSingleFile() on each filename.
func ProcessFiles(paths []string) Metrics {
	var total Metrics

	for _, p := range paths {
		m, _ := ProcessSingleFile(p)
		total.Add(&m)
	}

	return total
}

// Processes a single file given as a filename and returns the number of 
// lines, words, and characters in that file.
func ProcessSingleFile(path string) (m_ret Metrics, e error) {
	var m Metrics

	rd, err := os.Open(path)
	if err != nil {
		return m, err
	}
	defer rd.Close()
	m = processReader(rd, path)
	return m, nil
}

func processReader(rd io.Reader, name string) Metrics {
	m, _ := countAll(rd)
	return m
}

func countAll(rd io.Reader) (m_ret Metrics, e error) {
	brd := bufio.NewReader(rd)

	var m Metrics

	for {
		s, err := brd.ReadString('\n')
		if err != nil && err != io.EOF {
			return m, e
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

func countWords(s string) int {
	wasspace := true
	n := 0
	for _, ch := range s {
		if ch != ' ' && ch != '\t' && ch != '\n' {
			if wasspace {
				n++
			}
			wasspace = false
		} else {
			wasspace = true
		}
	}

	return n
}
