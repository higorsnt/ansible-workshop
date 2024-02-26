package br.edu.ufcg.middleware.dto.user;

import br.edu.ufcg.middleware.dto.address.AddressDto;

public record UserSaveDto(String name, String email, String document, AddressDto address) {
}
