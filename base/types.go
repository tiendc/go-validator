package base

import (
	"time"

	"github.com/tiendc/gofn"
)

// Int interface of int types and int-derived types
type Int gofn.IntExt

// UInt interface of uint types and uint-derived types
type UInt gofn.UIntExt

// Float interface of float types and float-derived types
type Float gofn.FloatExt

// Number interface of combined type of Int, UInt, and Float
type Number interface {
	Int | UInt | Float
}

// String interface of string type and string-derived types
type String gofn.StringExt

// Time interface of time type
type Time interface {
	Compare(time.Time) int
	IsZero() bool
}
