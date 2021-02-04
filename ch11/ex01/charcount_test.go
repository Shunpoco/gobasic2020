package charcount

import (
	"bufio"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestCharCount(t *testing.T) {
	var tests = []struct {
		input *strings.Reader
		wantC map[rune]int
		wantL [utf8.UTFMax + 1]int
		wantI int
	}{
		{
			strings.NewReader("hoge"), map[rune]int{
				'h': 1, 'o': 1, 'g': 1, 'e': 1,
			},
			[utf8.UTFMax + 1]int{0, 4, 0, 0, 0},
			0,
		},
		{
			strings.NewReader("hoge fuga, hoge"), map[rune]int{
				'h': 2, 'o': 2, 'g': 3, 'e': 2, 'f': 1, 'u': 1, 'a': 1, ' ': 2, ',': 1,
			},
			[utf8.UTFMax + 1]int{0, 15, 0, 0, 0},
			0,
		},
	}

	for idx, test := range tests {
		if c, l, i := CharCount(bufio.NewReader(test.input)); !compare(c, l, i, test.wantC, test.wantL, test.wantI) {
			t.Errorf("err %d", idx)
		}
	}
}

func compare(c map[rune]int, l [utf8.UTFMax + 1]int, i int, wc map[rune]int, wl [utf8.UTFMax + 1]int, wi int) bool {
	if i != wi {
		fmt.Println("I")
		return false
	}
	if !reflect.DeepEqual(l, wl) {
		fmt.Println(l)
		fmt.Println(wl)
		return false
	}

	for r, cou := range c {
		fmt.Printf("%s, %d == %d?\n", strconv.QuoteRune(r), cou, wc[r])
		if cou != wc[r] {
			return false
		}
	}
	return true
}
