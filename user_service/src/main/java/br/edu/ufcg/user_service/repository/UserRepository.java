package br.edu.ufcg.user_service.repository;

import br.edu.ufcg.user_service.models.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface UserRepository extends JpaRepository<User, Long> {
}
