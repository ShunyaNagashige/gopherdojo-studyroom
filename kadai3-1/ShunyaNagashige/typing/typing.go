package typing

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math/rand"
	"time"
)

func Input(r io.Reader) <-chan string {
	ch := make(chan string)

	//inputが返すチャネルchに，標準入力から(io.Stdinつまりio.Reader)の文字列を渡す
	go func(chan<- string) {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
	}(ch)

	return ch
}

func Start(ch <-chan string) {
	count := 0

	bc := context.Background()
	t := 10 * time.Second
	ctx, cancel := context.WithTimeout(bc, t)
	defer cancel()

LOOP:
	for {
		word := selectWord()

		select {
		case v := <-ch:
			if v == word {
				count++
			}
		case <-ctx.Done():
			fmt.Printf("\n%d問正解です！\n", count)
			break LOOP
		}
	}
}

func selectWord() string {
	words := []string{
		"Good morning",
		"semiconductor chips",
		"Hydrogen",
		"Helium",
		"Lithium",
		"Beryllium",
		"Boron",
		"Carbon",
		"Nitrogen",
		"Oxygen",
		"Fluorine",
		"Neon",
		"Sodium",
		"Magnesium",
		"Aluminum",
		"Silicon",
	}

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(words))

	fmt.Println(words[index] + ":")

	return words[index]
}
