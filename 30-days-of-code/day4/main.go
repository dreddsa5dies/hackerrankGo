package main

import "fmt"

type person struct {
	age int
}

// ---------------------
// объявление методов для типа person
// всего будут методы NewPerson, amIOld, yearPasses

func (p person) NewPerson(initialAge int) person {
	//Add some more code to run some checks on initialAge
	if initialAge <= 0 {
		fmt.Println("Age is not valid, setting age to 0.")
		p.age = 0
	}

	return p
}

// выполняет слудеющие условия:
// выводит "You are young." при age < 13 и т.д.
func (p person) amIOld() {
	//Do some computation in here and print out the correct statement to the console
	switch {
	case p.age < 13:
		fmt.Println("You are young.")
		break
	case p.age >= 13 && p.age < 18:
		fmt.Println("You are a teenager.")
		break
	case p.age >= 18:
		fmt.Println("You are old.")
		break
	}
}

// yearPasses() увеличивает age на 1
func (p person) yearPasses() person {
	//Increment the age of the person in here
	p.age++

	return p
}

// ---------------------
func main() {
	var T, age int

	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&age)
		p := person{age: age}
		p = p.NewPerson(age)
		p.amIOld()
		for j := 0; j < 3; j++ {
			p = p.yearPasses()
		}
		p.amIOld()
		fmt.Println()
	}
}
