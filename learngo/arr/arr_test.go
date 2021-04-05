package arr_test

import "testing"

const(
 R = 1 << iota
 W
 E
)

func TestComparaArray(t *testing.T) {
   var a [3]int
   a[0] = 1
   a[1] = 2
   a[2] = 3


   b := [4]int{1, 2, 3, 4 }
   c := [...]int{1,2,3,4}
   d := [...]int{1,3,4,5}
   e := [...]int{1,2,3,4,5}

   t.Log(e)

   t.Log(b==c)
   t.Log(b==d)

   t.Log(a)
}

func TestBitClear(t *testing.T) {
   a := 7
   a = a&^R
   t.Log(a&R == R, a&W == W, a&E == E)

   a = a &^ E
   t.Log(a&R == R, a&W == W, a&E == E)


}
