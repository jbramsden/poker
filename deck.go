package poker

import (
	"fmt"
	"math/rand"
	"time"
)

var fullDeck *Deck

func init() {
	fullDeck = &Deck{initializeFullCards()}
	rand.Seed(time.Now().UnixNano())
}

type Deck struct {
	cards []Card
}

//NewDeck - creates a new deck of cards and shuffles them.
func NewDeck() *Deck {
	deck := &Deck{}
	deck.Shuffle()
	return deck
}

//NewEmpty - creates an empty deck of cards
func NewEmpty() *Deck {
	deck := &Deck{}
	return deck
}

//Cards - Returns the cards in the Deck
func (deck *Deck) Cards() []Card {
	return deck.cards
}

//NumberOfCards - returns the number of cards in the deck.
func (deck *Deck) NumberOfCards() int {
	return len(deck.cards)
}

//Shuffle - mixes the cards order.
func (deck *Deck) Shuffle() {
	deck.cards = make([]Card, len(fullDeck.cards))
	copy(deck.cards, fullDeck.cards)
	rand.Shuffle(len(deck.cards), func(i, j int) {
		deck.cards[i], deck.cards[j] = deck.cards[j], deck.cards[i]
	})
}

//Draw - takes the number of cards from the deck and returns a slice of cards.
func (deck *Deck) Draw(n int) []Card {
	cards := make([]Card, n)
	copy(cards, deck.cards[:n])
	deck.cards = deck.cards[n:]
	return cards
}

//Empty - Returns true if the deck is empty
func (deck *Deck) Empty() bool {
	return len(deck.cards) == 0
}

func (deck *Deck) String() string {
	var str string
	for _, i := range deck.cards {
		str = fmt.Sprintf("%s%s", str, i.String())
	}
	return str
}

//ConCat - Concatenates two decks together and returns a new Deck.
func (deck *Deck) ConCat(second *Deck) *Deck {
	newD := &Deck{}
	newD.cards = append(second.cards, deck.cards...)
	return newD
}

//Deal - Takes cards from the deck and transfers them to the hands provided
func (d *Deck) Deal(cards int, hands ...*Deck) {
	for i := 0; i < cards; i++ {
		for _, hand := range hands {
			card := d.cards[0]
			d.cards = d.cards[1:]
			hand.cards = append(hand.cards, card)
		}
	}
}

func initializeFullCards() []Card {
	var cards []Card

	for _, rank := range strRanks {
		for suit := range charSuitToIntSuit {
			cards = append(cards, NewCard(string(rank)+string(suit)))
		}
	}

	return cards
}
