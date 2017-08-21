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

func BenchmarkInvocationFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		foo()
	}
}

func BenchmarkInvocationClosure(b *testing.B) {
	var i int
	f := func() {
		i++
	}
	for i := 0; i < b.N; i++ {
		f()
	}
}

func BenchmarkInvocationConcrete(b *testing.B) {
	a := new(A)
	for i := 0; i < b.N; i++ {
		a.A()
	}
}

func BenchmarkInvocationInterface(b *testing.B) {
	var a InterfaceA = new(A)
	for i := 0; i < b.N; i++ {
		a.A()
	}
}

func BenchmarkInvocationConcreteFunc(b *testing.B) {
	a := new(A)
	f := a.A
	for i := 0; i < b.N; i++ {
		f()
	}
}

func BenchmarkInvocationInterfaceFunc(b *testing.B) {
	var a InterfaceA = new(A)
	f := a.A
	for i := 0; i < b.N; i++ {
		f()
	}
}

func bar(x int) {
	i = x
}

type InterfaceB interface {
	B(j int)
}

type B struct {
	j int
}

func (b *B) B(j int) {
	b.j = j
}

func BenchmarkInvocationFunctionWithArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bar(i)
	}
}

func BenchmarkInvocationClosureWithArgs(b *testing.B) {
	var i int
	f := func(x int) {
		i = x
	}
	for i := 0; i < b.N; i++ {
		f(i)
	}
}

func BenchmarkInvocationConcreteWithArgs(b *testing.B) {
	bb := new(B)
	for i := 0; i < b.N; i++ {
		bb.B(i)
	}
}

func BenchmarkInvocationInterfaceWithArgs(b *testing.B) {
	var bb InterfaceB = new(B)
	for i := 0; i < b.N; i++ {
		bb.B(i)
	}
}

func BenchmarkInvocationConcreteFuncWithArgs(b *testing.B) {
	bb := new(B)
	f := bb.B
	for i := 0; i < b.N; i++ {
		f(i)
	}
}

func BenchmarkInvocationInterfaceFuncWithArgs(b *testing.B) {
	var bb InterfaceB = new(B)
	f := bb.B
	for i := 0; i < b.N; i++ {
		f(i)
	}
}

func functionFactory(j int) func() {
	return func() {
		i = j
	}
}

func BenchmarkInvocationFactoryFunctionWithArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f := functionFactory(i)
		f()
	}
}
