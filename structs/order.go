package structs

type Order struct {
	Id           int64
	Number       int64
	Reference    string
	Status       Status
	Created_at   string //changetotime
	Updated_at   string //changetotime
	Itens 	     []OrderItem //[OrderItem]
	Notes        string
	Transactions []Transaction
	Price        int16
	Valor        int16
}


type Status int

const (
	DRAFT Status = 1 + iota
	ENTERED
	CANCELED
	PAID
	APPROVED
	REJECTED
	REENTERED
	CLOSED
)

var status = [...]string{
	"DRAFT",
	"ENTERED",
	"CANCELED",
	"PAID",
	"APPROVED",
	"REJECTED",
	"REENTERED",
	"CLOSED",
}

func (m Status) String() string { return status[m-1] }


