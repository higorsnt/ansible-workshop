package br.edu.ufcg.user_service.models;

import jakarta.persistence.*;
import lombok.*;

@Entity
@Table(name = "USER")
@Data
public class User {

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    @Column(name = "ID")
    @Setter(AccessLevel.NONE)
    private Long id;

    @Column(name = "NAME")
    private String name;

    @Column(name = "EMAIL")
    private String email;

    @Column(name = "DOCUMENT")
    private String document;

    @OneToOne(cascade = CascadeType.ALL)
    @JoinColumn(name = "ADDRESS_ID", referencedColumnName = "ID")
    private Address address;

    public User() {
    }

    public User(String name, String email, String document, Address address) {
        this.name = name;
        this.email = email;
        this.document = document;
        this.address = address;
    }
}
