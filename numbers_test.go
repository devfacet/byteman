// Byteman
// For the full copyright and license information, please view the LICENSE.txt file.

package byteman_test

import (
	"bytes"
	"math"
	"testing"

	"github.com/devfacet/byteman"
)

func TestFromUint(t *testing.T) {
	table := []struct {
		arg0 interface{}
		arg1 byteman.ByteOrder
		out  []byte
	}{
		{uint(math.MaxInt), &byteman.BigEndian{}, []byte{}},
		{uint(math.MaxInt), &byteman.LittleEndian{}, []byte{}},
		{uint8(0), &byteman.BigEndian{}, []byte{0x0}},
		{uint16(12345), &byteman.BigEndian{}, []byte{0x30, 0x39}},
		{uint16(54321), &byteman.BigEndian{}, []byte{0xD4, 0x31}},
		{uint8(math.MaxUint8), &byteman.BigEndian{}, []byte{0xff}},
		{uint16(math.MaxUint16), &byteman.BigEndian{}, []byte{0xff, 0xff}},
		{uint32(math.MaxUint32), &byteman.BigEndian{}, []byte{0xff, 0xff, 0xff, 0xff}},
		{uint64(math.MaxUint64), &byteman.BigEndian{}, []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
		{uint8(0), &byteman.LittleEndian{}, []byte{0x0}},
		{uint16(12345), &byteman.LittleEndian{}, []byte{0x39, 0x30}},
		{uint16(54321), &byteman.LittleEndian{}, []byte{0x31, 0xD4}},
		{uint8(math.MaxUint8), &byteman.LittleEndian{}, []byte{0xff}},
		{uint16(math.MaxUint16), &byteman.LittleEndian{}, []byte{0xff, 0xff}},
		{uint32(math.MaxUint32), &byteman.LittleEndian{}, []byte{0xff, 0xff, 0xff, 0xff}},
		{uint64(math.MaxUint64), &byteman.LittleEndian{}, []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
	}
	if byteman.IntSize == 64 {
		table[0].out = []byte{0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
		table[1].out = []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f}
	} else {
		table[0].out = []byte{0xff, 0xff, 0xff, 0xff}
		table[1].out = []byte{0xff, 0xff, 0xff, 0xff}
	}
	for _, v := range table {
		b := byteman.FromUint(v.arg0, v.arg1)
		if !bytes.Equal(b, v.out) {
			t.Errorf("got %v, want %v", b, v.out)
		}
	}

	if b := byteman.FromUint(float64(0), &byteman.BigEndian{}); b != nil {
		t.Errorf("got %v, want nil", b)
	} else if b := byteman.FromUint(float64(0), &byteman.LittleEndian{}); b != nil {
		t.Errorf("got %v, want nil", b)
	}
}

func BenchmarkFromUintBigEndian(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteman.FromUint(uint64(math.MaxUint64), &byteman.BigEndian{})
	}
}

func BenchmarkFromUintLittleEndian(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteman.FromUint(uint64(math.MaxUint64), &byteman.LittleEndian{})
	}
}

