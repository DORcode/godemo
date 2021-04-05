package func_test

import "testing"
import "time"
import "fmt"

func returnMultiValues()(int, int, int) {
   return 10, 30,50
}

func timeSpend(inner func(op int) int) func(op int) int {
   return func(n int) int {
      start := time.Now()
      ret := inner(n)

      fmt.Println(time.Since(start).Seconds())

      return ret
   }
}

func TestFunc(t *testing.T) {
  a, b, c := returnMultiValues()

  t.Log(a, b, c)

  r := timeSpend(func(op int) int {return op * 10})(2)
  t.Log(r)
}

func Sum(ops ...int) int {
   sum := 0

   for _, v := range ops {
      sum += v
   }
   return sum

}

func TestVarParam(t *testing.T) {
   t.Log(Sum(1,2,3,4,5,6))
}

func TestDefer(t *testing.T) {
   defer func() {
      t.Log("Clear")
   }()

   t.Log("Start")

   panic("中断")



}

