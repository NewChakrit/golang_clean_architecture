package usecases

import "github.com/NewChakrit/golang_clean_architecture/entities"

type OrderRepository interface {
	Save(order entities.Order) error
}
