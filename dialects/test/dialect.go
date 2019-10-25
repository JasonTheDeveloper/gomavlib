// Autogenerated with dialgen, do not edit.
//
// Generated from revision https://github.com/mavlink/mavlink/tree/8a0c0cc708741e248ba53ac51bae2d4774a3f81b
//
package test

import (
	"github.com/aler9/gomavlib"
)

// Dialect contains the dialect object that can be passed to the library.
var Dialect = dialect

// dialect is not exposed directly such that it is not displayed in godoc.
var dialect = gomavlib.MustDialect(3, []gomavlib.Message{
	// test.xml
	&MessageTestTypes{},
})

// test.xml

// Test all field types
type MessageTestTypes struct {
	// char
	C string
	// string
	S string `mavlen:"10"`
	// uint8_t
	U8 uint8
	// uint16_t
	U16 uint16
	// uint32_t
	U32 uint32
	// uint64_t
	U64 uint64
	// int8_t
	S8 int8
	// int16_t
	S16 int16
	// int32_t
	S32 int32
	// int64_t
	S64 int64
	// float
	F float32
	// double
	D float64
	// uint8_t_array
	U8Array [3]uint8
	// uint16_t_array
	U16Array [3]uint16
	// uint32_t_array
	U32Array [3]uint32
	// uint64_t_array
	U64Array [3]uint64
	// int8_t_array
	S8Array [3]int8
	// int16_t_array
	S16Array [3]int16
	// int32_t_array
	S32Array [3]int32
	// int64_t_array
	S64Array [3]int64
	// float_array
	FArray [3]float32
	// double_array
	DArray [3]float64
}

func (*MessageTestTypes) GetId() uint32 {
	return 0
}
