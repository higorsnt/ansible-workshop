package service

import (
    "context"
    "net/http"
    "order-processing-service/pkg/model"
    productpb "order-processing-service/pkg/protobuf/product"
    "order-processing-service/pkg/repository"
)

type ProductService struct {
    productpb.UnimplementedProductServiceServer

    ProductRepository repository.ProductRepository
}

func (ps *ProductService) Save(ctx context.Context, req *productpb.ProductSaveRequest) (*productpb.ProductSaveResponse, error) {
    product := model.Product{
        Description: req.Description,
        Name:        req.Name,
        Price:       req.Price,
        Stock:       req.Stock,
    }

    err := ps.ProductRepository.Save(&product)
    if err != nil {
        return nil, err
    }

    return &productpb.ProductSaveResponse{
        Status: http.StatusCreated,
        Data: &productpb.ProductSaveResponse_Data{
            Id:          int64(product.ID),
            Name:        product.Name,
            Description: product.Description,
            Price:       product.Price,
            Stock:       product.Stock,
        },
    }, nil
}

func (ps *ProductService) List(ctx context.Context, req *productpb.ProductListRequest) (*productpb.ProductListResponse, error) {
    name := ""
    priceMin := 0.0
    priceMax := 0.0

    if req.Filters != nil {
        if req.Filters.Name != nil {
            name = *req.Filters.Name
        }

        if req.Filters.PriceMin != nil {
            priceMin = *req.Filters.PriceMin
        }

        if req.Filters.PriceMax != nil {
            priceMax = *req.Filters.PriceMax
        }
    }

    list, err := ps.ProductRepository.List(name, priceMin, priceMax)
    if err != nil {
        return nil, err
    }

    var result []*productpb.ProductListResponse_Product
    for _, e := range list {
        result = append(result, &productpb.ProductListResponse_Product{
            Id:          uint32(e.ID),
            Name:        e.Name,
            Description: e.Description,
            Price:       e.Price,
            Stock:       e.Stock,
        })
    }

    return &productpb.ProductListResponse{
        Products: result,
    }, nil
}

func (ps *ProductService) Update(ctx context.Context, req *productpb.ProductUpdateRequest) (*productpb.ProductUpdateResponse, error) {
    product, err := ps.ProductRepository.FindById(int64(req.Id))
    if err != nil {
        return nil, err
    }

    if req.Name != nil {
        product.Name = *req.Name
    }

    if req.Price != nil {
        product.Price = *req.Price
    }

    if req.Description != nil {
        product.Description = *req.Description
    }

    if req.Stock != nil {
        product.Stock = *req.Stock
    }

    err = ps.ProductRepository.Update(&product)
    if err != nil {
        return nil, err
    }

    return &productpb.ProductUpdateResponse{
        Id:          uint32(product.ID),
        Name:        product.Name,
        Description: product.Description,
        Price:       product.Price,
        Stock:       product.Stock,
    }, nil
}

func (ps *ProductService) Delete(ctx context.Context, req *productpb.ProductDeleteRequest) (*productpb.ProductDeleteResponse, error) {
    err := ps.ProductRepository.Delete(req.Id)
    if err != nil {
        return nil, err
    }

    return &productpb.ProductDeleteResponse{}, nil
}
