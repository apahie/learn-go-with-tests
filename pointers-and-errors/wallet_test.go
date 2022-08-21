package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposite", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposite(10)

		got := wallet.Balance()
		want := BitCoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: BitCoin(20)}
		wallet.Withdraw(BitCoin(10))

		got := wallet.Balance()
		want := BitCoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
