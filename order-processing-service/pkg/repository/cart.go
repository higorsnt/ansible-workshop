package repository

import (
    "gorm.io/gorm"
    "order-processing-service/pkg/model"
)

type CartRepository interface {
    SaveOrUpdate(item *model.Cart) error
    RemoveItem(productId int32) error
    FindItem(userId int64, productId int64) *model.Cart
    FindByUserId(userId int64) ([]*model.Cart, error)
}

type cartRepositoryImpl struct {
    Db *gorm.DB
}

func NewCartRepository(Db *gorm.DB) CartRepository {
    return &cartRepositoryImpl{
        Db: Db,
    }
}

func (c cartRepositoryImpl) SaveOrUpdate(item *model.Cart) error {
    result := c.Db.Save(&item)
    return result.Error
}

func (c cartRepositoryImpl) RemoveItem(productId int32) error {
    result := c.Db.Where("product_id = ?", productId).Delete(&model.Cart{})
    return result.Error
}

func (c cartRepositoryImpl) FindItem(userId int64, productId int64) *model.Cart {
    query := c.Db.Model(&model.Cart{})
    query = query.Where("user_id = ? AND product_id = ?", userId, productId)

    var cart model.Cart
    if err := query.First(&cart).Error; err != nil {
        return nil
    }

    return &cart
}

func (c cartRepositoryImpl) FindByUserId(userId int64) ([]*model.Cart, error) {
    query := c.Db.Model(&model.Cart{})
    query = query.Where("user_id = ?", userId)

    var cartItems []*model.Cart
    if err := query.Find(&cartItems).Error; err != nil {
        return nil, err
    }

    return cartItems, nil
}
