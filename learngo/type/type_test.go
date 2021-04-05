package type_test

import "testing"

type MyInt int64

func TestImpicit(t *testing.T) {
   var a int32 = 1
   var b int64

   b = int64(a)

   var c MyInt

   c = MyInt(a)

   t.Log(a, b, c, string(a))
}

func TestPoint(t *testing.T) {
    a := 1
    aPtr := &a
    *aPtr = *aPtr +1

    t.Log(a, *aPtr)

}


func TestString(t *testing.T) {
   var s string
   //是空字符串，不是null
   t.Log(s)
   t.Log(len(s))
}
