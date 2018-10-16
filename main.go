package main

import (
	"context"
	"fmt"
	"github.com/tredoe/osutil/user/crypt/sha512_crypt"
	"github.com/tredoe/osutil/user/crypt"
	"os"
	"sync"
)

const ALPHABET = "abcdefghijklmnopqrstuvwxz0123456789!@#$%^&*()"
const alphabetSize = len(ALPHABET)

const hash = "$6$FA5jtDxj$UyJD6JZ5UpWn47bJZQvxkGJb8xG9k4X4FyYAt1ZsPK5L54V1Eyk5eumea.d7jEdRdeHAoxaGpPxj5j6KOov500"
// const hash = "$6$FA5jtDxj$zmYqs3rVUlk9ZRnyq5b1HM53XXso6Cu4h4UElbCEK5z/v.VbyXLQsvZL4Kxj0vd2OfIIXnOA9XNgogrdNm8tb/"
const salt = "$6$FA5jtDxj$"

const prefix = "<:cti18:>"
var byteSalt []byte = []byte(salt)
var arg1 string = os.Args[1]

func wrapper() {

	var wg sync.WaitGroup
	c := sha512_crypt.New()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := 0; i < alphabetSize; i++ {

		wg.Add(alphabetSize)
		go Gen3(ctx, cancel, &wg, string(ALPHABET[i]), c)
	}

	wg.Wait()
}

func Gen3(ctx context.Context, cancel context.CancelFunc, wg *sync.WaitGroup, first string, c crypt.Crypter) {
	defer wg.Done()
	var s string

	for i := 0; i < alphabetSize; i++ {
		for j := 0; j < alphabetSize; j++ {
			for k := 0; k < alphabetSize; k++ {
				s = fmt.Sprintf("%s%s%s%s%s%sn", prefix, arg1, first, string(ALPHABET[i]), string(ALPHABET[j]), string(ALPHABET[k]))
				// fmt.Println(s)
				fmt.Println(".")

				hashed, err := c.Generate([]byte(s), byteSalt)
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
