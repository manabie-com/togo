package com.todo.service.task;

import com.todo.common.LocalDateConverter;
import com.todo.entity.TodoTask;
import com.todo.model.TodoTaskDTO;
import com.todo.repository.TodoTaskRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.time.LocalDate;
import java.util.List;
import java.util.Optional;

@Service
public class TodoTaskServiceImpl implements TodoTaskService {
    @Autowired
    private TodoTaskRepository todoTaskRepository;

    @Override
    public List<TodoTask> findTodoTaskByAppAccount_Username(String username) {
        return todoTaskRepository.findTodoTaskByAppAccount_Username(username);
    }

    @Override
    public int countByAppAccount_Username(String username) {
        return todoTaskRepository.countByAppAccount_Username(username);
    }

    @Override
    public TodoTask create(TodoTask todoTask) {
        return todoTaskRepository.save(todoTask);
    }

    @Override
    public TodoTask update(Long id, TodoTaskDTO dto) {
        Optional<TodoTask> taskOptional = todoTaskRepository.findById(id);
        if (taskOptional.isPresent()) {
            TodoTask task = taskOptional.get();
            task.setContent(dto.getContent());
            if (dto.isComplete()) {
                task.setIsComplete(true);
                task.setCompleteTime(new LocalDateConverter().convertToDatabaseColumn(LocalDate.now()));
            } else {
                task.setIsComplete(false);
                task.setCompleteTime(null);
            }
            return todoTaskRepository.save(task);
        }
        return null;
    }
}



