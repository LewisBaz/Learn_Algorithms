package digitsalgs

import (
	"errors"
	"math/rand"
)

func DiceThatTossACoin() (string, error) {
	diceValue := rand.Intn(6) + 1
	switch diceValue {
	case 1, 2, 3:
		return "heads", nil
	case 4, 5, 6:
		return "tails", nil
	default:
		return "", errors.New("")
	}
}