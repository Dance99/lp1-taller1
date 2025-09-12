package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Objetivo: Lanzar varias goroutines que imprimen mensajes y esperar a que todas terminen.
// Completa los pasos marcados con TODO para entender goroutines y WaitGroup.

func worker(id int, veces int, wg *sync.WaitGroup) {
	// asegurar que al finalizar la función se haga wg.Done()
	
	//Se añade un defer wg.Done, para avisar al WaitGroup que esta goroutine ha terminado
	defer wg.Done()

	for i := 1; i <= veces; i++ {
		fmt.Printf("[worker %d] hola %d\n", id, i)
		// dormir un poco para simular trabajo (p. ej. 100–300 ms)
		
		//Se añade time.sleep, para simular el trabajo
		time.Sleep(time.Duration(100+rand.Intn(200)) * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup

	// cambiar estos parámetros y observar el intercalado de salidas
	// numGoroutines
	// veces

	//parametro: cuantps gorountines lanzar y cuantas veces debe imprimir cada una
	numGoroutines := 5
	veces := 5

	//rand para que los sleeps sean diferentes en cada ejecución
	rand.Seed(time.Now().UnixNano())

	// lanzar varias goroutines, sumar al WG y esperar con wg.Wait()
	for id := 1; id <= numGoroutines; id++ {
		wg.Add(1) // Se añade una goroutine al WaitGroup,sumando 1 al waitgroup
		go worker(id, veces, &wg) // Se lanza la goroutine, pasandole el WaitGroup

	}

	// Esperar a que todas las goroutines terminen
	wg.Wait()
	
	fmt.Println("Listo: todas las goroutines terminaron.")
}
