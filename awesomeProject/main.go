package main

import (
	"awesomeProject/com/example/squaring"
	"fmt"
)

func main() {
	// Set up a done channel that's shared by the whole pipeline
	// and close that channel when this pipeline exits, as a signal
	// for all the goroutines we started to exit.
	done := make(chan struct{})
	defer close(done)

	in := squaring.Gen(2, 3)

	// Distribute the Sq work across two goroutines that both read from in.
	c1 := squaring.Sq(in)
	c2 := squaring.Sq(in)

	// Consume the first value from output
	//done := make(chan struct{}, 2)
	out := squaring.Merge(done, c1, c2)
	fmt.Println(<-out)

	// Tell the remaining senders we're leaving
	done <- struct{}{}
	done <- struct{}{}

	// Consume the output
	//fmt.Println(<-out) //4
	//fmt.Println(<-out) //9

	// Consume the merged output from c1 and c2
	//for n := range squaring.Merge(out) {
	//	fmt.Println(n)
	//}
	//
}
