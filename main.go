package main

import (
	"fmt"
	"io/ioutil"
)

type point struct{x, y int64 }

func (p point) String() string {
	return fmt.Sprintf("Point(%d,%d)", p.x, p.y)
}

func (p point) GoString() string {
	return fmt.Sprintf("Point(%d,%d)", p.x, p.y)
}

var (
	i = 100
	f = 3.14
	b = true
	s = "Clear is better than clever."
	p = point{252, 101}
)

var bi  int64 = -922337203685
var ui uint64 = 1844679551615

func main() {
	// fmt.Println(outputToWriter())
	// outputToStdout()
	fmt.Println(p)
	fmt.Printf("%#v\n", p)
}

func outputToWriter() (string, error) {
	file, err := ioutil.TempFile("", "example")

	if err != nil {
		return "", err
	}

	defer file.Close()

	fmt.Fprintf(file, "i = %#v\nf = %#v\nb = %#v\ns = %#v\nbi = %#v\nui = %#v\np = %#v\n", i, f, b, s, bi, ui, p)
	return file.Name(), nil
}

func outputToStdout() {
	fmt.Printf("i = %#v\nf = %#v\nb = %#v\ns = %#v\nbi = %#v\nui = %#v\np = %#v\n", i, f, b, s, bi, ui, p)
	fmt.Printf("i = %T\nf = %T\nb = %T\ns = %T\nbi = %T\nui = %T\np = %T\n", i, f, b, s, bi, ui, p)
	fmt.Printf("%#v\n", p)
}