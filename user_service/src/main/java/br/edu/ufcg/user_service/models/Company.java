package br.edu.ufcg.user_service.models;

import jakarta.persistence.*;
import lombok.AccessLevel;
import lombok.Data;
import lombok.Setter;

@Entity
@Table(name = "COMPANY")
@Data
public class Company {

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    @Column(name = "ID")
    @Setter(AccessLevel.NONE)
    private Long id;

    @Column(name = "NAME")
    private String name;

    @Column(name = "EMAIL")
    private String email;

    @OneToOne(cascade = CascadeType.ALL)
    @JoinColumn(name = "ADDRESS_ID", referencedColumnName = "ID")
    private Address address;

    public Company() {
    }

    public Company(String name, String email, Address address) {
        this.name = name;
        this.email = email;
        this.address = address;
    }
}
