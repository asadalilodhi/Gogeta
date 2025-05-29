package main

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWallet(t *testing.T) {

	Convey("Test Wallet", t, func() {
		wallet := Wallet{}
		wallet.deposit(Bitcoin(10))
		result := wallet.balance()
		fmt.Printf("address of balance in test is %p \n", &wallet._balance)
		Convey("Check balance after deposit", func() {
			So(result, ShouldEqual, 10)
		})
	})

	Convey("Test Wallet with withdraw", t, func() {
		wallet := Wallet{Bitcoin(20)}
		wallet.withdraw(Bitcoin(10))
		result := wallet.balance()
		Convey("Check balance after withdraw", func() {
			So(result, ShouldEqual, Bitcoin(10))
		})
	})

	Convey("withdrawing more than balance should fail", t, func() {
		starting_balance := Bitcoin(20)
		wallet := Wallet{starting_balance}
		intentional_invalid_withdraw := wallet.withdraw(Bitcoin(100))
		result := wallet.balance()
		Convey("The balance should remian unchanged", func() {
			So(result, ShouldEqual, starting_balance)
		})

		Convey("An error should be returned", func() {
			fmt.Println("Withdraw error:", intentional_invalid_withdraw)
			So(intentional_invalid_withdraw, ShouldEqual, ErrInsufficientFunds)
		})
	})
}