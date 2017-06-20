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

// Package ifnil provides a type-safe way to test variables for nil and substitute a default akin to SQL's coalesce function.
package ifnil

// Bool returns def if arg is nil, otherwise it returns the value pointed to by arg
func Bool(arg *bool, def bool) bool {
	if arg == nil {
		return def
	}
	return *arg
}

// Rune returns def if arg is nil, otherwise it returns the value pointed to by arg
func Rune(arg *rune, def rune) rune {
	if arg == nil {
		return def
	}
	return *arg
}

// String returns def if arg is nil, otherwise it returns the value pointed to by arg
func String(arg *string, def string) string {
	if arg == nil {
		return def
	}
	return *arg
}

// Int returns def if arg is nil, otherwise it returns the value pointed to by arg
func Int(arg *int, def int) int {
	if arg == nil {
		return def
	}
	return *arg
}

// Int8 returns def if arg is nil, otherwise it returns the value pointed to by arg
func Int8(arg *int8, def int8) int8 {
	if arg == nil {
		return def
	}
	return *arg
}

// Int16 returns def if arg is nil, otherwise it returns the value pointed to by arg
func Int16(arg *int16, def int16) int16 {
	if arg == nil {
		return def
	}
	return *arg
}

// Int32 returns def if arg is nil, otherwise it returns the value pointed to by arg
func Int32(arg *int32, def int32) int32 {
	if arg == nil {
		return def
	}
	return *arg
}

// Int64 returns def if arg is nil, otherwise it returns the value pointed to by arg
func Int64(arg *int64, def int64) int64 {
	if arg == nil {
		return def
	}
	return *arg
}

// Uint returns def if arg is nil, otherwise it returns the value pointed to by arg
func Uint(arg *uint, def uint) uint {
	if arg == nil {
		return def
	}
	return *arg
}

// Bool returns def if arg is nil, otherwise it returns the value pointed to by arg
func Uint8(arg *uint8, def uint8) uint8 {
	if arg == nil {
		return def
	}
	return *arg
}

// Uint16 returns def if arg is nil, otherwise it returns the value pointed to by arg
func Uint16(arg *uint16, def uint16) uint16 {
	if arg == nil {
		return def
	}
	return *arg
}

// Uint32 returns def if arg is nil, otherwise it returns the value pointed to by arg
func Uint32(arg *uint32, def uint32) uint32 {
	if arg == nil {
		return def
	}
	return *arg
}

// Uint64 returns def if arg is nil, otherwise it returns the value pointed to by arg
func Uint64(arg *uint64, def uint64) uint64 {
	if arg == nil {
		return def
	}
	return *arg
}

// Float32 returns def if arg is nil, otherwise it returns the value pointed to by arg
func Float32(arg *float32, def float32) float32 {
	if arg == nil {
		return def
	}
	return *arg
}

// Float64 returns def if arg is nil, otherwise it returns the value pointed to by arg
func Float64(arg *float64, def float64) float64 {
	if arg == nil {
		return def
	}
	return *arg
}
