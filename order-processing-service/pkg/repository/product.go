package repository

import (
    "gorm.io/gorm"
    "order-processing-service/pkg/model"
)

type ProductRepository interface {
    Save(p *model.Product) error
    FindById(id int64) (model.Product, error)
    List(name string, priceMin float64, priceMax float64) ([]model.Product, error)
    Update(p *model.Product) error
    Delete(id int32) error
}

type ProductRepositoryImpl struct {
    Db *gorm.DB
}

func NewProductRepository(Db *gorm.DB) ProductRepository {
    return &ProductRepositoryImpl{
        Db: Db,
    }
}

func (p ProductRepositoryImpl) Save(product *model.Product) error {
    result := p.Db.Save(&product)
    return result.Error
}

func (p ProductRepositoryImpl) FindById(id int64) (model.Product, error) {
    var result model.Product
    err := p.Db.Find(&result, id).Error
    return result, err
}

func (p ProductRepositoryImpl) List(name string, priceMin float64, priceMax float64) ([]model.Product, error) {
    query := p.Db.Model(&model.Product{})

    if name != "" {
        query = query.Where("name ILIKE ?", "%"+name+"%")
    }

    if priceMin > 0 {
        query = query.Where("price >= ?", priceMin)
    }

    if priceMax > 0 {
        query = query.Where("price <= ?", priceMax)
    }

    var products []model.Product
    if err := query.Find(&products).Error; err != nil {
        return nil, err
    }

    return products, nil
}

func (p ProductRepositoryImpl) Update(product *model.Product) error {
    result := p.Db.Save(&product)
    return result.Error
}

func (p ProductRepositoryImpl) Delete(id int32) error {
    result := p.Db.Delete(&model.Product{}, id)
    return result.Error
}
