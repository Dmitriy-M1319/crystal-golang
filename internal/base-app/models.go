package baseapp

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
}

type ProductApplication struct {
	ID        uint64
	ProductId uint64
	Count     int32
	Provider  string
	Price     float64
	Status    bool
}
