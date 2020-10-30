package main

import (
	"bytes"
	"fmt"
)

const bitSize = 32 << (^uint(0) >> 63)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/bitSize, uint(x%bitSize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/bitSize, uint(x%bitSize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bitSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", bitSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	s := &IntSet{}
	i := 3
	s.Add(i)
	i = 2
	s.Add(i)
	i = 100
	s.Add(i)
	fmt.Println(s.String())
	fmt.Println(s.Len())
	s.Remove(100)
	fmt.Println(s.String())
	fmt.Println(s.Len())

	s2 := s.Copy()
	fmt.Println(s2.String())
	s.Clear()
	fmt.Println(s.Len())
	fmt.Println(s2.String())
	fmt.Println(s2.Len())
}

func (s *IntSet) Remove(x int) {
	if s.Has(x) {
		word, bit := x/bitSize, uint(x%bitSize)
		s.words[word] ^= 1 << bit
	}
}

func (s *IntSet) Len() int {
	var counter int
	for _, word := range s.words {
		for j := 0; j < bitSize; j++ {
			if word&(1<<j) != 0 {
				counter++
			}
		}
	}
	return counter
}

func (s *IntSet) Clear() {
	s.words = []uint{}
}

func (s *IntSet) Copy() *IntSet {
	return &IntSet{words: s.words}
}
