// Copyright (c) 2019,CAO HONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bitutil_test

import (
	"fmt"
	"testing"

	"github.com/cnotch/bitutil"
	"github.com/stretchr/testify/assert"
)

func TestTest(t *testing.T) {
	bs := []byte{0xe9, 0xa3}
	for i := 0; i < len(bs)*8; i++ {
		assert.Equal(t, bitutil.Test8(bs[i>>3], i%8), bitutil.Test(bs, i))
	}
}

func TestSet(t *testing.T) {
	bs := []byte{0xe9, 0xa3}
	u8 := uint8(0xa3)
	for i := 0; i < 8; i++ {
		bitutil.Set(bs, i+8)
		u8 = bitutil.Set8(u8, i)
		assert.Equal(t, u8, bs[1])
	}
	var bsnil []byte
	bsnil = bitutil.Set(bsnil, 14)
	assert.Equal(t, uint8(0x02), bsnil[1])
}

func TestClear(t *testing.T) {
	bs := []byte{0xe9, 0xa3}
	u8 := uint8(0xa3)
	for i := 0; i < 8; i++ {
		bitutil.Clear(bs, i+8)
		u8 = bitutil.Clear8(u8, i)
		assert.Equal(t, u8, bs[1])
	}
}

func TestFlip(t *testing.T) {
	bs := []byte{0xe9, 0xa3}
	bitutil.Flip(bs, 14)
	assert.Equal(t, uint8(0xa1), bs[1])
	bitutil.Flip(bs, 14)
	assert.Equal(t, uint8(0xa3), bs[1])
}

func TestTruncate(t *testing.T) {
	bs := []byte{0xe9, 0xa3, 0xe9, 0xa3}
	bs2 := bitutil.Truncate(bs, 24)
	assert.Equal(t, []byte{0xe9, 0xa3, 0xe9}, bs2)
	bs2 = bitutil.Truncate(bs, 21)
	assert.Equal(t, []byte{0xe9, 0xa3, 0xe8}, bs2)
	bs2 = bitutil.Truncate(bs, 20)
	assert.Equal(t, []byte{0xe9, 0xa3, 0xe0}, bs2)
	bs2 = bitutil.Truncate(bs, 18)
	assert.Equal(t, []byte{0xe9, 0xa3, 0xc0}, bs2)
	bs2 = bitutil.Truncate(bs, 17)
	assert.Equal(t, []byte{0xe9, 0xa3, 0x80}, bs2)
	bs2 = bitutil.Truncate(bs, 5)
	assert.Equal(t, []byte{0xe8}, bs2)
}

func TestTest64(t *testing.T) {
	var v uint64 = 123456
	want := fmt.Sprintf("%064b", v)
	for i := 0; i < 64; i++ {
		b := bitutil.Test64(v, i)
		assert.Equal(t, want[i] == '1', b)
	}
}

func TestSet64(t *testing.T) {
	var v uint64 = 123456
	for i := 0; i < 64; i++ {
		rv := bitutil.Set64(v, i)
		want := fmt.Sprintf("%064b", rv)
		assert.True(t, want[i] == '1')
	}
}

func TestClear64(t *testing.T) {
	var v uint64 = 123456
	for i := 0; i < 64; i++ {
		rv := bitutil.Clear64(v, i)
		want := fmt.Sprintf("%064b", rv)
		assert.False(t, want[i] == '1')
	}
}

func TestFlip64(t *testing.T) {
	var v uint64 = 0x1234_5678_9ABC_DEF1
	fv := bitutil.Flip64(v, 0)
	assert.Equal(t, uint64(0x9234_5678_9ABC_DEF1), fv)
	fv = bitutil.Flip64(v, 63)
	assert.Equal(t, uint64(0x1234_5678_9ABC_DEF0), fv)
	fv = bitutil.Flip64(v, 17)
	assert.Equal(t, uint64(0x1234_1678_9ABC_DEF1), fv)
	fv = bitutil.Flip64(v, 35)
	assert.Equal(t, uint64(0x1234_5678_8ABC_DEF1), fv)
}

func TestSub64(t *testing.T) {
	var v uint64 = 0x1234_5678_9ABC_DEF1
	want := fmt.Sprintf("%064b", v)
	cases := []struct {
		begin int
		end   int
	}{
		{1, 4},
		{6, 14},
		{50, 64},
	}
	for _, c := range cases {
		ret := bitutil.Sub64(v, c.begin, c.end)
		var wantu64 uint64
		fmt.Sscanf(want[c.begin:c.end], "%b", &wantu64)
		assert.Equal(t, wantu64, ret)
	}
}

func TestLeft64(t *testing.T) {
	var v uint64 = 0x1234_5678_9ABC_DEF1
	cases := []struct {
		w    int
		want uint64
	}{
		{4, 0x1},
		{6, 0b000100},
		{18, 0b00010010_00110100_01},
	}
	for _, c := range cases {
		ret := bitutil.Left64(v, c.w)
		assert.Equal(t, c.want, ret)
	}
}

func TestRight64(t *testing.T) {
	var v uint64 = 0x1234_5678_9ABC_DEF1
	cases := []struct {
		w    int
		want uint64
	}{
		{4, 0x1},
		{6, 0b110001},
		{18, 0xDEF1},
	}
	for _, c := range cases {
		ret := bitutil.Right64(v, c.w)
		assert.Equal(t, c.want, ret)
	}
}

func TestReplace64(t *testing.T) {
	var v uint64 = 0x1234_5678_9ABC_DEF1
	want := fmt.Sprintf("%064b", v)
	cases := []struct {
		begin int
		end   int
		new   uint64
	}{
		{1, 4, 0xff},
		{6, 14, 0x80809912},
		{6, 20, 0x90},
	}
	for _, c := range cases {
		ret := bitutil.Replace64(v, c.begin, c.end, c.new)
		newS := fmt.Sprintf("%064b", c.new)
		wantStr := want[:c.begin] + newS[64-c.end+c.begin:] + want[c.end:]
		var wantu64 uint64
		fmt.Sscanf(wantStr, "%064b", &wantu64)
		assert.Equal(t, wantu64, ret)
	}
}
