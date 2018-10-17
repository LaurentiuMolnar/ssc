package main

import (
	"fmt"
	"github.com/tredoe/osutil/user/crypt/sha512_crypt"
	"os"
)

const ALPHABET = "abcdefghijklmnopqrstuvwxz0123456789!@#$%^&*()"
const alphabetSize = len(ALPHABET)

const hash = "$6$FA5jtDxj$OeF/yrnIG37pI5Jfz9k8bHKDDCapJFCvHqaRBg.d2XPkvhiXOzq7Yak5n2PX9L/vIHB3oN.TdFQmeA.nVLBwC0"
// const hash = "$6$FA5jtDxj$zmYqs3rVUlk9ZRnyq5b1HM53XXso6Cu4h4UElbCEK5z/v.VbyXLQsvZL4Kxj0vd2OfIIXnOA9XNgogrdNm8tb/"
const salt = "$6$FA5jtDxj$"

var byteSalt []byte = []byte(salt)
var arg1 string = os.Args[1]

func main() {

	c := sha512_crypt.New()

	for i := 0; i < alphabetSize; i++ {
		for j := 0; j < alphabetSize; j++ {
			for k := 0; k < alphabetSize; k++ {
				for l := 0; l < alphabetSize; l++ {
					s := "<:cti18:>" + string(arg1) + string(ALPHABET[i]) + string(ALPHABET[j]) + string(ALPHABET[k]) + string(ALPHABET[l]) + "n"

					// fmt.Println(s)

					hashed, err := c.Generate([]byte(s), byteSalt)

					if err != nil {
						fmt.Errorf("Some error occurred: %s", err)
						os.Exit(1) // we had an error
					}

					if hashed == hash {
						fmt.Printf("Found password: %s\n", s)
						os.Exit(33)
					}

				}
			}
		}
	}

}
