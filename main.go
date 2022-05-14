package main

import (
	"workshop-golang/solutions"
	"workshop-golang/utils"
)

func main() {
	begin()
}

func begin() {
	utils.SolutionsWithOneParam("Begin1", "perimeter", solutions.Begin1)
	utils.SolutionsWithOneParam("Begin2", "square", solutions.Begin2)
	utils.SolutionsWithOneParam("Begin4", "length", solutions.Begin4)

	utils.SolutionsWithTwoParams("Begin8", "average", solutions.Begin8)
	utils.SolutionsWithTwoParams("Begin9", "geometric average", solutions.Begin9)

	utils.SolutionsWithOneParamSomeResults("Begin5", []string{"volume", "square"}, solutions.Begin5)
	utils.SolutionsWithOneParamSomeResults("Begin7", []string{"length", "square"}, solutions.Begin7)

	utils.SolutionsWithTwoParamsSomeResults("Begin3", []string{"square", "perimeter"}, solutions.Begin3)
	utils.SolutionsWithTwoParamsSomeResults("Begin10", []string{"sum", "difference", "production", "quotient"}, solutions.Begin10)

	utils.SolutionsWithThreeParamsSomeResults("Begin6", []string{"volume", "square"}, solutions.Begin6)
}
