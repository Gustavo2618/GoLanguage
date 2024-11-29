package main

import "fmt"

const (
	x = iota
	y = iota
	_ = iota
	z = iota
)

func main() {
	fmt.Println(x, y, z)
	a := 1
	b := a << 1
	c := b << 1
	d := c << 1
	e := d << 1
	f := e >> 1
	g := f >> 1
	h := g >> 1
	i := h >> 1
	fmt.Println(a, b, c, d, e, f, g, h, i)

}
