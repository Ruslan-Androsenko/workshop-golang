package solutions

import (
	"fmt"
	"sort"
	"strings"
)

type Student struct {
	Fio      string
	Group    string
	Progress [5]int
}

func (item Student) AverageProgress() float64 {
	var sum int = 0

	for _, value := range item.Progress {
		sum += value
	}

	return float64(sum) / 5.0
}

func (item Student) PrintItem(index int) {
	rowItem := fmt.Sprintf("| %02d | %15s | %12s | %12.2f |", index, item.Fio, item.Group, item.AverageProgress())
	fmt.Println(rowItem)
}

func PrintHeaderTable() {
	line := strings.Repeat("-", 54)
	title := fmt.Sprintf("| %2s | %15s | %12s | %12s |", "#", "Фамилия", "Номер группы", "Средний бал")

	fmt.Println(line)
	fmt.Println(title)
	fmt.Println(line)
}

func PrintTableItems(students []Student) {
	PrintHeaderTable()

	for index, item := range students {
		item.PrintItem(index + 1)
	}
}

func GetSuccessfulItems(students []Student) []Student {
	var successfulStudents []Student

	for _, item := range students {
		if item.AverageProgress() > 4.0 {
			successfulStudents = append(successfulStudents, item)
		}
	}

	return successfulStudents
}

func SortItemsByFio(students []Student) {
	sort.Slice(students, func(i, j int) bool {
		return students[i].Fio < students[j].Fio
	})
}

func SortItemsByGroup(students []Student) {
	sort.Slice(students, func(i, j int) bool {
		return students[i].Group < students[j].Group
	})
}

func SortItemsByProgress(students []Student) {
	sort.Slice(students, func(i, j int) bool {
		return students[i].AverageProgress() < students[j].AverageProgress()
	})
}

func MakeMockData() []Student {
	var ivanov = Student{
		Fio:      "Иванов А.Д.",
		Group:    "КН-256",
		Progress: [5]int{4, 5, 5, 4, 3},
	}
	var petrova = Student{
		Fio:      "Петрова Е.С.",
		Group:    "ПК-115",
		Progress: [5]int{5, 5, 4, 4, 4},
	}
	var sidorov = Student{
		Fio:      "Сидоров Н.А.",
		Group:    "КН-255",
		Progress: [5]int{4, 3, 4, 4, 3},
	}
	var nikiforov = Student{
		Fio:      "Никифоров И.Н.",
		Group:    "ПК-115",
		Progress: [5]int{3, 5, 4, 3, 5},
	}
	var sergeeva = Student{
		Fio:      "Сергеева А.В.",
		Group:    "КН-256",
		Progress: [5]int{5, 5, 4, 5, 4},
	}

	return []Student{
		ivanov,
		petrova,
		sidorov,
		nikiforov,
		sergeeva,
	}
}
