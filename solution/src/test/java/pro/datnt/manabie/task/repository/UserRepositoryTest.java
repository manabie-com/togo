package pro.datnt.manabie.task.repository;

import com.google.gson.Gson;
import org.h2.jdbc.JdbcSQLIntegrityConstraintViolationException;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.jdbc.AutoConfigureTestDatabase;
import org.springframework.boot.test.autoconfigure.orm.jpa.DataJpaTest;
import org.springframework.dao.DataIntegrityViolationException;
import pro.datnt.manabie.task.model.UserDBO;

import java.util.Optional;

import static org.assertj.core.api.Assertions.assertThat;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;

@DataJpaTest
@AutoConfigureTestDatabase(replace = AutoConfigureTestDatabase.Replace.NONE)
class UserRepositoryTest {

    @Autowired
    UserRepository userRepository;


    @Test
    public void createUser() {
        // Setup
        UserDBO user = new UserDBO();
        user.setUsername("tester");
        user.setPassword("secret");
        // Execution
        UserDBO savedUser = userRepository.save(user);
        // Verification
        UserDBO dbUser = userRepository.getOne(savedUser.getId());
        assertThat(dbUser).isEqualTo(savedUser);
        assertThat(dbUser.getMaxTodo()).isEqualTo(5);
    }

    @Test
    public void createDuplicateUserFalse() {
        // Setup
        UserDBO user = new UserDBO();
        user.setUsername("tester");
        user.setPassword("secret");

        Gson gson = new Gson();
        UserDBO copyUser = gson.fromJson(gson.toJson(user), UserDBO.class);
        // Execution
        userRepository.save(user);
        // Verification
        assertThrows(DataIntegrityViolationException.class, () -> userRepository.save(copyUser));
    }

    @Test
    public void saveEmptyUsername() {
        // Setup
        UserDBO userDBO = new UserDBO();
        // Verification
        assertThrows(DataIntegrityViolationException.class, () -> userRepository.save(userDBO));
    }


    @Test
    public void findByUsername() {
        // Setup
        UserDBO user = new UserDBO();
        user.setUsername("tester");
        user.setPassword("secret");
        // Execution
        UserDBO savedUser = userRepository.save(user);
        Optional<UserDBO> dbUser = userRepository.findByUsername("tester");
        // Verification
        assertThat(savedUser).isEqualTo(dbUser.get());
    }
}