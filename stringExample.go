package main

import  (
	"fmt"
	"strings"
	"unicode"
	"bytes"
)

func main() {
	contains()
	trimFunc()
	replacer()
	join()
}

func join() {
	alphabet := "alpha      beta      gamma"
	fields := strings.Fields(alphabet)
	fmt.Printf("Fields: %q\n", fields)
	fmt.Printf("Joined: %q\n", strings.Join(fields, "|"))
}

func replacer() {
	r:= strings.NewReplacer("alpha","A","theta","φ", "delta", "0")
	fmt.Printf("%q\n", r.Replace("The alpha differs from the theta differs from the delta"))
}

func trimFunc() {
	f := func(c rune) bool  {
		return !unicode.IsLetter(c)
	}
	fmt.Printf("[%q]\n", strings.TrimFunc("     1234gophers unite!123445656656    ", f))
}

func contains() {
	fmt.Println(strings.Contains("Hello, world!", "gophers"))
	fmt.Println(bytes.Contains([]byte("abc"), []byte("b")))

	fmt.Println(bytes.Compare([]byte("ab"), []byte{'a', 'b'}))
	fmt.Println(bytes.Compare([]byte("ab"), []byte{'a', 'b', 'c'}))
	fmt.Println(bytes.Compare([]byte("ab"), []byte{'a'}))

	fmt.Println(strings.TrimSpace(" hello! "))

}