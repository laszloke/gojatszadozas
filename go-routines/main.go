
package main

import (
	"fmt"
	"sync"
	"runtime"
)



func main() {
	fmt.Println("Hello a main bol")

	counter:=0
	const gs = 10
	var wg sync.WaitGroup
	wg.Add(gs)


	for i := 0; i < gs; i++ {
		go func() {
			v:=counter

			
			runtime.Gosched()
			v++
			counter = v
			fmt.Println("Hello a ", i, " dik rutinbol")
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}

