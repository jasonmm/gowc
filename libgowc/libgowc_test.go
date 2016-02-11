package libgowc

import "testing"

func TestProcessSingleFile(t *testing.T) {
	var m Metrics
	var e error

	if m, e = ProcessSingleFile("test-file.txt"); e != nil {
		t.Errorf(e.Error())
	}
	if m.Lines != 3 {
		t.Errorf("Line count did not equal 3. Line count was %d instead.", m.Lines)
	}
	if m.Words != 38 {
		t.Errorf("Word count did not equal 38.  Word count was %d instead.", m.Words)
	}
	if m.Chars != 246 {
		t.Errorf("Char count did not equal 246.  Char count was %d instead.", m.Chars)
	}
}

func TestProcessFiles(t *testing.T) {
	paths := []string{"./test-file.txt"}
	m := ProcessFiles(paths)
	if m.Lines != 3 {
		t.Errorf("Line count did not equal 3. Line count was %d instead.", m.Lines)
	}
	if m.Words != 38 {
		t.Errorf("Word count did not equal 38.  Word count was %d instead.", m.Words)
	}
	if m.Chars != 246 {
		t.Errorf("Char count did not equal 246.  Char count was %d instead.", m.Chars)
	}
}
