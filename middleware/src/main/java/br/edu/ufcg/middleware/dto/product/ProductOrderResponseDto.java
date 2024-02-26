package br.edu.ufcg.middleware.dto.product;

import br.edu.ufcg.middleware.proto.order.OrderSaveResponse;

public record ProductOrderResponseDto(String name, Double price, Long quantity) {
    public ProductOrderResponseDto(OrderSaveResponse.Data.Product product) {
        this(
                product.getName(),
                product.getPrice(),
                product.getQuantity()
        );
    }
}
