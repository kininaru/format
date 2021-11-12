package format

import (
	"testing"
)

func TestSetDict(t *testing.T) {
	var newDict = []rune{'z', 'A', '0'}
	SetDict(newDict)
	if len(newDict) != len(dict) {
		t.Errorf("Failed. Length not match.")
		return
	}

	for i, r := range dict {
		if r != newDict[i] {
			t.Errorf("Failed. Dict not match.")
		}
	}
}

func TestAppendDict(t *testing.T) {
	SetDict(nil)
	AppendDict([]rune{'A'})
	if len(dict) != 1 {
		t.Errorf("Appending failed.")
		return
	}
	AppendDict([]rune{'b'})
	if len(dict) != 2 {
		t.Errorf("Appending failed.")
		return
	}
	AppendDict([]rune{'1'})
	if len(dict) != 3 {
		t.Errorf("Appending failed.")
		return
	}

	for i, r := range []rune{'A', 'b', '1'} {
		if dict[i] != r {
			t.Errorf("Dict data error after appending.")
			return
		}
	}
}
