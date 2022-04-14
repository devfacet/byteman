// Byteman
// For the full copyright and license information, please view the LICENSE.txt file.

package byteman

// ByteOrderType represents the byte order type.
type ByteOrderType uint8

const (
	// ByteOrderTypeLittleEndian represents the little-endian byte order type.
	ByteOrderTypeLittleEndian ByteOrderType = 0
	// ByteOrderTypeBigEndian represents the big-endian byte order type.
	ByteOrderTypeBigEndian ByteOrderType = 1
)

// ByteOrder provides interface for endianness.
type ByteOrder interface {
	Type() ByteOrderType
}

// LittleEndian represents the little-endian byte order.
type LittleEndian struct {
}

// Type returns the byte order type.
func (bo *LittleEndian) Type() ByteOrderType {
	return ByteOrderTypeLittleEndian
}

// BigEndian represents the big-endian byte order.
type BigEndian struct {
}

// Type returns the byte order type.
func (bo *BigEndian) Type() ByteOrderType {
	return ByteOrderTypeBigEndian
}
