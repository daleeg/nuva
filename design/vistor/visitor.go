
package main
import (
    "fmt"
    "encoding/json"
)

type Teacher struct {
    Name string
}

func NewTeacher(name string) *Teacher {
    return & Teacher {
        Name: name,
    }
}

func (t * Teacher) Data() map[string]string {
    result := make(map[string]string)
    result["name"] = t.Name
    return result
}

func (t * Teacher) DisplayJson(v VisitorInterface) {
    data := t.Data()
    fmt.Println(v.Tojson(data))
}

type VisitorInterface interface {
    Tojson(map[string]string) string
}

type Visitor struct {
    Name string
}

func NewVisitor(name string) *Visitor {
    return & Visitor {
        Name: name,
    }
}

func (v *Visitor) Tojson(data map[string]string) string {
    data ["visitor"] = v.Name
    jsonStr, err := json.Marshal(data)
    if err != nil {
        fmt.Println("MapToJsonDemo err: ", err)
        return ""
    }
    return string(jsonStr) 
}

 
func TestVisitor() {
    t := NewTeacher("lao wang")
    v := NewVisitor("json")
    t.DisplayJson(v)
}
    
func main() {
    TestVisitor()
}
