package baseapp

import "time"

type User struct {
	ID       uint64
	Name     string
	Surname  string
	Email    string
	Phone    string
	Password string
	IsAdmin  bool
}

type Product struct {
	ID            uint64
	Name          string
	Company       string
	ClientPrice   float64
	PurchasePrice float64
	Count         int32
}

type Order struct {
	ID          uint64
	ClientID    uint64
	TotalPrice  float64
	Address     string
	IsDelivery  bool
	PaymentType string
	Status      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type OrderWithProducts struct {
	order    Order
	user     User
	products []Product
	counts   []int32
}
