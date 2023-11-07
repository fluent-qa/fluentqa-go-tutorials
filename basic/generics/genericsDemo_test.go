package generics

import (
	"fmt"
	"os"
	"testing"
)

func TestNewPlayingCardDeck(t *testing.T) {
	deck := NewPlayingCardDeck()
	card := deck.RandomCard()
	fmt.Printf("card is: %s \n", card)
	playingCard, ok := card.(*PlayingCard) //cast to PlayingCard Type
	if !ok {
		fmt.Printf("card received wasn't a playing card!")
		os.Exit(1)
	}
	fmt.Printf("card suit: %s\n", playingCard.Suit)
	fmt.Printf("card rank: %s\n", playingCard.Rank)
}

func TestNewPlayingCardGenericDeck(t *testing.T) {
	deck := NewPlayingCardGenericDeck()

	fmt.Printf("--- drawing playing card ---\n")
	playingCard := deck.RandomCard()
	fmt.Printf("drew card: %s\n", playingCard)
	// Code removed
	fmt.Printf("card suit: %s\n", playingCard.Suit)
	fmt.Printf("card rank: %s\n", playingCard.Rank)
}
