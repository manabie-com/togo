package pro.datnt.manabie.task.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;
import pro.datnt.manabie.task.model.UserDBO;

import java.util.Optional;

@Repository
public interface UserRepository extends JpaRepository<UserDBO, Long>, CrudRepository<UserDBO, Long> {
    Optional<UserDBO> findByUsername(String username);
}
