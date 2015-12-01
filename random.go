// package random provides functions and methods useful for generating randomized short (16 bytes or less) values.
package random

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"time"

	mathrand "math/rand"
)

// Random is meant be used as a pointer field on a struct. Leaving the
// instance as a nil reference will cause any calls on the *Random to forward
// to the corresponding standalone functions in this package. This is meant
// to be the behavior used in production. In testing, set the field to a non-nil
// instance of a *Random to provide deterministic values based on the provided
// length parameter.
type Random struct{ Calls int }

func (this *Random) Bytes(length byte) []byte {
	if this == nil {
		return Bytes(length)
	} else {
		this.Calls++
		return _bytes(length)
	}
}
func (this *Random) Base64(length byte) string {
	if this == nil {
		return Base64(length)
	} else {
		this.Calls++
		return _base64(length)
	}
}
func (this *Random) Base62(length byte) string {
	if this == nil {
		return Base62(length)
	} else {
		this.Calls++
		return _base62(length)
	}
}
func (this *Random) Uint32(min, max uint32) uint32 {
	if min > max {
		temp := max
		max = min
		min = temp
	}

	if this == nil {
		return Uint32(min, max)
	} else {
		this.Calls++
		return _uint32(min+uint32(this.Calls)-1, max)
	}
}
func (this *Random) Hex(length byte) string {
	if this == nil {
		return Hex(length)
	} else {
		this.Calls++
		return _hex(length)
	}
}
func (this *Random) GUID() []byte {
	if this == nil {
		return GUID()
	} else {
		this.Calls++
		return _guid()
	}
}
func (this *Random) GUIDString() string {
	if this == nil {
		return GUIDString()
	} else {
		this.Calls++
		return _guidString()
	}
}

///////////////////////////////////////////////////////////////////////////////

// Production-grade randomized function:

func Bytes(length byte) []byte {
	all := make([]byte, length)
	rand.Read(all)
	return all
}
func Base64(length byte) string {
	return base64.StdEncoding.EncodeToString(Bytes(length))
}

func Uint32(min, max uint32) uint32 {
	if value, err := rand.Int(rand.Reader, big.NewInt(int64(max-min+1))); err == nil {
		return uint32(value.Int64()) + min
	} else {
		return min + uint32(mathrand.Intn(int(max-min+1))) // not enough entropy, use a pseudo-random number
	}
}

// Base62 removes the '+', '/', and '=' characters from a base64 encoded string. This is useful for
// generating url-safe values.
func Base62(length byte) string {
	value := Base64(length)
	value = strings.Replace(value, "+", "3", -1)
	value = strings.Replace(value, "/", "7", -1)
	value = strings.Replace(value, "=", "1", -1) // only necessary when length % 3 != 0
	return value
}
func Hex(length byte) string {
	raw := Bytes(length)
	return hex.EncodeToString(raw)
}
func GUID() []byte { return Bytes(16) }
func GUIDString() string {
	encoded := Hex(16)
	return format(encoded)
}

///////////////////////////////////////////////////////////////////////////////

// Faked functions:

func _bytes(length byte) []byte {
	buffer := make([]byte, length)
	for x := byte(0); x < length; x++ {
		buffer[x] = length
	}
	return buffer
}
func _base64(length byte) string {
	return base64.StdEncoding.EncodeToString(_bytes(length))
}
func _base62(length byte) string {
	value := _base64(length)
	value = strings.Replace(value, "+", "3", -1)
	value = strings.Replace(value, "/", "7", -1)
	value = strings.Replace(value, "=", "1", -1) // only necessary when length % 3 != 0
	return value
}
func _uint32(min, max uint32) uint32 {
	random := uint32(3214)

	if random < min {
		return min
	} else if random > max {
		return max
	}

	return random
}
func _hex(length byte) string {
	raw := _bytes(length)
	return hex.EncodeToString(raw)
}
func _guid() []byte { return _bytes(16) }
func _guidString() string {
	encoded := _hex(16)
	return format(encoded)
}

///////////////////////////////////////////////////////////////////////////////

func format(guid string) string {
	return fmt.Sprintf("%s-%s-%s-%s-%s",
		guid[:8],
		guid[8:12],
		guid[12:16],
		guid[16:20],
		guid[20:])
}

func init() {
	mathrand.Seed(time.Now().UnixNano())
}
