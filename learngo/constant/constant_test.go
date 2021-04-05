package constant

import "testing"

const(
   Monday = iota +1
   Tuesday
   Wednesday
)

const(
   Readable = 1 << iota
   Writeable
   Executeable
)


func TestConstantTry(t *testing.T) {
    t.Log(Monday, Wednesday)
}

func TestConstantTry2(t *testing.T) {
    a:=7
    t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executeable == Executeable)
}
