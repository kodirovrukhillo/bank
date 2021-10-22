package types

//Shows money
type Money int

// AddBonus

//Currency shows currency code
type Currency string

// Currency codes
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

//PAN shows card number
type PAN string

// ID shows identification number

//Car shows the card issue infomations
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
	MinBalance Money
}
