package main

import "github.com/dorcode/learngo/pkg2"
import "github.com/dorcode/learngo/pkg"

import "fmt"
import "os"

func init() {
   fmt.Println("打印")
}

func main() {
   if len(os.Args) > 1 {
       fmt.Println(os.Args[1])
   }	
   pkg.A()
   pkg.B()
   pkg2.B()
   pkg2.A()

   fmt.Println("hello world")


}
