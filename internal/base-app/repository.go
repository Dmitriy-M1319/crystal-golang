package baseapp

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type IUserRepository interface {
	UserById(id uint64) (User, error)
}

type IProductRepository interface {
	ProductById(id uint64) (Product, error)
	ProductsByOrder(orderId uint64) ([]Product, error)
	ProductCountByOrder(orderId, productId uint64) (int32, error)
}

type IOrderRepository interface {
	OrderById(id uint64) (Order, error)
	OrdersByPeriod(from time.Time, to time.Time) ([]Order, error)
}

type SqlProductRepository struct {
	connection *sqlx.DB
}

func NewSqlProductRepository(conn *sqlx.DB) *SqlProductRepository {
	return &SqlProductRepository{connection: conn}
}

func (repo *SqlProductRepository) ProductById(id uint64) (Product, error) {
	var result Product = Product{}
	err := repo.connection.Get(&result, "SELECT * FROM products WHERE id = $1", id)
	if err != nil {
		return result, fmt.Errorf("failed to get product: %s", err.Error())
	}
	return result, nil
}

func (repo *SqlProductRepository) ProductsByOrder(orderId uint64) ([]Product, error) {
	var result []Product
	err := repo.connection.Select(&result, "SELECT * FROM products WHERE id IN (SELECT DISTINCT product_id FROM orders_products WHERE order_id = $1)", orderId)
	if err != nil {
		return result, fmt.Errorf("failed to get products: %s", err.Error())
	}
	return result, nil
}

func (repo *SqlProductRepository) ProductCountByOrder(orderId, productId uint64) (int32, error) {
	var result int32
	err := repo.connection.Get(&result, "SELECT product_count FROM orders_products WHERE order_id = $1 AND product_id = $2", orderId, productId)
	if err != nil {
		return result, fmt.Errorf("failed to get products count: %s", err.Error())
	}
	return result, nil
}

type SqlOrderRepository struct {
	connection *sqlx.DB
}

func NewSqlOrderRepository(conn *sqlx.DB) *SqlOrderRepository {
	return &SqlOrderRepository{connection: conn}
}

func (repo *SqlOrderRepository) OrderById(id uint64) (Order, error) {
	var result Order = Order{}
	err := repo.connection.Get(&result, "SELECT * FROM orders WHERE id = $1", id)
	if err != nil {
		return result, fmt.Errorf("failed to get order: %s", err.Error())
	}
	return result, nil
}

func (repo *SqlOrderRepository) OrdersByPeriod(from time.Time, to time.Time) ([]Order, error) {
	var result []Order
	err := repo.connection.Select(&result, "SELECT * FROM orders WHERE updated_at BETWEEN $1 AND $2", from, to)
	if err != nil {
		return result, fmt.Errorf("failed to get products: %s", err.Error())
	}
	return result, nil
}

type SqlUserRepository struct {
	connection *sqlx.DB
}

func NewSqlUserRepository(conn *sqlx.DB) *SqlUserRepository {
	return &SqlUserRepository{connection: conn}
}

func (repo *SqlUserRepository) UserById(id uint64) (User, error) {
	var result User = User{}
	err := repo.connection.Get(&result, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return result, fmt.Errorf("failed to get user: %s", err.Error())
	}
	return result, nil
}
