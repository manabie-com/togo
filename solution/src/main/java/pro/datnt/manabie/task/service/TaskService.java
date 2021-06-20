package pro.datnt.manabie.task.service;

import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import pro.datnt.manabie.task.model.TaskDBO;
import pro.datnt.manabie.task.model.UserDBO;
import pro.datnt.manabie.task.repository.TaskRepository;
import pro.datnt.manabie.task.repository.UserRepository;

import java.util.List;

@Service
@RequiredArgsConstructor
public class TaskService {
    private final TaskRepository taskRepository;
    private final UserRepository userRepository;

    public TaskDBO createTask(String content, Long userId) {
        UserDBO userDBO = userRepository.getOne(userId);
        Integer totalTask = taskRepository.countTask(userId);
        if (totalTask >= userDBO.getMaxTodo()) {
            throw new IllegalArgumentException("User dont have permission to create task");
        }
        TaskDBO taskDBO = new TaskDBO();
        taskDBO.setContent(content);
        taskDBO.setUserId(userId);
        return taskRepository.save(taskDBO);
    }

    public List<TaskDBO> list(Long userId) {
        return taskRepository.findAllByUserId(userId);
    }
}
