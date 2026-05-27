package structuredstream

import (
	"encoding/binary"
	"io"
	"time"
)

// Reader wraps a reader with calls for reading binary data types
// If any error is encountered, all subsequent calls will fail
// error checking must be done in a separate call to error
type Reader struct {
	r         io.Reader
	err       error
	byteOrder binary.ByteOrder
}

// NewReader returns a new reader
func NewReader(r io.Reader) *Reader { _ = "STUB: not implemented"; return nil }

// Read takes a pointer to a data type (e.g. uint16, int64, []byte) and reads
// data from the wrapped reader, and advances the reader offset to the next value.
func (s *Reader) Read(x interface{}) { _ = "STUB: not implemented"; return }

// ReadInt8 returns a int8
func (s *Reader) ReadInt8() int8 { _ = "STUB: not implemented"; return 0 }

// ReadUint8 returns a uint8
func (s *Reader) ReadUint8() uint8 { _ = "STUB: not implemented"; return 0 }

// ReadInt16 returns an int16
func (s *Reader) ReadInt16() int16 { _ = "STUB: not implemented"; return 0 }

// ReadUint16 returns a uint16
func (s *Reader) ReadUint16() uint16 { _ = "STUB: not implemented"; return 0 }

// ReadInt32 returns an int32
func (s *Reader) ReadInt32() int32 { _ = "STUB: not implemented"; return 0 }

// ReadUint32 retursn a uint32
func (s *Reader) ReadUint32() uint32 { _ = "STUB: not implemented"; return 0 }

// ReadInt64 retursn a int64
func (s *Reader) ReadInt64() int64 { _ = "STUB: not implemented"; return 0 }

// ReadUint64 retursn a uint64
func (s *Reader) ReadUint64() uint64 { _ = "STUB: not implemented"; return 0 }

// ReadFloat64 retursn a float64
func (s *Reader) ReadFloat64() float64 { _ = "STUB: not implemented"; return 0 }

// ReadBytes returns l many bytes
func (s *Reader) ReadBytes(l int) []byte { _ = "STUB: not implemented"; return nil }

// ReadUint16PrefixedBytes first reads a uint16, then reads that many following bytes
func (s *Reader) ReadUint16PrefixedBytes() []byte { _ = "STUB: not implemented"; return nil }

// ReadUint16PrefixedString first reads a uint16, then reads that many following chars
func (s *Reader) ReadUint16PrefixedString() string { _ = "STUB: not implemented"; return "" }

// ReadUnixTime64UTC reads a uint64 representing unix epoch time in UTC and converts it to a time.time
func (s *Reader) ReadUnixTime64UTC() time.Time {
	_ = "STUB: not implemented"
	return *

	// can't use time.Unix which assumes timezone is local
	new(time.Time)
}

// Error returns the last encountered error
func (s *Reader) Error() error { _ = "STUB: not implemented"; return nil }
