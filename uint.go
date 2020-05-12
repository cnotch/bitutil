// Code generated using genbits.go; DO NOT EDIT.

// Copyright (c) 2019,CAO HONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bitutil

// Test8 returns the ith bit value of an uint8.
func Test8(v uint8, i int) bool {
	if i < 0 || i > 7 {
		return false
	}
	return v&(1<<(7-i)) != 0
}

// Set8 set the ith bit value of an uint8 to 1, and returns the new uint8.
func Set8(v uint8, i int) uint8 {
	if i < 0 || i > 7 {
		return v
	}
	return v | (1 << (7 - i))
}

// Clear8 clear the ith bit value of an uint8 to 0, and returns the new uint8.
func Clear8(v uint8, i int) uint8 {
	if i < 0 || i > 7 {
		return v
	}
	return v & ^(1 << (7 - i))
}

// SetTo8 set the ith bit value of an uint8, and returns the new uint8.
func SetTo8(v uint8, i int, b bool) uint8 {
	if b {
		return Set8(v, i)
	}
	return Clear8(v, i)
}

// Flip8 flip the ith bit value of an uint8, and returns the new uint8.
func Flip8(v uint8, i int) uint8 {
	if i < 0 || i > 7 {
		return v
	}
	return v ^(1 << (7 - i))
}

// Sub8 returns the copy uint8 of v[begin:end].
func Sub8(v uint8, begin, end int) uint8 {
	if begin < 0 || begin > 8 || end < 0 || end > 8 {
		return 0
	}
	return (v << begin) >> (8 - end + begin)
}

// Left8 returns the copy uint8 of v[:w].
func Left8(v uint8, w int) uint8 {
	if w < 0 || w > 8 {
		return 0
	}
	return v >> (8 - w)
}

// Right8 returns the copy uint8 of v[8-w:].
func Right8(v uint8, w int) uint8 {
	if w < 0 || w > 8 {
		return 0
	}
	return (v << (8 - w)) >> (8 - w)
}

// Replace8 returns a copy of the v with v[begin:end] instances of old replaced by new[8-end+begin:].
func Replace8(v uint8, begin, end int, new uint8) uint8 {
	if begin < 0 || begin > 8 || end < 0 || end > 8 {
		return v
	}
	mask := (uint8(0xff) << (8 - end + begin)) >> begin
	return (v & ^mask) | ((new << (8 - end)) & mask)
}

// Test16 returns the ith bit value of an uint16.
func Test16(v uint16, i int) bool {
	if i < 0 || i > 15 {
		return false
	}
	return v&(1<<(15-i)) != 0
}

// Set16 set the ith bit value of an uint16 to 1, and returns the new uint16.
func Set16(v uint16, i int) uint16 {
	if i < 0 || i > 15 {
		return v
	}
	return v | (1 << (15 - i))
}

// Clear16 clear the ith bit value of an uint16 to 0, and returns the new uint16.
func Clear16(v uint16, i int) uint16 {
	if i < 0 || i > 15 {
		return v
	}
	return v & ^(1 << (15 - i))
}

// SetTo16 set the ith bit value of an uint16, and returns the new uint16.
func SetTo16(v uint16, i int, b bool) uint16 {
	if b {
		return Set16(v, i)
	}
	return Clear16(v, i)
}

// Flip16 flip the ith bit value of an uint16, and returns the new uint16.
func Flip16(v uint16, i int) uint16 {
	if i < 0 || i > 15 {
		return v
	}
	return v ^(1 << (15 - i))
}

// Sub16 returns the copy uint16 of v[begin:end].
func Sub16(v uint16, begin, end int) uint16 {
	if begin < 0 || begin > 16 || end < 0 || end > 16 {
		return 0
	}
	return (v << begin) >> (16 - end + begin)
}

// Left16 returns the copy uint16 of v[:w].
func Left16(v uint16, w int) uint16 {
	if w < 0 || w > 16 {
		return 0
	}
	return v >> (16 - w)
}

// Right16 returns the copy uint16 of v[16-w:].
func Right16(v uint16, w int) uint16 {
	if w < 0 || w > 16 {
		return 0
	}
	return (v << (16 - w)) >> (16 - w)
}

// Replace16 returns a copy of the v with v[begin:end] instances of old replaced by new[16-end+begin:].
func Replace16(v uint16, begin, end int, new uint16) uint16 {
	if begin < 0 || begin > 16 || end < 0 || end > 16 {
		return v
	}
	mask := (uint16(0xffff) << (16 - end + begin)) >> begin
	return (v & ^mask) | ((new << (16 - end)) & mask)
}

