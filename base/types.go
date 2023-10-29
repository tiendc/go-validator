package base

import (
	"time"

	"github.com/tiendc/gofn"
)

type Int interface {
	gofn.Int | gofn.IntEx
}

type UInt interface {
	gofn.UInt | gofn.UIntEx
}

type Float interface {
	gofn.Float | gofn.FloatEx
}

type Number interface {
	Int | UInt | Float
}

type String interface {
	~string
}

type Time interface {
	Compare(time.Time) int
	IsZero() bool
}
