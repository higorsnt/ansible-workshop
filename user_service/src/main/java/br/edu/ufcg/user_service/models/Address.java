package br.edu.ufcg.user_service.models;

import jakarta.persistence.*;
import lombok.AccessLevel;
import lombok.Data;
import lombok.Setter;

@Entity
@Table(name = "ADDRESS")
@Data
public class Address {

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    @Column(name = "ID")
    @Setter(AccessLevel.NONE)
    private Long id;

    @Column(name = "STREET")
    private String street;

    @Column(name = "CITY")
    private String city;

    @Column(name = "STATE")
    private String state;

    @Column(name = "NUMBER")
    private String number;

    public Address() {
    }

    public Address(String street, String city, String state, String number) {
        this.street = street;
        this.city = city;
        this.state = state;
        this.number = number;
    }
}
