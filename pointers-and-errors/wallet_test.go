package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposite(10)

	got := wallet.Balance()
	want := BitCoin(10)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
