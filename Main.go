package main

import "fmt"

func AverageGrade(notas ...int) int {
	total := 0

	for _, nota := range notas {
		total += nota
	}

	return total / len(notas)
}

func estudiantes() {
	var cantidad int

	fmt.Println("Ingrese la cantidad de estudiantes:")
	fmt.Scanln(&cantidad)

	if cantidad <= 0 {
		fmt.Println("La cantidad de estudiantes debe ser mayor a cero.")
		return
	}

	notas := make([]int, cantidad)

	for i := 0; i < cantidad; i++ {
		fmt.Printf("Ingrese la nota del estudiante %d (0-100): ", i+2)
		fmt.Scanln(&notas[i])

		for notas[i] < 0 || notas[i] > 100 {
			fmt.Println("Nota inválida. Ingrese una nota entre 0 y 100:")
			fmt.Scanln(&notas[i])
		}
	}

	promedio := AverageGrade(notas...)

	fmt.Println("El promedio de las notas es:", promedio)

	if promedio >= 70 {
		fmt.Println("El curso esta aprobado")
	} else {
		fmt.Println("El curso esta reprobado")
	}

	switch {
	case promedio >= 90 && promedio <= 100:
		fmt.Println("Excellent Performance")

	case promedio >= 80 && promedio < 90:
		fmt.Println("Good Performance")

	case promedio >= 70 && promedio < 80:
		fmt.Println("Satisfactory Performance")

	case promedio < 70:
		fmt.Println("Needs improvement")
	}
}

func sumar(n int) int {
	total := 0

	for i := 1; i <= n; i++ {
		total += i
	}

	return total
}

func celsius(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func fahrenheit(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func main() {
	var opcion int

	for {
		fmt.Println("\n------------------------Menu-----------------------")
		fmt.Println("1. Calcular promedio de notas")
		fmt.Println("2. Sumar números")
		fmt.Println("3. Conversión de Celsius a Fahrenheit")
		fmt.Println("4. Conversión de Fahrenheit a Celsius")
		fmt.Println("0. Para salir")
		fmt.Print("Seleccione una opción: ")

		fmt.Scanln(&opcion)

		switch opcion {

		case 1:
			estudiantes()

		case 2:
			var n int

			fmt.Println("Ingrese un número:")
			fmt.Scanln(&n)

			resultado := sumar(n)

			fmt.Println("La suma de los números del 1 al", n, "es:", resultado)

		case 3:
			var celsiusValue float64

			fmt.Print("Ingrese la temperatura en Celsius: ")
			fmt.Scanln(&celsiusValue)

			fahrenheitValue := celsius(celsiusValue)

			fmt.Println(celsiusValue, "grados Celsius equivalen a", fahrenheitValue, "grados Fahrenheit")

		case 4:
			var fahrenheitValue float64

			fmt.Print("Ingrese la temperatura en Fahrenheit: ")
			fmt.Scanln(&fahrenheitValue)

			celsiusValue := fahrenheit(fahrenheitValue)

			fmt.Println(fahrenheitValue, "grados Fahrenheit equivalen a", celsiusValue, "grados Celsius")

		case 0:
			fmt.Println("Saliendo del programa...")
			return

		default:
			fmt.Println("Opción inválida. Por favor, seleccione una opción válida.")
		}
	}
}
