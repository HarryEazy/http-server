package poker_test

import (
	poker "http-server"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {

	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})

	// t.Run("do not read beyond the first newline", func(t *testing.T) {
	// 	in := failOnEndReader{
	// 		t,
	// 		strings.NewReader("Chris wins\n hello there"),
	// 	}

	// 	playerStore := &poker.StubPlayerStore{}

	// 	cli := poker.NewCLI(playerStore, in)
	// 	cli.PlayPoker()
	// })

}