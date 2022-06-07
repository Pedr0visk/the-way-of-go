package main

import "fmt"

type Engine interface {
	Start()
	Stop()
}

type Motor struct{}

func (m Motor) Start() {
	fmt.Println("Started!")
}

func (m Motor) Stop() {
	fmt.Println("Ended!")
}

type Car struct {
	Engine
}

func (c *Car) GoToWorkIn() {
	// get in car
	c.Start()
	// drive to work
	c.Stop()
	// get out of car
}

func main() {

}
