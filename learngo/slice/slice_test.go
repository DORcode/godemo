package slice_test

import "testing"


func TestSliceInit(t *testing.T) {
   s := []int{1,2,3,4,5}
   t.Log(s, len(s), cap(s))

   s = append(s, 6)

   t.Log(len(s), cap(s))

   var s2 [] int

   s2 = append(s2, 1)

   t.Log(s2, len(s2), cap(s2))


   s3 := make([]int, 3, 10)

   t.Log(s3, len(s3), cap(s3))
}


func TestSliceShareMemory(t *testing.T) {
   year := []string{"J", "F", "M"}

   Q := year[2:]
   t.Log(Q)

   Q[0] = "O"

   t.Log(year)
}

func TestSliceComp(t *testing.T) {
   s := []int{1,2,3}
   s2 := []int{1,2,3}

   t.Log(s==s2)
}
