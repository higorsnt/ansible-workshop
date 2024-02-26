package br.edu.ufcg.middleware.dto.user;

import br.edu.ufcg.middleware.dto.address.AddressDto;

public record UserDto(Long id, String name, String email, AddressDto address) {
}
