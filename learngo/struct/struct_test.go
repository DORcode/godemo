package struct_test

import "testing"

import "fmt"


type Employee struct {
   Id string
   Name string
   Age int
 
}

func (e *Employee) String() string {
   return fmt.Sprintf("%s%s%d", e.Id, e.Name, e.Age)
}

func TestStructInit(t *testing.T) {
   e := Employee {"1", "李小爽", 30}

   t.Log(e)

   el := Employee{Name: "张光", Age: 50}

   t.Log(el)

   e3 := new(Employee)

   e3.Name = "红"
   t.Log(e3)
   
   t.Logf("e %T", e3)
   
   t.Log(e3.String())

}
