/*
 * Copyright (c) 2010 Nicolas Thery (nthery@gmail.com)
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 */

/*
   wc(1) clone
*/
package main


import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)


var (
	wordFlag = flag.Bool("w", false, "count words")
	charFlag = flag.Bool("c", false, "count characters")
	lineFlag = flag.Bool("l", false, "count lines")
)


type metrics struct {
	nline, nword, nchar int
}


func (m metrics) String() string {
	s := ""
	if *lineFlag {
		s += fmt.Sprintf("\t%d", m.nline)
	}
	if *wordFlag {
		s += fmt.Sprintf("\t%d", m.nword)
	}
	if *charFlag {
		s += fmt.Sprintf("\t%d", m.nchar)
	}
	return s
}


func (lhs *metrics) Add(rhs *metrics) {
	lhs.nline += rhs.nline
	lhs.nword += rhs.nword
	lhs.nchar += rhs.nchar
}


func main() {
	parseFlags()
	if flag.NArg() == 0 {
		processReader(os.Stdin, "")
	} else {
		processFiles(flag.Args())
	}
}


func parseFlags() {
	flag.Parse()
	if !*wordFlag && !*lineFlag && !*charFlag {
		*wordFlag = true
		*lineFlag = true
		*charFlag = true
	}
}


func processFiles(paths []string) {
	var total metrics

	for _, p := range paths {
		m := processSingleFile(p)
		total.Add(&m)
	}

	if len(flag.Args()) > 1 {
		fmt.Printf("%v\ttotal\n", total)
	}
}


func processSingleFile(path string) metrics {
	rd, err := os.Open(path, os.O_RDONLY, 0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open error: %s: %s\n", path, err.String())
		os.Exit(1)
	}
	defer rd.Close()
	m := processReader(rd, path)
	return m
}


func processReader(rd io.Reader, name string) metrics {
	m := countAll(rd)
	fmt.Printf("%v\t%s\n", m, name)
	return m
}


func countAll(rd io.Reader) metrics {
	brd := bufio.NewReader(rd)

	var m metrics

	for {
		s, err := brd.ReadString('\n')
		if err != nil && err != os.EOF {
			fmt.Fprintf(os.Stderr, "read error: %s\n", err.String())
			os.Exit(1)
		}
		if len(s) == 0 {
			break
		}

		m.nline++
		m.nchar += len(s)
		m.nword += countWords(s)

	}

	return m
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
