package baseapp

import "time"

type OrderService struct {
	userRepo    IUserRepository
	orderRepo   IOrderRepository
	productRepo IProductRepository
}

func NewOrderService(uRepo IUserRepository, oRepo IOrderRepository, pRepo IProductRepository) *OrderService {
	return &OrderService{userRepo: uRepo, orderRepo: oRepo, productRepo: pRepo}
}

func (s *OrderService) GetOrdersInfo(from time.Time, to time.Time) ([]OrderWithProducts, error) {
	orders, err := s.orderRepo.OrdersByPeriod(from, to)
	if err != nil {
		return nil, err
	}

	result := make([]OrderWithProducts, 0)

	for _, order := range orders {
		products, err := s.productRepo.ProductsByOrder(order.ID)
		if err != nil {
			return nil, err
		}

		client, err := s.userRepo.UserById(order.ClientID)
		if err != nil {
			return nil, err
		}

		counts := make([]int32, 0)

		for _, product := range products {
			count, err := s.productRepo.ProductCountByOrder(order.ID, product.ID)
			if err != nil {
				return nil, err
			}

			counts = append(counts, count)
		}

		result = append(result, OrderWithProducts{order: order, products: products, user: client, counts: counts})
	}

	return result, nil
}
