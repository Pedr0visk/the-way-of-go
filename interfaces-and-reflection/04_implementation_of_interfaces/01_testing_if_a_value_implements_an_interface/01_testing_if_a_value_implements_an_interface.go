// type Stringer interface { String() string }
// if sv, ok := v.(Stringer); ok {
//   fmt.Printf("v implements String(): %s\n", sv.String()); // note: sv, not v
// }

package main

import "fmt"

type Stringer interface{ String() string }

type Text struct {
	msg string
}

func (u *Text) String() string {
	return fmt.Sprintf("hello: %s\n", u.msg)
}

func main() {
	var text = new(Text)
	text.msg = "Pedro"

	var v Stringer
	v = text

	if sv, ok := v.(Stringer); ok {
		fmt.Printf("v implements String(): %s\n", sv.String()) // note: sv, not v
	}

	// if sv, ok := text.(Stringer); ok {
	// 	fmt.Printf("v implements String(): %s\n", sv.String()) // note: sv, not v
	// } => invalid operation: text (variable of type *Text) is not an interfacecompiler
}
