package main

type IPerson interface {
	print()
}

type Person struct {
}

func main() {
	p := Person{}
	p.print()
}

func (person Person) print() {
	println("Hello")
}
