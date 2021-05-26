package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type secretAgent struct {
	person
}

func (p person) pSpeak() {
	fmt.Println(p.firstName, "says hello")
}
func (s secretAgent) saSpeak() {
	fmt.Println(s.firstName, s.lastName, "says something")

}

func main() {
	p := person{"Szop", "Junior"}
	sa := secretAgent{person{"James", "Bond"}}

	p.pSpeak()
	sa.saSpeak()

}
