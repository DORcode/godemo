package main

import "github.com/dorcode/learngo/pkg2"
import "github.com/dorcode/learngo/pkg"

import "fmt"

func init() {
   fmt.Println("打印")
}

func main() {
   pkg.A()
   pkg.B()
   pkg2.B()
   pkg2.A()

   fmt.Println("hello world")


}
