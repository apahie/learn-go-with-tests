package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want BitCoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("deposite", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposite(10)
		assertBalance(t, wallet, BitCoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: BitCoin(20)}
		wallet.Withdraw(BitCoin(10))
		assertBalance(t, wallet, BitCoin(10))
	})
}
