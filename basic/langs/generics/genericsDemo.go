package generics

import (
	"fmt"
	"math/rand"
	"time"
)

//https://www.digitalocean.com/community/tutorials/how-to-use-generics-in-go

type PlayingCard struct {
	Suit string
	Rank string
}

type Deck struct {
	cards []interface{}
}

func NewPlayingCard(suite string, rank string) *PlayingCard {
	return &PlayingCard{
		Suit: suite,
		Rank: rank,
	}
}
func (pc *PlayingCard) String() string {
	return fmt.Sprintf("%s of %s", pc.Rank, pc.Suit)
}
func (d *Deck) AddCard(card interface{}) {
	d.cards = append(d.cards, card)
}

func (d *Deck) RandomCard() interface{} {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	cardIdx := r.Intn(len(d.cards))
	return d.cards[cardIdx]
}

func NewPlayingCardDeck() *Deck {
	suits := []string{"Diamonds", "Hearts", "Clubs", "Spades"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	deck := &Deck{}
	for _, suit := range suits {
		for _, rank := range ranks {
			deck.AddCard(NewPlayingCard(suit, rank))
		}
	}
	return deck
}

type GenericDeck[C any] struct {
	cards []C
}

func (d *GenericDeck[C]) AddCard(card C) {
	d.cards = append(d.cards, card)
}

func (d *GenericDeck[C]) RandomCard() C {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	cardIdx := r.Intn(len(d.cards))
	return d.cards[cardIdx]
}

func NewPlayingCardGenericDeck() *GenericDeck[*PlayingCard] {
	suits := []string{"Diamonds", "Hearts", "Clubs", "Spades"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	deck := &GenericDeck[*PlayingCard]{}
	for _, suit := range suits {
		for _, rank := range ranks {
			deck.AddCard(NewPlayingCard(suit, rank))
		}
	}
	return deck
}

type TradingCard struct {
	CollectableName string
}

func NewTradingCard(collectableName string) *TradingCard {
	return &TradingCard{CollectableName: collectableName}
}

func (tc *TradingCard) String() string {
	return tc.CollectableName
}

func NewTradingCardDeck() *GenericDeck[*TradingCard] {
	collectables := []string{"Sammy", "Droplets", "Spaces", "App Platform"}

	deck := &GenericDeck[*TradingCard]{}
	for _, collectable := range collectables {
		deck.AddCard(NewTradingCard(collectable))
	}
	return deck
}
