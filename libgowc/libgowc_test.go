package libgowc

import "testing"

func TestProcessSingleFile(t *testing.T) {
	var m metrics
	var e error

	if m, e = ProcessSingleFile("test-file.txt"); e != nil {
		t.Errorf(e.Error())
	}
	if m.nline != 3 {
		t.Errorf("Line count did not equal 3. Line count was %d instead.", m.nline)
	}
	if m.nword != 38 {
		t.Errorf("Word count did not equal 38.  Word count was %d instead.", m.nword)
	}
	if m.nchar != 246 {
		t.Errorf("Char count did not equal 246.  Char count was %d instead.", m.nchar)
	}
}

