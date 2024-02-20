package service

import (
    "context"
    "order-processing-service/pkg/model"
    pbc "order-processing-service/pkg/protobuf/cart"
    "order-processing-service/pkg/repository"
)

type CartService struct {
    CartRepository    repository.CartRepository
    ProductRepository repository.ProductRepository
}

func (c *CartService) AddItem(ctx context.Context, req *pbc.AddItemRequest) (*pbc.AddItemResponse, error) {
    product, err := c.ProductRepository.FindById(req.ProductId)
    if err != nil {
        return nil, err
    }

    item := c.CartRepository.FindItem(req.UserId, req.ProductId)
    if item == nil {
        item = &model.Cart{
            Product:  product,
            UserId:   req.UserId,
            Quantity: 0,
        }
    }

    item.Quantity = item.Quantity + 1

    err = c.CartRepository.SaveOrUpdate(item)
    if err != nil {
        return nil, err
    }

    items, err := c.CartRepository.FindByUserId(req.UserId)
    if err != nil {
        return nil, err
    }

    return &pbc.AddItemResponse{
        Items: convertItems(items),
    }, nil
}

func (c *CartService) RemoveItem(ctx context.Context, req *pbc.RemoveItemRequest) (*pbc.RemoveItemResponse, error) {
    product, err := c.ProductRepository.FindById(req.ProductId)
    if err != nil {
        return nil, err
    }

    item := c.CartRepository.FindItem(req.UserId, req.ProductId)
    if err != nil {
        return nil, err
    }

    item.Quantity = item.Quantity - req.Quantity

    if item.Quantity <= 0 {
        err := c.CartRepository.RemoveItem(int32(product.ID))
        if err != nil {
            return nil, err
        }
    } else {
        err := c.CartRepository.SaveOrUpdate(item)
        if err != nil {
            return nil, err
        }
    }

    items, err := c.CartRepository.FindByUserId(req.UserId)
    if err != nil {
        return nil, err
    }

    return &pbc.RemoveItemResponse{
        Items: convertItems(items),
    }, nil
}

func convertItems(items []*model.Cart) []*pbc.CartItem {
    var result []*pbc.CartItem
    for _, item := range items {
        result = append(result, &pbc.CartItem{
            Id:       int64(item.ProductID),
            Price:    item.Product.Price,
            Quantity: item.Quantity,
            Name:     item.Product.Name,
        })
    }

    return result
}