func TestFromInt(t *testing.T) {
	table := []struct {
		arg0 interface{}
		arg1 byteman.ByteOrder
		out  []byte
	}{
		{int(math.MaxInt), &byteman.BigEndian{}, []byte{}},
		{int(math.MaxInt), &byteman.LittleEndian{}, []byte{}},
		{int8(0), &byteman.BigEndian{}, []byte{0x0}},
		{int16(12345), &byteman.BigEndian{}, []byte{0x30, 0x39}},
		{int16(12321), &byteman.BigEndian{}, []byte{0x30, 0x21}},
		{int8(math.MaxInt8), &byteman.BigEndian{}, []byte{0x7f}},
		{int16(math.MaxInt16), &byteman.BigEndian{}, []byte{0x7f, 0xff}},
		{int32(math.MaxInt32), &byteman.BigEndian{}, []byte{0x7f, 0xff, 0xff, 0xff}},
		{int64(math.MaxInt64), &byteman.BigEndian{}, []byte{0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
		{int8(0), &byteman.LittleEndian{}, []byte{0x0}},
		{int16(12345), &byteman.LittleEndian{}, []byte{0x39, 0x30}},
		{int16(12321), &byteman.LittleEndian{}, []byte{0x21, 0x30}},
		{int8(math.MaxInt8), &byteman.LittleEndian{}, []byte{0x7f}},
		{int16(math.MaxInt16), &byteman.LittleEndian{}, []byte{0xff, 0x7f}},
		{int32(math.MaxInt32), &byteman.LittleEndian{}, []byte{0xff, 0xff, 0xff, 0x7f}},
		{int64(math.MaxInt64), &byteman.LittleEndian{}, []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f}},
	}
	if byteman.IntSize == 64 {
		table[0].out = []byte{0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
		table[1].out = []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f}
	} else {
		table[0].out = []byte{0x7f, 0xff, 0xff, 0xff}
		table[1].out = []byte{0xff, 0xff, 0xff, 0x7f}
	}
	for _, v := range table {
		b := byteman.FromInt(v.arg0, v.arg1)
		if !bytes.Equal(b, v.out) {
			t.Errorf("got %v, want %v", b, v.out)
		}
	}

	if b := byteman.FromInt(float64(0), &byteman.BigEndian{}); b != nil {
		t.Errorf("got %v, want nil", b)
	} else if b := byteman.FromInt(float64(0), &byteman.LittleEndian{}); b != nil {
		t.Errorf("got %v, want nil", b)
	}
}

func BenchmarkFromIntBigEndian(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteman.FromInt(uint64(math.MaxUint64), &byteman.BigEndian{})
	}
}

func BenchmarkFromIntLittleEndian(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteman.FromInt(uint64(math.MaxUint64), &byteman.LittleEndian{})
	}
}

func TestUint(t *testing.T) {
	table := []struct {
		arg0 []byte
		arg1 byteman.ByteOrder
		out  uint
	}{
		{[]byte{0xff, 0xff, 0xff, 0xff}, &byteman.BigEndian{}, math.MaxUint32},
		{[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, &byteman.BigEndian{}, math.MaxUint},
		{[]byte{0x49, 0x96, 0x02, 0xd2}, &byteman.BigEndian{}, 1234567890},
		{[]byte{0xbf, 0x63, 0xc8, 0x86}, &byteman.BigEndian{}, 3210987654},
		{[]byte{0xff, 0xff, 0xff, 0xff}, &byteman.LittleEndian{}, math.MaxUint32},
		{[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, &byteman.LittleEndian{}, math.MaxUint},
		{[]byte{0xd2, 0x02, 0x96, 0x49}, &byteman.LittleEndian{}, 1234567890},
		{[]byte{0x86, 0xc8, 0x63, 0xbf}, &byteman.LittleEndian{}, 3210987654},
	}
	for _, v := range table {
		if i := byteman.Uint(v.arg0, v.arg1); i != v.out {
			t.Errorf("got %v, want %v", i, v.out)
		}
	}

	if i := byteman.Uint([]byte{}, &byteman.BigEndian{}); i != 0 {
		t.Errorf("got %v, want 0", i)
	} else if i := byteman.Uint([]byte{}, &byteman.LittleEndian{}); i != 0 {
		t.Errorf("got %v, want 0", i)
	}
}

func BenchmarkUintBigEndian(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteman.Uint([]byte{0xff, 0xff, 0xff, 0xff}, &byteman.BigEndian{})
	}
}

func BenchmarkUintLittleEndian(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteman.Uint([]byte{0xff, 0xff, 0xff, 0xff}, &byteman.LittleEndian{})
	}
}

