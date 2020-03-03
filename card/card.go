package card

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
)

type suit int

const (
	Heart suit = iota + 1
	Diamond
	Club
	Spade
)

type Card struct {
	num int
	suit suit
}

type Deck []Card

func (s suit) String () string  {
	switch s {
	case Heart:
		return "H"
	case Diamond:
		return "D"
	case Club:
		return "C"
	case Spade:
		return "S"
	default:
		return "Unknown Suit"
	}
}

func (c Card) String () string {
	return c.suit.String() + strconv.Itoa(c.num)
}

func AllDeck () Deck {
	var cs []Card
	for i := 1 ; i < 14; i++ {
		for j := Heart; j <= Spade; j++ {
			cs = append(cs, Card{i, j})
		}
	}
	return cs
}

func (dk Deck) Shuffle() Deck{
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(dk), func(i, j int) {
		dk[i], dk[j] = dk[j], dk[i]
	})
	return dk
}

func (c Card) Num () int {
	return c.num
}

func Draw (dk Deck) (Deck, *Card, error) {
	if len(dk) == 0 {
		return []Card{}, nil, errors.New("empty deck")
	}
	card := dk[0]
	dk = dk[1:]
	return dk, &card, nil
}

func (c Card) BjVals() []int {
	switch c.num {
	case 1:
		return []int{1, 10}
	case 11:
		return []int{10}
	case 12:
		return []int{10}
	case 13:
		return []int{10}
	default:
		return []int{c.num}
	}
}
