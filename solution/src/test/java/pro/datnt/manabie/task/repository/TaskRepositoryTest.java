package pro.datnt.manabie.task.repository;

import com.google.gson.Gson;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.jdbc.AutoConfigureTestDatabase;
import org.springframework.boot.test.autoconfigure.orm.jpa.DataJpaTest;
import pro.datnt.manabie.task.model.TaskDBO;
import pro.datnt.manabie.task.model.UserDBO;

import java.util.List;

import static org.assertj.core.api.Assertions.assertThat;


@DataJpaTest
@AutoConfigureTestDatabase(replace = AutoConfigureTestDatabase.Replace.NONE)
class TaskRepositoryTest {

    @Autowired
    TaskRepository taskRepository;

    @Autowired
    UserRepository userRepository;

    public Long userId;

    @BeforeEach
    public void setup() {
        UserDBO userDBO = new UserDBO();
        userDBO.setUsername("test");
        userDBO.setPassword("secret");
        UserDBO savedUser = userRepository.save(userDBO);
        userId = savedUser.getId();
    }

    @AfterEach
    public void cleanup() {
        userRepository.deleteAll();
        taskRepository.deleteAll();
    }

    @Test
    public void createTask() {
        // Setup
        TaskDBO task = new TaskDBO();
        task.setUserId(userId);
        task.setContent("Test task");
        // Execution
        TaskDBO createdTask = taskRepository.save(task);
        // Verification
        TaskDBO taskDb = taskRepository.getOne(createdTask.getId());
        assertThat(taskDb).isEqualTo(createdTask);
    }

    @Test
    public void countTaskByUserId() {
        // Setup
        for (int i = 0; i < 5; i++) {
            TaskDBO task = new TaskDBO();
            task.setUserId(userId);
            task.setContent("Test task " + i);
            taskRepository.save(task);
        }
        // Execution
        Integer count = taskRepository.countTask(userId);
        // Verification
        assertThat(count).isEqualTo(5);
    }

    @Test
    void findAllByUserId() {
        // Setup
        for (int i = 0; i < 5; i++) {
            TaskDBO task = new TaskDBO();
            task.setUserId(userId);
            task.setContent("Test task " + i);
            taskRepository.save(task);
        }
        // Execution
        List<TaskDBO> allByUserId = taskRepository.findAllByUserId(userId);
        // Verification
        assertThat(allByUserId.size()).isEqualTo(5);
    }
}