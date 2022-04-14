// Byteman
// For the full copyright and license information, please view the LICENSE.txt file.

package byteman

import (
	"encoding/hex"
)

// FromString returns a byte slice by the given string and desired byte size.
func FromString(str string, size int) []byte {
	if size == 0 {
		size = len(str)
	} else if size < 0 {
		if ns := len(str) + size; ns > 0 {
			size = ns
		} else {
			size = 0
		}
	}
	b := make([]byte, size)
	copy(b, []byte(str))
	return b
}

// FromHex returns a byte slice by the given ASCII Hex code and desired byte size.
func FromHex(hexcode string, size int) []byte {
	decoded, err := hex.DecodeString(hexcode)
	if err != nil {
		return nil
	}
	if size == 0 {
		size = len(decoded)
	} else if size < 0 {
		if ns := len(decoded) + size; ns > 0 {
			size = ns
		} else {
			size = 0
		}
	}
	b := make([]byte, size)
	copy(b, []byte(decoded))
	return b
}
