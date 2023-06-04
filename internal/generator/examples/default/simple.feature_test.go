package examples_test

import (
	"testing"

	"github.com/hedhyw/gherkingen/v3/pkg/bdd"
)

func TestGuessTheWord(t *testing.T) {
	t.Parallel()

	f := bdd.NewFeature(t, "Guess the word")

	/*
		The word guess game is a turn-based game for two players.
		The Maker makes a word for the Breaker to guess. The game
		is over when the Breaker guesses the Maker's word.
	*/

	f.Example("Maker starts a game", func(t *testing.T, f *bdd.Feature) {
		t.Parallel()

	})
}
