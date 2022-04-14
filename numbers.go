// Byteman
// For the full copyright and license information, please view the LICENSE.txt file.

package byteman

import (
	"encoding/binary"
)

const (
	// IntSize represents the supported integer value size.
	IntSize = 32 << (^uint(0) >> 63) // 32 or 64
)

// FromUint returns a byte slice by the given Uint and byte order (endianness).
func FromUint(number interface{}, bo ByteOrder) []byte {
	switch v := number.(type) {
	case uint:
		b := make([]byte, IntSize/8)
		if bo.Type() == ByteOrderTypeBigEndian {
			if IntSize == 64 {
				binary.BigEndian.PutUint64(b, uint64(v))
			} else {
				binary.BigEndian.PutUint32(b, uint32(v))
			}
		} else {
			if IntSize == 64 {
				binary.LittleEndian.PutUint64(b, uint64(v))
			} else {
				binary.LittleEndian.PutUint32(b, uint32(v))
			}
		}
		return b
	case uint8:
		b := make([]byte, 1)
		b[0] = v
		return b
	case uint16:
		b := make([]byte, 2)
		if bo.Type() == ByteOrderTypeBigEndian {
			binary.BigEndian.PutUint16(b, v)
		} else {
			binary.LittleEndian.PutUint16(b, v)
		}
		return b
	case uint32:
		b := make([]byte, 4)
		if bo.Type() == ByteOrderTypeBigEndian {
			binary.BigEndian.PutUint32(b, v)
		} else {
			binary.LittleEndian.PutUint32(b, v)
		}
		return b
	case uint64:
		b := make([]byte, 8)
		if bo.Type() == ByteOrderTypeBigEndian {
			binary.BigEndian.PutUint64(b, v)
		} else {
			binary.LittleEndian.PutUint64(b, v)
		}
		return b
	default:
		return nil
	}
}

// FromInt returns a byte slice by the given int and byte order (endianness).
func FromInt(number interface{}, bo ByteOrder) []byte {
	switch v := number.(type) {
	case int:
		return FromUint(uint(v), bo)
	case int8:
		return FromUint(uint8(v), bo)
	case int16:
		return FromUint(uint16(v), bo)
	case int32:
		return FromUint(uint32(v), bo)
	case int64:
		return FromUint(uint64(v), bo)
	default:
		return nil
	}
}

// Uint returns an uint value by the given byte slice and byte order (endianness).
func Uint(b []byte, bo ByteOrder) uint {
	if len(b) == 8 && IntSize == 64 {
		if bo.Type() == ByteOrderTypeBigEndian {
			return uint(binary.BigEndian.Uint64(b))
		}
		return uint(binary.LittleEndian.Uint64(b))
	} else if len(b) >= 4 {
		if bo.Type() == ByteOrderTypeBigEndian {
			return uint(binary.BigEndian.Uint32(b))
		}
		return uint(binary.LittleEndian.Uint32(b))
	}
	return 0
}

// Uint16 returns an uint16 value by the given byte slice and byte order (endianness).
func Uint16(b []byte, bo ByteOrder) uint16 {
	if len(b) == 2 {
		if bo.Type() == ByteOrderTypeBigEndian {
			return binary.BigEndian.Uint16(b)
		}
		return binary.LittleEndian.Uint16(b)
	}
	return 0
}

// Uint32 returns an uint32 value by the given byte slice and byte order (endianness).
func Uint32(b []byte, bo ByteOrder) uint32 {
	if len(b) == 4 {
		if bo.Type() == ByteOrderTypeBigEndian {
			return binary.BigEndian.Uint32(b)
		}
		return binary.LittleEndian.Uint32(b)
	}
	return 0
}

// Uint64 returns an uint64 value by the given byte slice and byte order (endianness).
func Uint64(b []byte, bo ByteOrder) uint64 {
	if len(b) == 8 {
		if bo.Type() == ByteOrderTypeBigEndian {
			return binary.BigEndian.Uint64(b)
		}
		return binary.LittleEndian.Uint64(b)
	}
	return 0
}

// Int returns an int value by the given byte slice and byte order (endianness).
func Int(b []byte, bo ByteOrder) int {
	if len(b) == 8 && IntSize == 64 {
		if bo.Type() == ByteOrderTypeBigEndian {
			return int(binary.BigEndian.Uint64(b))
		}
		return int(binary.LittleEndian.Uint64(b))
	} else if len(b) >= 4 {
		if bo.Type() == ByteOrderTypeBigEndian {
			return int(binary.BigEndian.Uint32(b))
		}
		return int(binary.LittleEndian.Uint32(b))
	}
	return 0
}

// Int16 returns an int16 value by the given byte slice and byte order (endianness).
func Int16(b []byte, bo ByteOrder) int16 {
	if len(b) == 2 {
		if bo.Type() == ByteOrderTypeBigEndian {
			return int16(binary.BigEndian.Uint16(b))
		}
		return int16(binary.LittleEndian.Uint16(b))
	}
	return 0
}

// Int32 returns an int32 value by the given byte slice and byte order (endianness).
func Int32(b []byte, bo ByteOrder) int32 {
	if len(b) == 4 {
		if bo.Type() == ByteOrderTypeBigEndian {
			return int32(binary.BigEndian.Uint32(b))
		}
		return int32(binary.LittleEndian.Uint32(b))
	}
	return 0
}

// Int64 returns an int64 value by the given byte slice and byte order (endianness).
func Int64(b []byte, bo ByteOrder) int64 {
	if len(b) == 8 {
		if bo.Type() == ByteOrderTypeBigEndian {
			return int64(binary.BigEndian.Uint64(b))
		}
		return int64(binary.LittleEndian.Uint64(b))
	}
	return 0
}
