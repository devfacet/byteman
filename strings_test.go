// Byteman
// For the full copyright and license information, please view the LICENSE.txt file.

package byteman_test

import (
	"bytes"
	"testing"

	"github.com/devfacet/byteman"
)

func TestFromString(t *testing.T) {
	table := []struct {
		arg0 string
		arg1 int
		out  []byte
	}{
		{"foo", 0, []byte("foo")},
		{"bar", 1, []byte("b")},
		{"baz", 2, []byte("ba")},
		{"qux", 3, []byte("qux")},
		{"foo", 4, []byte{0x66, 0x6f, 0x6f, 0x00}},
		{"foo", -1, []byte("fo")},
		{"bar", -2, []byte("b")},
		{"baz", -3, []byte{}},
		{"qux", -4, []byte{}},
		{"", 0, []byte("")},
		{"", 0, []byte{}},
	}
	for _, v := range table {
		b := byteman.FromString(v.arg0, v.arg1)
		if !bytes.Equal(b, v.out) {
			t.Errorf("got %v, want %v", b, v.out)
		}
	}
}

func BenchmarkFromString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteman.FromString("foo", 3)
	}
}

func TestFromHex(t *testing.T) {
	table := []struct {
		arg0 string
		arg1 int
		out  []byte
	}{
		{"414243", 0, []byte("ABC")},
		{"414243", 1, []byte("A")},
		{"414243", 2, []byte("AB")},
		{"414243", 3, []byte("ABC")},
		{"414243", 4, []byte{0x41, 0x42, 0x43, 0x00}},
		{"414243", -1, []byte("AB")},
		{"414243", -2, []byte("A")},
		{"414243", -3, []byte{}},
		{"414243", -4, []byte{}},
		{"", 0, []byte("")},
		{"", 0, []byte{}},
	}
	for _, v := range table {
		b := byteman.FromHex(v.arg0, v.arg1)
		if !bytes.Equal(b, v.out) {
			t.Errorf("got %v, want %v", b, v.out)
		}
	}

	if b := byteman.FromHex("9", 1); b != nil {
		t.Errorf("got %v, want nil", b)
	}
}

func BenchmarkFromHex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteman.FromHex("A", 1)
	}
}
