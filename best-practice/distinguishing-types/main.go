package main

type Foo string
type Bar string
type FooBar struct {
	foo Foo
	bar Bar
}

func NewFooBar(foo Foo, bar Bar) FooBar {
	return FooBar{
		foo: foo,
		bar: bar,
	}
}
