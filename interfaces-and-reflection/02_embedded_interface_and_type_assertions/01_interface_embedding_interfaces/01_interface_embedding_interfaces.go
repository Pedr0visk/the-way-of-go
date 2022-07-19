// An interface can contain the name of one (or more) interface(s), which is equivalent to explicitly enumerating the methods of the embedded interface in the containing interface. For example, the interface File contains all the methods of ReadWrite and Lock, in addition to a Close() method:

package main

// type ReadWrite interface {
// 	Read(b Buffer) bool
// 	Write(b Buffer) bool
// }

// type Lock interface {
// 	Lock()
// 	Unlock()
// }

// type File interface {
// 	ReadWrite
// 	Lock
// 	Close()
// }

func main() {}
