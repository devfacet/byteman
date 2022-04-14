// Byteman
// For the full copyright and license information, please view the LICENSE.txt file.

// Package byteman provides functions for bytes and bits.
package byteman

// Combine combines the given byte slices.
func Combine(bs ...[]byte) []byte {
	b := []byte{}
	for _, v := range bs {
		b = append(b, v...)
	}
	return b
}

// Resize resizes the given byte slice.
func Resize(b []byte, size int) []byte {
	if size == 0 {
		size = len(b)
	} else if size < 0 {
		if ns := len(b) + size; ns > 0 {
			size = ns
		} else {
			size = 0
		}
	}
	nb := make([]byte, size)
	if b != nil {
		copy(nb, b)
	}
	return nb
}
