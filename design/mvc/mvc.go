
package main
import (
    "fmt"
    "encoding/json"
)

type Student struct {
    Name string
}

func NewStudent(name string) *Student {
    return & Student {
        Name: name,
    }
}

func (s * Student) Data() map[string]string {
    result := make(map[string]string)
    result["name"] = s.Name
    return result
}


type View struct {
    Name string
}

func NewView(name string) *View {
    return & View{
        Name: name,
    }
}

func (v *View) DisplayJson(data map[string]string) {
    data ["visitor"] = v.Name
    jsonStr, err := json.Marshal(data)
    if err != nil {
        fmt.Println("MapToJsonDemo err: ", err)
        return
    }
    fmt.Println(string(jsonStr))
}

type Control struct {
    M *Student
    V *View
}

func NewControl(s * Student, v *View) *Control {
    return & Control{
        M: s,
        V: v,
    }
}

func (c *Control) DisplayStudentJson() {
    c.V.DisplayJson(c.M.Data())
}

 
func TestMVC() {
    s := NewStudent("xiao ming")
    v := NewView("json")
    c := NewControl(s, v)
    c.DisplayStudentJson()
}
    
func main() {
    TestMVC()
}
