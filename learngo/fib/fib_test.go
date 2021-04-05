package fib

import "testing"
import "fmt"

func TestFibList(t *testing.T) {
   var a int = 1
   var b int = 1

   for i := 0; i < 5; i++ {
      fmt.Print(" ", b)
      tmp := a
      a = b
      b = tmp + a
   }
}


func TestExchange(t *testing.T) {
   a:=1
   b:=2
   tmp:=b
   a=b
   b=tmp
   t.Log(a,b)

   a,b=b,a
   t.Log(a,b)
}
