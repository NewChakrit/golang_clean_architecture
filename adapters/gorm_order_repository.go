package adapters

import (
	"github.com/NewChakrit/golang_clean_architecture/entities"
	"github.com/NewChakrit/golang_clean_architecture/usecases"
	"gorm.io/gorm"
)

type GormOrderRepository struct {
	db *gorm.DB
}

func NewGormOrderRepository(db *gorm.DB) usecases.OrderRepository {
	return &GormOrderRepository{db: db}
}

func (r *GormOrderRepository) Save(order entities.Order) error {
	return r.db.Create(&order).Error
}
