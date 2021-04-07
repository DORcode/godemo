package err_test

import "testing"
import "errors"
import "os"

//var LessThanTwoError = errors.New("<2")

func GetFib(n int) ([]int64, error) {
   if n < 2 {
      return nil, errors.New("n should greater than 2")
   }

   if n > 93 {
      return nil, errors.New("n should less than 100")
   }
   arr := []int64{1, 1}

   for i := 2; i< n ; i++ {
      arr = append(arr, arr[i-2] + arr[i-1])
   }
   return arr, nil
}

func TestFib(t *testing.T) {
   if v, e := GetFib(90); e != nil {
      t.Error(e)
   } else {
      t.Log(v)
   }
}

// panic 用于不可恢复的错误
// panic 退出前会执行defer指定的内容

// os.Exit退出不用调用defer
// os.Exit不会打出堆栈信息

func TestPanicVxExit(t *testing.T) {
   defer func() {
     t.Log("defer")
     if err := recover(); err != nil {
        t.Log("recoverd from ", err)
     }

   }()
   t.Log("Start")
   panic(errors.New("Something wrong!"))
   os.Exit(-1)
}