func TestUint16(t *testing.T) {
	table := []struct {
		arg0 []byte
		arg1 byteman.ByteOrder
		out  uint16
	}{
		{[]byte{0xff, 0xff}, &byteman.BigEndian{}, math.MaxUint16},
		{[]byte{0x30, 0x39}, &byteman.BigEndian{}, 12345},
		{[]byte{0xD4, 0x31}, &byteman.BigEndian{}, 54321},
		{[]byte{0xff, 0xff}, &byteman.LittleEndian{}, math.MaxUint16},
		{[]byte{0x39, 0x30}, &byteman.LittleEndian{}, 12345},
		{[]byte{0x31, 0xD4}, &byteman.LittleEndian{}, 54321},
	}
	for _, v := range table {
		if i := byteman.Uint16(v.arg0, v.arg1); i != v.out {
			t.Errorf("got %v, want %v", i, v.out)
		}
	}

	if i := byteman.Uint16([]byte{}, &byteman.BigEndian{}); i != 0 {
		t.Errorf("got %v, want 0", i)
	} else if i := byteman.Uint16([]byte{}, &byteman.LittleEndian{}); i != 0 {
		t.Errorf("got %v, want 0", i)
	}
}

func BenchmarkUint16BigEndian(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteman.Uint16([]byte{0xff, 0xff}, &byteman.BigEndian{})
	}
}

func BenchmarkUint16LittleEndian(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteman.Uint16([]byte{0xff, 0xff}, &byteman.LittleEndian{})
	}
}

func TestUint32(t *testing.T) {
	table := []struct {
		arg0 []byte
		arg1 byteman.ByteOrder
		out  uint32
	}{
		{[]byte{0xff, 0xff, 0xff, 0xff}, &byteman.BigEndian{}, math.MaxUint32},
		{[]byte{0x49, 0x96, 0x02, 0xd2}, &byteman.BigEndian{}, 1234567890},
		{[]byte{0xbf, 0x63, 0xc8, 0x86}, &byteman.BigEndian{}, 3210987654},
		{[]byte{0xff, 0xff, 0xff, 0xff}, &byteman.LittleEndian{}, math.MaxUint32},
		{[]byte{0xd2, 0x02, 0x96, 0x49}, &byteman.LittleEndian{}, 1234567890},
		{[]byte{0x86, 0xc8, 0x63, 0xbf}, &byteman.LittleEndian{}, 3210987654},
	}
	for _, v := range table {
		if i := byteman.Uint32(v.arg0, v.arg1); i != v.out {
			t.Errorf("got %v, want %v", i, v.out)
		}
	}

	if i := byteman.Uint32([]byte{}, &byteman.BigEndian{}); i != 0 {
		t.Errorf("got %v, want 0", i)
	} else if i := byteman.Uint32([]byte{}, &byteman.LittleEndian{}); i != 0 {
		t.Errorf("got %v, want 0", i)
	}
}

func BenchmarkUint32BigEndian(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteman.Uint32([]byte{0xff, 0xff, 0xff, 0xff}, &byteman.BigEndian{})
	}
}

func BenchmarkUint32LittleEndian(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteman.Uint32([]byte{0xff, 0xff, 0xff, 0xff}, &byteman.LittleEndian{})
	}
}

