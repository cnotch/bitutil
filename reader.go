// Copyright (c) 2019,CAO HONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bitutil

import (
	"errors"
	"io"
)

var errInvalidBitsNumber = errors.New("bits.Reader: invalid n, negative or out of range of reader bits")

// Reader is a bits reader
type Reader struct {
	buf    []byte
	offset int // bit base
}

// NewReader retruns a new Reader.
func NewReader(buf []byte) *Reader {
	return &Reader{
		buf: buf,
	}
}

// Skip skip n bits.
func (r *Reader) Skip(n int) (err error) {
	if n < 0 {
		return errInvalidBitsNumber
	}

	if (r.offset + n) > len(r.buf)<<3 {
		return io.EOF
	}
	r.offset += n
	return
}

// Peek peek the int of n bits.
func (r *Reader) Peek(n int) (u64 uint64, err error) {
	offset := r.offset
	defer func() {
		r.offset = offset
	}()
	return r.ReadUint64(n)
}

// ReadBit read a bit.
func (r *Reader) ReadBit() (ret uint8, err error) {
	if (r.offset + 1) > len(r.buf)<<3 {
		return 0, io.EOF
	}

	ret = (r.buf[r.offset>>3] >> (7 - r.offset&0x7)) & 1
	r.offset++
	return
}

// ReadUint8 read the uint8 of n bits.
func (r *Reader) ReadUint8(n int) (u8 uint8, err error) {
	if n < 0 || n > 8 {
		return 0, errInvalidBitsNumber
	}
	var u64 uint64
	u64, err = r.readBits(n)
	u8 = uint8(u64)
	return
}

// ReadUint16 read the uint16 of n bits.
func (r *Reader) ReadUint16(n int) (u16 uint16, err error) {
	if n < 0 || n > 16 {
		return 0, errInvalidBitsNumber
	}
	var u64 uint64
	u64, err = r.readBits(n)
	u16 = uint16(u64)
	return
}

// ReadUint32 read the uint32 of n bits.
func (r *Reader) ReadUint32(n int) (u32 uint32, err error) {
	if n < 0 || n > 32 {
		return 0, errInvalidBitsNumber
	}
	var u64 uint64
	u64, err = r.readBits(n)
	u32 = uint32(u64)
	return
}

// ReadUint64 read the uint64 of n bits.
func (r *Reader) ReadUint64(n int) (u64 uint64, err error) {
	if n < 0 || n > 64 {
		return 0, errInvalidBitsNumber
	}
	return r.readBits(n)
}

// Offset returns the offset of bits.
func (r *Reader) Offset() int {
	return r.offset
}

// BitsLeft returns the number of left bits.
func (r *Reader) BitsLeft() int {
	return len(r.buf)<<3 - r.offset
}

// Bytes returns the left byte slice.
func (r *Reader) Bytes() []byte {
	return r.buf[r.offset>>3:]
}

var bitsMask = [9]byte{
	0x00,
	0x01, 0x03, 0x07, 0x0f,
	0x1f, 0x3f, 0x7f, 0xff,
}

func (r *Reader) readBits(n int) (u64 uint64, err error) {
	if (r.offset + n) > len(r.buf)<<3 {
		return 0, io.EOF
	}

	idx := r.offset >> 3
	validBits := 8 - r.offset&0x7
	r.offset += n

	for n >= validBits {
		n -= validBits
		u64 |= uint64(r.buf[idx]&bitsMask[validBits]) << n
		idx++
		validBits = 8
	}

	if n > 0 {
		u64 |= uint64((r.buf[idx] >> (validBits - n)) & bitsMask[n])
	}
	return
}
