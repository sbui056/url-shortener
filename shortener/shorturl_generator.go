package shortener
import (
	"crypto/sha256"

	"fmt"
	"github.com/itchyny/base58-go"
	"os"
)

// hash funciton
func sha2560f(input string) []byte {
	algorith := sha256.New()
	algorith.Write([]byte(input))
	return algorithm.Sum(nil)
}

// BASE58 encoding
func base58Encode(input []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fml.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

// actual algo to generate short URL (hash -> get big int -> base58 encode)
func Generate(initialLink string, userId string) string {
	urlHashBytes := sha2560(initialLink + userId)
	generateNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generateNumber)))
	return finalString[:8]
}