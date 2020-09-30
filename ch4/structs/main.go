package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

type Point struct {
	X, Y int
}

type Circle struct {
	Center Point
	Radius int
}

type Wheel struct {
	Circle Circle
	Spokes int
}

type Circle2 struct {
	Point
	Radius int
}

type Wheel2 struct {
	Circle2
	Spokes int
}

func main() {
	var dilbert Employee
	dilbert.Salary -= 5000 // demoted, for writing too few lines of code

	position := &dilbert.Position
	*position = "Senior " + *position // promoted, for outsourcing to Elbonia

	fmt.Println(dilbert)

	var emplyeeOfTheMonth *Employee = &dilbert
	emplyeeOfTheMonth.Position += " (proactive team player"
	fmt.Println(*emplyeeOfTheMonth)
	fmt.Println(dilbert)

	fmt.Println(Bonus(&Employee{Salary: 5000}, 200))

	AwardAnnualRaise(&dilbert)
	fmt.Println(dilbert)

	p := Point{1, 2}
	p2 := Point{1, 2}
	q := Point{2, 1}
	fmt.Println(p.X == q.X && p.Y == q.Y) // false
	fmt.Println(p == q)                   // false
	fmt.Println(p == p2)                  // true

	var w Wheel
	w.Circle.Center.X = 8
	w.Circle.Center.Y = 8
	w.Circle.Radius = 5
	w.Spokes = 20

	// 無名
	var w2 Wheel2
	w2.X = 8
	w2.Y = 8
	w2.Radius = 5
	w2.Spokes = 20
	// w3 := Wheel2{8, 8, 5, 20} // compile error
	// w3 = Wheel2{X: 8, Y: 8, Radius: 5, Spokes: 20} // compile error

	w3 := Wheel2{
		Circle2: Circle2{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}

	fmt.Println(w3)
	fmt.Printf("%#v\n", w3)
	fmt.Printf("%v\n", w3)

	w3.X = 42
	fmt.Printf("%#v\n", w3)
}

func Bonus(e *Employee, percent int) int {
	return e.Salary * percent / 100
}

func AwardAnnualRaise(e *Employee) {
	e.Salary = e.Salary * 105 / 100
}
