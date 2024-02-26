package br.edu.ufcg.middleware.dto.company;

import br.edu.ufcg.middleware.dto.address.AddressDto;

public record CompanySaveDto(String name, String email, AddressDto addressDto) {
}
