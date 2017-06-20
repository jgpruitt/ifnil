/*
MIT License

Copyright (c) 2017 John Pruitt

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package ifnil_test

import (
	"ifnil"
	"testing"
)

func TestBool1(t *testing.T) {
	var x *bool = nil
	if ifnil.Bool(x, true) == false {
		t.Fail()
	}
}

func TestBool2(t *testing.T) {
	var x *bool = nil
	if ifnil.Bool(x, false) == true {
		t.Fail()
	}
}

func TestBool3(t *testing.T) {
	x := true
	if ifnil.Bool(&x, false) == false {
		t.Fail()
	}
}

func TestBool4(t *testing.T) {
	x := false
	if ifnil.Bool(&x, true) == true {
		t.Fail()
	}
}

func TestRune1(t *testing.T) {
	var x *rune = nil
	if ifnil.Rune(x, 'a') != 'a' {
		t.Fail()
	}
}

func TestRune2(t *testing.T) {
	x := 'a'
	if ifnil.Rune(&x, 'b') != 'a' {
		t.Fail()
	}
}

func TestString1(t *testing.T) {
	var x *string = nil
	if ifnil.String(x, "a") != "a" {
		t.Fail()
	}
}

func TestString2(t *testing.T) {
	x := "a"
	if ifnil.String(&x, "b") != "a" {
		t.Fail()
	}
}

func TestFloat321(t *testing.T) {
	var (
		x *float32 = nil
		y float32 = 1.0
	)
	if ifnil.Float32(x, y) != y {
		t.Fail()
	}
}

func TestFloat322(t *testing.T) {
	var (
		x float32 = 1.0
		y float32 = 2.0
	)
	if ifnil.Float32(&x, y) != x {
		t.Fail()
	}
}

func TestFloat641(t *testing.T) {
	var (
		x *float64 = nil
		y float64 = 1.0
	)
	if ifnil.Float64(x, y) != y {
		t.Fail()
	}
}

func TestFloat642(t *testing.T) {
	var (
		x float64 = 1.0
		y float64 = 2.0
	)
	if ifnil.Float64(&x, y) != x {
		t.Fail()
	}
}

func TestInt1(t *testing.T) {
	var (
		x *int = nil
		y int = 1
	)
	if ifnil.Int(x, y) != y {
		t.Fail()
	}
}

func TestInt2(t *testing.T) {
	var (
		x int = 1
		y int = 2
	)
	if ifnil.Int(&x, y) != x {
		t.Fail()
	}
}

func TestInt81(t *testing.T) {
	var (
		x *int8 = nil
		y int8 = 1
	)
	if ifnil.Int8(x, y) != y {
		t.Fail()
	}
}

func TestInt82(t *testing.T) {
	var (
		x int8 = 1
		y int8 = 2
	)
	if ifnil.Int8(&x, y) != x {
		t.Fail()
	}
}

func TestInt161(t *testing.T) {
	var (
		x *int16 = nil
		y int16 = 1
	)
	if ifnil.Int16(x, y) != y {
		t.Fail()
	}
}

func TestInt162(t *testing.T) {
	var (
		x int16 = 1
		y int16 = 2
	)
	if ifnil.Int16(&x, y) != x {
		t.Fail()
	}
}

func TestInt321(t *testing.T) {
	var (
		x *int32 = nil
		y int32 = 1
	)
	if ifnil.Int32(x, y) != y {
		t.Fail()
	}
}

func TestInt322(t *testing.T) {
	var (
		x int32 = 1
		y int32 = 2
	)
	if ifnil.Int32(&x, y) != x {
		t.Fail()
	}
}

func TestInt641(t *testing.T) {
	var (
		x *int64 = nil
		y int64 = 1
	)
	if ifnil.Int64(x, y) != y {
		t.Fail()
	}
}

func TestInt642(t *testing.T) {
	var (
		x int64 = 1
		y int64 = 2
	)
	if ifnil.Int64(&x, y) != x {
		t.Fail()
	}
}

func TestUint1(t *testing.T) {
	var (
		x *uint = nil
		y uint = 1
	)
	if ifnil.Uint(x, y) != y {
		t.Fail()
	}
}

func TestUint2(t *testing.T) {
	var (
		x uint = 1
		y uint = 2
	)
	if ifnil.Uint(&x, y) != x {
		t.Fail()
	}
}

func TestUint81(t *testing.T) {
	var (
		x *uint8 = nil
		y uint8 = 1
	)
	if ifnil.Uint8(x, y) != y {
		t.Fail()
	}
}

func TestUint82(t *testing.T) {
	var (
		x uint8 = 1
		y uint8 = 2
	)
	if ifnil.Uint8(&x, y) != x {
		t.Fail()
	}
}

func TestUint161(t *testing.T) {
	var (
		x *uint16 = nil
		y uint16 = 1
	)
	if ifnil.Uint16(x, y) != y {
		t.Fail()
	}
}

func TestUint162(t *testing.T) {
	var (
		x uint16 = 1
		y uint16 = 2
	)
	if ifnil.Uint16(&x, y) != x {
		t.Fail()
	}
}

func TestUint321(t *testing.T) {
	var (
		x *uint32 = nil
		y uint32 = 1
	)
	if ifnil.Uint32(x, y) != y {
		t.Fail()
	}
}

func TestUint322(t *testing.T) {
	var (
		x uint32 = 1
		y uint32 = 2
	)
	if ifnil.Uint32(&x, y) != x {
		t.Fail()
	}
}

func TestUint641(t *testing.T) {
	var (
		x *uint64 = nil
		y uint64 = 1
	)
	if ifnil.Uint64(x, y) != y {
		t.Fail()
	}
}

func TestUint642(t *testing.T) {
	var (
		x uint64 = 1
		y uint64 = 2
	)
	if ifnil.Uint64(&x, y) != x {
		t.Fail()
	}
}