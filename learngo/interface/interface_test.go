package interface_test

import (
   "testing"
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
