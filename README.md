# random
--
    import "github.com/smartystreets/random"

package random provides functions and methods useful for generating randomized
short (16 bytes or less) values.

## Usage

#### func  Base62

```go
func Base62(length byte) string
```
Base62 removes the '+', '/', and '=' characters from a base64 encoded string.
This is useful for generating url-safe values.

#### func  Base64

```go
func Base64(length byte) string
```

#### func  Bytes

```go
func Bytes(length byte) []byte
```

#### func  GUID

```go
func GUID() []byte
```

#### func  GUIDString

```go
func GUIDString() string
```

#### func  Hex

```go
func Hex(length byte) string
```

#### type Random

```go
type Random struct{ Calls int }
```

Random is meant be used as a pointer field on a struct. Leaving the instance as
a nil reference will cause any calls on the *Random to forward to the
corresponding standalone functions in this package. This is meant to be the
behavior used in production. In testing, set the field to a non-nil instance of
a *Random to provide deterministic values based on the provided length
parameter.

#### func (*Random) Base62

```go
func (this *Random) Base62(length byte) string
```

#### func (*Random) Base64

```go
func (this *Random) Base64(length byte) string
```

#### func (*Random) Bytes

```go
func (this *Random) Bytes(length byte) []byte
```

#### func (*Random) GUID

```go
func (this *Random) GUID() []byte
```

#### func (*Random) GUIDString

```go
func (this *Random) GUIDString() string
```

#### func (*Random) Hex

```go
func (this *Random) Hex(length byte) string
```
