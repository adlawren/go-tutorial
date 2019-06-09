package main

import (
	"fmt"
	"math"
	"strconv"
	"sync"
	"time"
)

// Global variable declaration
var (
	TestBool bool = false
	TestInt  int  = 10
)

const (
	ConstInt = 20
)

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

func sliceSum(slic []int, c chan int) {
	sum := 0
	for i := 0; i < len(slic); i++ {
		sum += slic[i]
	}

	c <- sum
}

func fib(n int, ch chan int) {
	n0, n1 := 0, 1
	for i := 0; i < n; i++ {
		ch <- n0
		n0, n1 = n1, n1+n0
	}

	close(ch)
}

func fib2(ch1 chan int, ch2 chan int) {
	n0, n1 := 0, 1
	for {
		select {
		case ch1 <- n0:
			n0, n1 = n1, n1+n0
		case <-ch2:
			fmt.Println("Exiting fib2")
			return
		}
	}
}

func wait(ch chan int) {
	for {
		select {
		case <-ch:
			fmt.Println("Exiting wait")
			return
		default:
			fmt.Println("Waiting..")
		}
	}
}

type GlobalIntStruct struct {
	v   int
	mut sync.Mutex
}

func (gis *GlobalIntStruct) inc() {
	defer gis.mut.Unlock()
	gis.mut.Lock()
	gis.v += 1
}

