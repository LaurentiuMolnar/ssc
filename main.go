package main

import (
	"context"
	"fmt"
	"github.com/amoghe/go-crypt"
	"sync"
	"time"
)

const ALPHABET = "abcdefghijklmnopqrstuvwxz0123456789!@#$%^&*()"
const alphabetSize = len(ALPHABET)

// const hash = "$6$SRM3J9B1$Fk7jQICjeGcPWbNM8FsHCoSQPQ/SjzK/dtzy14oT62haJji6539o9qfD7oMpdkZgfajsQSThHvEvhhATZtIb00"
// const salt = "$6$SRM3J9B1$"

const hash = "$6$4LDtKhWO$T1.Bnzi/k/XYJsNA26FJ2wDOOcEea/ZeojzO8/QHEe5mV5A84/uvmImYt0rGb0LWaINULcQL7OIfdUp90vbHU/"
const salt = "$6$4LDtKhWO$"

const prefix = "<:cti18:>"

func timer(f func()) time.Duration {

	start := time.Now()
	fmt.Printf("Algorithm started at %s\n", start)
	f()
	elapsed := time.Since(start)
	return elapsed
}

func wrapper() {

	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := 0; i < alphabetSize; i++ {

		wg.Add(1)
		go Gen3(ctx, cancel, &wg, string(ALPHABET[i]))
	}

	wg.Wait()
}

func Gen3(ctx context.Context, cancel context.CancelFunc, wg *sync.WaitGroup, first string) {
	defer wg.Done()
	var s string

	for i := 0; i < alphabetSize; i++ {
		for j := 0; j < alphabetSize; j++ {
			for k := 0; k < alphabetSize; k++ {
				s = fmt.Sprintf("%s%s%s%s%s", prefix, first, string(ALPHABET[i]), string(ALPHABET[j]), string(ALPHABET[k]))
				fmt.Println(s)

				hashed, err := crypt.Crypt(s, salt)
				if err != nil {
					fmt.Errorf("Some error occurred: %s", err)
					return
				}

				if hashed == hash {
					fmt.Printf("Found password: %s\n", s)
					cancel()
					return
				}

				select {
				case <-ctx.Done():
					return
				default:
				}
			}
		}
	}
}

func main() {

	// fmt.Printf("Ran %s combinations\n", string(wrapper(runs)))
	fmt.Printf("Algorithm cracked password in %s\n", timer(wrapper))

}
