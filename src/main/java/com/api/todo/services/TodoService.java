package com.api.todo.services;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.api.todo.entities.Task;
import com.api.todo.repositories.TaskRepository;

@Service
public class TodoService {
    @Autowired
    private TaskRepository taskRepository;

    public TodoService() {}
    public TodoService(TaskRepository taskRepository) {
		this.taskRepository = taskRepository;
	}

	public Task createTask(Task task) {
        return taskRepository.save(task);
    }

    public int countTaskOfOneUser(long user_id, String createdDate) {
        return taskRepository.getCountTaskOfUser(user_id, createdDate);
    }
}