func TestUint64(t *testing.T) {
	table := []struct {
		arg0 []byte
		arg1 byteman.ByteOrder
		out  uint64
	}{
		{[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, &byteman.BigEndian{}, math.MaxUint64},
		{[]byte{0xab, 0x54, 0xa9, 0x8c, 0xeb, 0x1f, 0x0a, 0xd2}, &byteman.BigEndian{}, 12345678901234567890},
		{[]byte{0xf5, 0x00, 0xbf, 0x80, 0xb2, 0x98, 0xf5, 0x2d}, &byteman.BigEndian{}, 17654321098765432109},
		{[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, &byteman.LittleEndian{}, math.MaxUint64},
		{[]byte{0xd2, 0x0a, 0x1f, 0xeb, 0x8c, 0xa9, 0x54, 0xab}, &byteman.LittleEndian{}, 12345678901234567890},
		{[]byte{0x2d, 0xf5, 0x98, 0xb2, 0x80, 0xbf, 0x00, 0xf5}, &byteman.LittleEndian{}, 17654321098765432109},
	}
	for _, v := range table {
		if i := byteman.Uint64(v.arg0, v.arg1); i != v.out {
			t.Errorf("got %v, want %v", i, v.out)
		}
	}
	if i := byteman.Uint64([]byte{}, &byteman.BigEndian{}); i != 0 {
		t.Errorf("got %v, want 0", i)
	} else if i := byteman.Uint64([]byte{}, &byteman.LittleEndian{}); i != 0 {
		t.Errorf("got %v, want 0", i)
	}
}

func BenchmarkUint64BigEndian(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteman.Uint64([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, &byteman.BigEndian{})
	}
}

func BenchmarkUint64LittleEndian(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteman.Uint64([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, &byteman.LittleEndian{})
	}
}

func TestInt(t *testing.T) {
	table := []struct {
		arg0 []byte
		arg1 byteman.ByteOrder
		out  int
	}{
		{[]byte{0x7f, 0xff, 0xff, 0xff}, &byteman.BigEndian{}, math.MaxInt32},
		{[]byte{0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, &byteman.BigEndian{}, math.MaxInt},
		{[]byte{0x49, 0x96, 0x02, 0xd2}, &byteman.BigEndian{}, 1234567890},
		{[]byte{0x7d, 0xc2, 0x29, 0x3f}, &byteman.BigEndian{}, 2109876543},
		{[]byte{0xff, 0xff, 0xff, 0x7f}, &byteman.LittleEndian{}, math.MaxInt32},
		{[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f}, &byteman.LittleEndian{}, math.MaxInt},
		{[]byte{0xd2, 0x02, 0x96, 0x49}, &byteman.LittleEndian{}, 1234567890},
		{[]byte{0x3f, 0x29, 0xc2, 0x7d}, &byteman.LittleEndian{}, 2109876543},
	}
	for _, v := range table {
		if i := byteman.Int(v.arg0, v.arg1); i != v.out {
			t.Errorf("got %v, want %v", i, v.out)
		}
	}

	if i := byteman.Int([]byte{}, &byteman.BigEndian{}); i != 0 {
		t.Errorf("got %v, want 0", i)
	} else if i := byteman.Int([]byte{}, &byteman.LittleEndian{}); i != 0 {
		t.Errorf("got %v, want 0", i)
	}
}

func BenchmarkIntBigEndian(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteman.Int([]byte{0x7f, 0xff, 0xff, 0xff}, &byteman.BigEndian{})
	}
}

func BenchmarkIntLittleEndian(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteman.Int([]byte{0x7f, 0xff, 0xff, 0xff}, &byteman.LittleEndian{})
	}
}

func TestInt16(t *testing.T) {
	table := []struct {
		arg0 []byte
		arg1 byteman.ByteOrder
		out  int16
	}{
		{[]byte{0x7f, 0xff}, &byteman.BigEndian{}, math.MaxInt16},
		{[]byte{0x30, 0x39}, &byteman.BigEndian{}, 12345},
		{[]byte{0x30, 0x21}, &byteman.BigEndian{}, 12321},
		{[]byte{0xff, 0x7f}, &byteman.LittleEndian{}, math.MaxInt16},
		{[]byte{0x39, 0x30}, &byteman.LittleEndian{}, 12345},
		{[]byte{0x21, 0x30}, &byteman.LittleEndian{}, 12321},
	}
	for _, v := range table {
		if i := byteman.Int16(v.arg0, v.arg1); i != v.out {
			t.Errorf("got %v, want %v", i, v.out)
		}
	}

	if i := byteman.Int16([]byte{}, &byteman.BigEndian{}); i != 0 {
		t.Errorf("got %v, want 0", i)
	} else if i := byteman.Int16([]byte{}, &byteman.LittleEndian{}); i != 0 {
		t.Errorf("got %v, want 0", i)
	}
}

func BenchmarkInt16BigEndian(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteman.Int16([]byte{0xff, 0xff}, &byteman.BigEndian{})
	}
}

