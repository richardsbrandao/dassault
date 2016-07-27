package structs

type Transaction struct {
	Sku                int64
	Unit_price         int16
	Quantity           int16
	Id                 int64
	External_id        int64
	Amount             int16
	Pay_type           Pay_type
	Authorization_code string
	Card_brand         string
	Card_bin           string
	Card_last          string
}


type Pay_type int

const (
	PAYMENT Status = 1 + iota
	CANCEL
)

var pay_type = [...]string{
	"PAYMENT",
	"CANCEL",
}

func (p Pay_type) String() string { return status[p-1] }
