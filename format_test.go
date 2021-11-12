package format

import (
	"testing"
)

func TestFormat(t *testing.T) {
	data := []struct {
		src  int64
		dest string
	}{
		{5577006791947779410, "16ddfwutxgjte"},
		{8674665223082153551, "1twm86kqk9z67"},
		{6129484611666145821, "1akhcyde76ot9"},
		{4037200794235010051, "uo7w8oxnzncz"},
		{3916589616287113937, "tr8b5apa57f5"},
		{6334824724549167320, "1c4n7y6zh4t08"},
		{605394647632969758, "4lkyhppi2lni"},
		{1443635317331776148, "ayum0d8s0dxw"},
		{894385949183117216, "6smhcwuuqly8"},
		{2775422040480279449, "l33x03nwf9uh"},
	}

	for _, p := range data {
		if str, _ := IntToString(p.src, 36); str != p.dest {
			t.Errorf("Wrong output. src: %d, dest: %s, expected: %s", p.src, str, p.dest)
		}
	}
}
