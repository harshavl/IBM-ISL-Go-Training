/* generate the first count prime numbers from the given start using a goroutine and print them */
//  generatePrimes(start, count int)

package main

import (
	"fmt"
	"time"
)

func main() {
	count := 20
	primeNoCh := generatePrimes(3, count)
	for i := 0; i < count; i++ {
		fmt.Println(<-primeNoCh)
	}
	fmt.Println("Done")
}

func generatePrimes(start, count int) chan int {
	ch := make(chan int)
	go func() {
		no := start
		for count > 0 {
			if isPrime(no) {
				ch <- no
				count--
				time.Sleep(500 * time.Millisecond)
			}
			no++
		}
	}()
	return ch
}

func isPrime(no int) bool {
	for i := 2; i <= no/2; i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
