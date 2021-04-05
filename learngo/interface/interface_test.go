package interface_test

import (
   "testing"
   "fmt"
 )


 type Progrommer interface {
     WriteHello() string
 
 }

 type GoProgrammer struct {
 
 }

 func (g *GoProgrammer) WriteHello() string {
     return "Hello World"

 }

 func TestGo(t *testing.T) {
    g := &GoProgrammer{}

    t.Log(g.WriteHello())
 
 }


 func DoSomething(p interface{}) {
    if i, ok := p.(int); ok {
       fmt.Println("Int", i)
       return
    } 

    if s, ok := p.(string); ok {
       fmt.Println("str", s)
       return
    }
 }

 func TestEmptyInterface(t *testing.T) {
    DoSomething(1)
    DoSomething("1")
 }
