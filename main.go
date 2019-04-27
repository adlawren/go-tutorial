package main

import (
	"fmt"
	"math"
	"strconv"
)

var (
	TestBool bool = false
	TestInt  int  = 10
)

const (
	ConstInt = 20
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func increment(x, y int) (a, b int) {
	a = x + 1
	b = y + 2
	return
}

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.001
}

func containsDefers(x int) {
	defer fmt.Printf("x + 2: %v\n", x+2)
	defer fmt.Printf("x + 1: %v\n", x+1)
	fmt.Printf("x: %v\n", x)
}

type Vertex struct {
	X, Y int
}

type MyFloat float64

func fn1(fn func(float64, float64) float64) float64 {
	return fn(1, 2)
}

func createClosure() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func (v *Vertex) myPointerMethod(s int) *Vertex {
	v.X *= s
	v.Y *= s
	return v
}

func (v Vertex) myValueMethod(s int) Vertex {
	v.X *= s
	v.Y *= s
	return v
}

func (f MyFloat) myFloatMethod(other float64) MyFloat {
	f2 := MyFloat(other)
	return f * f2
}

type MyInterface interface {
	Length() float64
}

func (v *Vertex) Length() float64 {
	if v == nil {
		var zeroFloat64 float64
		return zeroFloat64
	}

	xFloat := float64(v.X)
	yFloat := float64(v.Y)
	return math.Sqrt(xFloat*xFloat + yFloat*yFloat)
}

type MyEmptyInterface interface{}

func main() {
	fmt.Printf("%v\n", math.Sqrt(4))
	fmt.Printf("%v\n", add(2, 2))

	var val int = 9
	fmt.Printf("%v\n", val)

	val2 := 8.0
	fmt.Printf("%T\n", val2)

	fmt.Printf("%v\n", needFloat(ConstInt))
	fmt.Printf("%v\n", needInt(ConstInt))

	for v := 10; v > 0; v -= 1 {
		fmt.Printf("%v\n", v)
	}

	val3 := 20
	for val3 > 10 {
		fmt.Printf("%v\n", val3)
		val3 -= 1
	}

	val4 := 0
	for {
		if val4 > 10 {
			break
		}
		fmt.Printf("%v\n", val4)
		val4 += 1
	}

	if v := 1; val4 < 10 {
		fmt.Printf("%v\n", v)
	} else if val4 > 0 {
		fmt.Printf("%v\n", v+1)
	}

	val5 := 5
	switch val5 {
	case 0:
		fmt.Println("Nope")
	case 5:
		fmt.Printf("%v\n", val5)
	default:
		fmt.Println("No")
	}

	val6 := 99
	switch {
	case val6 > 100:
		fmt.Println("Not here")
	case val6 > 50:
		fmt.Println("Here")
	default:
		fmt.Println("Or here")
	}

	containsDefers(0)

	var p *int
	val7 := 10
	p = &val7
	*p = 21
	fmt.Println(*p)

	fmt.Println(Vertex{1, 2})
	fmt.Println(Vertex{Y: 1, X: 2})
	fmt.Println(Vertex{})
	fmt.Println(Vertex{}.X)

	var p2 *Vertex
	s := Vertex{3, 4}
	p2 = &s
	fmt.Println((*p2).X)
	fmt.Println(p2.X)

	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	slic := arr[1:3]
	slic[0] = 10
	fmt.Println(slic)
	fmt.Println(arr)
	slic2 := []int{10, 11, 12}
	fmt.Println(slic2)
	slic3 := arr[1:]
	fmt.Println(slic3)
	slic4 := arr[:3]
	fmt.Println(slic4)
	fmt.Printf("len: %v, cap: %v\n", len(slic4), cap(slic4))
	slic4 = slic4[:5]
	fmt.Printf("len: %v, cap: %v\n", len(slic4), cap(slic4))
	slic5 := make([]int, 3, 5)
	fmt.Printf("len: %v, cap: %v\n", len(slic5), cap(slic5))
	slic6 := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Println(slic6)
	slic7 := append(slic6, []int{10, 11, 12})
	fmt.Println(slic7)
	for i, v := range arr {
		fmt.Printf("index: %v, value: %v\n", i, v)
	}
	for _, v := range arr {
		fmt.Printf("value: %v\n", v)
	}
	for i, _ := range arr {
		fmt.Printf("index: %v\n", i)
	}
	for i := range arr {
		fmt.Printf("index: %v\n", i)
	}
	var m = map[int]string{
		0: "str 1",
		1: "str 2",
		2: "str 3",
	}
	fmt.Println(m)

	m[1] = "str 2 updated"
	delete(m, 2)
	fmt.Println(m)

	elem, exists := m[10]
	fmt.Printf("elem: %v, exists: %v\n", elem, exists)

	fn2 := func(x, y float64) float64 {
		return x * y
	}
	fmt.Println(fn1(fn2))
	fmt.Println(fn2(2, 3))

	clo := createClosure()
	fmt.Println(clo(1))
	fmt.Println(clo(2))

	vertex1 := Vertex{X: 1, Y: 1}
	fmt.Println(vertex1.myPointerMethod(2))
	fmt.Println(vertex1.myValueMethod(2))

	p3 := &vertex1
	fmt.Println(p3.myPointerMethod(2))
	fmt.Println(p3.myValueMethod(2))

	var mf MyFloat = 1.0
	fmt.Println(mf.myFloatMethod(2.5))

	vertex2 := Vertex{X: 3, Y: 4}
	fmt.Println(vertex2.Length())

	var i1 MyInterface = &vertex2
	fmt.Println(i1.Length())
	fmt.Printf("value: %v, type: %T\n", i1, i1)

	var vertex3 *Vertex
	var i2 MyInterface = vertex3
	fmt.Println(i2.Length())
	fmt.Printf("value: %v, type: %T\n", i2, i2)

	var i3 MyInterface = nil
	//fmt.Println(i3.Length()) // Throws segfault
	fmt.Printf("value: %v, type: %T\n", i3, i3)

	var i4 MyEmptyInterface
	i4 = 4
	fmt.Printf("value: %v, type: %T\n", i4, i4)
	i4 = "test"
	fmt.Printf("value: %v, type: %T\n", i4, i4)

	i4 = 10
	i4IntValue := i4.(int)
	fmt.Println(i4IntValue)
	//i4StringValue := i4.(string) // Throws panic
	i4StringValue, ok := i4.(string)
	fmt.Printf("value: %v, ok: %v\n", i4StringValue, ok)

	//i4 = "test again"
	//i4 = 3.0
	switch i4Value := i4.(type) {
	case int:
		fmt.Printf("value: %v, type: %T\n", i4Value, i4Value)
	case string:
		fmt.Printf("value: %v, type: %T\n", i4Value, i4Value)
	default:
		fmt.Printf("value: %v, type: %T\n", i4Value, i4Value)
	}

	i, err := strconv.Atoi("42")
	//i, err := strconv.Atoi("test")
	if err != nil {
		fmt.Println("Error occurred")
	} else {
		fmt.Println(i)
	}
}
