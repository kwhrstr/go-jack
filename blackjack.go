package main

import (
	"bufio"
	"fmt"
	. "go-jack/card"
	"os"
	"sort"
	)

type PlayerBurst struct {}

func (e *PlayerBurst) Error() string  {
	return "player is burned !! dealer won"
}

type DealerBurst struct {}

func (e *DealerBurst) Error() string {
	return "dealer is burned !! player won"
}

type Game struct {
	player []Card
	dealer []Card
	deck Deck
}

func initialGame () Game {
	dk := AllDeck().Shuffle()
	var player []Card
	var dealer []Card
	for i := 0; i < 2; i ++ {
		p := new(Card)
		d := new(Card)
		dk, p, _ = Draw(dk)
		dk, d, _ = Draw(dk)
		player = append(player, *p)
		dealer = append(dealer, *d)
	}
	return Game{player, dealer, dk}
}

func calcBjVals(cs []Card) []int {
	var nss [][] int
	for _, c := range cs {
		nss = append(nss, c.BjVals())
	}
	sums := []int {0}
	for _, ns := range nss {
		sums = someEach(ns, sums)
	}
	isNotBurst := func(i int) bool {
		return i <= 21
	}
	filtered := filter(sums, isNotBurst)

	return filtered
}

func someEach(ms []int, ns []int ) []int {
	var sums []int
	for _, m := range ms {
		for _, n := range ns {
			sums = append(sums, m + n)
		}
	}
	return sums
}

func filter(xs []int, f func(int) bool)[]int  {
	var ys []int
	for _, x := range xs {
		if f(x) {
			ys = append(ys, x)
		}
	}
	return ys
}


func playerTurn (game *Game) (*int, error) {
	bjVal := new(int)
	for ; ; {
		bjVals := calcBjVals(game.player)
		if len(bjVals) == 0 {
			return nil, &PlayerBurst{}
		}
		reader := bufio.NewReaderSize(os.Stdin, 1)
		fmt.Println("Hit(h) or Stand(s)")
		ch, _ := reader.ReadByte()
		if ch == 's' {
			sort.Ints(bjVals)
			bjVal = &bjVals[len(bjVals) - 1]
			break
		}
		if ch == 'h' {
			rest, c, e := Draw(game.deck)
			if e != nil {
				return nil, e
			}
			game.deck = rest
			game.player = append(game.player, *c)
			fmt.Println(*c)
			continue
		}
		fmt.Println("Please Input h or s !!")
	}
	return bjVal, nil
}

func dealerTurn (game *Game) (*int, error) {
	bjVal := new(int)
	for ; ;  {
		bjVals := calcBjVals(game.dealer)
		if len(bjVals) == 0 {
			fmt.Println(game.dealer)
			return nil, &DealerBurst{}
		}
		sort.Ints(bjVals)
		maximum := bjVals[len(bjVals) - 1]
		if maximum >= 17 {
			bjVal = &maximum
			break
		}
		rest, c, e := Draw(game.deck)
		if e != nil {
			return  nil, e
		}
		game.deck = rest
		game.dealer = append(game.dealer, *c)
	}
	fmt.Println(game.dealer)
	return bjVal, nil
}


func main() {
	game := initialGame()
	fmt.Println("dealer hand is:")
	fmt.Println(game.dealer[0])
	fmt.Println("Player hand is:")
	fmt.Println(game.player)
	pVal, err := playerTurn(&game)
	if err != nil {
		fmt.Println(err)
		return
	}
	dVal, err := dealerTurn(&game)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Player point is:", *pVal)
	fmt.Println("Dealer point is:",  *dVal)
	if *pVal > *dVal {
		fmt.Println("player won")
	} else {
		fmt.Println("dealer won")
	}
	return
}
