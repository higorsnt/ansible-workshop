package br.edu.ufcg.middleware.controller;

import br.edu.ufcg.middleware.dto.product.ProductDto;
import br.edu.ufcg.middleware.dto.product.ProductSaveDto;
import br.edu.ufcg.middleware.dto.product.ProductUpdateDto;
import br.edu.ufcg.middleware.proto.product.ProductUpdateRequest;
import br.edu.ufcg.middleware.proto.product.ProductUpdateResponse;
import br.edu.ufcg.middleware.proto.product.*;
import br.edu.ufcg.middleware.proto.product.ProductServiceGrpc.ProductServiceBlockingStub;
import net.devh.boot.grpc.client.inject.GrpcClient;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/product")
public class ProductController {

    @GrpcClient("order-service")
    private ProductServiceBlockingStub productServiceStub;

    @PostMapping
    @ResponseStatus(HttpStatus.CREATED)
    public ProductDto save(@RequestBody ProductSaveDto productDto) {
        ProductSaveRequest req = ProductSaveRequest.newBuilder()
                .setName(productDto.name())
                .setDescription(productDto.description())
                .setPrice(productDto.price())
                .setStock(productDto.stock())
                .build();
        ProductSaveResponse.Data data = this.productServiceStub.save(req).getData();
        return new ProductDto(data);
    }

    @GetMapping
    @ResponseStatus(HttpStatus.OK)
    public List<ProductDto> list(
            @RequestParam(value = "name", required = false) String name,
            @RequestParam(value = "priceMax", required = false) Double priceMin,
            @RequestParam(value = "priceMax", required = false) Double priceMax
    ) {
        ProductListRequest.Filters filters = ProductListRequest.Filters.newBuilder()
                .setName(name)
                .setPriceMin(priceMin)
                .setPriceMax(priceMax)
                .build();
        ProductListRequest request = ProductListRequest.newBuilder()
                .setFilters(filters)
                .build();
        return this.productServiceStub.list(request).getProductsList()
                .stream()
                .map(ProductDto::new)
                .toList();
    }

    @PatchMapping
    @ResponseStatus(HttpStatus.OK)
    public ProductDto update(@RequestBody ProductUpdateDto productDto) {
        ProductUpdateRequest req = ProductUpdateRequest.newBuilder()
                .setName(productDto.name())
                .setDescription(productDto.description())
                .setPrice(productDto.price())
                .setStock(productDto.stock())
                .build();
        ProductUpdateResponse data = this.productServiceStub.update(req);
        return new ProductDto(
                (long) data.getId(),
                data.getName(),
                data.getDescription(),
                data.getPrice(),
                data.getStock()
        );
    }

    @DeleteMapping("/{id}")
    @ResponseStatus(HttpStatus.OK)
    public void delete(@PathVariable Integer id) {
        ProductDeleteRequest req = ProductDeleteRequest.newBuilder()
                .setId(id)
                .build();
        this.productServiceStub.delete(req);
    }
}
