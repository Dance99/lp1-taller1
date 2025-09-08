package main

import (
	"fmt"
	"sync"
	"time"
)

// Objetivo: Lanzar varias goroutines que imprimen mensajes y esperar a que todas terminen.
// TODO: Completa los pasos marcados con TODO para entender goroutines y WaitGroup.

func worker(id int, veces int, wg *sync.WaitGroup) {
	// TODO: asegurar que al finalizar la función se haga wg.Done()
	

	for i := 1; i <= veces; i++ {
		fmt.Printf("[worker %d] hola %d\n", id, i)
		// TODO: dormir un poco para simular trabajo (p. ej. 100–300 ms)
		
	}
}

func main() {
	var wg sync.WaitGroup

	// TODO: cambiar estos parámetros y observar el intercalado de salidas
	// numGoroutines
	// veces

	// TODO: lanzar varias goroutines, sumar al WG y esperar con wg.Wait()
	for id := 1; id <= numGoroutines; id++ {

	}

	// Esperar a que todas las goroutines terminen
	
	fmt.Println("Listo: todas las goroutines terminaron.")
}
