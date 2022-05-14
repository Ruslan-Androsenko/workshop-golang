package utils

import "fmt"

// SolutionsWithOneParam Тестирование задач для расчетов с одним заданным параметром
func SolutionsWithOneParam(title, caption string, operation func(float64) float64) {
	var valuesBySide = []float64{2.7, 3.14, 9.2, 5.38, 7.25}
	var results []float64

	for _, value := range valuesBySide {
		results = append(results, operation(value))
	}

	DisplayFloatItems(title, caption, results)
}

// SolutionsWithTwoParams Тестирование задач для расчетов с двумя заданными параметрами
func SolutionsWithTwoParams(title, caption string, operation func(float64, float64) float64) {
	var valuesBySides = [][]float64{
		{2.7, 5.3},
		{3.14, 8.38},
		{4.56, 9.2},
		{5.38, 12.32},
		{7.25, 14.54},
	}
	var results []float64

	for _, values := range valuesBySides {
		results = append(results, operation(values[0], values[1]))
	}

	DisplayFloatItems(title, caption, results)
}

// SolutionsWithOneParamSomeResults Тестирование задач для расчетов с одним заданным параметром и несколькими результатами
func SolutionsWithOneParamSomeResults(title string, captions []string, operation func(float64) map[string]float64) {
	var valuesBySide = []float64{2.7, 3.14, 9.2, 5.38, 7.25}
	var results []map[string]float64

	for _, value := range valuesBySide {
		results = append(results, operation(value))
	}

	DisplaySomeFloatItems(title, captions, results)
}

// SolutionsWithTwoParamsSomeResults Тестирование задач для расчетов с двумя заданными параметрами и несколькими результатами
func SolutionsWithTwoParamsSomeResults(title string, captions []string, operation func(float64, float64) map[string]float64) {
	var valuesBySides = [][]float64{
		{2.7, 5.3},
		{3.14, 8.38},
		{4.56, 9.2},
		{5.38, 12.32},
		{7.25, 14.54},
	}
	var results []map[string]float64

	for _, values := range valuesBySides {
		results = append(results, operation(values[0], values[1]))
	}

	DisplaySomeFloatItems(title, captions, results)
}

// SolutionsWithThreeParamsSomeResults Тестирование задач для расчетов с тремя заданными параметрами и несколькими результатами
func SolutionsWithThreeParamsSomeResults(title string, captions []string, operation func(float64, float64, float64) map[string]float64) {
	var valuesBySides = [][]float64{
		{2.7, 5.3, 4.2},
		{3.14, 8.38, 6.23},
		{4.56, 9.2, 7.1},
		{5.38, 12.32, 4.87},
		{7.25, 14.54, 6.8},
	}
	var results []map[string]float64

	for _, values := range valuesBySides {
		results = append(results, operation(values[0], values[1], values[2]))
	}

	DisplaySomeFloatItems(title, captions, results)
}

// DisplayFloatItems Вывод результирующего массива из вещественных чисел
func DisplayFloatItems(title, caption string, items []float64) {
	fmt.Println("\n" + title)

	for key, value := range items {
		captionValue := fmt.Sprintf("%s%d:", caption, key+1)
		fmt.Println(captionValue, value)
	}
}

// DisplaySomeFloatItems Вывод результирующего массива из нескольких вещественных чисел с заданными индексами
func DisplaySomeFloatItems(title string, captions []string, items []map[string]float64) {
	fmt.Println("\n" + title)

	for index, item := range items {
		for _, key := range captions {
			if value, ok := item[key]; ok {
				captionValue := fmt.Sprintf("%s%d:", key, index+1)
				fmt.Println(captionValue, value)
			}
		}

		fmt.Println()
	}
}
