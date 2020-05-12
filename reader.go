// Copyright (c) 2019,CAO HONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bitutil

import (
	"encoding/binary"
	"errors"
	"io"
)

var errInvalidWidth = errors.New("bits.Reader: invalid w, negative or out of range of uint bits")

// Reader is a bits reader
type Reader struct {
	buf        []byte
	offset     int
	byteOffset int
}

// NewReader retruns a new Reader.
func NewReader(buf []byte) *Reader {
	return &Reader{
		buf: buf,
	}
}

// Skip skip w bits.
func (r *Reader) Skip(w int) (err error) {
	if w < 0 {
		return errInvalidWidth
	}

	bits := w + r.byteOffset
	if r.offset+bits>>3 >= len(r.buf) {
		return io.EOF
	}
	r.offset += bits >> 3
	r.byteOffset = bits - (bits>>3)<<3
	return
}

// ReadBit read a bit.
func (r *Reader) ReadBit() (ret uint8, err error) {
	if r.offset >= len(r.buf) {
		return 0, io.EOF
	}

	ret = (r.buf[r.offset] >> (7 - r.byteOffset)) & 1
	r.byteOffset++
	if r.byteOffset == 8 {
		r.offset++
		r.byteOffset = 0
	}
	return
}

// ReadUint8 read the uint8 of w bits.
func (r *Reader) ReadUint8(w int) (u8 uint8, err error) {
	if w < 0 || w > 8 {
		return 0, errInvalidWidth
	}
	bits := w + r.byteOffset
	if r.offset+bits>>3 >= len(r.buf) {
		return 0, io.EOF
	}

	if bits < 8 {
		if r.byteOffset == 0 {
			u8 = r.buf[r.offset] >> (8 - bits)
		} else {
			u8 = ((r.buf[r.offset] << r.byteOffset) >> (8 - bits + r.byteOffset))
		}
		r.byteOffset = bits
		return
	}

	if r.byteOffset == 0 {
		u8 = r.buf[r.offset]
	} else {
		u8 = (r.buf[r.offset] << r.byteOffset) >> r.byteOffset
	}
	r.offset++
	r.byteOffset = bits - 8

	if r.byteOffset > 0 {
		u8 = (u8 << r.byteOffset) | (r.buf[r.offset] >> (8 - r.byteOffset))
	}
	return
}

// ReadUint16 read the uint16 of w bits.
func (r *Reader) ReadUint16(w int) (u16 uint16, err error) {
	if w < 0 || w > 16 {
		return 0, errInvalidWidth
	}

	bits := w + r.byteOffset
	if r.offset+bits>>3 >= len(r.buf) {
		return 0, io.EOF
	}

	if bits >= 16 {
		u16 = binary.BigEndian.Uint16(r.buf[r.offset:])
		if r.byteOffset > 0 {
			u16 = (u16 << r.byteOffset) >> r.byteOffset
		}
		r.offset += 2
		r.byteOffset = 0
		bits -= 16
		goto LESS8
	}
	if bits >= 8 {
		if r.byteOffset == 0 {
			u16 = (u16 << 8) | uint16(r.buf[r.offset])
		} else {
			u16 = (u16 << 8) | uint16((r.buf[r.offset]<<r.byteOffset)>>r.byteOffset)
		}
		r.offset++
		r.byteOffset = 0
		bits -= 8
	}
LESS8:
	if bits > 0 {
		if r.byteOffset == 0 {
			u16 = (u16 << bits) | uint16(r.buf[r.offset]>>(8-bits))
		} else {
			u16 = (u16 << bits) | uint16((r.buf[r.offset]<<r.byteOffset)>>(8-bits+r.byteOffset))
		}
		r.byteOffset = bits
	}
	return
}

