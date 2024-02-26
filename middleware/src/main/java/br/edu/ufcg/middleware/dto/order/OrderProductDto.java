package br.edu.ufcg.middleware.dto.order;

import br.edu.ufcg.middleware.proto.order.ListUserOrdersResponse;

public record OrderProductDto(Integer id, String name, String description, Double price, Long quantity) {
    public OrderProductDto(ListUserOrdersResponse.Product product) {
        this(
                product.getId(),
                product.getName(),
                product.getDescription(),
                product.getPrice(),
                product.getQuantity()
        );
    }
}
