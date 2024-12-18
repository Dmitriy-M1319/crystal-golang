package baseapp

import "time"

type User struct {
	ID       uint64
	Name     string
	Surname  string
	Email    string
	Phone    string `db:"phone_number"`
	Password string
	IsAdmin  bool `db:"is_admin"`
}

type Product struct {
	ID            uint64
	Name          string `db:"product_name"`
	Company       string
	ClientPrice   float64 `db:"client_price"`
	PurchasePrice float64 `db:"purchase_price"`
	Count         int32
}

type Order struct {
	ID          uint64
	ClientID    uint64  `db:"client_id"`
	TotalPrice  float64 `db:"total_price"`
	Address     string
	IsDelivery  bool      `db:"is_delivery"`
	PaymentType string    `db:"payment_type"`
	Status      bool      `db:"order_status"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type OrderWithProducts struct {
	Order    Order
	User     User
	Products []Product
	Counts   []int32
}
