package main

import (
	"fmt"
	"workshop-golang/solutions"
)

func main() {
	fmt.Println("Практикум для решение задач на языке GO.")
	fmt.Println("Вариант 1.")

	var students = solutions.MakeMockData()
	solutions.SortItemsByGroup(students)
	solutions.PrintTableItems(students)

	var successfulStudents = solutions.GetSuccessfulItems(students)
	solutions.PrintTableItems(successfulStudents)
}
