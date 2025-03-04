package main

import "fmt"

type Person struct {
	Name string
}

var defaultPerson = Person{
	Name: "Mike",
}

// Можно так сделать или нет?
func changeName(person *Person) {
	// person.Name = "Alice"	// 1

	/*
		person = &Person{		// 2
			Name: "Alice",
		}
	*/

	*person = defaultPerson // 3
}

func main() {
	person := &Person{
		Name: "Bob",
	}
	fmt.Println(person.Name) // Bob
	changeName(person)
	// fmt.Println(person.Name) // 1. Alice
	// fmt.Println(person.Name) // 2. Alice
	fmt.Println(person.Name) // 3. Mike
}
