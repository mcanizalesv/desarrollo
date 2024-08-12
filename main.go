package main

import (
	"fmt"
)
func main () {
	
	// Dada una array de enteros, invierte el orden de los elementos.
	
	
	// Ejemplo de uso
	numeros := []int{10, 20, 30, 40, 50}
	suma, promedio, max := Estadisticas(numeros)
		
	fmt.Printf("Suma: %d\n", suma)
	fmt.Printf("Promedio: %.2f\n", promedio)
	fmt.Printf("Máximo: %d\n", max)
	
}

func InvertirArray(arr [5]int) [5]int {
	var result [5]int
	for i := 0; i < len(arr); i++ {
		result[i] = arr[len(arr)-1-i]
	}
	return result
}

// Calcula la suma, el promedio y el valor máximo de un array de enteros.
// Función para calcular suma, promedio y máximo
func Estadisticas(arr []int) (suma int, promedio float64, max int) {
// Inicializa la suma y el valor máximo
	suma = 0
	if len(arr) > 0 {
		max = arr[0]
	}
	
	// Recorre el slice para calcular la suma y encontrar el valor máximo
	for _, v := range arr {
		suma += v
		if v > max {
			max = v
		}
	}
	
	// Calcula el promedio
	if len(arr) > 0 {
		promedio = float64(suma) / float64(len(arr))
	}
	
	return
}