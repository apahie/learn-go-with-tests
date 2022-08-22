package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposite", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposite(10)
		assertBalance(t, wallet, BitCoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: BitCoin(20)}
		got := wallet.Withdraw(BitCoin(10))

		assertNoError(t, got)
		assertBalance(t, wallet, BitCoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := BitCoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(BitCoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want BitCoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
