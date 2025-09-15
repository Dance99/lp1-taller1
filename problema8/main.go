package main

import (
	"fmt"
	"time"
)

// Objetivo: Simular "futuros" en Go usando canales. Una función lanza trabajo asíncrono
// y retorna un canal de solo lectura con el resultado futuro.
// completa las funciones y experimenta con varios futuros a la vez.

func asyncCuadrado(x int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		//simular trabajo
		time.Sleep(500 * time.Millisecond)

		ch <- x * x
	}()
	return ch
}

// fanIn combina varios canales de enteros en uno solo
func fanIn(chans ...<-chan int) <-chan int {
	out := make(chan int)

	go func() {
		// contador de goroutines
		done := make(chan bool)

		for _, ch := range chans {
			// por cada canal, lanzamos un lector
			go func(c <-chan int) {
				for v := range c {
					out <- v
				}
				done <- true
			}(ch)
		}

		// esperamos que todos terminen
		for i := 0; i < len(chans); i++ {
			<-done
		}
		close(out)
	}()

	return out
}
func main() {
	fmt.Println("===Simulación de futuro con Go ===")
	//crea varios futuros y recolecta sus resultados: f1, f2, f3
	f1 := asyncCuadrado(2)
	f2 := asyncCuadrado(3)
	f3 := asyncCuadrado(4)

	//  Opción 1: esperar cada futuro secuencialmente
	// Opción 1: esperar cada futuro secuencialmente
	fmt.Println("Secuencial:")
	fmt.Println("Resultado f1:", <-f1)
	fmt.Println("Resultado f2:", <-f2)
	fmt.Println("Resultado f3:", <-f3)

	// Volvemos a crearlos porque los anteriores ya se consumieron
	f1 = asyncCuadrado(5)
	f2 = asyncCuadrado(6)
	f3 = asyncCuadrado(7)

	
	// Opción 2: fan-in (combinar múltiples canales)
	// Pista: crea una función fanIn que recibe múltiples <-chan int y retorna un único <-chan int
	// que emita todos los valores. Requiere goroutines y cerrar el canal de salida cuando todas terminen.
	
	// Opción 2: usar fan-in
	fmt.Println("\nCon fan-in:")
	merged := fanIn(f1, f2, f3)
	for v := range merged {
		fmt.Println("Recibido:", v)
	}

	fmt.Println("Fin del programa.")
}
