package main

import "fmt"

type T struct{ a, b int }

type myStruct struct {
	i int
	j float32
	k string
}

func main() {
	var s T
	s.a = 10
	s.b = 20

	var t *T
	t = new(T)

	fmt.Println(t)

	// -

	var v myStruct  // v has struct type
	var p *myStruct // p is a pointer to a struct
	fmt.Println("print v p", v, p)

	mt := &myStruct{10, 15.5, "Chris"} // this means that `v` is of type *myStruct
	var mp myStruct
	mp = myStruct{20, 30.2, "Pedro"}

	fmt.Println(mt.i, mp.i)

}
