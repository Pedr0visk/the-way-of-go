package main

import "fmt"

/* basic data structure upon which we'll define methods */
type employee struct {
	salary float32
}

/* a method which will add a specified percent to an
   employees salary */
func (this *employee) giveRaise(pct float32) {
	this.salary += this.salary * pct
}

func (e employee) raiseSalary(pct float32) {
	e.salary += e.salary * pct
}

func main() {
	/* create an employee instance */
	e := employee{10000}
	// or
	// var e = new(employee)
	// e.salary = 100000;
	/* call our method */
	e.giveRaise(0.04)
	fmt.Printf("Employee now makes %f\n", e.salary)

	e2 := employee{10000}
	e2.raiseSalary(0.04)
	fmt.Printf("Employee now makse %f\n", e.salary)
}
