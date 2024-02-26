package br.edu.ufcg.middleware.dto.product;

public record ProductUpdateDto(Long id, String name, String description, Double price, Long stock) {
}
