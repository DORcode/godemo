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
