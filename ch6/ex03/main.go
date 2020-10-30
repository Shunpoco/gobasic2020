package main

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
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
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	s := &IntSet{}
	s.AddAll(1, 2, 3, 4, 5, 100, 102)
	t := &IntSet{}
	t.AddAll(1, 3, 7, 8, 100, 102)

	fmt.Println(s.String())

	s.IntersectWith(t)
	fmt.Println(s.String())

	s = &IntSet{}
	s.AddAll(1, 2, 3, 4, 5, 100, 102)

	s.DifferenceWith(t)
	fmt.Println(s.String())

	s = &IntSet{}
	s.AddAll(1, 2, 3, 4, 5, 100, 102)

	s.SymmetricDifference(t)
	fmt.Println(s.String())

}

func (s *IntSet) Remove(x int) {
	if s.Has(x) {
		word, bit := x/64, uint(x%64)
		s.words[word] ^= 1 << bit
	}
}

func (s *IntSet) Len() int {
	var counter int
	for _, word := range s.words {
		for j := 0; j < 64; j++ {
			if word&(1<<j) != 0 {
				counter++
			}
		}
	}
	return counter
}

func (s *IntSet) Clear() {
	s.words = []uint64{}
}

func (s *IntSet) Copy() *IntSet {
	return &IntSet{words: s.words}
}

func (s *IntSet) AddAll(vals ...int) {
	for _, val := range vals {
		s.Add(val)
	}
}

// 共通集合
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// sにあってtにないものの集合
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			and := s.words[i] & tword
			s.words[i] ^= and
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// sのみにあるものおよびtのみにあるものの集合
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}
