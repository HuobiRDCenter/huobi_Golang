package wallet

type QueryDepositWithdrawOptionalRequest struct {
	Currency string
	From     string
	Size     string
	Direct   string
}
