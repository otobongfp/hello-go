package main

import "fmt"

type humanoid interface {
	speak()
	walk()
}

type person struct{ name string }

func (p person) speak() { fmt.Printf("%s speaking...", p.name) }
func (p person) walk()  { fmt.Printf("%s walking... \n", p.name) }

type dog struct{}

func (d dog) walk() { fmt.Printf("%s is walking...", d) }

func main() {

	p := person{name: "Bolt"}
	doHumanoid(p)

}

func doHumanoid(h humanoid) {
	h.speak()
	h.walk()
}
