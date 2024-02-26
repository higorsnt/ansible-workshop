package br.edu.ufcg.user_service.dto.address;

import br.edu.ufcg.user_service.models.Address;

public record AddressDto(String street, String city, String state, String number) {
    public AddressDto(Address address) {
        this(
                address.getStreet(),
                address.getCity(),
                address.getState(),
                address.getNumber()
        );
    }

    public Address toEntity() {
        return new Address(street, city, state, number);
    }
}
