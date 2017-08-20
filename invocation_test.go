package bench_test

import "testing"

type InterfaceA interface {
	A()
}

type A struct {
	i int
}

func (a *A) A() {
	a.i++
}

var i int

func foo() {
	i++
}

func BenchmarkFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		foo()
	}
}

func BenchmarkClosure(b *testing.B) {
	var i int
	f := func() {
		i++
	}
	for i := 0; i < b.N; i++ {
		f()
	}
}

func BenchmarkConcrete(b *testing.B) {
	a := new(A)
	for i := 0; i < b.N; i++ {
		a.A()
	}
}

func BenchmarkInterface(b *testing.B) {
	var a InterfaceA = new(A)
	for i := 0; i < b.N; i++ {
		a.A()
	}
}

func BenchmarkConcreteFunc(b *testing.B) {
	a := new(A)
	f := a.A
	for i := 0; i < b.N; i++ {
		f()
	}
}

func BenchmarkInterfaceFunc(b *testing.B) {
	var a InterfaceA = new(A)
	f := a.A
	for i := 0; i < b.N; i++ {
		f()
	}
}
