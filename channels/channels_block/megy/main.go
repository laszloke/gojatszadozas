package main

import (
	"fmt"
	//"time"
)

func main() {
	ca := make(chan int)

	go func() {
		
		/*time.Sleep(3 * time.Second)
		fmt.Println(" hello a go rutin bol")
		*/

		ca <- 42
	}()

	fmt.Println(<-ca)

}
