package main

import (
	"fmt"
	"github.com/tredoe/osutil/user/crypt/sha512_crypt"
	"os"
	"sync"
	"time"
)

const ALPHABET = "abcdefghijklmnopqrstuvwxz0123456789!@#$%^&*()"
const alphabetSize = len(ALPHABET)

//const hash = "$6$4zChM1o9$gXJPtj/HPo6siamKw38.BgnwrdfSM/mWa0UZsmH9SEm8pi5mVKSmZnD2UWbv9RFTTIShiMDxVwrd6epALRFDs/"
const hash = "$6$4zChM1o9$oC0AEfjL2XoNkv/X.sRhM8y0H0EWYzQpT9qUA0DxaA7PnIzRZ3QsXrBJGKhiLnHIHG0QNH0orobW2o2E3atLp0"
const salt = "$6$4zChM1o9$"

var byteSalt []byte = []byte(salt)
var arg1 string = os.Args[1]

func main() {
	var wg sync.WaitGroup
	wg.Add(4)
	t := time.Now()
	fmt.Printf("Started at %s\n", t)
	go func() {
		defer wg.Done()
		for i := 0; i < 12; i++ {
			c := sha512_crypt.New()
			for j := 0; j < alphabetSize; j++ {
				for k := 0; k < alphabetSize; k++ {
					for l := 0; l < alphabetSize; l++ {
						s := "<:cti18:>" + string(arg1) + string(ALPHABET[i]) + string(ALPHABET[j]) + string(ALPHABET[k]) + string(ALPHABET[l]) + "#n8"

						//fmt.Println(s)

						hashed, err := c.Generate([]byte(s), byteSalt)

						if err != nil {
							fmt.Errorf("Some error occurred: %s", err)
							os.Exit(1) // we had an error
						}

						if hashed == hash {
							fmt.Printf("Found password: %s\n", s)
							fmt.Printf("Finished in %s\n", time.Since(t))
							os.Exit(33)
						}

					}
				}
			}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 12; i < 23; i++ {
			c := sha512_crypt.New()
			for j := 0; j < alphabetSize; j++ {
				for k := 0; k < alphabetSize; k++ {
					for l := 0; l < alphabetSize; l++ {
						s := "<:cti18:>" + string(arg1) + string(ALPHABET[i]) + string(ALPHABET[j]) + string(ALPHABET[k]) + string(ALPHABET[l]) + "#n8"

						//fmt.Println(s)

						hashed, err := c.Generate([]byte(s), byteSalt)

						if err != nil {
							fmt.Errorf("Some error occurred: %s", err)
							os.Exit(1) // we had an error
						}

						if hashed == hash {
							fmt.Printf("Found password: %s\n", s)
							fmt.Printf("Finished in %s\n", time.Since(t))
							os.Exit(33)
						}

					}
				}
			}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 24; i < 36; i++ {
			c := sha512_crypt.New()
			for j := 0; j < alphabetSize; j++ {
				for k := 0; k < alphabetSize; k++ {
					for l := 0; l < alphabetSize; l++ {
						s := "<:cti18:>" + string(arg1) + string(ALPHABET[i]) + string(ALPHABET[j]) + string(ALPHABET[k]) + string(ALPHABET[l]) + "#n8"

						//fmt.Println(s)

						hashed, err := c.Generate([]byte(s), byteSalt)

						if err != nil {
							fmt.Errorf("Some error occurred: %s", err)
							os.Exit(1) // we had an error
						}

						if hashed == hash {
							fmt.Printf("Found password: %s\n", s)
							fmt.Printf("Finished in %s\n", time.Since(t))
							os.Exit(33)
						}

					}
				}
			}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 36; i < alphabetSize; i++ {
			c := sha512_crypt.New()
			for j := 0; j < alphabetSize; j++ {
				for k := 0; k < alphabetSize; k++ {
					for l := 0; l < alphabetSize; l++ {
						s := "<:cti18:>" + string(arg1) + string(ALPHABET[i]) + string(ALPHABET[j]) + string(ALPHABET[k]) + string(ALPHABET[l]) + "#n8"

						//fmt.Println(s)

						hashed, err := c.Generate([]byte(s), byteSalt)

						if err != nil {
							fmt.Errorf("Some error occurred: %s", err)
							os.Exit(1) // we had an error
						}

						if hashed == hash {
							fmt.Printf("Found password: %s\n", s)
							fmt.Printf("Finished in %s\n", time.Since(t))
							os.Exit(33)
						}

					}
				}
			}
		}
	}()
	wg.Wait()
}
