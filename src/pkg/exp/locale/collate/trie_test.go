// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collate

import (
	"testing"
)

// We take the smallest, largest and an arbitrary value for each 
// of the UTF-8 sequence lengths.
var testRunes = []rune{
	0x01, 0x0C, 0x7F, // 1-byte sequences
	0x80, 0x100, 0x7FF, // 2-byte sequences
	0x800, 0x999, 0xFFFF, // 3-byte sequences
	0x10000, 0x10101, 0x10FFFF, // 4-byte sequences
	0x200, 0x201, 0x202, 0x210, 0x215, // five entries in one sparse block
}

// Test cases for illegal runes.
type trietest struct {
	size  int
	bytes []byte
}

var tests = []trietest{
	// illegal runes
	{1, []byte{0x80}},
	{1, []byte{0xFF}},
	{1, []byte{t2, tx - 1}},
	{1, []byte{t2, t2}},
	{2, []byte{t3, tx, tx - 1}},
	{2, []byte{t3, tx, t2}},
	{1, []byte{t3, tx - 1, tx}},
	{3, []byte{t4, tx, tx, tx - 1}},
	{3, []byte{t4, tx, tx, t2}},
	{1, []byte{t4, t2, tx, tx - 1}},
	{2, []byte{t4, tx, t2, tx - 1}},

	// short runes
	{0, []byte{t2}},
	{0, []byte{t3, tx}},
	{0, []byte{t4, tx, tx}},

	// we only support UTF-8 up to utf8.UTFMax bytes (4 bytes)
	{1, []byte{t5, tx, tx, tx, tx}},
	{1, []byte{t6, tx, tx, tx, tx, tx}},
}

func TestLookupTrie(t *testing.T) {
	for i, r := range testRunes {
		b := []byte(string(r))
		v, sz := testTrie.lookup(b)
		if int(v) != i {
			t.Errorf("lookup(%U): found value %#x, expected %#x", r, v, i)
		}
		if sz != len(b) {
			t.Errorf("lookup(%U): found size %d, expected %d", r, sz, len(b))
		}
	}
	for i, tt := range tests {
		v, sz := testTrie.lookup(tt.bytes)
		if int(v) != 0 {
			t.Errorf("lookup of illegal rune, case %d: found value %#x, expected 0", i, v)
		}
		if sz != tt.size {
			t.Errorf("lookup of illegal rune, case %d: found size %d, expected %d", i, sz, tt.size)
		}
	}
}

// test data is taken from exp/collate/locale/build/trie_test.go
var testValues = [832]uint32{
	0x000c: 0x00000001,
	0x007f: 0x00000002,
	0x00c0: 0x00000003,
	0x0100: 0x00000004,
	0x0140: 0x0000000c, 0x0141: 0x0000000d, 0x0142: 0x0000000e,
	0x0150: 0x0000000f,
	0x0155: 0x00000010,
	0x01bf: 0x00000005,
	0x01c0: 0x00000006,
	0x0219: 0x00000007,
	0x027f: 0x00000008,
	0x0280: 0x00000009,
	0x02c1: 0x0000000a,
	0x033f: 0x0000000b,
}

var testLookup = [640]uint16{
	0x0c2: 0x03, 0x0c4: 0x04,
	0x0c8: 0x05,
	0x0df: 0x06,
	0x0e0: 0x04,
	0x0ef: 0x05,
	0x0f0: 0x07, 0x0f4: 0x09,
	0x120: 0x07, 0x126: 0x08,
	0x17f: 0x09,
	0x180: 0x0a, 0x184: 0x0b,
	0x1d0: 0x06,
	0x23f: 0x0c,
	0x24f: 0x08,
}

var testTrie = trie{testLookup[:], testValues[:]}
