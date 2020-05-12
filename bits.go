// Copyright (c) 2019,CAO HONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:generate go run genbits.go

package bitutil

// Test whether the ith bit is set of a byte slice.
func Test(s []byte, i int) bool {
	if i < 0 || i >= len(s)<<3 {
		return false
	}
	return s[i>>3]&(1<<(7&^i)) != 0
}

// Set set the ith bit value to 1.
func Set(s []byte, i int) []byte {
	if i < 0 {
		return s
	}

	n := i>>3 + 1
	if n > len(s) {
		if cap(s) < n {
			buf := make([]byte, n, 2*cap(s)+n-len(s))
			copy(buf, s)
			s = buf
		} else {
			s = s[:n]
		}
	}
	s[i>>3] |= (1 << (7 &^ i))
	return s
}

// Clear clear the ith bit value to 0.
func Clear(s []byte, i int) []byte {
	if i < 0 || i >= len(s)<<3 {
		return s
	}

	s[i>>3] &^= (1 << (7 &^ i))
	return s
}

// SetTo set the ith bit value.
func SetTo(s []byte, i int, value bool) []byte {
	if value {
		return Set(s, i)
	}
	return Clear(s, i)
}

// Flip flip the ith bit value.
func Flip(s []byte, i int) []byte {
	if i < 0 {
		return s
	}
	if i >= len(s)<<3 {
		return Set(s, i)
	}

	s[i>>3] ^= 1 << (7 &^ i)
	return s
}

// Truncate discards all but the first n bitsbits from the s,
// returns the new byte slict but continues to use the same allocated storage.
func Truncate(s []byte, n int) []byte {
	m := (n-1)>>3 + 1
	if m > len(s) {
		return s
	}

	s = s[:m]
	clearSet := s[m:]
	for i := 0; i < len(clearSet); i++ {
		clearSet[i] = 0
	}
	if m == 0 {
		return s
	}

	// last byte
	if 7&n > 0 {
		s[m-1] &= 0xff << (8 - (7 & n))
	}
	return s
}
