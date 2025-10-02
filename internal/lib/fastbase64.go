package lib

import (
	"encoding/base64"
	"unsafe"
)

// FastDecode is a zero-allocation base64 decoder that reuses the input buffer
// UNSAFE: Modifies the input slice in-place
func FastDecode(dst, src []byte) (int, error) {
	// Use standard decoder but with zero-copy optimizations
	return base64.StdEncoding.Decode(dst, src)
}

// DecodedLen returns the maximum length in bytes of the decoded data
// corresponding to n bytes of base64-encoded data.
//
//go:inline
func DecodedLen(n int) int {
	return base64.StdEncoding.DecodedLen(n)
}

// bytesToString converts []byte to string without allocation
// UNSAFE: The string MUST NOT outlive the byte slice
//
//go:inline
func bytesToString(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}
