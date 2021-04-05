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

func TestArrInit(t *testing.T) {
    var arr [3]int
    arr[0] = 1
    t.Log(arr)

    var arr2 [5]string
    arr2[0] = "a" 

    t.Log(arr2)


    arr3 := [3]int{1, 2, 3}
    t.Log(arr3)

    arr4 := [...]string{"a", "b", "c", "d", "e"}
    t.Log(arr4)
}

func TestArrLoop(t *testing.T) {
   arr := [...]int{1,2,3,4,5,6,7}
   for i:=0;i<len(arr);i++ {
      t.Log(arr[i])
   }

   for i,v := range arr {
      t.Log(i,v)
   }
}

func TestArrSlice(t *testing.T) {
   arr := [...]string{"a", "b", "c", "d", "e", "f"}
   t.Log(arr)
   
   a := arr[:3]

   t.Log(a)
   b := arr[2:]
   t.Log(b)
   t.Log(arr[0:len(arr)])
}
