package main

type Point struct{ x, y int }

type Rect1 struct{ Min, Max Point }
type Rect2 struct{ Min, Max *Point }

func main() {}
