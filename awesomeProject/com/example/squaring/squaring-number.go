package squaring

import "sync"

func Gen(nums ...int) <-chan int {
	out := make(chan int, len(nums))
	for _, n := range nums {
		out <- n
	}
	close(out)
	return out
}

func Sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func Merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int, 1) // enough space for the unread inputs

	// Start an output goroutine for each input channel in cs.output
	// copies values from c to out until c is closed, then calls wg.Done.
	// from done, then output calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			select {
			case out <- n:
			case <-done:
			}
		}
		wg.Done()
	}

	for _, c := range cs {
		go output(c)
		wg.Add(len(cs))
	}

	// Start a goroutine to close out once all the output goroutines are
	// done . This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
