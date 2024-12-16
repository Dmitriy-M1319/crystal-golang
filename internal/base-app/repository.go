package baseapp

type IProductRepository interface {
	AllProducts() ([]Product, error)
	ProductById(id uint64) (Product, error)
	ProductsByOrder(orderId uint64) ([]Product, error)
	InsertProduct(p *Product) error
	UpdateProduct(id uint64, p *Product) error
	DeleteProduct(id uint64) error
}

type IOrderRepository interface {
	AllOrders() ([]Order, error)
	OrderById(id uint64) (Order, error)
	OrdersByUser(userId uint64) ([]Order, error)
	InsertOrder(o *Order) error
	UpdateOrder(id uint64, o *Order) error
	DeleteOrder(id uint64) error
}

type IProductAppRepository interface {
	AllApplications() ([]ProductApplication, error)
	ApplicationById(id uint64) (ProductApplication, error)
	InsertApplication(pa *ProductApplication) error
	UpdateApplication(id uint64, pa *ProductApplication) error
	DeleteApplication(id uint64) error
}
