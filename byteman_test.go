// Byteman
// For the full copyright and license information, please view the LICENSE.txt file.

package byteman_test

import (
	"bytes"
	"testing"

	"github.com/devfacet/byteman"
)

func TestCombine(t *testing.T) {
	table := []struct {
		arg0 []byte
		arg1 []byte
		out  []byte
	}{
		{[]byte("foo"), []byte("bar"), []byte("foobar")},
		{[]byte("foo"), []byte(""), []byte("foo")},
		{[]byte(""), []byte("bar"), []byte("bar")},
		{[]byte(""), []byte(""), []byte("")},
		{[]byte(""), []byte(""), []byte{}},
		{[]byte{}, []byte{}, []byte{}},
		{[]byte{}, []byte{}, []byte("")},
	}
	for _, v := range table {
		b := byteman.Combine(v.arg0, v.arg1)
		if !bytes.Equal(b, v.out) {
			t.Errorf("got %v, want %v", b, v.out)
		}
	}
}

func BenchmarkCombine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteman.Combine([]byte("a"), []byte("b"))
	}
}

func TestResize(t *testing.T) {
	table := []struct {
		arg0 []byte
		arg1 int
		out  []byte
	}{
		{[]byte("foo"), 0, []byte{0x66, 0x6f, 0x6f}},
		{[]byte("foo"), 1, []byte{0x66}},
		{[]byte("foo"), 2, []byte{0x66, 0x6f}},
		{[]byte("foo"), 3, []byte{0x66, 0x6f, 0x6f}},
		{[]byte("foo"), -1, []byte{0x66, 0x6f}},
		{[]byte("foo"), -2, []byte{0x66}},
		{[]byte("foo"), -3, []byte{}},
		{[]byte("foo"), -4, []byte{}},
		{[]byte("foo"), 4, []byte{0x66, 0x6f, 0x6f, 0x00}},
	}
	for _, v := range table {
		b := byteman.Resize(v.arg0, v.arg1)
		if !bytes.Equal(b, v.out) {
			t.Errorf("got %v, want %v", b, v.out)
		}
	}
}

func BenchmarkResize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteman.Resize([]byte("a"), 1)
	}
}
