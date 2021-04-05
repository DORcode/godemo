package string_test

import "testing"
import "strings"
import "strconv"

func TestString(t *testing.T) {
  s := "中"
  t.Log(len(s))
  c:= []rune(s)
  t.Log(c)

  t.Log(len(c))

  t.Logf("中 unicode %x", c[0])
  t.Logf("中 utf-8 %x", s)

  str := "中华人民共和国"

  for _, v := range str {
     t.Logf("%[1]c %[1]d", v)
  }
}

func TestStringFunc(t *testing.T) {
   s := "a,b,c"
   ss := strings.Split(s, ",")
   t.Log(ss)

   t.Log(strings.Join(ss, "-"))
}

func TestStrconv(t *testing.T) {
   s := strconv.Itoa(10)
   t.Log(s)


   if i, err := strconv.Atoi("10"); err == nil {
      t.Log(i + 50)
   }
}
