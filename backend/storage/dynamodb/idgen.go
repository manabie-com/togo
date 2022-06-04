package dynamodb

import (
	"errors"
	"fmt"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

type Alphabet string

const (
	Lowercase    Alphabet = "abcdefghijklmnopqrstuvwxyz"
	Uppercase    Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numbers      Alphabet = "0123456789"
	Symbols      Alphabet = "_-"
	Alphanumeric Alphabet = Numbers + Uppercase + Lowercase
	Uppernumeric Alphabet = Uppercase + Numbers
	Fullalphabet Alphabet = Alphanumeric + Symbols
)

// Generates a unique ID string using nanoid with length equal to 'size'. This handles retry if ID collision happens. To trigger the retry, the parameter f must return an ErrIDExists error.
// The parameter size ranges from 5 to 25, while retry ranges from 0 to 10.
func generateID(f func(string) error, size, retry int, alphabet Alphabet) (*string, error) {
	if size < 5 || size > 25 {
		// TODO: See if this is the correct implementation for developer errors
		panic("size is out of range")
	}
	if retry < 0 || retry > 10 {
		// TODO: See if this is the correct implementation for developer errors
		panic("retry is out of range")
	}

	for i := 0; i <= retry; i++ {
		id, err := gonanoid.Generate(string(alphabet), size)
		if err != nil {
			return nil, fmt.Errorf("failed to generate id: %w", err)
		}
		if err = f(id); err != nil {
			if errors.Is(err, ErrIDExists) && i < retry {
				continue
			}
			return nil, err
		}
		return &id, nil
	}
	// Function should not reach here
	panic("unexpected error: function should not have reached here")
}
