package main

import (
	"context"
	"fmt"
	"github.com/amoghe/go-crypt"
	"os"
	"sync"
)

const ALPHABET = "abcdefghijklmnopqrstuvwxz0123456789!@#$%^&*()"
const alphabetSize = len(ALPHABET)


// const hash = "$6$FA5jtDxj$UyJD6JZ5UpWn47bJZQvxkGJb8xG9k4X4FyYAt1ZsPK5L54V1Eyk5eumea.d7jEdRdeHAoxaGpPxj5j6KOov500"
const hash = "$6$FA5jtDxj$zmYqs3rVUlk9ZRnyq5b1HM53XXso6Cu4h4UElbCEK5z/v.VbyXLQsvZL4Kxj0vd2OfIIXnOA9XNgogrdNm8tb/"
const salt = "$6$FA5jtDxj$"

const prefix = "<:cti18:>"

// func timer(f func()) time.Duration {
//
// 	start := time.Now()
// 	fmt.Printf("Algorithm started at %s\n", start)
// 	f()
// 	elapsed := time.Since(start)
// 	return elapsed
// }

func wrapper() {

	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := 0; i < alphabetSize; i++ {

		wg.Add(alphabetSize)
		go Gen3(ctx, cancel, &wg, string(ALPHABET[i]))
	}

	wg.Wait()
}

func Gen3(ctx context.Context, cancel context.CancelFunc, wg *sync.WaitGroup, first string) {
	defer wg.Done()
	var s string

	arg1 := os.Args[1]

	for i := 0; i < alphabetSize; i++ {
		for j := 0; j < alphabetSize; j++ {
			for k := 0; k < alphabetSize; k++ {
				s = fmt.Sprintf("%s%s%s%s%s%sn", prefix, arg1, first, string(ALPHABET[i]), string(ALPHABET[j]), string(ALPHABET[k]))
				// fmt.Println(s)
				fmt.Println(".")

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

	wrapper()

}
