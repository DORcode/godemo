package map_test

import "testing"


func TestMapInit(t *testing.T) {
   m := map[string]int{"one":1, "two":2,"three":3}
   t.Log(m)

   t.Log(m["one"])

   m2 := map[string]int{}

   m2["one"] = 1

   t.Log(m2)


   m3 := make(map[string]int, 20)

   t.Log(len(m3))

   t.Log(m["a"])

   for k,v := range m {
      t.Log(k,v)
   }
}

func TestNotExistingKey(t *testing.T) {
   m1 := map[int]int{}

   if v, ok := m1[3]; ok {
     t.Log(v)
   } else {
     t.Log("key is not")
   }

}

func TestMapFeature(t *testing.T) {
   m := map[int]func(op int)int{}

   m[1] = func(op int) int {return op}
   m[2] = func(op int) int {return op * op}
   m[3] = func(op int) int {return op * op * op}

   t.Log(m[2](2))

}

func TestMapForSet(t *testing.T) {
    set := map[int]bool{}

    set[1] = true

    if set[2] {
       t.Log("存在")
    } else {
       t.Log("不存在")
    }

    delete(set, 1)

    t.Log(set)

}

