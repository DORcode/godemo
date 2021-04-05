package extension

import "fmt"
import "testing"

type Pet struct {

}

func (p *Pet) Speak() {
    fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
   p.Speak()
   fmt.Println(host)
}

type Dog struct {
   // 扩展	
   p *Pet
}

func (d *Dog) Speak() {
   fmt.Print("dog")
   d.p.Speak()
}

func (d *Dog) SpeakTo(host string) {
   d.p.SpeakTo(host)
}

type Cat struct {	
   Pet
}

func TestDog(t *testing.T) {
   dog := new(Dog)
   dog.SpeakTo("Chao")

   cat := new(Cat)

   cat.SpeakTo("He")


}
