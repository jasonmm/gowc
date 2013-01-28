package libgowc

import (
	"bufio"
	"io"
	"os"
)

type metrics struct {
	nline, nword, nchar int
}

func (lhs *metrics) Add(rhs *metrics) {
	lhs.nline += rhs.nline
	lhs.nword += rhs.nword
	lhs.nchar += rhs.nchar
}

func ProcessFiles(paths []string) {
	var total metrics

	for _, p := range paths {
		m, _ := ProcessSingleFile(p)
		total.Add(&m)
	}
}

func ProcessSingleFile(path string) (m_ret metrics, e error) {
	var m metrics

	rd, err := os.Open(path)
	if err != nil {
		return m, err
	}
	defer rd.Close()
	m = processReader(rd, path)
	return m, nil
}

func processReader(rd io.Reader, name string) metrics {
	m, _ := countAll(rd)
	return m
}

func countAll(rd io.Reader) (m_ret metrics, e error) {
	brd := bufio.NewReader(rd)

	var m metrics

	for {
		s, err := brd.ReadString('\n')
		if err != nil && err != io.EOF {
			return m, e
		}
		if len(s) == 0 {
			break
		}

		m.nline++
		m.nchar += len(s)
		m.nword += countWords(s)

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