func BenchmarkInt16LittleEndian(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteman.Int16([]byte{0xff, 0xff}, &byteman.LittleEndian{})
	}
}

func TestInt32(t *testing.T) {
	table := []struct {
		arg0 []byte
		arg1 byteman.ByteOrder
		out  int32
	}{
		{[]byte{0x7f, 0xff, 0xff, 0xff}, &byteman.BigEndian{}, math.MaxInt32},
		{[]byte{0x49, 0x96, 0x02, 0xd2}, &byteman.BigEndian{}, 1234567890},
		{[]byte{0x7d, 0xc2, 0x29, 0x3f}, &byteman.BigEndian{}, 2109876543},
		{[]byte{0xff, 0xff, 0xff, 0x7f}, &byteman.LittleEndian{}, math.MaxInt32},
		{[]byte{0xd2, 0x02, 0x96, 0x49}, &byteman.LittleEndian{}, 1234567890},
		{[]byte{0x3f, 0x29, 0xc2, 0x7d}, &byteman.LittleEndian{}, 2109876543},
	}
	for _, v := range table {
		if i := byteman.Int32(v.arg0, v.arg1); i != v.out {
			t.Errorf("got %v, want %v", i, v.out)
		}
	}

	if i := byteman.Int32([]byte{}, &byteman.BigEndian{}); i != 0 {
		t.Errorf("got %v, want 0", i)
	} else if i := byteman.Int32([]byte{}, &byteman.LittleEndian{}); i != 0 {
		t.Errorf("got %v, want 0", i)
	}
}

func BenchmarkInt32BigEndian(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteman.Int32([]byte{0xff, 0xff}, &byteman.BigEndian{})
	}
}

func BenchmarkInt32LittleEndian(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteman.Int32([]byte{0xff, 0xff}, &byteman.LittleEndian{})
	}
}

func TestInt64(t *testing.T) {
	table := []struct {
		arg0 []byte
		arg1 byteman.ByteOrder
		out  int64
	}{
		{[]byte{0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, &byteman.BigEndian{}, math.MaxInt64},
		{[]byte{0x11, 0x22, 0x10, 0xf4, 0x7d, 0xe9, 0x81, 0x15}, &byteman.BigEndian{}, 1234567890123456789},
		{[]byte{0x4b, 0x62, 0xbb, 0x26, 0xf4, 0x3f, 0xc5, 0xeb}, &byteman.BigEndian{}, 5432109876543210987},
		{[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f}, &byteman.LittleEndian{}, math.MaxInt64},
		{[]byte{0x15, 0x81, 0xe9, 0x7d, 0xf4, 0x10, 0x22, 0x11}, &byteman.LittleEndian{}, 1234567890123456789},
		{[]byte{0xeb, 0xc5, 0x3f, 0xf4, 0x26, 0xbb, 0x62, 0x4b}, &byteman.LittleEndian{}, 5432109876543210987},
	}
	for _, v := range table {
		if i := byteman.Int64(v.arg0, v.arg1); i != v.out {
			t.Errorf("got %v, want %v", i, v.out)
		}
	}

	if i := byteman.Int64([]byte{}, &byteman.BigEndian{}); i != 0 {
		t.Errorf("got %v, want 0", i)
	} else if i := byteman.Int64([]byte{}, &byteman.LittleEndian{}); i != 0 {
		t.Errorf("got %v, want 0", i)
	}
}

func BenchmarkInt64BigEndian(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteman.Int64([]byte{0xff, 0xff}, &byteman.BigEndian{})
	}
}

func BenchmarkInt64LittleEndian(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteman.Int64([]byte{0xff, 0xff}, &byteman.LittleEndian{})
	}
}
