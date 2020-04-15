# Poker

[![CircleCI](https://circleci.com/gh/chehsunliu/poker/tree/master.svg?style=shield&circle-token=abebd63b852ce8ecdcdf3f7e597be743d07402e4)](https://circleci.com/gh/chehsunliu/poker/tree/master) [![GoDoc](https://godoc.org/github.com/jbramsden/poker?status.svg)](https://godoc.org/github.com/jbramsden/poker) [![codecov](https://codecov.io/gh/chehsunliu/poker/branch/master/graph/badge.svg)](https://codecov.io/gh/chehsunliu/poker)

Poker is ported from the Python library [worldveil/deuces](https://github.com/worldveil/deuces).

## Changes on this fork

1) Added NewEmpty to Deck to be able to add an empty deck of cards. Use case when you want players to have a container ready to be dealt cards.
2) Added Deal to Deck to be able to deal X number of cards to other decks. 
4) Added ConCat to Deck which will concatenate of deck of cards to another deck of cards
5) Added Cards function to Deck which returns a slice of cards. E.G Exposing the card attribute.
6) Changed String in Card so that it uses prettyprint on the suite.

## Installation

Use `go get` to install Poker:

```sh
$ go get github.com/jbramsden/poker
```

## Usage

Support 5-, 6-, and 7-card evalutions:

```go
package main

import (
	"fmt"

	"github.com/jbramsden/poker"
)

func main() {

	//Texas Hold'em example with 2 hands.
	deck := poker.NewDeck()
	handA := poker.NewEmpty()
	handB := poker.NewEmpty()
	table := poker.NewEmpty()

	//Deal 2 cards to handA and handB
	deck.Deal(2, handA, handB)
	fmt.Printf("HandA - %s, HandB - %s\n", handA, handB)

	//Deal 3 cards to the table
	deck.Deal(3, table)
	fmt.Printf("Table - %s\n", table)

	//Deal 1 card to the table
	deck.Deal(1, table)
	fmt.Printf("Table - %s\n", table)

	//Deal final card to the table
	deck.Deal(1, table)
	fmt.Printf("Table - %s\n", table)

	handAFinal := table.ConCat(handA)
	handBFinal := table.ConCat(handB)

	fmt.Printf("A - %s | B - %s\n", handAFinal, handBFinal)
	handAScore := poker.Evaluate(handAFinal.Cards())
	handBScore := poker.Evaluate(handBFinal.Cards())

	fmt.Printf("hand A score %d, Hand B Score %d\n", handAScore, handBScore)
	if handAScore < handBScore {
		fmt.Printf("Hand A is winner with %s\n", poker.RankString(handAScore) )
	}else if handBScore < handAScore {
		fmt.Printf("Hand B is winner with %s\n", poker.RankString(handBScore) )
	}else{
		fmt.Printf("It is a draw!")
	}

}
```

```sh
$ go run main.go
HandA - 6♠9❤, HandB - 2♦6♣
Table - J♠8♠5♠
Table - J♠8♠5♠5♦
Table - J♠8♠5♠5♦9♠
A - 6♠9❤J♠8♠5♠5♦9♠ | B - 2♦6♣J♠8♠5♠5♦9♠
hand A score 1414, Hand B Score 5449
Hand A is winner with Flush

$ go run main.go
HandA - 5❤5♣, HandB - 7❤Q♠
Table - 8♦K♣J♠
Table - 8♦K♣J♠9❤
Table - 8♦K♣J♠9❤J♣
A - 5❤5♣8♦K♣J♠9❤J♣ | B - 7❤Q♠8♦K♣J♠9❤J♣
hand A score 2887, Hand B Score 4042
Hand A is winner with Two Pair


$ go run main.go
HandA - 8♣K❤, HandB - J♣9♣
Table - J❤A♠Q❤
Table - J❤A♠Q❤4♣
Table - J❤A♠Q❤4♣3♠
A - 8♣K❤J❤A♠Q❤4♣3♠ | B - J♣9♣J❤A♠Q❤4♣3♠
hand A score 6187, Hand B Score 3997
Hand B is winner with Pair
```

## Performance

Compared with [notnil/joker](https://github.com/notnil/joker), Poker is 160x faster on 5-card evaluation, and drops to 40x faster on 7-card evaluation.

```sh
go test -bench=. -benchtime 5s
goos: darwin
goarch: amd64
pkg: github.com/chehsunliu/poker
BenchmarkFivePoker-4    	23396181	       253 ns/op
BenchmarkFiveJoker-4    	  141036	     41662 ns/op
BenchmarkSixPoker-4     	 3037298	      1949 ns/op
BenchmarkSixJoker-4     	   28158	    211533 ns/op
BenchmarkSevenPoker-4   	  356448	     16357 ns/op
BenchmarkSevenJoker-4   	    7143	    759394 ns/op
PASS
ok  	github.com/chehsunliu/poker	40.111s
```
