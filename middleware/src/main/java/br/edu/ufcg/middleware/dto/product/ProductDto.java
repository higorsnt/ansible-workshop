package br.edu.ufcg.middleware.dto.product;

import br.edu.ufcg.middleware.proto.product.ProductListResponse;
import br.edu.ufcg.middleware.proto.product.ProductSaveResponse;

public record ProductDto(Long id, String name, String description, Double price, Long stock) {
    public ProductDto(ProductSaveResponse.Data data) {
        this(
                data.getId(),
                data.getName(),
                data.getDescription(),
                data.getPrice(),
                data.getStock()
        );
    }

    public ProductDto(ProductListResponse.Product product) {
        this(
                (long) product.getId(),
                product.getName(),
                product.getDescription(),
                product.getPrice(),
                product.getStock()
        );
    }
}
