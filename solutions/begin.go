package solutions

import "math"

// Begin1 Дана сторона квадрата a.
// Найти его периметр P = 4 * a.
func Begin1(a float64) float64 {
	return a * 4
}

// Begin2 Дана сторона квадрата a.
// Найти его площадь S = a^2.
func Begin2(a float64) float64 {
	return a * a
}

// Begin3 Даны стороны прямоугольника a и b.
// Найти его площадь S = a * b и периметр P = 2 * (a + b).
func Begin3(a, b float64) map[string]float64 {
	square := a * b
	perimeter := 2 * (a + b)

	return map[string]float64{
		"square":    square,
		"perimeter": perimeter,
	}
}

// Begin4 Дан диаметр окружности d.
// Найти ее длину L = pi * d.
func Begin4(d float64) float64 {
	return 3.14 * d
}

// Begin5 Дана длина ребра куба a.
// Найти объем куба V = a^3 и площадь его поверхности S = 6 * a^2.
func Begin5(a float64) map[string]float64 {
	volume := a * a * a
	square := 6 * a * a

	return map[string]float64{
		"volume": volume,
		"square": square,
	}
}

// Begin6 Даны длины ребер a, b, c прямоугольного параллелепипеда.
// Найти его объем V = a * b * c и площадь поверхности S = 2 * (a * b + b * c + a * c).
func Begin6(a, b, c float64) map[string]float64 {
	volume := a * b * c
	square := 2 * (a*b + b*c + a*c)

	return map[string]float64{
		"volume": volume,
		"square": square,
	}
}

// Begin7 Найти длину окружности L и площадь круга S заданного радиуса R.
// L = 2 * pi * R, S = pi * R^2.
func Begin7(r float64) map[string]float64 {
	length := 2 * 3.14 * r
	square := 3.14 * r * r

	return map[string]float64{
		"length": length,
		"square": square,
	}
}

// Begin8 Даны два числа a и b.
// Найти их среднее арифметическое: (a + b) / 2.
func Begin8(a, b float64) float64 {
	return (a + b) / 2
}

// Begin9 Даны два неотрицательных числа a и b.
//Найти их среднее геометрическое, то есть квадратный корень из их произведения: a * b.
func Begin9(a, b float64) float64 {
	var result float64

	if a > 0 && b > 0 {
		result = math.Sqrt(a * b)
	}

	return result
}

// Begin10 Даны два ненулевых числа.
// Найти сумму, разность, произведение и частное их квадратов.
func Begin10(a, b float64) map[string]float64 {
	var result = make(map[string]float64)

	if a != 0 && b != 0 {
		a *= a
		b *= b

		sum := a + b
		difference := a - b
		production := a * b
		quotient := a / b

		result = map[string]float64{
			"sum":        sum,
			"difference": difference,
			"production": production,
			"quotient":   quotient,
		}
	}

	return result
}