// ReadUint32 read the uint32 of w bits.
func (r *Reader) ReadUint32(w int) (u32 uint32, err error) {
	if w < 0 || w > 32 {
		return 0, errInvalidWidth
	}
	bits := w + r.byteOffset
	if r.offset+bits>>3 >= len(r.buf) {
		return 0, io.EOF
	}

	if bits >= 32 {
		u32 = binary.BigEndian.Uint32(r.buf[r.offset:])
		if r.byteOffset > 0 {
			u32 = (u32 << r.byteOffset) >> r.byteOffset
		}
		r.offset += 4
		r.byteOffset = 0
		bits -= 32
		goto LESS8
	}
	if bits >= 16 {
		u16 := binary.BigEndian.Uint16(r.buf[r.offset:])
		if r.byteOffset == 0 {
			u32 = (u32 << 16) | uint32(u16)
		} else {
			u32 = (u32 << 16) | uint32((u16<<r.byteOffset)>>r.byteOffset)
		}
		r.offset += 2
		r.byteOffset = 0
		bits -= 16
	}
	if bits >= 8 {
		if r.byteOffset == 0 {
			u32 = (u32 << 8) | uint32(r.buf[r.offset])
		} else {
			u32 = (u32 << 8) | uint32((r.buf[r.offset]<<r.byteOffset)>>r.byteOffset)
		}
		r.offset++
		r.byteOffset = 0
		bits -= 8
	}
LESS8:
	if bits > 0 {
		if r.byteOffset == 0 {
			u32 = (u32 << bits) | uint32(r.buf[r.offset]>>(8-bits))
		} else {
			u32 = (u32 << bits) | uint32((r.buf[r.offset]<<r.byteOffset)>>(8-bits+r.byteOffset))
		}
		r.byteOffset = bits
	}
	return
}

// ReadUint64 read the uint64 of w bits.
func (r *Reader) ReadUint64(w int) (u64 uint64, err error) {
	if w < 0 || w > 64 {
		return 0, errInvalidWidth
	}

	bits := w + r.byteOffset
	if r.offset+bits>>3 >= len(r.buf) {
		return 0, io.EOF
	}

	if bits >= 64 {
		u64 = binary.BigEndian.Uint64(r.buf[r.offset:])
		if r.byteOffset > 0 {
			u64 = (u64 << r.byteOffset) >> r.byteOffset
		}
		r.offset += 8
		r.byteOffset = 0
		bits -= 64
		goto LESS8
	}
	if bits >= 32 {
		u32 := binary.BigEndian.Uint32(r.buf[r.offset:])
		if r.byteOffset == 0 {
			u64 = uint64(u32)
		} else {
			u64 = uint64((u32 << r.byteOffset) >> r.byteOffset)
		}
		r.offset += 4
		r.byteOffset = 0
		bits -= 32
	}
	if bits >= 16 {
		u16 := binary.BigEndian.Uint16(r.buf[r.offset:])
		if r.byteOffset == 0 {
			u64 = (u64 << 16) | uint64(u16)
		} else {
			u64 = (u64 << 16) | uint64((u16<<r.byteOffset)>>r.byteOffset)
		}
		r.offset += 2
		r.byteOffset = 0
		bits -= 16
	}
	if bits >= 8 {
		if r.byteOffset == 0 {
			u64 = (u64 << 8) | uint64(r.buf[r.offset])
		} else {
			u64 = (u64 << 8) | uint64((r.buf[r.offset]<<r.byteOffset)>>r.byteOffset)
		}
		r.offset++
		r.byteOffset = 0
		bits -= 8
	}
LESS8:
	if bits > 0 {
		if r.byteOffset == 0 {
			u64 = (u64 << bits) | uint64(r.buf[r.offset]>>(8-bits))
		} else {
			u64 = (u64 << bits) | uint64((r.buf[r.offset]<<r.byteOffset)>>(8-bits+r.byteOffset))
		}
		r.byteOffset = bits
	}
	return
}

// Len returns the number of left bits.
func (r *Reader) Len() int {
	return (len(r.buf)-r.offset)<<3 - r.byteOffset
}

// Bytes returns the left byte slice.
func (r *Reader) Bytes() []byte {
	return r.buf[r.offset:]
}
