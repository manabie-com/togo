package com.uuhnaut69.app.user.repository;

import com.uuhnaut69.app.user.model.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

/**
 * @author uuhnaut
 */
@Repository
public interface UserRepository extends JpaRepository<User, Long> {

}
