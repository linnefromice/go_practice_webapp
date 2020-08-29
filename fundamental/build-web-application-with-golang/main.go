package main

import (
	"fmt"
	"math"
	"runtime"
	"errors"
)

func confirmError() {
	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Println(err)
	}
}

func confirmFlow() {
	x := 20
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than 10")
	}
	y := 0
Here:
	y++
	if y >= 5 {
		fmt.Printf("%d ... Finish\n", y)
	} else {
		fmt.Printf("%d ... Continue\n", y)
		goto Here
	}
}

type person struct {
	name string
	age int
}
func confirmStruct() {
	var P person
	P.name = "Asataxie"
	P.age = 25
	fmt.Printf("The person's name is %s, age is %d.\n", P.name, P.age)
}

type Rectangle struct {
	width, height float64
}
type Circle struct {
	radius float64
}
func (r Rectangle) area() float64 {
	return r.width * r.height
}
func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}
func confirmObject() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}

	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())
	fmt.Println("Area of c1 is: ", c1.area())
	fmt.Println("Area of c2 is: ", c2.area())
}

type Human struct {
	name string
	age int
	phone string
}
type Student struct {
	Human
	school string
}
type Employee struct {
	Human
	company string
}
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}
func confirmExtend() {
	mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := Employee{Human{"Sam", 45, "111-888-YYYY"}, "Golang Inc"}
	mark.SayHi()
	sam.SayHi()
}

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}
func confirmGoroutines() {
	fmt.Println("START: use Goroutines")
	go say("world")
	say("hello")
}

func main() {
	confirmError()
	confirmFlow()
	confirmStruct()
	confirmObject()
	confirmExtend()
	confirmGoroutines()
	// channels
	// Buffered Channels
	// Range&Close
	// Select
	// タイムアウト
	// runtime goroutine
	fmt.Printf("Hello, world or こんにちわ、世界")
}