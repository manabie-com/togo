package pro.datnt.manabie.task.service;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.ArgumentCaptor;
import org.mockito.Captor;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;
import pro.datnt.manabie.task.model.TaskDBO;
import pro.datnt.manabie.task.model.UserDBO;
import pro.datnt.manabie.task.repository.TaskRepository;
import pro.datnt.manabie.task.repository.UserRepository;

import static org.assertj.core.api.Assertions.assertThat;
import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
class TaskServiceTest {

    @Mock
    TaskRepository taskRepository;

    @Mock
    UserRepository userRepository;

    @InjectMocks
    TaskService taskService;

    @Captor
    ArgumentCaptor<TaskDBO> taskDBOArgumentCaptor;

    @Test
    void createTask() {
        // Setup
        UserDBO userDBO = new UserDBO();
        userDBO.setId(1L);
        userDBO.setMaxTodo(5);
        when(userRepository.getOne(1L)).thenReturn(userDBO);
        when(taskRepository.countTask(1L)).thenReturn(2);
        taskService.createTask("content", 1L);
        // Verification
        verify(taskRepository).save(taskDBOArgumentCaptor.capture());
        TaskDBO value = taskDBOArgumentCaptor.getValue();
        assertThat(value.getContent()).isEqualTo("content");
        assertThat(value.getUserId()).isEqualTo(1L);
    }
}