func main() {
	//// Packages, variables & functions
	// Functions
	{
		s1, s2 := swap("a", "b")
		fmt.Printf("s1: %v\n", s1)
		fmt.Printf("s2: %v\n", s2)

		x, y := increment(1, 2)
		fmt.Printf("x: %v\n", x)
		fmt.Printf("y: %v\n", y)
	}

	// Variables
	{
		var val int = 9
		fmt.Printf("val: %v\n", val)

		val2 := 8.0
		fmt.Printf("val2: %T\n", val2)
	}

	// Untyped constant casting
	fmt.Printf("needFloat: %v\n", needFloat(ConstInt))
	fmt.Printf("needInt: %v\n", needInt(ConstInt))

	//// Flow control
	// For loops
	{
		for v := 10; v > 0; v -= 1 {
			fmt.Printf("Ordinary for loop: %v\n", v)
		}

		val := 20
		for val > 10 {
			fmt.Printf("Single-clause for loop: %v\n", val)
			val -= 1
		}

		val2 := 0
		for {
			if val2 > 10 {
				break
			}
			fmt.Printf("Infinite for loop: %v\n", val2)
			val2 += 1
		}

	}

	// If statement
	{
		val := 15
		if v := 1; val < 10 {
			fmt.Printf("If val < 10: %v\n", v)
		} else if val > 0 {
			fmt.Printf("If val > 0: %v\n", v+1)
		}
	}

	// Switch statements
	{
		val := 5
		switch val {
		case 0:
			fmt.Println("Nope")
		case 5:
			fmt.Printf("%v\n", val)
		default:
			fmt.Println("No")
		}

		val2 := 99
		switch {
		case val2 > 100:
			fmt.Println("Not here")
		case val2 > 50:
			fmt.Println("Here")
		default:
			fmt.Println("Or here")
		}
	}

	// Defer example
	containsDefers(0)

	//// Structs, slices & maps
	// Pointers & references
	{
		var p *int
		val := 10
		p = &val
		*p = 21
		fmt.Printf("Dereferenced pointer: %v\n", *p)
	}

	// Struct initializations
	{
		fmt.Println("Struct initializations:")
		fmt.Println(Vertex{1, 2})
		fmt.Println(Vertex{Y: 1, X: 2})
		fmt.Println(Vertex{})
		fmt.Println(Vertex{}.X)
	}

	// Struct pointers
	{
		var p *Vertex
		s := Vertex{3, 4}
		p = &s
		fmt.Printf("Dereferenced pointer field: %v\n", (*p).X)
		fmt.Printf("Pointer field shorthand: %v\n", p.X)
	}

	// Arrays, slices & range loops
	{
		arr := [5]int{1, 2, 3, 4, 5}
		fmt.Printf("Array: %v\n", arr)

		slic := arr[1:3]
		slic[0] = 10
		fmt.Printf("Slice: %v\n", slic)
		fmt.Printf("Unaltered array: %v\n", arr)

		slic2 := []int{10, 11, 12}
		fmt.Printf("Inline slice: %v\n", slic2)

		slic3 := arr[1:]
		fmt.Printf("Implicit upper-bound slice: %v\n", slic3)

		slic4 := arr[:3]
		fmt.Printf("Implicit lower-bound slice: %v\n", slic4)
		fmt.Printf("Slice len, from array: %v, cap: %v\n",
			len(slic4), cap(slic4))
		slic4 = slic4[:5]
		fmt.Printf("Slice len, from other slice: %v, cap: %v\n",
			len(slic4), cap(slic4))

		slic5 := make([]int, 3, 5)
		fmt.Printf("Make-ed slice len: %v, cap: %v\n",
			len(slic5), cap(slic5))

		slic6 := [][]int{
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]int{7, 8, 9},
		}
		fmt.Printf("2-d slice: %v\n", slic6)

		slic7 := append(slic6, []int{10, 11, 12})
		fmt.Printf("Slice after append: %v\n", slic7)

		for i, v := range arr {
			fmt.Printf("Ordinary range: index: %v, value: %v\n", i, v)
		}
		for _, v := range arr {
			fmt.Printf("Range without index: value: %v\n", v)
		}
		for i, _ := range arr {
			fmt.Printf("Range without value: index: %v\n", i)
		}
		for i := range arr {
			fmt.Printf("Range without value placeholder: index: %v\n", i)
		}

	}

	// Map
	var m = map[int]string{
		0: "str 1",
		1: "str 2",
		2: "str 3",
	}
	fmt.Printf("Map: %v\n", m)

	m[1] = "str 2 updated"
	delete(m, 2)
	fmt.Printf("Updated map: %v\n", m)

	elem, exists := m[10]
	fmt.Printf("Map elem: %v, exists: %v\n", elem, exists)

	// Function value, closure
	fn2 := func(x, y float64) float64 {
		return x * y
	}
	fmt.Printf("fn1: %v\n", fn1(fn2))
	fmt.Printf("fn2: %v\n", fn2(2, 3))

	clo := createClosure()
	fmt.Printf("clo(1): %v\n", clo(1))
	fmt.Printf("clo(2): %v\n", clo(2))

	//// Methods & interfaces
	// Methods
	{
		vertex := Vertex{X: 1, Y: 1}
		fmt.Printf("Vertex value, pointer method: %v\n",
			vertex.myPointerMethod(2))
		fmt.Printf("Vertex value, value method: %v\n",
			vertex.myValueMethod(2))

		p := &vertex
		fmt.Printf("Vertex pointer, pointer method: %v\n",
			p.myPointerMethod(2))
		fmt.Printf("Vertex pointer, value method: %v\n",
			p.myValueMethod(2))

		var mf MyFloat = 1.0
		fmt.Printf("MyFloat method: %v\n", mf.myFloatMethod(2.5))
	}

	// Interfaces
	{
		vertex := Vertex{X: 3, Y: 4}
		fmt.Printf("Vertex length: %v\n", vertex.Length())

		var i1 MyInterface = &vertex
		fmt.Printf("Vector interface: %v\n", i1.Length())
		fmt.Printf("Interface value: %v, type: %T\n", i1, i1)

		var vertex2 *Vertex
		var i2 MyInterface = vertex2
		fmt.Printf("Uninitialized vector interface: %v\n", i2.Length())
		fmt.Printf("Uninitialized interface value: %v, type: %T\n", i2, i2)

		var i3 MyInterface = nil
		//fmt.Println(i3.Length()) // Throws segfault
		fmt.Printf("Nil interface value: %v, type: %T\n", i3, i3)

		var i4 MyEmptyInterface
		i4 = 4
		fmt.Printf("Int interface value: %v, type: %T\n", i4, i4)
		i4 = "test"
		fmt.Printf("String interface value: %v, type: %T\n", i4, i4)

		i4 = 10
		i4IntValue := i4.(int)
		fmt.Printf("Type-cast int interface int value: %v\n", i4IntValue)
		//i4StringValue := i4.(string) // Throws panic
		i4StringValue, ok := i4.(string)
		fmt.Printf("Type-cast int interface string value: %v, ok: %v\n", i4StringValue, ok)

		//i4 = "test again"
		//i4 = 3.0
		switch i4Value := i4.(type) {
		case int:
			fmt.Printf("Type switch int value: %v, type: %T\n", i4Value, i4Value)
		case string:
			fmt.Printf("Type switch string value: %v, type: %T\n", i4Value, i4Value)
		default:
			fmt.Printf("Type switch default value: %v, type: %T\n", i4Value, i4Value)
		}
	}

	// Error pattern example
	i, err := strconv.Atoi("42")
	//i, err := strconv.Atoi("test")
	if err != nil {
		fmt.Println("Error occurred")
	} else {
		fmt.Printf("Atoi: %v\n", i)
	}

	//// Concurrency
	// Channel example
	{
		var slic [30]int
		for i := 0; i < len(slic); i++ {
			slic[i] = i
		}

		ch := make(chan int)
		go sliceSum(slic[:len(slic)/2], ch)
		go sliceSum(slic[len(slic)/2:], ch)
		sum1 := <-ch
		sum2 := <-ch
		fmt.Printf("slice sum: %v\n", sum1+sum2)
	}

	// Buffered channel
	{
		ch := make(chan int, 2)
		ch <- 1
		ch <- 2
		fmt.Printf("First value from channel: %v\n", <-ch)
		fmt.Printf("Second value from channel: %v\n", <-ch)
	}

	// Range & close
	{
		ch := make(chan int)
		go fib(10, ch)
		for val := range ch {
			fmt.Printf("Next fib: %v\n", val)
		}

		_, test := <-ch
		fmt.Printf("Channel open: %v\n", test)
	}

	// Select
	{
		ch1 := make(chan int)
		ch2 := make(chan int)
		go fib2(ch1, ch2)
		for i := 0; i < 10; i++ {
			fmt.Printf("Next fib: %v\n", <-ch1)
		}
		ch2 <- 0
	}

	// Select w/ default
	{
		ch := make(chan int)
		go wait(ch)
		time.Sleep(15 * time.Microsecond)
		ch <- 0
	}

	// Mutex
	{
		var gis GlobalIntStruct
		for i := 0; i < 10; i++ {
			go gis.inc()
		}

		for gis.v != 10 {
		}

		fmt.Printf("Global int: %v\n", gis.v)
	}
}
