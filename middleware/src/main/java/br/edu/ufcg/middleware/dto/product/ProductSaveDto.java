package br.edu.ufcg.middleware.dto.product;

public record ProductSaveDto(String name, String description, Double price, Integer stock) {
}
