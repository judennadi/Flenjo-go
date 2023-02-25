package bank_test

import (
	"testing"

	"github.com/judennadi/flenjo-go/bank"
)

func TestCreateAcct(t *testing.T) {
	fidelityAcct := bank.NewFidelity()
	zenithAcct := bank.NewZenith()

	expectedFidelityAcct := &bank.Fidelity{Balance: 0}
	expectedZenithAcct := &bank.Zenith{Balance: 0}

	if *fidelityAcct != *expectedFidelityAcct {
		t.Errorf("Expected an Empty Fidelity acct")
	}

	if *zenithAcct != *expectedZenithAcct {
		t.Errorf("Expected an Empty Zenith acct")
	}

}

func TestDeposit(t *testing.T) {
	amount := 500
	fidelityAcct := &bank.Fidelity{Balance: 0}
	zenithAcct := &bank.Fidelity{Balance: 0}

	fidelityAcct.Deposit(amount)
	fidelityAcct.Deposit(amount)
	zenithAcct.Deposit(amount)
	zenithAcct.Deposit(amount)

	if fidelityAcct.Balance != amount*2 {
		t.Errorf("Expected Fidelity Balance (%v) to be equal to amount (%v)", fidelityAcct.Balance, amount*2)
	}

	if zenithAcct.Balance != amount*2 {
		t.Errorf("Expected Zenith Balance (%v) to be equal to amount (%v)", zenithAcct.Balance, amount*2)
	}
}

func TestWithdrawal(t *testing.T) {
	amount := 500
	fidelityAcct := &bank.Fidelity{Balance: amount}
	zenithAcct := &bank.Fidelity{Balance: amount}

	fidelityAcct.Withdraw(amount)
	zenithAcct.Withdraw(amount)

	if fidelityAcct.Balance != 0 {
		t.Errorf("Expected Fidelity Balance (%v) to be equal to zero", fidelityAcct.Balance)
	}

	if zenithAcct.Balance != 0 {
		t.Errorf("Expected Zenith Balance (%v) to be equal to zero", zenithAcct.Balance)
	}
}