// Test32 returns the ith bit value of an uint32.
func Test32(v uint32, i int) bool {
	if i < 0 || i > 31 {
		return false
	}
	return v&(1<<(31-i)) != 0
}

// Set32 set the ith bit value of an uint32 to 1, and returns the new uint32.
func Set32(v uint32, i int) uint32 {
	if i < 0 || i > 31 {
		return v
	}
	return v | (1 << (31 - i))
}

// Clear32 clear the ith bit value of an uint32 to 0, and returns the new uint32.
func Clear32(v uint32, i int) uint32 {
	if i < 0 || i > 31 {
		return v
	}
	return v & ^(1 << (31 - i))
}

// SetTo32 set the ith bit value of an uint32, and returns the new uint32.
func SetTo32(v uint32, i int, b bool) uint32 {
	if b {
		return Set32(v, i)
	}
	return Clear32(v, i)
}

// Flip32 flip the ith bit value of an uint32, and returns the new uint32.
func Flip32(v uint32, i int) uint32 {
	if i < 0 || i > 31 {
		return v
	}
	return v ^(1 << (31 - i))
}

// Sub32 returns the copy uint32 of v[begin:end].
func Sub32(v uint32, begin, end int) uint32 {
	if begin < 0 || begin > 32 || end < 0 || end > 32 {
		return 0
	}
	return (v << begin) >> (32 - end + begin)
}

// Left32 returns the copy uint32 of v[:w].
func Left32(v uint32, w int) uint32 {
	if w < 0 || w > 32 {
		return 0
	}
	return v >> (32 - w)
}

// Right32 returns the copy uint32 of v[32-w:].
func Right32(v uint32, w int) uint32 {
	if w < 0 || w > 32 {
		return 0
	}
	return (v << (32 - w)) >> (32 - w)
}

// Replace32 returns a copy of the v with v[begin:end] instances of old replaced by new[32-end+begin:].
func Replace32(v uint32, begin, end int, new uint32) uint32 {
	if begin < 0 || begin > 32 || end < 0 || end > 32 {
		return v
	}
	mask := (uint32(0xffffffff) << (32 - end + begin)) >> begin
	return (v & ^mask) | ((new << (32 - end)) & mask)
}

// Test64 returns the ith bit value of an uint64.
func Test64(v uint64, i int) bool {
	if i < 0 || i > 63 {
		return false
	}
	return v&(1<<(63-i)) != 0
}

// Set64 set the ith bit value of an uint64 to 1, and returns the new uint64.
func Set64(v uint64, i int) uint64 {
	if i < 0 || i > 63 {
		return v
	}
	return v | (1 << (63 - i))
}

// Clear64 clear the ith bit value of an uint64 to 0, and returns the new uint64.
func Clear64(v uint64, i int) uint64 {
	if i < 0 || i > 63 {
		return v
	}
	return v & ^(1 << (63 - i))
}

// SetTo64 set the ith bit value of an uint64, and returns the new uint64.
func SetTo64(v uint64, i int, b bool) uint64 {
	if b {
		return Set64(v, i)
	}
	return Clear64(v, i)
}

// Flip64 flip the ith bit value of an uint64, and returns the new uint64.
func Flip64(v uint64, i int) uint64 {
	if i < 0 || i > 63 {
		return v
	}
	return v ^(1 << (63 - i))
}

// Sub64 returns the copy uint64 of v[begin:end].
func Sub64(v uint64, begin, end int) uint64 {
	if begin < 0 || begin > 64 || end < 0 || end > 64 {
		return 0
	}
	return (v << begin) >> (64 - end + begin)
}

// Left64 returns the copy uint64 of v[:w].
func Left64(v uint64, w int) uint64 {
	if w < 0 || w > 64 {
		return 0
	}
	return v >> (64 - w)
}

// Right64 returns the copy uint64 of v[64-w:].
func Right64(v uint64, w int) uint64 {
	if w < 0 || w > 64 {
		return 0
	}
	return (v << (64 - w)) >> (64 - w)
}

// Replace64 returns a copy of the v with v[begin:end] instances of old replaced by new[64-end+begin:].
func Replace64(v uint64, begin, end int, new uint64) uint64 {
	if begin < 0 || begin > 64 || end < 0 || end > 64 {
		return v
	}
	mask := (uint64(0xffffffffffffffff) << (64 - end + begin)) >> begin
	return (v & ^mask) | ((new << (64 - end)) & mask)
}
