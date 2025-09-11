package main

import (
	"fmt"
	"sync"
	"time"
)

// Objetivo: Simular tareas que toman tiempo con time.Sleep y comparar
// ejecución secuencial vs concurrente midiendo la duración total.
// TODO: completa las secciones marcadas para observar la mejora.

func tarea(id int, dur time.Duration) {
	fmt.Printf("[tarea %d] iniciando, dur=%v\n", id, dur)
	time.Sleep(dur)
	fmt.Printf("[tarea %d] finalizada\n", id)
}

func secuencial(durs []time.Duration) time.Duration {
	inicio := time.Now()
	// TODO: ejecutar las tareas en orden, sin goroutines
	for i, d := range durs {
		//ejecuta las tareaas en orden, sin goroutines
		tarea(i+1, d) //se usa en +1 para que el id inicie
	}
	return time.Since(inicio)
}

func concurrente(durs []time.Duration) time.Duration {
	inicio := time.Now()
	var wg sync.WaitGroup
	// TODO: lanzar cada tarea en su propia goroutine y esperar con WaitGroup
	for i, d := range durs {
		wg.Add(1)
		go func(id int, dur time.Duration){ //lanza cada tarea en su propia goroutine
			defer wg.Done()
			tarea(id, dur)
		}(i+1, d) //se usa en +1 para que el id inicie

	}
	wg.Wait()
	return time.Since(inicio)
}

func main() {
	// TODO: experimenta con diferentes duraciones
	// durs := []time.Duration{700 * time.Millisecond, 500 * time.Millisecond, 1 * time.Second}
	durs := []time.Duration{
		700 * time.Millisecond,
		500 * time.Millisecond,
		1 * time.Second,
	}
	// d1 := 
	d1 := secuencial(durs)
	fmt.Println("Duración SEC:", d1)

	// d2 := 
	d2 := concurrente(durs)
	fmt.Println("Duración CONC:", d2)

	fmt.Println("Nota: la ejecución concurrente debería ser ~max(durs). Cambia valores y observa.")
}
