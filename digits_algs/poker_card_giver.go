package digitsalgs

import (
	"errors"
	"fmt"
	"main/simple"
)

func Poker_card_giver(numberOfPlayers int) ([][]string, error) {
	if numberOfPlayers > 8 {
		return nil, errors.New("too many players")
	}

	hearts := []string{}
	diamonds := []string{}
	clubs:= []string{}
	spades := []string{}

	for i := 1; i <= 13; i++ {
		hearts = append(hearts, fmt.Sprint("h", i))
		diamonds = append(diamonds, fmt.Sprint("d", i))
		clubs = append(clubs, fmt.Sprint("c", i))
		spades = append(spades, fmt.Sprint("s", i))
	}

	merged := hearts
	merged = append(merged, diamonds...)
	merged = append(merged, clubs...)
	merged = append(merged, spades...)

	randomized := simple.Randomize_String_array(merged)

	chunked := simple.ChunkSlicer(randomized, 5)

	result := [][]string{}

	for i := 0; i <= numberOfPlayers; i++ {
		result = append(result, chunked[i])
	}

	return result, nil
}