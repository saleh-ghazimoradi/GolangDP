package facade

import "fmt"

type Account struct {
	Name string
}

func (a *Account) CheckAccount(accountName string) error {
	if a.Name != accountName {
		return fmt.Errorf("account name is not correct")
	}
	fmt.Println("Account is verified")
	return nil
}

func NewAccount(accountName string) *Account {
	return &Account{Name: accountName}
}

type SecurityCode struct {
	Code int
}

func (s *SecurityCode) CheckCode(incomingCode int) error {
	if s.Code != incomingCode {
		return fmt.Errorf("security Code is incorrect")
	}
	fmt.Println("SecurityCode Verified")
	return nil
}

func NewSecurityCode(code int) *SecurityCode {
	return &SecurityCode{
		Code: code,
	}
}

type Wallet struct {
	Balance int
}

func (w *Wallet) CreditBalance(amount int) {
	w.Balance += amount
	fmt.Println("Wallet balance added successfully")
	return
}

func (w *Wallet) DebitBalance(amount int) error {
	if w.Balance < amount {
		return fmt.Errorf("balance is not sufficient")
	}
	fmt.Println("Wallet Balance is Sufficient")
	w.Balance = w.Balance - amount
	return nil
}

func NewWallet() *Wallet {
	return &Wallet{
		Balance: 0,
	}
}

type Ledger struct{}

func (s *Ledger) MakeEntry(accountID, txnType string, amount int) {
	fmt.Printf("Make ledger entry for accountId %s with txnType %s for amount %d\n", accountID, txnType, amount)
	return
}

type Notification struct{}

func (n *Notification) SendWalletCreditNotification() {
	fmt.Println("Sending wallet credit notification")
}

func (n *Notification) SendWalletDebitNotification() {
	fmt.Println("Sending wallet debit notification")
}

type WalletFacade struct {
	Account      *Account
	Wallet       *Wallet
	SecurityCode *SecurityCode
	Notification *Notification
	Ledger       *Ledger
}

func (w *WalletFacade) AddMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.Account.CheckAccount(accountID)
	if err != nil {
		return err
	}
	err = w.SecurityCode.CheckCode(securityCode)
	if err != nil {
		return err
	}
	w.Wallet.CreditBalance(amount)
	w.Notification.SendWalletCreditNotification()
	w.Ledger.MakeEntry(accountID, "credit", amount)
	return nil
}

func (w *WalletFacade) DeductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting debit money from wallet")
	err := w.Account.CheckAccount(accountID)
	if err != nil {
		return err
	}

	err = w.SecurityCode.CheckCode(securityCode)
	if err != nil {
		return err
	}
	err = w.Wallet.DebitBalance(amount)
	if err != nil {
		return err
	}
	w.Notification.SendWalletDebitNotification()
	w.Ledger.MakeEntry(accountID, "debit", amount)
	return nil
}

func NewWalletFacade(accountID string, code int) *WalletFacade {
	fmt.Println("Starting create account")
	walletFacade := &WalletFacade{
		Account:      NewAccount(accountID),
		SecurityCode: NewSecurityCode(code),
		Wallet:       NewWallet(),
		Notification: &Notification{},
		Ledger:       &Ledger{},
	}
	fmt.Println("Account created")
	return walletFacade
}
