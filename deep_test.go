package deep

import (
	"reflect"
	"testing"
)

func TestCopy_Bool(t *testing.T) {
	doCopyAndCheck(t, true, false)
}

func TestCopy_Int(t *testing.T) {
	doCopyAndCheck(t, 42, false)
}

func TestCopy_Int8(t *testing.T) {
	doCopyAndCheck(t, int8(42), false)
}

func TestCopy_Int16(t *testing.T) {
	doCopyAndCheck(t, int16(42), false)
}

func TestCopy_Int32(t *testing.T) {
	doCopyAndCheck(t, int32(42), false)
}

func TestCopy_Int64(t *testing.T) {
	doCopyAndCheck(t, int64(42), false)
}

func TestCopy_Uint(t *testing.T) {
	doCopyAndCheck(t, uint(42), false)
}

func TestCopy_Uint8(t *testing.T) {
	doCopyAndCheck(t, uint8(42), false)
}

func TestCopy_Uint16(t *testing.T) {
	doCopyAndCheck(t, uint16(42), false)
}

func TestCopy_Uint32(t *testing.T) {
	doCopyAndCheck(t, uint32(42), false)
}

func TestCopy_Uint64(t *testing.T) {
	doCopyAndCheck(t, uint64(42), false)
}

func TestCopy_Uintptr(t *testing.T) {
	doCopyAndCheck(t, uintptr(42), false)
}

func TestCopy_Float32(t *testing.T) {
	doCopyAndCheck(t, float32(42), false)
}

func TestCopy_Float64(t *testing.T) {
	doCopyAndCheck(t, float64(42), false)
}

func TestCopy_Complex64(t *testing.T) {
	doCopyAndCheck(t, complex64(42), false)
}

func TestCopy_Complex128(t *testing.T) {
	doCopyAndCheck(t, complex128(42), false)
}

func TestCopy_String(t *testing.T) {
	doCopyAndCheck(t, "42", false)
}

func TestCopy_Array(t *testing.T) {
	doCopyAndCheck(t, [4]int{42, 43, 44, 45}, false)
}

func TestCopy_Map(t *testing.T) {
	doCopyAndCheck(t, map[int]string{42: "42", 43: "43", 44: "44", 45: "45"}, false)
}

func TestCopy_Ptr(t *testing.T) {
	value := 42
	doCopyAndCheck(t, &value, false)
}

func TestCopy_Ptr_Nil(t *testing.T) {
	value := (*int)(nil)
	doCopyAndCheck(t, value, false)
}

func TestCopy_Slice(t *testing.T) {
	doCopyAndCheck(t, []int{42, 43, 44, 45}, false)
}

func TestCopy_Struct(t *testing.T) {
	type S struct {
		A int
		B string
	}
	doCopyAndCheck(t, S{42, "42"}, false)
}

func TestCopy_Struct_Loop(t *testing.T) {
	type S struct {
		A int
		B *S
	}

	// Create a loop.
	src := S{A: 1}
	src.B = &src

	doCopyAndCheck(t, src, false)
}

func TestCopy_Struct_Unexported(t *testing.T) {
	type S struct {
		a int
		b string
	}

	doCopyAndCheck(t, S{42, "42"}, false)
}

func TestCopy_Func_Error(t *testing.T) {
	doCopyAndCheck(t, func() {}, true)
}

func TestCopy_Chan_Error(t *testing.T) {
	doCopyAndCheck(t, make(chan struct{}), true)
}

func doCopyAndCheck[T any](t *testing.T, src T, expectError bool) {
	t.Helper()

	dst, err := Copy(src)
	if err != nil {
		if !expectError {
			t.Errorf("Copy failed: %v", err)
		}
		return
	}
	if !reflect.DeepEqual(dst, src) {
		t.Errorf("Copy failed: expected %v, got %v", src, dst)
	}
}
