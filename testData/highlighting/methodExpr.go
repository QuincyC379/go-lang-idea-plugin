package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func main() {
	v := Vertex{3, 4}
	fmt.Print(v.Abs())
}

func (v *Vertex) Abs() float64 {
	return (*v).X * v.X + v.Y * v.Y
}

type T struct {
	a int
}
func (tv  T) Mv(a int) int { return 0 }  // value receiver
func (tp *T) Mp(f float32) float32 { return 1 }  // pointer receiver

var t T

func _() {
	t.Mv(7)
	T.Mv<error>(t, 7)</error> // shouldn't be a error
	(T).Mv(t, 7)
	f1 := T.Mv; f1(t, 7)
	f2 := (T).Mv; f2(t, 7)
	fmt.Println((*T).Mp)
	fmt.Println((*T).Mv)
}

type Type interface {
    Size() int64
}

type ArrayType struct {
    Type Type
}

func _(t *Type) {
    at, _ := (*t).(*ArrayType)
    println(at.Type)
    println(at.Type.Size)
}