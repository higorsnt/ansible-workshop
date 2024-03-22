package repository

import (
    "gorm.io/gorm"
    "order-processing-service/pkg/model"
)

type OrderRepository interface {
    Save(order *model.Order) error
    SaveOrderProduct(op []*model.OrderProduct) error
    ListByUserId(userId int32) ([]model.OrderProduct, error)
}

type orderRepositoryImpl struct {
    Db *gorm.DB
}

func NewOrderRepository(Db *gorm.DB) OrderRepository {
    return &orderRepositoryImpl{
        Db: Db,
    }
}

func (o orderRepositoryImpl) Save(order *model.Order) error {
    result := o.Db.Save(&order)
    return result.Error
}

func (o orderRepositoryImpl) SaveOrderProduct(orderProducts []*model.OrderProduct) error {
    result := o.Db.Save(&orderProducts)
    return result.Error
}

func (o orderRepositoryImpl) ListByUserId(userId int32) ([]model.OrderProduct, error) {
    query := o.Db.Model(&model.OrderProduct{})
    query = query.Joins("join orders o on order_products.order_id = o.id").Where("o.user_id = ?", userId)

    var orderProducts []model.OrderProduct
    if err := query.Find(&orderProducts).Error; err != nil {
        return nil, err
    }

    return orderProducts, nil
}
