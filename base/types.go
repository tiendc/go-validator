package base

import (
	"time"

	"github.com/tiendc/gofn"
)

// Int interface of int types and int-derived types
type Int interface {
	gofn.Int | gofn.IntEx
}

// UInt interface of uint types and uint-derived types
type UInt interface {
	gofn.UInt | gofn.UIntEx
}

// Float interface of float types and float-derived types
type Float interface {
	gofn.Float | gofn.FloatEx
}

// Number interface of combined type of Int, UInt, and Float
type Number interface {
	Int | UInt | Float
}

// String interface of string type and string-derived types
type String interface {
	~string
}

// Time interface of time type
type Time interface {
	Compare(time.Time) int
	IsZero() bool
}
