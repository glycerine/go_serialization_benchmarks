package goserbench

import (
	"time"
)

//go:generate msgp -o msgp_gen.go -io=false -tests=false

//go:generate zebrapack -fast -fast-strings -o zebrapack_gen.go -io=false -tests=false

type A struct {
	Name     string    `zid:"0"`
	BirthDay time.Time `zid:"1"`
	Phone    string    `zid:"2"`
	Siblings int       `zid:"3"`
	Spouse   bool      `zid:"4"`
	Money    float64   `zid:"5"`
}
